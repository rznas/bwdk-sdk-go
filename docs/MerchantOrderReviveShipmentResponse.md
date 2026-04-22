# MerchantOrderReviveShipmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | پیام موفق‌‌ | 
**OrderUuid** | **string** | شناسه منحصر به فرد سفارش | 
**Status** | [**OrderStatusEnum**](OrderStatusEnum.md) | وضعیت فعلی سفارش  * &#x60;1&#x60; - اولیه * &#x60;2&#x60; - شروع در * &#x60;3&#x60; - در انتظار * &#x60;4&#x60; - در انتظار درگاه * &#x60;5&#x60; - منقضی شده * &#x60;6&#x60; - لغو شده * &#x60;7&#x60; - ممنوع شده توسط ما * &#x60;8&#x60; - ناموفق در پرداخت * &#x60;9&#x60; - تأیید شده توسط فروشنده * &#x60;10&#x60; - ناموفق در تأیید توسط فروشنده * &#x60;11&#x60; - فروشگاه * &#x60;12&#x60; - لغو شده توسط فروشنده * &#x60;13&#x60; - درخواست بازگرداندن وجه به مشتری به دلیل درخواست مشتری * &#x60;14&#x60; - درخواست بازگرداندن وجه به فروشنده پس از ناموفقی در تأیید توسط فروشنده * &#x60;15&#x60; - درخواست بازگرداندن وجه به مشتری پس از ناموفقی توسط فروشنده * &#x60;16&#x60; - بازگردانده شده به فروشنده پس از لغو توسط فروشنده * &#x60;17&#x60; - بازگرداندن وجه تکمیل شد * &#x60;18&#x60; - زمان مجاز برای منقضی کردن گذشته است * &#x60;19&#x60; - تحویل نشده * &#x60;20&#x60; - مرسوله | 
**StatusDisplay** | **string** | وضعیت قابل‌خواندن | 

## Methods

### NewMerchantOrderReviveShipmentResponse

`func NewMerchantOrderReviveShipmentResponse(message string, orderUuid string, status OrderStatusEnum, statusDisplay string, ) *MerchantOrderReviveShipmentResponse`

NewMerchantOrderReviveShipmentResponse instantiates a new MerchantOrderReviveShipmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantOrderReviveShipmentResponseWithDefaults

`func NewMerchantOrderReviveShipmentResponseWithDefaults() *MerchantOrderReviveShipmentResponse`

NewMerchantOrderReviveShipmentResponseWithDefaults instantiates a new MerchantOrderReviveShipmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *MerchantOrderReviveShipmentResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MerchantOrderReviveShipmentResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MerchantOrderReviveShipmentResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOrderUuid

`func (o *MerchantOrderReviveShipmentResponse) GetOrderUuid() string`

GetOrderUuid returns the OrderUuid field if non-nil, zero value otherwise.

### GetOrderUuidOk

`func (o *MerchantOrderReviveShipmentResponse) GetOrderUuidOk() (*string, bool)`

GetOrderUuidOk returns a tuple with the OrderUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderUuid

`func (o *MerchantOrderReviveShipmentResponse) SetOrderUuid(v string)`

SetOrderUuid sets OrderUuid field to given value.


### GetStatus

`func (o *MerchantOrderReviveShipmentResponse) GetStatus() OrderStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MerchantOrderReviveShipmentResponse) GetStatusOk() (*OrderStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MerchantOrderReviveShipmentResponse) SetStatus(v OrderStatusEnum)`

SetStatus sets Status field to given value.


### GetStatusDisplay

`func (o *MerchantOrderReviveShipmentResponse) GetStatusDisplay() string`

GetStatusDisplay returns the StatusDisplay field if non-nil, zero value otherwise.

### GetStatusDisplayOk

`func (o *MerchantOrderReviveShipmentResponse) GetStatusDisplayOk() (*string, bool)`

GetStatusDisplayOk returns a tuple with the StatusDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDisplay

`func (o *MerchantOrderReviveShipmentResponse) SetStatusDisplay(v string)`

SetStatusDisplay sets StatusDisplay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


