# ShippingMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** | نام روش/گزینه بسته‌بندی | 
**Description** | Pointer to **string** | شناسه روش ارسال برای استفاده در سفارش | [optional] 
**ShippingType** | Pointer to [**ShippingTypeEnum**](ShippingTypeEnum.md) | شناسه وضعیت ارسال از دیجی اکسپرس  * &#x60;1&#x60; - سایر * &#x60;2&#x60; - دیجی اکسپرس | [optional] 
**GetShippingTypeDisplay** | **string** |  | [readonly] 
**ShippingTypeDisplay** | **string** |  | [readonly] 
**Cost** | Pointer to **int32** | هزینه ارسال برای منطقه اصلی (مثلاً تهران) به تومان | [optional] 
**SecondaryCost** | Pointer to **int32** | هزینه ارسال برای مناطق دیگر به تومان | [optional] 
**MinimumTimeSending** | Pointer to **int32** | حداقل تعداد روز از تاریخ سفارش تا تحویل | [optional] 
**MaximumTimeSending** | Pointer to **int32** | Maximum number of days from order date to delivery | [optional] 
**DeliveryTimeDisplay** | **string** |  | [readonly] 
**DeliveryTimeRangeDisplay** | [**DeliveryTimeRangeDisplay**](DeliveryTimeRangeDisplay.md) |  | [readonly] 
**InventoryAddress** | [**BusinessAddress**](BusinessAddress.md) |  | [readonly] 
**IsPayAtDestination** | Pointer to **bool** | آیا روش ارسال پرداخت در مقصد است | [optional] 

## Methods

### NewShippingMethod

`func NewShippingMethod(id int32, name string, getShippingTypeDisplay string, shippingTypeDisplay string, deliveryTimeDisplay string, deliveryTimeRangeDisplay DeliveryTimeRangeDisplay, inventoryAddress BusinessAddress, ) *ShippingMethod`

NewShippingMethod instantiates a new ShippingMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingMethodWithDefaults

`func NewShippingMethodWithDefaults() *ShippingMethod`

NewShippingMethodWithDefaults instantiates a new ShippingMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShippingMethod) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingMethod) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingMethod) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ShippingMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShippingMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShippingMethod) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ShippingMethod) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShippingMethod) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShippingMethod) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ShippingMethod) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShippingType

`func (o *ShippingMethod) GetShippingType() ShippingTypeEnum`

GetShippingType returns the ShippingType field if non-nil, zero value otherwise.

### GetShippingTypeOk

`func (o *ShippingMethod) GetShippingTypeOk() (*ShippingTypeEnum, bool)`

GetShippingTypeOk returns a tuple with the ShippingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingType

`func (o *ShippingMethod) SetShippingType(v ShippingTypeEnum)`

SetShippingType sets ShippingType field to given value.

### HasShippingType

`func (o *ShippingMethod) HasShippingType() bool`

HasShippingType returns a boolean if a field has been set.

### GetGetShippingTypeDisplay

`func (o *ShippingMethod) GetGetShippingTypeDisplay() string`

GetGetShippingTypeDisplay returns the GetShippingTypeDisplay field if non-nil, zero value otherwise.

### GetGetShippingTypeDisplayOk

`func (o *ShippingMethod) GetGetShippingTypeDisplayOk() (*string, bool)`

GetGetShippingTypeDisplayOk returns a tuple with the GetShippingTypeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetShippingTypeDisplay

`func (o *ShippingMethod) SetGetShippingTypeDisplay(v string)`

SetGetShippingTypeDisplay sets GetShippingTypeDisplay field to given value.


### GetShippingTypeDisplay

`func (o *ShippingMethod) GetShippingTypeDisplay() string`

GetShippingTypeDisplay returns the ShippingTypeDisplay field if non-nil, zero value otherwise.

### GetShippingTypeDisplayOk

`func (o *ShippingMethod) GetShippingTypeDisplayOk() (*string, bool)`

GetShippingTypeDisplayOk returns a tuple with the ShippingTypeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTypeDisplay

`func (o *ShippingMethod) SetShippingTypeDisplay(v string)`

SetShippingTypeDisplay sets ShippingTypeDisplay field to given value.


### GetCost

`func (o *ShippingMethod) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ShippingMethod) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ShippingMethod) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ShippingMethod) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetSecondaryCost

`func (o *ShippingMethod) GetSecondaryCost() int32`

GetSecondaryCost returns the SecondaryCost field if non-nil, zero value otherwise.

### GetSecondaryCostOk

`func (o *ShippingMethod) GetSecondaryCostOk() (*int32, bool)`

GetSecondaryCostOk returns a tuple with the SecondaryCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCost

`func (o *ShippingMethod) SetSecondaryCost(v int32)`

SetSecondaryCost sets SecondaryCost field to given value.

### HasSecondaryCost

`func (o *ShippingMethod) HasSecondaryCost() bool`

HasSecondaryCost returns a boolean if a field has been set.

### GetMinimumTimeSending

`func (o *ShippingMethod) GetMinimumTimeSending() int32`

GetMinimumTimeSending returns the MinimumTimeSending field if non-nil, zero value otherwise.

### GetMinimumTimeSendingOk

`func (o *ShippingMethod) GetMinimumTimeSendingOk() (*int32, bool)`

GetMinimumTimeSendingOk returns a tuple with the MinimumTimeSending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTimeSending

`func (o *ShippingMethod) SetMinimumTimeSending(v int32)`

SetMinimumTimeSending sets MinimumTimeSending field to given value.

### HasMinimumTimeSending

`func (o *ShippingMethod) HasMinimumTimeSending() bool`

HasMinimumTimeSending returns a boolean if a field has been set.

### GetMaximumTimeSending

`func (o *ShippingMethod) GetMaximumTimeSending() int32`

GetMaximumTimeSending returns the MaximumTimeSending field if non-nil, zero value otherwise.

### GetMaximumTimeSendingOk

`func (o *ShippingMethod) GetMaximumTimeSendingOk() (*int32, bool)`

GetMaximumTimeSendingOk returns a tuple with the MaximumTimeSending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumTimeSending

`func (o *ShippingMethod) SetMaximumTimeSending(v int32)`

SetMaximumTimeSending sets MaximumTimeSending field to given value.

### HasMaximumTimeSending

`func (o *ShippingMethod) HasMaximumTimeSending() bool`

HasMaximumTimeSending returns a boolean if a field has been set.

### GetDeliveryTimeDisplay

`func (o *ShippingMethod) GetDeliveryTimeDisplay() string`

GetDeliveryTimeDisplay returns the DeliveryTimeDisplay field if non-nil, zero value otherwise.

### GetDeliveryTimeDisplayOk

`func (o *ShippingMethod) GetDeliveryTimeDisplayOk() (*string, bool)`

GetDeliveryTimeDisplayOk returns a tuple with the DeliveryTimeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTimeDisplay

`func (o *ShippingMethod) SetDeliveryTimeDisplay(v string)`

SetDeliveryTimeDisplay sets DeliveryTimeDisplay field to given value.


### GetDeliveryTimeRangeDisplay

`func (o *ShippingMethod) GetDeliveryTimeRangeDisplay() DeliveryTimeRangeDisplay`

GetDeliveryTimeRangeDisplay returns the DeliveryTimeRangeDisplay field if non-nil, zero value otherwise.

### GetDeliveryTimeRangeDisplayOk

`func (o *ShippingMethod) GetDeliveryTimeRangeDisplayOk() (*DeliveryTimeRangeDisplay, bool)`

GetDeliveryTimeRangeDisplayOk returns a tuple with the DeliveryTimeRangeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTimeRangeDisplay

`func (o *ShippingMethod) SetDeliveryTimeRangeDisplay(v DeliveryTimeRangeDisplay)`

SetDeliveryTimeRangeDisplay sets DeliveryTimeRangeDisplay field to given value.


### GetInventoryAddress

`func (o *ShippingMethod) GetInventoryAddress() BusinessAddress`

GetInventoryAddress returns the InventoryAddress field if non-nil, zero value otherwise.

### GetInventoryAddressOk

`func (o *ShippingMethod) GetInventoryAddressOk() (*BusinessAddress, bool)`

GetInventoryAddressOk returns a tuple with the InventoryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryAddress

`func (o *ShippingMethod) SetInventoryAddress(v BusinessAddress)`

SetInventoryAddress sets InventoryAddress field to given value.


### GetIsPayAtDestination

`func (o *ShippingMethod) GetIsPayAtDestination() bool`

GetIsPayAtDestination returns the IsPayAtDestination field if non-nil, zero value otherwise.

### GetIsPayAtDestinationOk

`func (o *ShippingMethod) GetIsPayAtDestinationOk() (*bool, bool)`

GetIsPayAtDestinationOk returns a tuple with the IsPayAtDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPayAtDestination

`func (o *ShippingMethod) SetIsPayAtDestination(v bool)`

SetIsPayAtDestination sets IsPayAtDestination field to given value.

### HasIsPayAtDestination

`func (o *ShippingMethod) HasIsPayAtDestination() bool`

HasIsPayAtDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


