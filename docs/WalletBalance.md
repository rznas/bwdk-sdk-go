# WalletBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | Current wallet balance in Tomans | [optional] 
**NegativeSettlementDeadline** | **NullableString** | Deadline for settling negative balance | [readonly] 

## Methods

### NewWalletBalance

`func NewWalletBalance(negativeSettlementDeadline NullableString, ) *WalletBalance`

NewWalletBalance instantiates a new WalletBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletBalanceWithDefaults

`func NewWalletBalanceWithDefaults() *WalletBalance`

NewWalletBalanceWithDefaults instantiates a new WalletBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *WalletBalance) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WalletBalance) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WalletBalance) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *WalletBalance) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetNegativeSettlementDeadline

`func (o *WalletBalance) GetNegativeSettlementDeadline() string`

GetNegativeSettlementDeadline returns the NegativeSettlementDeadline field if non-nil, zero value otherwise.

### GetNegativeSettlementDeadlineOk

`func (o *WalletBalance) GetNegativeSettlementDeadlineOk() (*string, bool)`

GetNegativeSettlementDeadlineOk returns a tuple with the NegativeSettlementDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeSettlementDeadline

`func (o *WalletBalance) SetNegativeSettlementDeadline(v string)`

SetNegativeSettlementDeadline sets NegativeSettlementDeadline field to given value.


### SetNegativeSettlementDeadlineNil

`func (o *WalletBalance) SetNegativeSettlementDeadlineNil(b bool)`

 SetNegativeSettlementDeadlineNil sets the value for NegativeSettlementDeadline to be an explicit nil

### UnsetNegativeSettlementDeadline
`func (o *WalletBalance) UnsetNegativeSettlementDeadline()`

UnsetNegativeSettlementDeadline ensures that no value is present for NegativeSettlementDeadline, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


