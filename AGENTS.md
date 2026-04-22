# BWDK Go SDK — Integration Guide for Agents

You are integrating **BWDK (Buy With DigiKala)** into a Go project (net/http, Gin, Echo, Fiber, Chi) via this SDK. Read this file **first**, then consult the companion references below.

## BWDK constants

- **Host:** `https://bwdk-backend.digify.shop`
- **Auth scheme:** `MerchantAPIKeyAuth` — header `Authorization: <api_key>`.
- **Module path:** import alias `bwdk_sdk` — the generated code lives at the module path defined in `go.mod` (typically `github.com/rznas/bwdk-sdk-go`; see `go.mod` at the SDK root).
- **Main client:** `bwdk_sdk.APIClient` with field `DefaultAPI` exposing every endpoint method.

## Companion references

| File                   | When to read                                                |
|------------------------|-------------------------------------------------------------|
| `README.md`            | Install (`go get` / `go mod tidy`), server / proxy / context configuration, and "Getting Started" example. Follow it verbatim. |
| `FLOWCHART.md`         | Canonical order state machine. All callback/webhook branching MUST match these state names. |
| `docs/DefaultAPI.md`   | Exact method names and signatures per endpoint.             |
| `docs/*.md`            | Per-model reference (e.g. `docs/OrderCreate.md`).           |

Do **not** duplicate install or method-signature details here — fetch `README.md`.

## Minimal wrapper pattern

```go
import (
    "context"
    "os"

    bwdk_sdk "github.com/rznas/bwdk-sdk-go"
)

cfg := bwdk_sdk.NewConfiguration()
cfg.Servers = bwdk_sdk.ServerConfigurations{
    {URL: "https://bwdk-backend.digify.shop"},
}
client := bwdk_sdk.NewAPIClient(cfg)

ctx := context.WithValue(context.Background(),
    bwdk_sdk.ContextAPIKeys,
    map[string]bwdk_sdk.APIKey{
        "MerchantAPIKeyAuth": {Key: os.Getenv("BWDK_API_KEY")},
    })

resp, httpResp, err := client.DefaultAPI.OrderApiV1CreateOrderCreate(ctx).
    OrderCreate(payload).Execute()
```

Method names are PascalCase and OpenAPI-generated (e.g. `OrderApiV1CreateOrderCreate`, `OrderApiV1ManagerVerifyCreate`). Look them up in `docs/DefaultAPI.md`; do **not** guess.

## Integration invariants

1. **SDK only.** Never call BWDK with `net/http` directly, `resty`, `go-resty`, or `gorequest`.
2. **Callback flow.** After payment, BWDK redirects the customer to your `callback_url`. Load the order (`OrderApiV1ManagerRetrieve`), switch on `order.Status` per `FLOWCHART.md`, then call `OrderApiV1ManagerVerifyCreate` — `verify` is mandatory before `SHIPPED`.
3. **Errors.** All methods return `(response, *http.Response, error)`. Check `err` first, then `httpResp.StatusCode`. Retry only on network errors (`net.OpError`, timeouts), never on 4xx.
4. **Context:** always pass a request-scoped `context.Context`; do not use `context.Background()` in handlers — you'll leak goroutines on cancellation.
5. **Pinning.** Pin a concrete tag in `go.mod` (`require ... vX.Y.Z`); `go get` → `latest` is forbidden in production `go.mod`.

## Project conventions

- One `*APIClient` per process; it's safe for concurrent use. Stash it in your DI container (or a top-level `var`).
- Pointer fields in generated models (`*string`, `*int64`) are optional — use the `*Order.SetField(...)` / `GetField()` / `GetFieldOk()` helpers rather than dereferencing directly.
- Use `context.WithTimeout` for outbound calls from handlers (e.g. 5s) — the SDK respects the context deadline.
