# \MerchantOrdersAPI

All URIs are relative to *https://bwdk-backend.digify.shop*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrderApiV1CreateOrderCreate**](MerchantOrdersAPI.md#OrderApiV1CreateOrderCreate) | **Post** /order/api/v1/create-order/ | ساخت سفارش
[**OrderApiV1ManagerList**](MerchantOrdersAPI.md#OrderApiV1ManagerList) | **Get** /order/api/v1/manager/ | لیست سفارشات
[**OrderApiV1ManagerPaidList**](MerchantOrdersAPI.md#OrderApiV1ManagerPaidList) | **Get** /order/api/v1/manager/paid/ | سفارش پرداخت‌شده و تایید‌نشده
[**OrderApiV1ManagerRefundCreate**](MerchantOrdersAPI.md#OrderApiV1ManagerRefundCreate) | **Post** /order/api/v1/manager/{order_uuid}/refund/ | بازگشت سفارش
[**OrderApiV1ManagerRetrieve**](MerchantOrdersAPI.md#OrderApiV1ManagerRetrieve) | **Get** /order/api/v1/manager/{order_uuid}/ | دریافت سفارش
[**OrderApiV1ManagerUpdateStatusUpdate**](MerchantOrdersAPI.md#OrderApiV1ManagerUpdateStatusUpdate) | **Put** /order/api/v1/manager/{order_uuid}/update-status/ | Update Order Status
[**OrderApiV1ManagerVerifyCreate**](MerchantOrdersAPI.md#OrderApiV1ManagerVerifyCreate) | **Post** /order/api/v1/manager/{order_uuid}/verify/ | تایید سفارش



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
	resp, r, err := apiClient.MerchantOrdersAPI.OrderApiV1CreateOrderCreate(context.Background()).OrderCreate(orderCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantOrdersAPI.OrderApiV1CreateOrderCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1CreateOrderCreate`: OrderCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `MerchantOrdersAPI.OrderApiV1CreateOrderCreate`: %v\n", resp)
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
	status := int32(56) // int32 | * `1` - اولیه * `2` - شروع شده * `3` - در انتظار * `4` - در انتظار درگاه * `5` - منقضی شده * `6` - لغو شده * `7` - پرداخت‌شده توسط کاربر * `8` - پرداخت موفیت آمیز نبود * `9` - تأیید شده توسط فروشگاه * `10` - تأیید توسط فروشگاه ناموفق بود * `11` - ناموفق از سوی فروشگاه * `12` - لغوشده توسط فروشگاه * `13` - درخواست بازگشت وجه به مشتری به دلیل درخواست مشتری * `14` - درخواست بازگشت وجه به فروشگاه پس از عدم تأیید توسط فروشگاه * `15` - درخواست بازگشت وجه به مشتری پس از ناموفق بودن توسط فروشگاه * `16` - بازپرداخت به فروشگاه پس از لغو توسط فروشگاه * `17` - بازپرداخت تکمیل شد * `18` - زمان انقضا گذشته است * `19` - تحویل شده * `20` - جمع اوری شده و در حال ارسال (optional)
	statuses := "statuses_example" // string |  (optional)
	todayPickup := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantOrdersAPI.OrderApiV1ManagerList(context.Background()).Cities(cities).CreatedAt(createdAt).Cursor(cursor).OrderIds(orderIds).Ordering(ordering).PaymentTypes(paymentTypes).Provinces(provinces).ReferenceCode(referenceCode).Search(search).ShippingTypes(shippingTypes).Status(status).Statuses(statuses).TodayPickup(todayPickup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantOrdersAPI.OrderApiV1ManagerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1ManagerList`: PaginatedOrderDetailList
	fmt.Fprintf(os.Stdout, "Response from `MerchantOrdersAPI.OrderApiV1ManagerList`: %v\n", resp)
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
 **status** | **int32** | * &#x60;1&#x60; - اولیه * &#x60;2&#x60; - شروع شده * &#x60;3&#x60; - در انتظار * &#x60;4&#x60; - در انتظار درگاه * &#x60;5&#x60; - منقضی شده * &#x60;6&#x60; - لغو شده * &#x60;7&#x60; - پرداخت‌شده توسط کاربر * &#x60;8&#x60; - پرداخت موفیت آمیز نبود * &#x60;9&#x60; - تأیید شده توسط فروشگاه * &#x60;10&#x60; - تأیید توسط فروشگاه ناموفق بود * &#x60;11&#x60; - ناموفق از سوی فروشگاه * &#x60;12&#x60; - لغوشده توسط فروشگاه * &#x60;13&#x60; - درخواست بازگشت وجه به مشتری به دلیل درخواست مشتری * &#x60;14&#x60; - درخواست بازگشت وجه به فروشگاه پس از عدم تأیید توسط فروشگاه * &#x60;15&#x60; - درخواست بازگشت وجه به مشتری پس از ناموفق بودن توسط فروشگاه * &#x60;16&#x60; - بازپرداخت به فروشگاه پس از لغو توسط فروشگاه * &#x60;17&#x60; - بازپرداخت تکمیل شد * &#x60;18&#x60; - زمان انقضا گذشته است * &#x60;19&#x60; - تحویل شده * &#x60;20&#x60; - جمع اوری شده و در حال ارسال | 
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
	status := int32(56) // int32 | * `1` - اولیه * `2` - شروع شده * `3` - در انتظار * `4` - در انتظار درگاه * `5` - منقضی شده * `6` - لغو شده * `7` - پرداخت‌شده توسط کاربر * `8` - پرداخت موفیت آمیز نبود * `9` - تأیید شده توسط فروشگاه * `10` - تأیید توسط فروشگاه ناموفق بود * `11` - ناموفق از سوی فروشگاه * `12` - لغوشده توسط فروشگاه * `13` - درخواست بازگشت وجه به مشتری به دلیل درخواست مشتری * `14` - درخواست بازگشت وجه به فروشگاه پس از عدم تأیید توسط فروشگاه * `15` - درخواست بازگشت وجه به مشتری پس از ناموفق بودن توسط فروشگاه * `16` - بازپرداخت به فروشگاه پس از لغو توسط فروشگاه * `17` - بازپرداخت تکمیل شد * `18` - زمان انقضا گذشته است * `19` - تحویل شده * `20` - جمع اوری شده و در حال ارسال (optional)
	statuses := "statuses_example" // string |  (optional)
	todayPickup := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantOrdersAPI.OrderApiV1ManagerPaidList(context.Background()).Cities(cities).CreatedAt(createdAt).Cursor(cursor).OrderIds(orderIds).Ordering(ordering).PaymentTypes(paymentTypes).Provinces(provinces).ReferenceCode(referenceCode).Search(search).ShippingTypes(shippingTypes).Status(status).Statuses(statuses).TodayPickup(todayPickup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantOrdersAPI.OrderApiV1ManagerPaidList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1ManagerPaidList`: PaginatedMerchantPaidOrderListList
	fmt.Fprintf(os.Stdout, "Response from `MerchantOrdersAPI.OrderApiV1ManagerPaidList`: %v\n", resp)
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
 **status** | **int32** | * &#x60;1&#x60; - اولیه * &#x60;2&#x60; - شروع شده * &#x60;3&#x60; - در انتظار * &#x60;4&#x60; - در انتظار درگاه * &#x60;5&#x60; - منقضی شده * &#x60;6&#x60; - لغو شده * &#x60;7&#x60; - پرداخت‌شده توسط کاربر * &#x60;8&#x60; - پرداخت موفیت آمیز نبود * &#x60;9&#x60; - تأیید شده توسط فروشگاه * &#x60;10&#x60; - تأیید توسط فروشگاه ناموفق بود * &#x60;11&#x60; - ناموفق از سوی فروشگاه * &#x60;12&#x60; - لغوشده توسط فروشگاه * &#x60;13&#x60; - درخواست بازگشت وجه به مشتری به دلیل درخواست مشتری * &#x60;14&#x60; - درخواست بازگشت وجه به فروشگاه پس از عدم تأیید توسط فروشگاه * &#x60;15&#x60; - درخواست بازگشت وجه به مشتری پس از ناموفق بودن توسط فروشگاه * &#x60;16&#x60; - بازپرداخت به فروشگاه پس از لغو توسط فروشگاه * &#x60;17&#x60; - بازپرداخت تکمیل شد * &#x60;18&#x60; - زمان انقضا گذشته است * &#x60;19&#x60; - تحویل شده * &#x60;20&#x60; - جمع اوری شده و در حال ارسال | 
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
	resp, r, err := apiClient.MerchantOrdersAPI.OrderApiV1ManagerRefundCreate(context.Background(), orderUuid).RefundOrder(refundOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantOrdersAPI.OrderApiV1ManagerRefundCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1ManagerRefundCreate`: MerchantOrderRefundResponse
	fmt.Fprintf(os.Stdout, "Response from `MerchantOrdersAPI.OrderApiV1ManagerRefundCreate`: %v\n", resp)
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
	resp, r, err := apiClient.MerchantOrdersAPI.OrderApiV1ManagerRetrieve(context.Background(), orderUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantOrdersAPI.OrderApiV1ManagerRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1ManagerRetrieve`: OrderDetail
	fmt.Fprintf(os.Stdout, "Response from `MerchantOrdersAPI.OrderApiV1ManagerRetrieve`: %v\n", resp)
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


## OrderApiV1ManagerUpdateStatusUpdate

> OrderDetail OrderApiV1ManagerUpdateStatusUpdate(ctx, orderUuid).UpdateOrderStatus(updateOrderStatus).Execute()

Update Order Status



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
	resp, r, err := apiClient.MerchantOrdersAPI.OrderApiV1ManagerUpdateStatusUpdate(context.Background(), orderUuid).UpdateOrderStatus(updateOrderStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantOrdersAPI.OrderApiV1ManagerUpdateStatusUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1ManagerUpdateStatusUpdate`: OrderDetail
	fmt.Fprintf(os.Stdout, "Response from `MerchantOrdersAPI.OrderApiV1ManagerUpdateStatusUpdate`: %v\n", resp)
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
	resp, r, err := apiClient.MerchantOrdersAPI.OrderApiV1ManagerVerifyCreate(context.Background(), orderUuid).VerifyOrder(verifyOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantOrdersAPI.OrderApiV1ManagerVerifyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderApiV1ManagerVerifyCreate`: OrderDetail
	fmt.Fprintf(os.Stdout, "Response from `MerchantOrdersAPI.OrderApiV1ManagerVerifyCreate`: %v\n", resp)
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

