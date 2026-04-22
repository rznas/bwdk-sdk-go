# MerchantOrderRefundResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**OrderUuid** | **string** |  | 
**Status** | [**OrderStatusEnum**](OrderStatusEnum.md) |  | 
**StatusDisplay** | **string** |  | 
**RefundReason** | **string** |  | 

## Methods

### NewMerchantOrderRefundResponse

`func NewMerchantOrderRefundResponse(message string, orderUuid string, status OrderStatusEnum, statusDisplay string, refundReason string, ) *MerchantOrderRefundResponse`

NewMerchantOrderRefundResponse instantiates a new MerchantOrderRefundResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantOrderRefundResponseWithDefaults

`func NewMerchantOrderRefundResponseWithDefaults() *MerchantOrderRefundResponse`

NewMerchantOrderRefundResponseWithDefaults instantiates a new MerchantOrderRefundResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *MerchantOrderRefundResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MerchantOrderRefundResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MerchantOrderRefundResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOrderUuid

`func (o *MerchantOrderRefundResponse) GetOrderUuid() string`

GetOrderUuid returns the OrderUuid field if non-nil, zero value otherwise.

### GetOrderUuidOk

`func (o *MerchantOrderRefundResponse) GetOrderUuidOk() (*string, bool)`

GetOrderUuidOk returns a tuple with the OrderUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderUuid

`func (o *MerchantOrderRefundResponse) SetOrderUuid(v string)`

SetOrderUuid sets OrderUuid field to given value.


### GetStatus

`func (o *MerchantOrderRefundResponse) GetStatus() OrderStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MerchantOrderRefundResponse) GetStatusOk() (*OrderStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MerchantOrderRefundResponse) SetStatus(v OrderStatusEnum)`

SetStatus sets Status field to given value.


### GetStatusDisplay

`func (o *MerchantOrderRefundResponse) GetStatusDisplay() string`

GetStatusDisplay returns the StatusDisplay field if non-nil, zero value otherwise.

### GetStatusDisplayOk

`func (o *MerchantOrderRefundResponse) GetStatusDisplayOk() (*string, bool)`

GetStatusDisplayOk returns a tuple with the StatusDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDisplay

`func (o *MerchantOrderRefundResponse) SetStatusDisplay(v string)`

SetStatusDisplay sets StatusDisplay field to given value.


### GetRefundReason

`func (o *MerchantOrderRefundResponse) GetRefundReason() string`

GetRefundReason returns the RefundReason field if non-nil, zero value otherwise.

### GetRefundReasonOk

`func (o *MerchantOrderRefundResponse) GetRefundReasonOk() (*string, bool)`

GetRefundReasonOk returns a tuple with the RefundReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundReason

`func (o *MerchantOrderRefundResponse) SetRefundReason(v string)`

SetRefundReason sets RefundReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


