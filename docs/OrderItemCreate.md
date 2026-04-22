# OrderItemCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | نام کامل محصول شامل تمام مشخصات | 
**PrimaryAmount** | Pointer to **int32** | قیمت اولیه برای هر واحد بدون تخفیف (به تومان) | [optional] 
**Amount** | Pointer to **int32** | قیمت نهایی برای تمام واحدها بعد از تخفیف (به تومان) | [optional] 
**Count** | **int32** | تعداد واحدهای این کالا در سفارش | 
**DiscountAmount** | Pointer to **int32** | مبلغ کل تخفیف برای این کالا (به تومان) | [optional] 
**TaxAmount** | Pointer to **int32** | مبلغ کل مالیات برای این کالا (به تومان) | [optional] 
**ImageLink** | Pointer to **NullableString** | آدرس تصویر محصول | [optional] 
**Options** | [**[]Option**](Option.md) |  | 
**PreparationTime** | Pointer to **int32** | Preparation time for the item (in days) | [optional] [default to 2]
**Weight** | Pointer to **float64** | Weight of the item (in grams) | [optional] 

## Methods

### NewOrderItemCreate

`func NewOrderItemCreate(name string, count int32, options []Option, ) *OrderItemCreate`

NewOrderItemCreate instantiates a new OrderItemCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderItemCreateWithDefaults

`func NewOrderItemCreateWithDefaults() *OrderItemCreate`

NewOrderItemCreateWithDefaults instantiates a new OrderItemCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrderItemCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderItemCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderItemCreate) SetName(v string)`

SetName sets Name field to given value.


### GetPrimaryAmount

`func (o *OrderItemCreate) GetPrimaryAmount() int32`

GetPrimaryAmount returns the PrimaryAmount field if non-nil, zero value otherwise.

### GetPrimaryAmountOk

`func (o *OrderItemCreate) GetPrimaryAmountOk() (*int32, bool)`

GetPrimaryAmountOk returns a tuple with the PrimaryAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAmount

`func (o *OrderItemCreate) SetPrimaryAmount(v int32)`

SetPrimaryAmount sets PrimaryAmount field to given value.

### HasPrimaryAmount

`func (o *OrderItemCreate) HasPrimaryAmount() bool`

HasPrimaryAmount returns a boolean if a field has been set.

### GetAmount

`func (o *OrderItemCreate) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrderItemCreate) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrderItemCreate) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *OrderItemCreate) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCount

`func (o *OrderItemCreate) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OrderItemCreate) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OrderItemCreate) SetCount(v int32)`

SetCount sets Count field to given value.


### GetDiscountAmount

`func (o *OrderItemCreate) GetDiscountAmount() int32`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *OrderItemCreate) GetDiscountAmountOk() (*int32, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *OrderItemCreate) SetDiscountAmount(v int32)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *OrderItemCreate) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetTaxAmount

`func (o *OrderItemCreate) GetTaxAmount() int32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *OrderItemCreate) GetTaxAmountOk() (*int32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *OrderItemCreate) SetTaxAmount(v int32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *OrderItemCreate) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetImageLink

`func (o *OrderItemCreate) GetImageLink() string`

GetImageLink returns the ImageLink field if non-nil, zero value otherwise.

### GetImageLinkOk

`func (o *OrderItemCreate) GetImageLinkOk() (*string, bool)`

GetImageLinkOk returns a tuple with the ImageLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageLink

`func (o *OrderItemCreate) SetImageLink(v string)`

SetImageLink sets ImageLink field to given value.

### HasImageLink

`func (o *OrderItemCreate) HasImageLink() bool`

HasImageLink returns a boolean if a field has been set.

### SetImageLinkNil

`func (o *OrderItemCreate) SetImageLinkNil(b bool)`

 SetImageLinkNil sets the value for ImageLink to be an explicit nil

### UnsetImageLink
`func (o *OrderItemCreate) UnsetImageLink()`

UnsetImageLink ensures that no value is present for ImageLink, not even an explicit nil
### GetOptions

`func (o *OrderItemCreate) GetOptions() []Option`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *OrderItemCreate) GetOptionsOk() (*[]Option, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *OrderItemCreate) SetOptions(v []Option)`

SetOptions sets Options field to given value.


### GetPreparationTime

`func (o *OrderItemCreate) GetPreparationTime() int32`

GetPreparationTime returns the PreparationTime field if non-nil, zero value otherwise.

### GetPreparationTimeOk

`func (o *OrderItemCreate) GetPreparationTimeOk() (*int32, bool)`

GetPreparationTimeOk returns a tuple with the PreparationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparationTime

`func (o *OrderItemCreate) SetPreparationTime(v int32)`

SetPreparationTime sets PreparationTime field to given value.

### HasPreparationTime

`func (o *OrderItemCreate) HasPreparationTime() bool`

HasPreparationTime returns a boolean if a field has been set.

### GetWeight

`func (o *OrderItemCreate) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *OrderItemCreate) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *OrderItemCreate) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *OrderItemCreate) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


