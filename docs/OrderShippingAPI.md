# \OrderShippingAPI

All URIs are relative to *https://bwdk-backend.digify.shop*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrderApiV1ManagerCancelShipmentCreate**](OrderShippingAPI.md#OrderApiV1ManagerCancelShipmentCreate) | **Post** /order/api/v1/manager/{order_uuid}/cancel-shipment/ | Cancel Shipment
[**OrderApiV1ManagerChangeShippingMethodUpdate**](OrderShippingAPI.md#OrderApiV1ManagerChangeShippingMethodUpdate) | **Put** /order/api/v1/manager/{order_uuid}/change-shipping-method/ | Change Shipping Method
[**OrderApiV1ManagerReviveShipmentCreate**](OrderShippingAPI.md#OrderApiV1ManagerReviveShipmentCreate) | **Post** /order/api/v1/manager/{order_uuid}/revive-shipment/ | Revive Shipment



## OrderApiV1ManagerCancelShipmentCreate

> MerchantOrderCancelShipmentResponse OrderApiV1ManagerCancelShipmentCreate(ctx, orderUuid).Execute()

Cancel Shipment



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
	orderUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderShippingAPI.OrderApiV1ManagerCancelShipmentCreate(context.Background(), orderUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderShippingAPI.OrderApiV1ManagerCancelShipmentCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1ManagerCancelShipmentCreate`: MerchantOrderCancelShipmentResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderShippingAPI.OrderApiV1ManagerCancelShipmentCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderApiV1ManagerCancelShipmentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MerchantOrderCancelShipmentResponse**](MerchantOrderCancelShipmentResponse.md)

### Authorization

[MerchantAPIKeyAuth](../README.md#MerchantAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderApiV1ManagerChangeShippingMethodUpdate

> OrderDetail OrderApiV1ManagerChangeShippingMethodUpdate(ctx, orderUuid).OrderDetail(orderDetail).Execute()

Change Shipping Method



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	orderUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	orderDetail := *openapiclient.NewOrderDetail(int32(123), time.Now(), "OrderUuid_example", NullableInt32(123), "MerchantOrderId_example", openapiclient.OrderStatusEnum(1), "StatusDisplay_example", int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), "CallbackUrl_example", *openapiclient.NewMerchant("Domain_example", "Logo_example"), []openapiclient.OrderItemCreate{*openapiclient.NewOrderItemCreate("Name_example", int32(123), []openapiclient.Option{*openapiclient.NewOption(openapiclient.TypeNameEnum("color"), "Name_example", "Value_example")})}, interface{}(123), interface{}(123), *openapiclient.NewShippingMethod(int32(123), "Name_example", "GetShippingTypeDisplay_example", "ShippingTypeDisplay_example", "DeliveryTimeDisplay_example", *openapiclient.NewDeliveryTimeRangeDisplay("MinDate_example", "MaxDate_example"), *openapiclient.NewBusinessAddress(int32(123), "Address_example", "CityName_example", "StateName_example", time.Now(), time.Now())), time.Now(), time.Now(), int32(123), time.Now(), *openapiclient.NewPacking(int32(123), "Name_example"), false, false, false, false, false, *openapiclient.NewOrderUser(), *openapiclient.NewPaymentOrder(int32(123), openapiclient.GatewayTypeEnum(1), "GatewayTypeDisplay_example", "PaidAt_example", "GatewayName_example", "InvoiceId_example", "SettlementDate_example", int32(123), "SettlementStatusDisplay_example"), int32(123), float64(123), map[string]interface{}{"key": interface{}(123)}, "ReferenceCode_example", float64(123), map[string]interface{}{"key": interface{}(123)}, int32(123), int32(123), "TODO", "PreviousStatusDisplay_example") // OrderDetail | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderShippingAPI.OrderApiV1ManagerChangeShippingMethodUpdate(context.Background(), orderUuid).OrderDetail(orderDetail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderShippingAPI.OrderApiV1ManagerChangeShippingMethodUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1ManagerChangeShippingMethodUpdate`: OrderDetail
	fmt.Fprintf(os.Stdout, "Response from `OrderShippingAPI.OrderApiV1ManagerChangeShippingMethodUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderApiV1ManagerChangeShippingMethodUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderDetail** | [**OrderDetail**](OrderDetail.md) |  | 

### Return type

[**OrderDetail**](OrderDetail.md)

### Authorization

[MerchantAPIKeyAuth](../README.md#MerchantAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderApiV1ManagerReviveShipmentCreate

> MerchantOrderReviveShipmentResponse OrderApiV1ManagerReviveShipmentCreate(ctx, orderUuid).ReviveShipment(reviveShipment).Execute()

Revive Shipment



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
	orderUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	reviveShipment := *openapiclient.NewReviveShipment() // ReviveShipment |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderShippingAPI.OrderApiV1ManagerReviveShipmentCreate(context.Background(), orderUuid).ReviveShipment(reviveShipment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderShippingAPI.OrderApiV1ManagerReviveShipmentCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1ManagerReviveShipmentCreate`: MerchantOrderReviveShipmentResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderShippingAPI.OrderApiV1ManagerReviveShipmentCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderApiV1ManagerReviveShipmentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reviveShipment** | [**ReviveShipment**](ReviveShipment.md) |  | 

### Return type

[**MerchantOrderReviveShipmentResponse**](MerchantOrderReviveShipmentResponse.md)

### Authorization

[MerchantAPIKeyAuth](../README.md#MerchantAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

