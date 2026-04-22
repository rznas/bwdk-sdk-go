# VerifyOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantUniqueId** | **string** | شناسه منحصر به فرد ارسال شده هنگام ساخت سفارش برای تأیید اصالت سفارش | 

## Methods

### NewVerifyOrder

`func NewVerifyOrder(merchantUniqueId string, ) *VerifyOrder`

NewVerifyOrder instantiates a new VerifyOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyOrderWithDefaults

`func NewVerifyOrderWithDefaults() *VerifyOrder`

NewVerifyOrderWithDefaults instantiates a new VerifyOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantUniqueId

`func (o *VerifyOrder) GetMerchantUniqueId() string`

GetMerchantUniqueId returns the MerchantUniqueId field if non-nil, zero value otherwise.

### GetMerchantUniqueIdOk

`func (o *VerifyOrder) GetMerchantUniqueIdOk() (*string, bool)`

GetMerchantUniqueIdOk returns a tuple with the MerchantUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantUniqueId

`func (o *VerifyOrder) SetMerchantUniqueId(v string)`

SetMerchantUniqueId sets MerchantUniqueId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


