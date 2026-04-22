# Packing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** | نام روش/گزینه بسته‌بندی | 
**Price** | Pointer to **int32** | هزینه بسته بندی به تومان | [optional] 

## Methods

### NewPacking

`func NewPacking(id int32, name string, ) *Packing`

NewPacking instantiates a new Packing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackingWithDefaults

`func NewPackingWithDefaults() *Packing`

NewPackingWithDefaults instantiates a new Packing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Packing) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Packing) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Packing) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Packing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Packing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Packing) SetName(v string)`

SetName sets Name field to given value.


### GetPrice

`func (o *Packing) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Packing) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Packing) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Packing) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


