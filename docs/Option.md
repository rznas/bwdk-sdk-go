# Option

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | [**TypeNameEnum**](TypeNameEnum.md) | نوع ویژگی محصول: رنگ، سایز، گارانتی، وزن یا سایر  * &#x60;color&#x60; - رنگ * &#x60;size&#x60; - اندازه * &#x60;warranty&#x60; - گارانتی * &#x60;weight&#x60; - وزن * &#x60;other&#x60; - سایر | 
**Name** | **string** | نام ویژگی (مثلاً &#39;قرمز&#39; برای رنگ قرمز، &#39;XL&#39; برای سایز) | 
**Value** | **string** | مقدار ویژگی (مثلاً &#39;#FF0000&#39; برای کد رنگ یا مقدار دیگر) | 
**IsColor** | Pointer to **bool** | اگر نوع ویژگی رنگ است درست است، وگرنه غلط | [optional] [default to false]

## Methods

### NewOption

`func NewOption(typeName TypeNameEnum, name string, value string, ) *Option`

NewOption instantiates a new Option object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionWithDefaults

`func NewOptionWithDefaults() *Option`

NewOptionWithDefaults instantiates a new Option object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *Option) GetTypeName() TypeNameEnum`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *Option) GetTypeNameOk() (*TypeNameEnum, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *Option) SetTypeName(v TypeNameEnum)`

SetTypeName sets TypeName field to given value.


### GetName

`func (o *Option) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Option) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Option) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *Option) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Option) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Option) SetValue(v string)`

SetValue sets Value field to given value.


### GetIsColor

`func (o *Option) GetIsColor() bool`

GetIsColor returns the IsColor field if non-nil, zero value otherwise.

### GetIsColorOk

`func (o *Option) GetIsColorOk() (*bool, bool)`

GetIsColorOk returns a tuple with the IsColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsColor

`func (o *Option) SetIsColor(v bool)`

SetIsColor sets IsColor field to given value.

### HasIsColor

`func (o *Option) HasIsColor() bool`

HasIsColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


