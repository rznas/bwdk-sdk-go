# OrderDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**OrderUuid** | **string** |  | [readonly] 
**ReservationExpiredAt** | **NullableInt32** | Unix timestamp تا زمانی که سفارش برای پرداخت رزرو شده است | [readonly] 
**MerchantOrderId** | **string** | شناسه منحصر به فرد سفارش در سیستم فروشنده | [readonly] 
**Status** | [**OrderStatusEnum**](OrderStatusEnum.md) |  | [readonly] 
**StatusDisplay** | **string** |  | [readonly] 
**MainAmount** | **int32** | مجموع قیمت اولیه تمام کالاهای سفارش بدون تخفیف (به تومان) | [readonly] 
**FinalAmount** | **int32** | قیمت نهایی قابل پرداخت توسط مشتری: مبلغ_اصلی - مبلغ_تخفیف + مبلغ_مالیات (به تومان) | [readonly] 
**TotalPaidAmount** | **int32** | مبلغ کل پرداخت شده توسط کاربر: مبلغ_نهایی + هزینه_ارسال (به تومان) | [readonly] 
**DiscountAmount** | **int32** | مبلغ کل تخفیف اعمال شده بر سفارش (به تومان) | [readonly] 
**TaxAmount** | **int32** | مبلغ کل مالیات برای سفارش (به تومان) | [readonly] 
**ShippingAmount** | **int32** | هزینه ارسال برای سفارش (به تومان) | [readonly] 
**LoyaltyAmount** | **int32** | مقدار تخفیف از برنامه باشگاه مشتریان/پاداش (به تومان) | [readonly] 
**CallbackUrl** | **string** | آدرسی برای دریافت اطلاع رسانی وضعیت پرداخت پس از تکمیل سفارش | [readonly] 
**Merchant** | [**Merchant**](Merchant.md) |  | 
**Items** | [**[]OrderItemCreate**](OrderItemCreate.md) |  | 
**SourceAddress** | **interface{}** |  | [readonly] 
**DestinationAddress** | **interface{}** |  | [readonly] 
**SelectedShippingMethod** | [**ShippingMethod**](ShippingMethod.md) |  | [readonly] 
**ShippingSelectedAt** | **NullableTime** |  | [readonly] 
**AddressSelectedAt** | **NullableTime** |  | [readonly] 
**PackingAmount** | **int32** | هزینه روش بسته‌بندی انتخاب‌شده (به تومان) | [readonly] 
**PackingSelectedAt** | **NullableTime** |  | [readonly] 
**SelectedPacking** | [**Packing**](Packing.md) |  | [readonly] 
**CanSelectPacking** | **bool** |  | [readonly] 
**CanSelectShipping** | **bool** |  | [readonly] 
**CanSelectAddress** | **bool** |  | [readonly] 
**CanProceedToPayment** | **bool** |  | [readonly] 
**IsPaid** | **bool** |  | [readonly] 
**User** | [**OrderUser**](OrderUser.md) |  | [readonly] 
**Payment** | [**PaymentOrder**](PaymentOrder.md) |  | [readonly] 
**PreparationTime** | **int32** | Preparation time for the order (in days) | [readonly] 
**Weight** | **float64** | Total weight of the order (in grams) | [readonly] 
**SelectedShippingData** | **map[string]interface{}** |  | [readonly] 
**ReferenceCode** | **string** | کد مرجع یکتا برای پیگیری سفارش مشتری (قالب: BD-XXXXXXXX) | [readonly] 
**PromotionDiscountAmount** | **float64** |  | [readonly] 
**PromotionData** | **map[string]interface{}** |  | [readonly] 
**DigipayMarkupAmount** | **int32** | Markup amount for the order (in Tomans) | [readonly] 
**MarkupCommissionPercentage** | **int32** | Markup commission percentage for the order (in percent) | [readonly] 
**PreviousStatus** | [**NullableOrderStatusEnum**](OrderStatusEnum.md) |  | [readonly] 
**PreviousStatusDisplay** | **string** |  | [readonly] 

## Methods

### NewOrderDetail

`func NewOrderDetail(id int32, createdAt time.Time, orderUuid string, reservationExpiredAt NullableInt32, merchantOrderId string, status OrderStatusEnum, statusDisplay string, mainAmount int32, finalAmount int32, totalPaidAmount int32, discountAmount int32, taxAmount int32, shippingAmount int32, loyaltyAmount int32, callbackUrl string, merchant Merchant, items []OrderItemCreate, sourceAddress interface{}, destinationAddress interface{}, selectedShippingMethod ShippingMethod, shippingSelectedAt NullableTime, addressSelectedAt NullableTime, packingAmount int32, packingSelectedAt NullableTime, selectedPacking Packing, canSelectPacking bool, canSelectShipping bool, canSelectAddress bool, canProceedToPayment bool, isPaid bool, user OrderUser, payment PaymentOrder, preparationTime int32, weight float64, selectedShippingData map[string]interface{}, referenceCode string, promotionDiscountAmount float64, promotionData map[string]interface{}, digipayMarkupAmount int32, markupCommissionPercentage int32, previousStatus NullableOrderStatusEnum, previousStatusDisplay string, ) *OrderDetail`

NewOrderDetail instantiates a new OrderDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDetailWithDefaults

`func NewOrderDetailWithDefaults() *OrderDetail`

NewOrderDetailWithDefaults instantiates a new OrderDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderDetail) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *OrderDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrderDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrderDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetOrderUuid

`func (o *OrderDetail) GetOrderUuid() string`

GetOrderUuid returns the OrderUuid field if non-nil, zero value otherwise.

### GetOrderUuidOk

`func (o *OrderDetail) GetOrderUuidOk() (*string, bool)`

GetOrderUuidOk returns a tuple with the OrderUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderUuid

`func (o *OrderDetail) SetOrderUuid(v string)`

SetOrderUuid sets OrderUuid field to given value.


### GetReservationExpiredAt

`func (o *OrderDetail) GetReservationExpiredAt() int32`

GetReservationExpiredAt returns the ReservationExpiredAt field if non-nil, zero value otherwise.

### GetReservationExpiredAtOk

`func (o *OrderDetail) GetReservationExpiredAtOk() (*int32, bool)`

GetReservationExpiredAtOk returns a tuple with the ReservationExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationExpiredAt

`func (o *OrderDetail) SetReservationExpiredAt(v int32)`

SetReservationExpiredAt sets ReservationExpiredAt field to given value.


### SetReservationExpiredAtNil

`func (o *OrderDetail) SetReservationExpiredAtNil(b bool)`

 SetReservationExpiredAtNil sets the value for ReservationExpiredAt to be an explicit nil

### UnsetReservationExpiredAt
`func (o *OrderDetail) UnsetReservationExpiredAt()`

UnsetReservationExpiredAt ensures that no value is present for ReservationExpiredAt, not even an explicit nil
### GetMerchantOrderId

`func (o *OrderDetail) GetMerchantOrderId() string`

GetMerchantOrderId returns the MerchantOrderId field if non-nil, zero value otherwise.

### GetMerchantOrderIdOk

`func (o *OrderDetail) GetMerchantOrderIdOk() (*string, bool)`

GetMerchantOrderIdOk returns a tuple with the MerchantOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderId

`func (o *OrderDetail) SetMerchantOrderId(v string)`

SetMerchantOrderId sets MerchantOrderId field to given value.


### GetStatus

`func (o *OrderDetail) GetStatus() OrderStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderDetail) GetStatusOk() (*OrderStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderDetail) SetStatus(v OrderStatusEnum)`

SetStatus sets Status field to given value.


### GetStatusDisplay

`func (o *OrderDetail) GetStatusDisplay() string`

GetStatusDisplay returns the StatusDisplay field if non-nil, zero value otherwise.

### GetStatusDisplayOk

`func (o *OrderDetail) GetStatusDisplayOk() (*string, bool)`

GetStatusDisplayOk returns a tuple with the StatusDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDisplay

`func (o *OrderDetail) SetStatusDisplay(v string)`

SetStatusDisplay sets StatusDisplay field to given value.


### GetMainAmount

`func (o *OrderDetail) GetMainAmount() int32`

GetMainAmount returns the MainAmount field if non-nil, zero value otherwise.

### GetMainAmountOk

`func (o *OrderDetail) GetMainAmountOk() (*int32, bool)`

GetMainAmountOk returns a tuple with the MainAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainAmount

`func (o *OrderDetail) SetMainAmount(v int32)`

SetMainAmount sets MainAmount field to given value.


### GetFinalAmount

`func (o *OrderDetail) GetFinalAmount() int32`

GetFinalAmount returns the FinalAmount field if non-nil, zero value otherwise.

### GetFinalAmountOk

`func (o *OrderDetail) GetFinalAmountOk() (*int32, bool)`

GetFinalAmountOk returns a tuple with the FinalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalAmount

`func (o *OrderDetail) SetFinalAmount(v int32)`

SetFinalAmount sets FinalAmount field to given value.


### GetTotalPaidAmount

`func (o *OrderDetail) GetTotalPaidAmount() int32`

GetTotalPaidAmount returns the TotalPaidAmount field if non-nil, zero value otherwise.

### GetTotalPaidAmountOk

`func (o *OrderDetail) GetTotalPaidAmountOk() (*int32, bool)`

GetTotalPaidAmountOk returns a tuple with the TotalPaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaidAmount

`func (o *OrderDetail) SetTotalPaidAmount(v int32)`

SetTotalPaidAmount sets TotalPaidAmount field to given value.


### GetDiscountAmount

`func (o *OrderDetail) GetDiscountAmount() int32`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *OrderDetail) GetDiscountAmountOk() (*int32, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *OrderDetail) SetDiscountAmount(v int32)`

SetDiscountAmount sets DiscountAmount field to given value.


### GetTaxAmount

`func (o *OrderDetail) GetTaxAmount() int32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *OrderDetail) GetTaxAmountOk() (*int32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *OrderDetail) SetTaxAmount(v int32)`

SetTaxAmount sets TaxAmount field to given value.


### GetShippingAmount

`func (o *OrderDetail) GetShippingAmount() int32`

GetShippingAmount returns the ShippingAmount field if non-nil, zero value otherwise.

### GetShippingAmountOk

`func (o *OrderDetail) GetShippingAmountOk() (*int32, bool)`

GetShippingAmountOk returns a tuple with the ShippingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAmount

`func (o *OrderDetail) SetShippingAmount(v int32)`

SetShippingAmount sets ShippingAmount field to given value.


### GetLoyaltyAmount

`func (o *OrderDetail) GetLoyaltyAmount() int32`

GetLoyaltyAmount returns the LoyaltyAmount field if non-nil, zero value otherwise.

### GetLoyaltyAmountOk

`func (o *OrderDetail) GetLoyaltyAmountOk() (*int32, bool)`

GetLoyaltyAmountOk returns a tuple with the LoyaltyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyAmount

`func (o *OrderDetail) SetLoyaltyAmount(v int32)`

SetLoyaltyAmount sets LoyaltyAmount field to given value.


### GetCallbackUrl

`func (o *OrderDetail) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *OrderDetail) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *OrderDetail) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetMerchant

`func (o *OrderDetail) GetMerchant() Merchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *OrderDetail) GetMerchantOk() (*Merchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *OrderDetail) SetMerchant(v Merchant)`

SetMerchant sets Merchant field to given value.


### GetItems

`func (o *OrderDetail) GetItems() []OrderItemCreate`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OrderDetail) GetItemsOk() (*[]OrderItemCreate, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OrderDetail) SetItems(v []OrderItemCreate)`

SetItems sets Items field to given value.


### GetSourceAddress

`func (o *OrderDetail) GetSourceAddress() interface{}`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *OrderDetail) GetSourceAddressOk() (*interface{}, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *OrderDetail) SetSourceAddress(v interface{})`

SetSourceAddress sets SourceAddress field to given value.


### SetSourceAddressNil

`func (o *OrderDetail) SetSourceAddressNil(b bool)`

 SetSourceAddressNil sets the value for SourceAddress to be an explicit nil

### UnsetSourceAddress
`func (o *OrderDetail) UnsetSourceAddress()`

UnsetSourceAddress ensures that no value is present for SourceAddress, not even an explicit nil
### GetDestinationAddress

`func (o *OrderDetail) GetDestinationAddress() interface{}`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *OrderDetail) GetDestinationAddressOk() (*interface{}, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *OrderDetail) SetDestinationAddress(v interface{})`

SetDestinationAddress sets DestinationAddress field to given value.


### SetDestinationAddressNil

`func (o *OrderDetail) SetDestinationAddressNil(b bool)`

 SetDestinationAddressNil sets the value for DestinationAddress to be an explicit nil

### UnsetDestinationAddress
`func (o *OrderDetail) UnsetDestinationAddress()`

UnsetDestinationAddress ensures that no value is present for DestinationAddress, not even an explicit nil
### GetSelectedShippingMethod

`func (o *OrderDetail) GetSelectedShippingMethod() ShippingMethod`

GetSelectedShippingMethod returns the SelectedShippingMethod field if non-nil, zero value otherwise.

### GetSelectedShippingMethodOk

`func (o *OrderDetail) GetSelectedShippingMethodOk() (*ShippingMethod, bool)`

GetSelectedShippingMethodOk returns a tuple with the SelectedShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedShippingMethod

`func (o *OrderDetail) SetSelectedShippingMethod(v ShippingMethod)`

SetSelectedShippingMethod sets SelectedShippingMethod field to given value.


### GetShippingSelectedAt

`func (o *OrderDetail) GetShippingSelectedAt() time.Time`

GetShippingSelectedAt returns the ShippingSelectedAt field if non-nil, zero value otherwise.

### GetShippingSelectedAtOk

`func (o *OrderDetail) GetShippingSelectedAtOk() (*time.Time, bool)`

GetShippingSelectedAtOk returns a tuple with the ShippingSelectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingSelectedAt

`func (o *OrderDetail) SetShippingSelectedAt(v time.Time)`

SetShippingSelectedAt sets ShippingSelectedAt field to given value.


### SetShippingSelectedAtNil

`func (o *OrderDetail) SetShippingSelectedAtNil(b bool)`

 SetShippingSelectedAtNil sets the value for ShippingSelectedAt to be an explicit nil

### UnsetShippingSelectedAt
`func (o *OrderDetail) UnsetShippingSelectedAt()`

UnsetShippingSelectedAt ensures that no value is present for ShippingSelectedAt, not even an explicit nil
### GetAddressSelectedAt

`func (o *OrderDetail) GetAddressSelectedAt() time.Time`

GetAddressSelectedAt returns the AddressSelectedAt field if non-nil, zero value otherwise.

### GetAddressSelectedAtOk

`func (o *OrderDetail) GetAddressSelectedAtOk() (*time.Time, bool)`

GetAddressSelectedAtOk returns a tuple with the AddressSelectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressSelectedAt

`func (o *OrderDetail) SetAddressSelectedAt(v time.Time)`

SetAddressSelectedAt sets AddressSelectedAt field to given value.


### SetAddressSelectedAtNil

`func (o *OrderDetail) SetAddressSelectedAtNil(b bool)`

 SetAddressSelectedAtNil sets the value for AddressSelectedAt to be an explicit nil

### UnsetAddressSelectedAt
`func (o *OrderDetail) UnsetAddressSelectedAt()`

UnsetAddressSelectedAt ensures that no value is present for AddressSelectedAt, not even an explicit nil
### GetPackingAmount

`func (o *OrderDetail) GetPackingAmount() int32`

GetPackingAmount returns the PackingAmount field if non-nil, zero value otherwise.

### GetPackingAmountOk

`func (o *OrderDetail) GetPackingAmountOk() (*int32, bool)`

GetPackingAmountOk returns a tuple with the PackingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackingAmount

`func (o *OrderDetail) SetPackingAmount(v int32)`

SetPackingAmount sets PackingAmount field to given value.


### GetPackingSelectedAt

`func (o *OrderDetail) GetPackingSelectedAt() time.Time`

GetPackingSelectedAt returns the PackingSelectedAt field if non-nil, zero value otherwise.

### GetPackingSelectedAtOk

`func (o *OrderDetail) GetPackingSelectedAtOk() (*time.Time, bool)`

GetPackingSelectedAtOk returns a tuple with the PackingSelectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackingSelectedAt

`func (o *OrderDetail) SetPackingSelectedAt(v time.Time)`

SetPackingSelectedAt sets PackingSelectedAt field to given value.


### SetPackingSelectedAtNil

`func (o *OrderDetail) SetPackingSelectedAtNil(b bool)`

 SetPackingSelectedAtNil sets the value for PackingSelectedAt to be an explicit nil

### UnsetPackingSelectedAt
`func (o *OrderDetail) UnsetPackingSelectedAt()`

UnsetPackingSelectedAt ensures that no value is present for PackingSelectedAt, not even an explicit nil
### GetSelectedPacking

`func (o *OrderDetail) GetSelectedPacking() Packing`

GetSelectedPacking returns the SelectedPacking field if non-nil, zero value otherwise.

### GetSelectedPackingOk

`func (o *OrderDetail) GetSelectedPackingOk() (*Packing, bool)`

GetSelectedPackingOk returns a tuple with the SelectedPacking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPacking

`func (o *OrderDetail) SetSelectedPacking(v Packing)`

SetSelectedPacking sets SelectedPacking field to given value.


### GetCanSelectPacking

`func (o *OrderDetail) GetCanSelectPacking() bool`

GetCanSelectPacking returns the CanSelectPacking field if non-nil, zero value otherwise.

### GetCanSelectPackingOk

`func (o *OrderDetail) GetCanSelectPackingOk() (*bool, bool)`

GetCanSelectPackingOk returns a tuple with the CanSelectPacking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSelectPacking

`func (o *OrderDetail) SetCanSelectPacking(v bool)`

SetCanSelectPacking sets CanSelectPacking field to given value.


### GetCanSelectShipping

`func (o *OrderDetail) GetCanSelectShipping() bool`

GetCanSelectShipping returns the CanSelectShipping field if non-nil, zero value otherwise.

### GetCanSelectShippingOk

`func (o *OrderDetail) GetCanSelectShippingOk() (*bool, bool)`

GetCanSelectShippingOk returns a tuple with the CanSelectShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSelectShipping

`func (o *OrderDetail) SetCanSelectShipping(v bool)`

SetCanSelectShipping sets CanSelectShipping field to given value.


### GetCanSelectAddress

`func (o *OrderDetail) GetCanSelectAddress() bool`

GetCanSelectAddress returns the CanSelectAddress field if non-nil, zero value otherwise.

### GetCanSelectAddressOk

`func (o *OrderDetail) GetCanSelectAddressOk() (*bool, bool)`

GetCanSelectAddressOk returns a tuple with the CanSelectAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSelectAddress

`func (o *OrderDetail) SetCanSelectAddress(v bool)`

SetCanSelectAddress sets CanSelectAddress field to given value.


### GetCanProceedToPayment

`func (o *OrderDetail) GetCanProceedToPayment() bool`

GetCanProceedToPayment returns the CanProceedToPayment field if non-nil, zero value otherwise.

### GetCanProceedToPaymentOk

`func (o *OrderDetail) GetCanProceedToPaymentOk() (*bool, bool)`

GetCanProceedToPaymentOk returns a tuple with the CanProceedToPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanProceedToPayment

`func (o *OrderDetail) SetCanProceedToPayment(v bool)`

SetCanProceedToPayment sets CanProceedToPayment field to given value.


### GetIsPaid

`func (o *OrderDetail) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *OrderDetail) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *OrderDetail) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.


### GetUser

`func (o *OrderDetail) GetUser() OrderUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrderDetail) GetUserOk() (*OrderUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrderDetail) SetUser(v OrderUser)`

SetUser sets User field to given value.


### GetPayment

`func (o *OrderDetail) GetPayment() PaymentOrder`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *OrderDetail) GetPaymentOk() (*PaymentOrder, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *OrderDetail) SetPayment(v PaymentOrder)`

SetPayment sets Payment field to given value.


### GetPreparationTime

`func (o *OrderDetail) GetPreparationTime() int32`

GetPreparationTime returns the PreparationTime field if non-nil, zero value otherwise.

### GetPreparationTimeOk

`func (o *OrderDetail) GetPreparationTimeOk() (*int32, bool)`

GetPreparationTimeOk returns a tuple with the PreparationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparationTime

`func (o *OrderDetail) SetPreparationTime(v int32)`

SetPreparationTime sets PreparationTime field to given value.


### GetWeight

`func (o *OrderDetail) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *OrderDetail) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *OrderDetail) SetWeight(v float64)`

SetWeight sets Weight field to given value.


### GetSelectedShippingData

`func (o *OrderDetail) GetSelectedShippingData() map[string]interface{}`

GetSelectedShippingData returns the SelectedShippingData field if non-nil, zero value otherwise.

### GetSelectedShippingDataOk

`func (o *OrderDetail) GetSelectedShippingDataOk() (*map[string]interface{}, bool)`

GetSelectedShippingDataOk returns a tuple with the SelectedShippingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedShippingData

`func (o *OrderDetail) SetSelectedShippingData(v map[string]interface{})`

SetSelectedShippingData sets SelectedShippingData field to given value.


### GetReferenceCode

`func (o *OrderDetail) GetReferenceCode() string`

GetReferenceCode returns the ReferenceCode field if non-nil, zero value otherwise.

### GetReferenceCodeOk

`func (o *OrderDetail) GetReferenceCodeOk() (*string, bool)`

GetReferenceCodeOk returns a tuple with the ReferenceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCode

`func (o *OrderDetail) SetReferenceCode(v string)`

SetReferenceCode sets ReferenceCode field to given value.


### GetPromotionDiscountAmount

`func (o *OrderDetail) GetPromotionDiscountAmount() float64`

GetPromotionDiscountAmount returns the PromotionDiscountAmount field if non-nil, zero value otherwise.

### GetPromotionDiscountAmountOk

`func (o *OrderDetail) GetPromotionDiscountAmountOk() (*float64, bool)`

GetPromotionDiscountAmountOk returns a tuple with the PromotionDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionDiscountAmount

`func (o *OrderDetail) SetPromotionDiscountAmount(v float64)`

SetPromotionDiscountAmount sets PromotionDiscountAmount field to given value.


### GetPromotionData

`func (o *OrderDetail) GetPromotionData() map[string]interface{}`

GetPromotionData returns the PromotionData field if non-nil, zero value otherwise.

### GetPromotionDataOk

`func (o *OrderDetail) GetPromotionDataOk() (*map[string]interface{}, bool)`

GetPromotionDataOk returns a tuple with the PromotionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionData

`func (o *OrderDetail) SetPromotionData(v map[string]interface{})`

SetPromotionData sets PromotionData field to given value.


### GetDigipayMarkupAmount

`func (o *OrderDetail) GetDigipayMarkupAmount() int32`

GetDigipayMarkupAmount returns the DigipayMarkupAmount field if non-nil, zero value otherwise.

### GetDigipayMarkupAmountOk

`func (o *OrderDetail) GetDigipayMarkupAmountOk() (*int32, bool)`

GetDigipayMarkupAmountOk returns a tuple with the DigipayMarkupAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigipayMarkupAmount

`func (o *OrderDetail) SetDigipayMarkupAmount(v int32)`

SetDigipayMarkupAmount sets DigipayMarkupAmount field to given value.


### GetMarkupCommissionPercentage

`func (o *OrderDetail) GetMarkupCommissionPercentage() int32`

GetMarkupCommissionPercentage returns the MarkupCommissionPercentage field if non-nil, zero value otherwise.

### GetMarkupCommissionPercentageOk

`func (o *OrderDetail) GetMarkupCommissionPercentageOk() (*int32, bool)`

GetMarkupCommissionPercentageOk returns a tuple with the MarkupCommissionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupCommissionPercentage

`func (o *OrderDetail) SetMarkupCommissionPercentage(v int32)`

SetMarkupCommissionPercentage sets MarkupCommissionPercentage field to given value.


### GetPreviousStatus

`func (o *OrderDetail) GetPreviousStatus() OrderStatusEnum`

GetPreviousStatus returns the PreviousStatus field if non-nil, zero value otherwise.

### GetPreviousStatusOk

`func (o *OrderDetail) GetPreviousStatusOk() (*OrderStatusEnum, bool)`

GetPreviousStatusOk returns a tuple with the PreviousStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousStatus

`func (o *OrderDetail) SetPreviousStatus(v OrderStatusEnum)`

SetPreviousStatus sets PreviousStatus field to given value.


### SetPreviousStatusNil

`func (o *OrderDetail) SetPreviousStatusNil(b bool)`

 SetPreviousStatusNil sets the value for PreviousStatus to be an explicit nil

### UnsetPreviousStatus
`func (o *OrderDetail) UnsetPreviousStatus()`

UnsetPreviousStatus ensures that no value is present for PreviousStatus, not even an explicit nil
### GetPreviousStatusDisplay

`func (o *OrderDetail) GetPreviousStatusDisplay() string`

GetPreviousStatusDisplay returns the PreviousStatusDisplay field if non-nil, zero value otherwise.

### GetPreviousStatusDisplayOk

`func (o *OrderDetail) GetPreviousStatusDisplayOk() (*string, bool)`

GetPreviousStatusDisplayOk returns a tuple with the PreviousStatusDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousStatusDisplay

`func (o *OrderDetail) SetPreviousStatusDisplay(v string)`

SetPreviousStatusDisplay sets PreviousStatusDisplay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


