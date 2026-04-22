# \SellerProfileManagementAPI

All URIs are relative to *https://bwdk-backend.digify.shop*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantApiV1AuthStatusRetrieve**](SellerProfileManagementAPI.md#MerchantApiV1AuthStatusRetrieve) | **Get** /merchant/api/v1/auth/status/ | وضعیت لاگین بودن



## MerchantApiV1AuthStatusRetrieve

> AuthStatusResponse MerchantApiV1AuthStatusRetrieve(ctx).Execute()

وضعیت لاگین بودن



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
	resp, r, err := apiClient.SellerProfileManagementAPI.MerchantApiV1AuthStatusRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SellerProfileManagementAPI.MerchantApiV1AuthStatusRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantApiV1AuthStatusRetrieve`: AuthStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `SellerProfileManagementAPI.MerchantApiV1AuthStatusRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMerchantApiV1AuthStatusRetrieveRequest struct via the builder pattern


### Return type

[**AuthStatusResponse**](AuthStatusResponse.md)

### Authorization

[MerchantAPIKeyAuth](../README.md#MerchantAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

