# \MerchantWalletAPI

All URIs are relative to *https://bwdk-backend.digify.shop*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WalletsApiV1WalletBalanceRetrieve**](MerchantWalletAPI.md#WalletsApiV1WalletBalanceRetrieve) | **Get** /wallets/api/v1/wallet-balance/ | Get Wallet Balance



## WalletsApiV1WalletBalanceRetrieve

> WalletBalance WalletsApiV1WalletBalanceRetrieve(ctx).Execute()

Get Wallet Balance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantWalletAPI.WalletsApiV1WalletBalanceRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantWalletAPI.WalletsApiV1WalletBalanceRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletsApiV1WalletBalanceRetrieve`: WalletBalance
	fmt.Fprintf(os.Stdout, "Response from `MerchantWalletAPI.WalletsApiV1WalletBalanceRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWalletsApiV1WalletBalanceRetrieveRequest struct via the builder pattern


### Return type

[**WalletBalance**](WalletBalance.md)

### Authorization

[MerchantAPIKeyAuth](../README.md#MerchantAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

