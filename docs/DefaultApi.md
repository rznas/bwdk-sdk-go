# \DefaultApi

All URIs are relative to *https://bwdk-backend.digify.shop*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantApiV1AuthStatusRetrieve**](DefaultApi.md#MerchantApiV1AuthStatusRetrieve) | **Get** /merchant/api/v1/auth/status/ | وضعیت لاگین بودن
[**OrderApiV1CreateOrderCreate**](DefaultApi.md#OrderApiV1CreateOrderCreate) | **Post** /order/api/v1/create-order/ | ساخت سفارش
[**OrderApiV1ManagerCancelShipmentCreate**](DefaultApi.md#OrderApiV1ManagerCancelShipmentCreate) | **Post** /order/api/v1/manager/{order_uuid}/cancel-shipment/ | لغو ارسال سفارش
[**OrderApiV1ManagerChangeShippingMethodUpdate**](DefaultApi.md#OrderApiV1ManagerChangeShippingMethodUpdate) | **Put** /order/api/v1/manager/{order_uuid}/change-shipping-method/ | تغییر روش ارسال
[**OrderApiV1ManagerList**](DefaultApi.md#OrderApiV1ManagerList) | **Get** /order/api/v1/manager/ | لیست سفارشات
[**OrderApiV1ManagerPaidList**](DefaultApi.md#OrderApiV1ManagerPaidList) | **Get** /order/api/v1/manager/paid/ | سفارش پرداخت‌شده و تایید‌نشده
[**OrderApiV1ManagerRefundCreate**](DefaultApi.md#OrderApiV1ManagerRefundCreate) | **Post** /order/api/v1/manager/{order_uuid}/refund/ | بازگشت سفارش
[**OrderApiV1ManagerRetrieve**](DefaultApi.md#OrderApiV1ManagerRetrieve) | **Get** /order/api/v1/manager/{order_uuid}/ | دریافت سفارش
[**OrderApiV1ManagerReviveShipmentCreate**](DefaultApi.md#OrderApiV1ManagerReviveShipmentCreate) | **Post** /order/api/v1/manager/{order_uuid}/revive-shipment/ | احیای ارسال سفارش
[**OrderApiV1ManagerUpdateStatusUpdate**](DefaultApi.md#OrderApiV1ManagerUpdateStatusUpdate) | **Put** /order/api/v1/manager/{order_uuid}/update-status/ | به‌روزرسانی وضعیت سفارش
[**OrderApiV1ManagerVerifyCreate**](DefaultApi.md#OrderApiV1ManagerVerifyCreate) | **Post** /order/api/v1/manager/{order_uuid}/verify/ | تایید سفارش
[**WalletsApiV1WalletBalanceRetrieve**](DefaultApi.md#WalletsApiV1WalletBalanceRetrieve) | **Get** /wallets/api/v1/wallet-balance/ | دریافت موجودی کیف پول



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
	resp, r, err := apiClient.DefaultApi.MerchantApiV1AuthStatusRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MerchantApiV1AuthStatusRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantApiV1AuthStatusRetrieve`: AuthStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MerchantApiV1AuthStatusRetrieve`: %v\n", resp)
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


## OrderApiV1CreateOrderCreate

> OrderCreateResponse OrderApiV1CreateOrderCreate(ctx).OrderCreate(orderCreate).Execute()

ساخت سفارش



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
	orderCreate := *openapiclient.NewOrderCreate("MerchantOrderId_example", "MerchantUniqueId_example", "CallbackUrl_example", interface{}(123), []openapiclient.OrderItemCreate{*openapiclient.NewOrderItemCreate("Name_example", int32(123), []openapiclient.Option{*openapiclient.NewOption(openapiclient.TypeNameEnum("color"), "Name_example", "Value_example")})}, NullableInt32(123), "ReferenceCode_example") // OrderCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.OrderApiV1CreateOrderCreate(context.Background()).OrderCreate(orderCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrderApiV1CreateOrderCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1CreateOrderCreate`: OrderCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrderApiV1CreateOrderCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderApiV1CreateOrderCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderCreate** | [**OrderCreate**](OrderCreate.md) |  | 

### Return type

[**OrderCreateResponse**](OrderCreateResponse.md)

### Authorization

[MerchantAPIKeyAuth](../README.md#MerchantAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderApiV1ManagerCancelShipmentCreate

> MerchantOrderCancelShipmentResponse OrderApiV1ManagerCancelShipmentCreate(ctx, orderUuid).Execute()

لغو ارسال سفارش



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
	resp, r, err := apiClient.DefaultApi.OrderApiV1ManagerCancelShipmentCreate(context.Background(), orderUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrderApiV1ManagerCancelShipmentCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1ManagerCancelShipmentCreate`: MerchantOrderCancelShipmentResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrderApiV1ManagerCancelShipmentCreate`: %v\n", resp)
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

تغییر روش ارسال



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
	resp, r, err := apiClient.DefaultApi.OrderApiV1ManagerChangeShippingMethodUpdate(context.Background(), orderUuid).OrderDetail(orderDetail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrderApiV1ManagerChangeShippingMethodUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1ManagerChangeShippingMethodUpdate`: OrderDetail
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrderApiV1ManagerChangeShippingMethodUpdate`: %v\n", resp)
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


## OrderApiV1ManagerList

> PaginatedOrderDetailList OrderApiV1ManagerList(ctx).Cities(cities).CreatedAt(createdAt).Cursor(cursor).OrderIds(orderIds).Ordering(ordering).PaymentTypes(paymentTypes).Provinces(provinces).ReferenceCode(referenceCode).Search(search).ShippingTypes(shippingTypes).Status(status).Statuses(statuses).TodayPickup(todayPickup).Execute()

لیست سفارشات



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
	cities := "cities_example" // string |  (optional)
	createdAt := time.Now() // string |  (optional)
	cursor := "cursor_example" // string | مقدار نشانگر صفحه‌بندی. (optional)
	orderIds := "orderIds_example" // string |  (optional)
	ordering := "ordering_example" // string | کدام فیلد باید هنگام مرتب‌سازی نتایج استفاده شود. (optional)
	paymentTypes := "paymentTypes_example" // string |  (optional)
	provinces := "provinces_example" // string |  (optional)
	referenceCode := "referenceCode_example" // string |  (optional)
	search := "search_example" // string | یک عبارت جستجو. (optional)
	shippingTypes := "shippingTypes_example" // string |  (optional)
	status := int32(56) // int32 | * `1` - اولیه * `2` - شروع در * `3` - در انتظار * `4` - در انتظار درگاه * `5` - منقضی شده * `6` - لغو شده * `7` - ممنوع شده توسط ما * `8` - ناموفق در پرداخت * `9` - تأیید شده توسط فروشنده * `10` - ناموفق در تأیید توسط فروشنده * `11` - فروشگاه * `12` - لغو شده توسط فروشنده * `13` - درخواست بازگرداندن وجه به مشتری به دلیل درخواست مشتری * `14` - درخواست بازگرداندن وجه به فروشنده پس از ناموفقی در تأیید توسط فروشنده * `15` - درخواست بازگرداندن وجه به مشتری پس از ناموفقی توسط فروشنده * `16` - بازگردانده شده به فروشنده پس از لغو توسط فروشنده * `17` - بازگرداندن وجه تکمیل شد * `18` - زمان مجاز برای منقضی کردن گذشته است * `19` - تحویل نشده * `20` - مرسوله (optional)
	statuses := "statuses_example" // string |  (optional)
	todayPickup := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.OrderApiV1ManagerList(context.Background()).Cities(cities).CreatedAt(createdAt).Cursor(cursor).OrderIds(orderIds).Ordering(ordering).PaymentTypes(paymentTypes).Provinces(provinces).ReferenceCode(referenceCode).Search(search).ShippingTypes(shippingTypes).Status(status).Statuses(statuses).TodayPickup(todayPickup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrderApiV1ManagerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1ManagerList`: PaginatedOrderDetailList
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrderApiV1ManagerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderApiV1ManagerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cities** | **string** |  | 
 **createdAt** | **string** |  | 
 **cursor** | **string** | مقدار نشانگر صفحه‌بندی. | 
 **orderIds** | **string** |  | 
 **ordering** | **string** | کدام فیلد باید هنگام مرتب‌سازی نتایج استفاده شود. | 
 **paymentTypes** | **string** |  | 
 **provinces** | **string** |  | 
 **referenceCode** | **string** |  | 
 **search** | **string** | یک عبارت جستجو. | 
 **shippingTypes** | **string** |  | 
 **status** | **int32** | * &#x60;1&#x60; - اولیه * &#x60;2&#x60; - شروع در * &#x60;3&#x60; - در انتظار * &#x60;4&#x60; - در انتظار درگاه * &#x60;5&#x60; - منقضی شده * &#x60;6&#x60; - لغو شده * &#x60;7&#x60; - ممنوع شده توسط ما * &#x60;8&#x60; - ناموفق در پرداخت * &#x60;9&#x60; - تأیید شده توسط فروشنده * &#x60;10&#x60; - ناموفق در تأیید توسط فروشنده * &#x60;11&#x60; - فروشگاه * &#x60;12&#x60; - لغو شده توسط فروشنده * &#x60;13&#x60; - درخواست بازگرداندن وجه به مشتری به دلیل درخواست مشتری * &#x60;14&#x60; - درخواست بازگرداندن وجه به فروشنده پس از ناموفقی در تأیید توسط فروشنده * &#x60;15&#x60; - درخواست بازگرداندن وجه به مشتری پس از ناموفقی توسط فروشنده * &#x60;16&#x60; - بازگردانده شده به فروشنده پس از لغو توسط فروشنده * &#x60;17&#x60; - بازگرداندن وجه تکمیل شد * &#x60;18&#x60; - زمان مجاز برای منقضی کردن گذشته است * &#x60;19&#x60; - تحویل نشده * &#x60;20&#x60; - مرسوله | 
 **statuses** | **string** |  | 
 **todayPickup** | **bool** |  | 

### Return type

[**PaginatedOrderDetailList**](PaginatedOrderDetailList.md)

### Authorization

[MerchantAPIKeyAuth](../README.md#MerchantAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderApiV1ManagerPaidList

> PaginatedMerchantPaidOrderListList OrderApiV1ManagerPaidList(ctx).Cities(cities).CreatedAt(createdAt).Cursor(cursor).OrderIds(orderIds).Ordering(ordering).PaymentTypes(paymentTypes).Provinces(provinces).ReferenceCode(referenceCode).Search(search).ShippingTypes(shippingTypes).Status(status).Statuses(statuses).TodayPickup(todayPickup).Execute()

سفارش پرداخت‌شده و تایید‌نشده



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
	cities := "cities_example" // string |  (optional)
	createdAt := time.Now() // string |  (optional)
	cursor := "cursor_example" // string | مقدار نشانگر صفحه‌بندی. (optional)
	orderIds := "orderIds_example" // string |  (optional)
	ordering := "ordering_example" // string | کدام فیلد باید هنگام مرتب‌سازی نتایج استفاده شود. (optional)
	paymentTypes := "paymentTypes_example" // string |  (optional)
	provinces := "provinces_example" // string |  (optional)
	referenceCode := "referenceCode_example" // string |  (optional)
	search := "search_example" // string | یک عبارت جستجو. (optional)
	shippingTypes := "shippingTypes_example" // string |  (optional)
	status := int32(56) // int32 | * `1` - اولیه * `2` - شروع در * `3` - در انتظار * `4` - در انتظار درگاه * `5` - منقضی شده * `6` - لغو شده * `7` - ممنوع شده توسط ما * `8` - ناموفق در پرداخت * `9` - تأیید شده توسط فروشنده * `10` - ناموفق در تأیید توسط فروشنده * `11` - فروشگاه * `12` - لغو شده توسط فروشنده * `13` - درخواست بازگرداندن وجه به مشتری به دلیل درخواست مشتری * `14` - درخواست بازگرداندن وجه به فروشنده پس از ناموفقی در تأیید توسط فروشنده * `15` - درخواست بازگرداندن وجه به مشتری پس از ناموفقی توسط فروشنده * `16` - بازگردانده شده به فروشنده پس از لغو توسط فروشنده * `17` - بازگرداندن وجه تکمیل شد * `18` - زمان مجاز برای منقضی کردن گذشته است * `19` - تحویل نشده * `20` - مرسوله (optional)
	statuses := "statuses_example" // string |  (optional)
	todayPickup := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.OrderApiV1ManagerPaidList(context.Background()).Cities(cities).CreatedAt(createdAt).Cursor(cursor).OrderIds(orderIds).Ordering(ordering).PaymentTypes(paymentTypes).Provinces(provinces).ReferenceCode(referenceCode).Search(search).ShippingTypes(shippingTypes).Status(status).Statuses(statuses).TodayPickup(todayPickup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrderApiV1ManagerPaidList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1ManagerPaidList`: PaginatedMerchantPaidOrderListList
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrderApiV1ManagerPaidList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderApiV1ManagerPaidListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cities** | **string** |  | 
 **createdAt** | **string** |  | 
 **cursor** | **string** | مقدار نشانگر صفحه‌بندی. | 
 **orderIds** | **string** |  | 
 **ordering** | **string** | کدام فیلد باید هنگام مرتب‌سازی نتایج استفاده شود. | 
 **paymentTypes** | **string** |  | 
 **provinces** | **string** |  | 
 **referenceCode** | **string** |  | 
 **search** | **string** | یک عبارت جستجو. | 
 **shippingTypes** | **string** |  | 
 **status** | **int32** | * &#x60;1&#x60; - اولیه * &#x60;2&#x60; - شروع در * &#x60;3&#x60; - در انتظار * &#x60;4&#x60; - در انتظار درگاه * &#x60;5&#x60; - منقضی شده * &#x60;6&#x60; - لغو شده * &#x60;7&#x60; - ممنوع شده توسط ما * &#x60;8&#x60; - ناموفق در پرداخت * &#x60;9&#x60; - تأیید شده توسط فروشنده * &#x60;10&#x60; - ناموفق در تأیید توسط فروشنده * &#x60;11&#x60; - فروشگاه * &#x60;12&#x60; - لغو شده توسط فروشنده * &#x60;13&#x60; - درخواست بازگرداندن وجه به مشتری به دلیل درخواست مشتری * &#x60;14&#x60; - درخواست بازگرداندن وجه به فروشنده پس از ناموفقی در تأیید توسط فروشنده * &#x60;15&#x60; - درخواست بازگرداندن وجه به مشتری پس از ناموفقی توسط فروشنده * &#x60;16&#x60; - بازگردانده شده به فروشنده پس از لغو توسط فروشنده * &#x60;17&#x60; - بازگرداندن وجه تکمیل شد * &#x60;18&#x60; - زمان مجاز برای منقضی کردن گذشته است * &#x60;19&#x60; - تحویل نشده * &#x60;20&#x60; - مرسوله | 
 **statuses** | **string** |  | 
 **todayPickup** | **bool** |  | 

### Return type

[**PaginatedMerchantPaidOrderListList**](PaginatedMerchantPaidOrderListList.md)

### Authorization

[MerchantAPIKeyAuth](../README.md#MerchantAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderApiV1ManagerRefundCreate

> MerchantOrderRefundResponse OrderApiV1ManagerRefundCreate(ctx, orderUuid).RefundOrder(refundOrder).Execute()

بازگشت سفارش



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
	refundOrder := *openapiclient.NewRefundOrder() // RefundOrder |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.OrderApiV1ManagerRefundCreate(context.Background(), orderUuid).RefundOrder(refundOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrderApiV1ManagerRefundCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1ManagerRefundCreate`: MerchantOrderRefundResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrderApiV1ManagerRefundCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderApiV1ManagerRefundCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refundOrder** | [**RefundOrder**](RefundOrder.md) |  | 

### Return type

[**MerchantOrderRefundResponse**](MerchantOrderRefundResponse.md)

### Authorization

[MerchantAPIKeyAuth](../README.md#MerchantAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderApiV1ManagerRetrieve

> OrderDetail OrderApiV1ManagerRetrieve(ctx, orderUuid).Execute()

دریافت سفارش



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
	resp, r, err := apiClient.DefaultApi.OrderApiV1ManagerRetrieve(context.Background(), orderUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrderApiV1ManagerRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1ManagerRetrieve`: OrderDetail
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrderApiV1ManagerRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderApiV1ManagerRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrderDetail**](OrderDetail.md)

### Authorization

[MerchantAPIKeyAuth](../README.md#MerchantAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderApiV1ManagerReviveShipmentCreate

> MerchantOrderReviveShipmentResponse OrderApiV1ManagerReviveShipmentCreate(ctx, orderUuid).ReviveShipment(reviveShipment).Execute()

احیای ارسال سفارش



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
	resp, r, err := apiClient.DefaultApi.OrderApiV1ManagerReviveShipmentCreate(context.Background(), orderUuid).ReviveShipment(reviveShipment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrderApiV1ManagerReviveShipmentCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1ManagerReviveShipmentCreate`: MerchantOrderReviveShipmentResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrderApiV1ManagerReviveShipmentCreate`: %v\n", resp)
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


## OrderApiV1ManagerUpdateStatusUpdate

> OrderDetail OrderApiV1ManagerUpdateStatusUpdate(ctx, orderUuid).UpdateOrderStatus(updateOrderStatus).Execute()

به‌روزرسانی وضعیت سفارش



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
	updateOrderStatus := *openapiclient.NewUpdateOrderStatus(openapiclient.OrderStatusEnum(1)) // UpdateOrderStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.OrderApiV1ManagerUpdateStatusUpdate(context.Background(), orderUuid).UpdateOrderStatus(updateOrderStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrderApiV1ManagerUpdateStatusUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1ManagerUpdateStatusUpdate`: OrderDetail
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrderApiV1ManagerUpdateStatusUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderApiV1ManagerUpdateStatusUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrderStatus** | [**UpdateOrderStatus**](UpdateOrderStatus.md) |  | 

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


## OrderApiV1ManagerVerifyCreate

> OrderDetail OrderApiV1ManagerVerifyCreate(ctx, orderUuid).VerifyOrder(verifyOrder).Execute()

تایید سفارش



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
	verifyOrder := *openapiclient.NewVerifyOrder("MerchantUniqueId_example") // VerifyOrder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.OrderApiV1ManagerVerifyCreate(context.Background(), orderUuid).VerifyOrder(verifyOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrderApiV1ManagerVerifyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1ManagerVerifyCreate`: OrderDetail
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrderApiV1ManagerVerifyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderApiV1ManagerVerifyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verifyOrder** | [**VerifyOrder**](VerifyOrder.md) |  | 

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


## WalletsApiV1WalletBalanceRetrieve

> WalletBalance WalletsApiV1WalletBalanceRetrieve(ctx).Execute()

دریافت موجودی کیف پول



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
	resp, r, err := apiClient.DefaultApi.WalletsApiV1WalletBalanceRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.WalletsApiV1WalletBalanceRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletsApiV1WalletBalanceRetrieve`: WalletBalance
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.WalletsApiV1WalletBalanceRetrieve`: %v\n", resp)
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

