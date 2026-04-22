# MerchantPaidOrderList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderUuid** | **string** |  | [readonly] 
**MerchantOrderId** | **string** | شناسه منحصر به فرد سفارش در سیستم فروشنده | [readonly] 
**MerchantUniqueId** | **string** | شناسه منحصر به فرد برای پذیرنده برای تأیید سفارش | [readonly] 
**PaidAt** | **NullableTime** |  | [readonly] 
**RefundsAt** | **NullableTime** |  | [readonly] 

## Methods

### NewMerchantPaidOrderList

`func NewMerchantPaidOrderList(orderUuid string, merchantOrderId string, merchantUniqueId string, paidAt NullableTime, refundsAt NullableTime, ) *MerchantPaidOrderList`

NewMerchantPaidOrderList instantiates a new MerchantPaidOrderList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantPaidOrderListWithDefaults

`func NewMerchantPaidOrderListWithDefaults() *MerchantPaidOrderList`

NewMerchantPaidOrderListWithDefaults instantiates a new MerchantPaidOrderList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderUuid

`func (o *MerchantPaidOrderList) GetOrderUuid() string`

GetOrderUuid returns the OrderUuid field if non-nil, zero value otherwise.

### GetOrderUuidOk

`func (o *MerchantPaidOrderList) GetOrderUuidOk() (*string, bool)`

GetOrderUuidOk returns a tuple with the OrderUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderUuid

`func (o *MerchantPaidOrderList) SetOrderUuid(v string)`

SetOrderUuid sets OrderUuid field to given value.


### GetMerchantOrderId

`func (o *MerchantPaidOrderList) GetMerchantOrderId() string`

GetMerchantOrderId returns the MerchantOrderId field if non-nil, zero value otherwise.

### GetMerchantOrderIdOk

`func (o *MerchantPaidOrderList) GetMerchantOrderIdOk() (*string, bool)`

GetMerchantOrderIdOk returns a tuple with the MerchantOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderId

`func (o *MerchantPaidOrderList) SetMerchantOrderId(v string)`

SetMerchantOrderId sets MerchantOrderId field to given value.


### GetMerchantUniqueId

`func (o *MerchantPaidOrderList) GetMerchantUniqueId() string`

GetMerchantUniqueId returns the MerchantUniqueId field if non-nil, zero value otherwise.

### GetMerchantUniqueIdOk

`func (o *MerchantPaidOrderList) GetMerchantUniqueIdOk() (*string, bool)`

GetMerchantUniqueIdOk returns a tuple with the MerchantUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantUniqueId

`func (o *MerchantPaidOrderList) SetMerchantUniqueId(v string)`

SetMerchantUniqueId sets MerchantUniqueId field to given value.


### GetPaidAt

`func (o *MerchantPaidOrderList) GetPaidAt() time.Time`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *MerchantPaidOrderList) GetPaidAtOk() (*time.Time, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *MerchantPaidOrderList) SetPaidAt(v time.Time)`

SetPaidAt sets PaidAt field to given value.


### SetPaidAtNil

`func (o *MerchantPaidOrderList) SetPaidAtNil(b bool)`

 SetPaidAtNil sets the value for PaidAt to be an explicit nil

### UnsetPaidAt
`func (o *MerchantPaidOrderList) UnsetPaidAt()`

UnsetPaidAt ensures that no value is present for PaidAt, not even an explicit nil
### GetRefundsAt

`func (o *MerchantPaidOrderList) GetRefundsAt() time.Time`

GetRefundsAt returns the RefundsAt field if non-nil, zero value otherwise.

### GetRefundsAtOk

`func (o *MerchantPaidOrderList) GetRefundsAtOk() (*time.Time, bool)`

GetRefundsAtOk returns a tuple with the RefundsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundsAt

`func (o *MerchantPaidOrderList) SetRefundsAt(v time.Time)`

SetRefundsAt sets RefundsAt field to given value.


### SetRefundsAtNil

`func (o *MerchantPaidOrderList) SetRefundsAtNil(b bool)`

 SetRefundsAtNil sets the value for RefundsAt to be an explicit nil

### UnsetRefundsAt
`func (o *MerchantPaidOrderList) UnsetRefundsAt()`

UnsetRefundsAt ensures that no value is present for RefundsAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


