# OrderCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantOrderId** | **string** | شناسه منحصر به فرد این سفارش در سیستم فروشنده | 
**MerchantUniqueId** | **string** | شناسه منحصر به فرد برای تأیید اصالت سفارش | 
**MainAmount** | Pointer to **int32** | مجموع قیمت‌های اولیه تمام کالاها بدون تخفیف (به تومان) | [optional] 
**FinalAmount** | Pointer to **int32** | مبلغ نهایی: مبلغ_اصلی - مبلغ_تخفیف + مبلغ_مالیات (به تومان) | [optional] 
**TotalPaidAmount** | Pointer to **int32** | مبلغ کل پرداخت شده توسط کاربر: مبلغ_نهایی + هزینه_ارسال (به تومان) | [optional] 
**DiscountAmount** | Pointer to **int32** | مبلغ کل تخفیف برای تمام سفارش (به تومان) | [optional] 
**TaxAmount** | Pointer to **int32** | مبلغ کل مالیات برای تمام سفارش (به تومان) | [optional] 
**ShippingAmount** | Pointer to **int32** | هزینه ارسال برای سفارش (به تومان) | [optional] 
**LoyaltyAmount** | Pointer to **int32** | مبلغ تخفیف باشگاه مشتریان/پاداش (به تومان) | [optional] 
**CallbackUrl** | **string** | آدرس وب‌هوک برای دریافت اطلاع رسانی وضعیت پرداخت | 
**DestinationAddress** | **interface{}** |  | [readonly] 
**Items** | [**[]OrderItemCreate**](OrderItemCreate.md) |  | 
**Merchant** | Pointer to **int32** | مقدار توسط سیستم جایگذاری می شود | [optional] 
**SourceAddress** | Pointer to **interface{}** | مقدار توسط سیستم جایگذاری می شود | [optional] 
**User** | **NullableInt32** |  | [readonly] 
**ReservationExpiredAt** | Pointer to **NullableInt32** | مهلت پرداخت (به عنوان Unix timestamp) قبل از اتمام سفارش | [optional] 
**ReferenceCode** | **string** | کد مرجع منحصر به فرد برای پیگیری سفارش مشتری (فرمت: BD-XXXXXXXX) | [readonly] 
**PreparationTime** | Pointer to **int32** | زمان آمادهسازی سفارش (به روز) | [optional] [default to 2]
**Weight** | Pointer to **float64** | وزن کل سفارش (بر حسب گرم) | [optional] 

## Methods

### NewOrderCreate

`func NewOrderCreate(merchantOrderId string, merchantUniqueId string, callbackUrl string, destinationAddress interface{}, items []OrderItemCreate, user NullableInt32, referenceCode string, ) *OrderCreate`

NewOrderCreate instantiates a new OrderCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateWithDefaults

`func NewOrderCreateWithDefaults() *OrderCreate`

NewOrderCreateWithDefaults instantiates a new OrderCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantOrderId

`func (o *OrderCreate) GetMerchantOrderId() string`

GetMerchantOrderId returns the MerchantOrderId field if non-nil, zero value otherwise.

### GetMerchantOrderIdOk

`func (o *OrderCreate) GetMerchantOrderIdOk() (*string, bool)`

GetMerchantOrderIdOk returns a tuple with the MerchantOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderId

`func (o *OrderCreate) SetMerchantOrderId(v string)`

SetMerchantOrderId sets MerchantOrderId field to given value.


### GetMerchantUniqueId

`func (o *OrderCreate) GetMerchantUniqueId() string`

GetMerchantUniqueId returns the MerchantUniqueId field if non-nil, zero value otherwise.

### GetMerchantUniqueIdOk

`func (o *OrderCreate) GetMerchantUniqueIdOk() (*string, bool)`

GetMerchantUniqueIdOk returns a tuple with the MerchantUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantUniqueId

`func (o *OrderCreate) SetMerchantUniqueId(v string)`

SetMerchantUniqueId sets MerchantUniqueId field to given value.


### GetMainAmount

`func (o *OrderCreate) GetMainAmount() int32`

GetMainAmount returns the MainAmount field if non-nil, zero value otherwise.

### GetMainAmountOk

`func (o *OrderCreate) GetMainAmountOk() (*int32, bool)`

GetMainAmountOk returns a tuple with the MainAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainAmount

`func (o *OrderCreate) SetMainAmount(v int32)`

SetMainAmount sets MainAmount field to given value.

### HasMainAmount

`func (o *OrderCreate) HasMainAmount() bool`

HasMainAmount returns a boolean if a field has been set.

### GetFinalAmount

`func (o *OrderCreate) GetFinalAmount() int32`

GetFinalAmount returns the FinalAmount field if non-nil, zero value otherwise.

### GetFinalAmountOk

`func (o *OrderCreate) GetFinalAmountOk() (*int32, bool)`

GetFinalAmountOk returns a tuple with the FinalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalAmount

`func (o *OrderCreate) SetFinalAmount(v int32)`

SetFinalAmount sets FinalAmount field to given value.

### HasFinalAmount

`func (o *OrderCreate) HasFinalAmount() bool`

HasFinalAmount returns a boolean if a field has been set.

### GetTotalPaidAmount

`func (o *OrderCreate) GetTotalPaidAmount() int32`

GetTotalPaidAmount returns the TotalPaidAmount field if non-nil, zero value otherwise.

### GetTotalPaidAmountOk

`func (o *OrderCreate) GetTotalPaidAmountOk() (*int32, bool)`

GetTotalPaidAmountOk returns a tuple with the TotalPaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaidAmount

`func (o *OrderCreate) SetTotalPaidAmount(v int32)`

SetTotalPaidAmount sets TotalPaidAmount field to given value.

### HasTotalPaidAmount

`func (o *OrderCreate) HasTotalPaidAmount() bool`

HasTotalPaidAmount returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *OrderCreate) GetDiscountAmount() int32`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *OrderCreate) GetDiscountAmountOk() (*int32, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *OrderCreate) SetDiscountAmount(v int32)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *OrderCreate) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetTaxAmount

`func (o *OrderCreate) GetTaxAmount() int32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *OrderCreate) GetTaxAmountOk() (*int32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *OrderCreate) SetTaxAmount(v int32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *OrderCreate) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetShippingAmount

`func (o *OrderCreate) GetShippingAmount() int32`

GetShippingAmount returns the ShippingAmount field if non-nil, zero value otherwise.

### GetShippingAmountOk

`func (o *OrderCreate) GetShippingAmountOk() (*int32, bool)`

GetShippingAmountOk returns a tuple with the ShippingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAmount

`func (o *OrderCreate) SetShippingAmount(v int32)`

SetShippingAmount sets ShippingAmount field to given value.

### HasShippingAmount

`func (o *OrderCreate) HasShippingAmount() bool`

HasShippingAmount returns a boolean if a field has been set.

### GetLoyaltyAmount

`func (o *OrderCreate) GetLoyaltyAmount() int32`

GetLoyaltyAmount returns the LoyaltyAmount field if non-nil, zero value otherwise.

### GetLoyaltyAmountOk

`func (o *OrderCreate) GetLoyaltyAmountOk() (*int32, bool)`

GetLoyaltyAmountOk returns a tuple with the LoyaltyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyAmount

`func (o *OrderCreate) SetLoyaltyAmount(v int32)`

SetLoyaltyAmount sets LoyaltyAmount field to given value.

### HasLoyaltyAmount

`func (o *OrderCreate) HasLoyaltyAmount() bool`

HasLoyaltyAmount returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *OrderCreate) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *OrderCreate) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *OrderCreate) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetDestinationAddress

`func (o *OrderCreate) GetDestinationAddress() interface{}`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *OrderCreate) GetDestinationAddressOk() (*interface{}, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *OrderCreate) SetDestinationAddress(v interface{})`

SetDestinationAddress sets DestinationAddress field to given value.


### SetDestinationAddressNil

`func (o *OrderCreate) SetDestinationAddressNil(b bool)`

 SetDestinationAddressNil sets the value for DestinationAddress to be an explicit nil

### UnsetDestinationAddress
`func (o *OrderCreate) UnsetDestinationAddress()`

UnsetDestinationAddress ensures that no value is present for DestinationAddress, not even an explicit nil
### GetItems

`func (o *OrderCreate) GetItems() []OrderItemCreate`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OrderCreate) GetItemsOk() (*[]OrderItemCreate, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OrderCreate) SetItems(v []OrderItemCreate)`

SetItems sets Items field to given value.


### GetMerchant

`func (o *OrderCreate) GetMerchant() int32`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *OrderCreate) GetMerchantOk() (*int32, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *OrderCreate) SetMerchant(v int32)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *OrderCreate) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetSourceAddress

`func (o *OrderCreate) GetSourceAddress() interface{}`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *OrderCreate) GetSourceAddressOk() (*interface{}, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *OrderCreate) SetSourceAddress(v interface{})`

SetSourceAddress sets SourceAddress field to given value.

### HasSourceAddress

`func (o *OrderCreate) HasSourceAddress() bool`

HasSourceAddress returns a boolean if a field has been set.

### SetSourceAddressNil

`func (o *OrderCreate) SetSourceAddressNil(b bool)`

 SetSourceAddressNil sets the value for SourceAddress to be an explicit nil

### UnsetSourceAddress
`func (o *OrderCreate) UnsetSourceAddress()`

UnsetSourceAddress ensures that no value is present for SourceAddress, not even an explicit nil
### GetUser

`func (o *OrderCreate) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrderCreate) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrderCreate) SetUser(v int32)`

SetUser sets User field to given value.


### SetUserNil

`func (o *OrderCreate) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *OrderCreate) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetReservationExpiredAt

`func (o *OrderCreate) GetReservationExpiredAt() int32`

GetReservationExpiredAt returns the ReservationExpiredAt field if non-nil, zero value otherwise.

### GetReservationExpiredAtOk

`func (o *OrderCreate) GetReservationExpiredAtOk() (*int32, bool)`

GetReservationExpiredAtOk returns a tuple with the ReservationExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationExpiredAt

`func (o *OrderCreate) SetReservationExpiredAt(v int32)`

SetReservationExpiredAt sets ReservationExpiredAt field to given value.

### HasReservationExpiredAt

`func (o *OrderCreate) HasReservationExpiredAt() bool`

HasReservationExpiredAt returns a boolean if a field has been set.

### SetReservationExpiredAtNil

`func (o *OrderCreate) SetReservationExpiredAtNil(b bool)`

 SetReservationExpiredAtNil sets the value for ReservationExpiredAt to be an explicit nil

### UnsetReservationExpiredAt
`func (o *OrderCreate) UnsetReservationExpiredAt()`

UnsetReservationExpiredAt ensures that no value is present for ReservationExpiredAt, not even an explicit nil
### GetReferenceCode

`func (o *OrderCreate) GetReferenceCode() string`

GetReferenceCode returns the ReferenceCode field if non-nil, zero value otherwise.

### GetReferenceCodeOk

`func (o *OrderCreate) GetReferenceCodeOk() (*string, bool)`

GetReferenceCodeOk returns a tuple with the ReferenceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCode

`func (o *OrderCreate) SetReferenceCode(v string)`

SetReferenceCode sets ReferenceCode field to given value.


### GetPreparationTime

`func (o *OrderCreate) GetPreparationTime() int32`

GetPreparationTime returns the PreparationTime field if non-nil, zero value otherwise.

### GetPreparationTimeOk

`func (o *OrderCreate) GetPreparationTimeOk() (*int32, bool)`

GetPreparationTimeOk returns a tuple with the PreparationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparationTime

`func (o *OrderCreate) SetPreparationTime(v int32)`

SetPreparationTime sets PreparationTime field to given value.

### HasPreparationTime

`func (o *OrderCreate) HasPreparationTime() bool`

HasPreparationTime returns a boolean if a field has been set.

### GetWeight

`func (o *OrderCreate) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *OrderCreate) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *OrderCreate) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *OrderCreate) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


