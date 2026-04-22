# PaymentOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinalAmount** | **int32** |  | [readonly] 
**GatewayType** | [**GatewayTypeEnum**](GatewayTypeEnum.md) |  | [readonly] 
**GatewayTypeDisplay** | **string** |  | [readonly] 
**PaidAt** | **string** |  | [readonly] 
**GatewayName** | **NullableString** |  | [readonly] 
**InvoiceId** | **NullableString** |  | [readonly] 
**SettlementDate** | **string** |  | [readonly] 
**SettlementStatus** | **int32** |  | [readonly] 
**SettlementStatusDisplay** | **string** |  | [readonly] 

## Methods

### NewPaymentOrder

`func NewPaymentOrder(finalAmount int32, gatewayType GatewayTypeEnum, gatewayTypeDisplay string, paidAt string, gatewayName NullableString, invoiceId NullableString, settlementDate string, settlementStatus int32, settlementStatusDisplay string, ) *PaymentOrder`

NewPaymentOrder instantiates a new PaymentOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentOrderWithDefaults

`func NewPaymentOrderWithDefaults() *PaymentOrder`

NewPaymentOrderWithDefaults instantiates a new PaymentOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinalAmount

`func (o *PaymentOrder) GetFinalAmount() int32`

GetFinalAmount returns the FinalAmount field if non-nil, zero value otherwise.

### GetFinalAmountOk

`func (o *PaymentOrder) GetFinalAmountOk() (*int32, bool)`

GetFinalAmountOk returns a tuple with the FinalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalAmount

`func (o *PaymentOrder) SetFinalAmount(v int32)`

SetFinalAmount sets FinalAmount field to given value.


### GetGatewayType

`func (o *PaymentOrder) GetGatewayType() GatewayTypeEnum`

GetGatewayType returns the GatewayType field if non-nil, zero value otherwise.

### GetGatewayTypeOk

`func (o *PaymentOrder) GetGatewayTypeOk() (*GatewayTypeEnum, bool)`

GetGatewayTypeOk returns a tuple with the GatewayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayType

`func (o *PaymentOrder) SetGatewayType(v GatewayTypeEnum)`

SetGatewayType sets GatewayType field to given value.


### GetGatewayTypeDisplay

`func (o *PaymentOrder) GetGatewayTypeDisplay() string`

GetGatewayTypeDisplay returns the GatewayTypeDisplay field if non-nil, zero value otherwise.

### GetGatewayTypeDisplayOk

`func (o *PaymentOrder) GetGatewayTypeDisplayOk() (*string, bool)`

GetGatewayTypeDisplayOk returns a tuple with the GatewayTypeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTypeDisplay

`func (o *PaymentOrder) SetGatewayTypeDisplay(v string)`

SetGatewayTypeDisplay sets GatewayTypeDisplay field to given value.


### GetPaidAt

`func (o *PaymentOrder) GetPaidAt() string`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *PaymentOrder) GetPaidAtOk() (*string, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *PaymentOrder) SetPaidAt(v string)`

SetPaidAt sets PaidAt field to given value.


### GetGatewayName

`func (o *PaymentOrder) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *PaymentOrder) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *PaymentOrder) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.


### SetGatewayNameNil

`func (o *PaymentOrder) SetGatewayNameNil(b bool)`

 SetGatewayNameNil sets the value for GatewayName to be an explicit nil

### UnsetGatewayName
`func (o *PaymentOrder) UnsetGatewayName()`

UnsetGatewayName ensures that no value is present for GatewayName, not even an explicit nil
### GetInvoiceId

`func (o *PaymentOrder) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *PaymentOrder) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *PaymentOrder) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### SetInvoiceIdNil

`func (o *PaymentOrder) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *PaymentOrder) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetSettlementDate

`func (o *PaymentOrder) GetSettlementDate() string`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *PaymentOrder) GetSettlementDateOk() (*string, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *PaymentOrder) SetSettlementDate(v string)`

SetSettlementDate sets SettlementDate field to given value.


### GetSettlementStatus

`func (o *PaymentOrder) GetSettlementStatus() int32`

GetSettlementStatus returns the SettlementStatus field if non-nil, zero value otherwise.

### GetSettlementStatusOk

`func (o *PaymentOrder) GetSettlementStatusOk() (*int32, bool)`

GetSettlementStatusOk returns a tuple with the SettlementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementStatus

`func (o *PaymentOrder) SetSettlementStatus(v int32)`

SetSettlementStatus sets SettlementStatus field to given value.


### GetSettlementStatusDisplay

`func (o *PaymentOrder) GetSettlementStatusDisplay() string`

GetSettlementStatusDisplay returns the SettlementStatusDisplay field if non-nil, zero value otherwise.

### GetSettlementStatusDisplayOk

`func (o *PaymentOrder) GetSettlementStatusDisplayOk() (*string, bool)`

GetSettlementStatusDisplayOk returns a tuple with the SettlementStatusDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementStatusDisplay

`func (o *PaymentOrder) SetSettlementStatusDisplay(v string)`

SetSettlementStatusDisplay sets SettlementStatusDisplay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


