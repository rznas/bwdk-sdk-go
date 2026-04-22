# PaginatedMerchantPaidOrderListList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]MerchantPaidOrderList**](MerchantPaidOrderList.md) |  | 

## Methods

### NewPaginatedMerchantPaidOrderListList

`func NewPaginatedMerchantPaidOrderListList(results []MerchantPaidOrderList, ) *PaginatedMerchantPaidOrderListList`

NewPaginatedMerchantPaidOrderListList instantiates a new PaginatedMerchantPaidOrderListList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedMerchantPaidOrderListListWithDefaults

`func NewPaginatedMerchantPaidOrderListListWithDefaults() *PaginatedMerchantPaidOrderListList`

NewPaginatedMerchantPaidOrderListListWithDefaults instantiates a new PaginatedMerchantPaidOrderListList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *PaginatedMerchantPaidOrderListList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedMerchantPaidOrderListList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedMerchantPaidOrderListList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedMerchantPaidOrderListList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedMerchantPaidOrderListList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedMerchantPaidOrderListList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedMerchantPaidOrderListList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedMerchantPaidOrderListList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedMerchantPaidOrderListList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedMerchantPaidOrderListList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedMerchantPaidOrderListList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedMerchantPaidOrderListList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedMerchantPaidOrderListList) GetResults() []MerchantPaidOrderList`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedMerchantPaidOrderListList) GetResultsOk() (*[]MerchantPaidOrderList, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedMerchantPaidOrderListList) SetResults(v []MerchantPaidOrderList)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


