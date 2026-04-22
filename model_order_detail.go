/*
BWDK API

<div dir=\"rtl\" style=\"text-align: right;\">  # مستندات فروشندگان در سرویس خرید با دیجی‌کالا  این پلتفرم برای فروشندگان (مرچنت‌ها) جهت یکپارچه‌سازی خدمات پرداخت و تجارت الکترونیکی با سیستم خرید با دیجی‌کالا. شامل مدیریت سفارشات، ارسال، و احراز هویت فروشندگان است.   <div dir=\"rtl\" style=\"text-align: right;\">  <!-- ## توضیح وضعیت‌های سفارش  ### ۱. INITIAL — ایجاد اولیه سفارش  **معنا:** سفارش توسط بک‌اند مرچنت ساخته شده ولی هنوز هیچ کاربری به آن اختصاص داده نشده است.  **چگونه اتفاق می‌افتد:** مرچنت با ارسال درخواست `POST /api/v1/orders/create` و ارائه اطلاعات کالاها، مبلغ و `callback_url`، یک سفارش جدید ایجاد می‌کند. BWDK یک `order_uuid` منحصربه‌فرد و لینک شروع سفارش (`order_start_url`) برمی‌گرداند.  **وابستگی‌ها:** نیازی به کاربر یا پرداخت ندارد. فقط اطلاعات کالا از سمت مرچنت کافی است.  ---  ### ۲. STARTED — آغاز جریان خرید  **معنا:** مشتری روی لینک شروع سفارش کلیک کرده و وارد محیط BWDK شده است، اما هنوز لاگین نکرده.  **چگونه اتفاق می‌افتد:** وقتی مشتری به `order_start_url` هدایت می‌شود، BWDK وضعیت سفارش را از `INITIAL` به `STARTED` تغییر می‌دهد. در این مرحله فرآیند احراز هویت (SSO) آغاز می‌شود.  **وابستگی‌ها:** مشتری باید به لینک شروع هدایت شده باشد.  ---  ### ۳. PENDING — انتظار برای تکمیل سفارش  **معنا:** مشتری با موفقیت وارد سیستم شده و سفارش به حساب او اختصاص یافته. مشتری در حال انتخاب آدرس، روش ارسال، بسته‌بندی یا تخفیف است.  **چگونه اتفاق می‌افتد:** پس از تکمیل ورود به سیستم (SSO)، BWDK سفارش را به کاربر وصل کرده و وضعیت را به `PENDING` تغییر می‌دهد.  **وابستگی‌ها:** ورود موفق کاربر به سیستم (SSO). در این مرحله مشتری می‌تواند آدرس، شیپینگ، پکینگ و تخفیف را انتخاب کند.  ---  ### ۴. WAITING_FOR_GATEWAY — انتظار برای پرداخت  **معنا:** مشتری اطلاعات سفارش را تأیید کرده و به درگاه پرداخت هدایت شده است.  **چگونه اتفاق می‌افتد:** مشتری دکمه «پرداخت» را می‌زند (`POST /api/v1/orders/submit`)، سیستم یک رکورد پرداخت ایجاد می‌کند و کاربر به درگاه Digipay هدایت می‌شود. وضعیت سفارش به `WAITING_FOR_GATEWAY` تغییر می‌کند.  **وابستگی‌ها:** انتخاب آدرس، روش ارسال و بسته‌بندی الزامی است. پرداخت باید ایجاد شده باشد.  ---  ### ۷. PAID_BY_USER — پرداخت موفق  **معنا:** تراکنش پرداخت با موفقیت انجام شده و وجه از حساب مشتری کسر شده است.  **چگونه اتفاق می‌افتد:** درگاه پرداخت نتیجه موفق را به BWDK اطلاع می‌دهد. سیستم پرداخت را تأیید و وضعیت سفارش را به `PAID_BY_USER` تغییر می‌دهد. در این لحظه مشتری به `callback_url` مرچنت هدایت می‌شود.  **وابستگی‌ها:** تأیید موفق تراکنش از سوی درگاه پرداخت (Digipay).  ---  ### ۹. VERIFIED_BY_MERCHANT — تأیید توسط مرچنت  **معنا:** مرچنت سفارش را بررسی کرده و موجودی کالا و صحت اطلاعات را تأیید نموده است. سفارش آماده ارسال است.  **چگونه اتفاق می‌افتد:** مرچنت با ارسال درخواست `POST /api/v1/orders/manager/{uuid}/verify` سفارش را تأیید می‌کند. این مرحله **اجباری** است و باید پس از پرداخت موفق انجام شود.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` باشد. مرچنت باید موجودی کالا را بررسی کند.  ---  ### ۲۰. SHIPPED — ارسال شد  **معنا:** سفارش از انبار خارج شده و در حال ارسال به مشتری است.  **چگونه اتفاق می‌افتد:** مرچنت پس از ارسال کالا، وضعیت سفارش را از طریق API به `SHIPPED` تغییر می‌دهد.  **وابستگی‌ها:** سفارش باید در وضعیت `VERIFIED_BY_MERCHANT` باشد.  ---  ### ۱۹. DELIVERED — تحویل داده شد  **معنا:** سفارش به دست مشتری رسیده و فرآیند خرید به پایان رسیده است.  **چگونه اتفاق می‌افتد:** مرچنت پس از تحویل موفق، وضعیت را به `DELIVERED` تغییر می‌دهد.  **وابستگی‌ها:** سفارش باید در وضعیت `SHIPPED` باشد.  ---  ### ۵. EXPIRED — منقضی شد  **معنا:** زمان رزرو سفارش به پایان رسیده و سفارش به صورت خودکار لغو شده است.  **چگونه اتفاق می‌افتد:** یک Task دوره‌ای به طور خودکار سفارش‌هایی که `reservation_expired_at` آن‌ها گذشته را پیدا کرده و وضعیتشان را به `EXPIRED` تغییر می‌دهد. این مکانیزم مانع بلوکه شدن موجودی کالا می‌شود.  **وابستگی‌ها:** سفارش باید در یکی از وضعیت‌های `INITIAL`، `STARTED`، `PENDING` یا `WAITING_FOR_GATEWAY` باشد و زمان رزرو آن گذشته باشد.  ---  ### ۱۸. EXPIRATION_TIME_EXCEEDED — زمان انقضا گذشت  **معنا:** در لحظه ثبت نهایی یا پرداخت، مشخص شد که زمان مجاز سفارش تمام شده است.  **چگونه اتفاق می‌افتد:** هنگام ارسال درخواست پرداخت (`submit_order`)، سیستم بررسی می‌کند که `expiration_time` سفارش هنوز معتبر است یا خیر. در صورت گذشتن زمان، وضعیت به `EXPIRATION_TIME_EXCEEDED` تغییر می‌کند.  **وابستگی‌ها:** سفارش در وضعیت `PENDING` یا `WAITING_FOR_GATEWAY` است و فیلد `expiration_time` سپری شده.  ---  ### ۶. CANCELLED — لغو توسط مشتری  **معنا:** مشتری در حین فرآیند خرید (قبل از پرداخت) سفارش را لغو کرده یا از صفحه خارج شده است.  **چگونه اتفاق می‌افتد:** مشتری در صفحه checkout دکمه «انصراف» را می‌زند یا پرداخت ناموفق بوده و سفارش به حالت لغو درمی‌آید.  **وابستگی‌ها:** سفارش باید در وضعیت `PENDING` یا `WAITING_FOR_GATEWAY` باشد. پرداختی انجام نشده است.  ---  ### ۸. FAILED_TO_PAY — پرداخت ناموفق  **معنا:** تراکنش پرداخت انجام نشد یا با خطا مواجه شد.  **چگونه اتفاق می‌افتد:** درگاه پرداخت نتیجه ناموفق برمی‌گرداند یا فرآیند بازگشت وجه در مرحله پرداخت با شکست مواجه می‌شود.  **وابستگی‌ها:** سفارش باید در وضعیت `WAITING_FOR_GATEWAY` بوده باشد.  ---  ### ۱۰. FAILED_TO_VERIFY_BY_MERCHANT — تأیید ناموفق توسط مرچنت  **معنا:** مرچنت سفارش را رد کرده است؛ معمولاً به دلیل ناموجود بودن کالا یا مغایرت اطلاعات.  **چگونه اتفاق می‌افتد:** مرچنت در پاسخ به درخواست verify، خطا برمی‌گرداند یا API آن وضعیت ناموفق تنظیم می‌کند. پس از این وضعیت، فرآیند استرداد وجه آغاز می‌شود.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` باشد.  ---  ### ۱۱. FAILED_BY_MERCHANT — خطا از سمت مرچنت  **معنا:** مرچنت پس از تأیید اولیه، اعلام می‌کند که قادر به انجام سفارش نیست (مثلاً به دلیل اتمام موجودی).  **چگونه اتفاق می‌افتد:** مرچنت وضعیت را به `FAILED_BY_MERCHANT` تغییر می‌دهد. وجه پرداختی مشتری مسترد خواهد شد.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` باشد.  ---  ### ۱۲. CANCELLED_BY_MERCHANT — لغو توسط مرچنت  **معنا:** مرچنت پس از پرداخت، سفارش را به هر دلیلی لغو کرده است.  **چگونه اتفاق می‌افتد:** مرچنت درخواست لغو سفارش را ارسال می‌کند. وجه پرداختی مشتری به او بازگردانده می‌شود.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` یا `VERIFIED_BY_MERCHANT` باشد.  ---  ### ۱۳. REQUEST_TO_REFUND — درخواست استرداد توسط مشتری  **معنا:** مشتری درخواست بازگشت وجه داده و سیستم در حال پردازش استرداد است.  **چگونه اتفاق می‌افتد:** مرچنت از طریق API درخواست استرداد را ثبت می‌کند (`POST /api/v1/orders/manager/{uuid}/refund`). سفارش وارد صف پردازش استرداد می‌شود.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` یا `VERIFIED_BY_MERCHANT` باشد.  ---  ### ۱۴، ۱۵، ۱۶. سایر وضعیت‌های درخواست استرداد  این وضعیت‌ها بر اساس دلیل استرداد از هم تفکیک می‌شوند:  - **۱۴ — REQUEST_TO_REFUND_TO_MERCHANT_AFTER_FAILED_TO_VERIFY:** استرداد پس از ناموفق بودن تأیید مرچنت؛ وجه به حساب مرچنت بازمی‌گردد. - **۱۵ — REQUEST_TO_REFUND_TO_CUSTOMER_AFTER_FAILED_BY_MERCHANT:** استرداد پس از خطای مرچنت؛ وجه به مشتری بازمی‌گردد. - **۱۶ — REQUEST_TO_REFUND_TO_MERCHANT_AFTER_CANCELLED_BY_MERCHANT:** استرداد پس از لغو توسط مرچنت؛ وجه به حساب مرچنت برمی‌گردد.  **چگونه اتفاق می‌افتد:** به صورت خودکار پس از رسیدن به وضعیت‌های ناموفق/لغو مربوطه توسط سیستم تنظیم می‌شود.  ---  ### ۱۷. REFUND_COMPLETED — استرداد تکمیل شد  **معنا:** وجه با موفقیت به صاحب آن (مشتری یا مرچنت بسته به نوع استرداد) بازگردانده شده است.  **چگونه اتفاق می‌افتد:** Task پردازش استرداد (`process_order_refund`) پس از تأیید موفق بازگشت وجه از سوی Digipay، وضعیت سفارش را به `REFUND_COMPLETED` تغییر می‌دهد.  **وابستگی‌ها:** یکی از وضعیت‌های درخواست استرداد (۱۳، ۱۴، ۱۵ یا ۱۶) باید فعال باشد و Digipay تراکنش استرداد را تأیید کرده باشد.  --> </div> 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bwdk_sdk

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the OrderDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetail{}

// OrderDetail struct for OrderDetail
type OrderDetail struct {
	Id int32 `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	OrderUuid string `json:"order_uuid"`
	// Unix timestamp تا زمانی که سفارش برای پرداخت رزرو شده است
	ReservationExpiredAt NullableInt32 `json:"reservation_expired_at"`
	// شناسه منحصر به فرد سفارش در سیستم فروشنده
	MerchantOrderId string `json:"merchant_order_id"`
	Status OrderStatusEnum `json:"status"`
	StatusDisplay string `json:"status_display"`
	// مجموع قیمت اولیه تمام کالاهای سفارش بدون تخفیف (به تومان)
	MainAmount int32 `json:"main_amount"`
	// قیمت نهایی قابل پرداخت توسط مشتری: مبلغ_اصلی - مبلغ_تخفیف + مبلغ_مالیات (به تومان)
	FinalAmount int32 `json:"final_amount"`
	// مبلغ کل پرداخت شده توسط کاربر: مبلغ_نهایی + هزینه_ارسال (به تومان)
	TotalPaidAmount int32 `json:"total_paid_amount"`
	// مبلغ کل تخفیف اعمال شده بر سفارش (به تومان)
	DiscountAmount int32 `json:"discount_amount"`
	// مبلغ کل مالیات برای سفارش (به تومان)
	TaxAmount int32 `json:"tax_amount"`
	// هزینه ارسال برای سفارش (به تومان)
	ShippingAmount int32 `json:"shipping_amount"`
	// مقدار تخفیف از برنامه باشگاه مشتریان/پاداش (به تومان)
	LoyaltyAmount int32 `json:"loyalty_amount"`
	// آدرسی برای دریافت اطلاع رسانی وضعیت پرداخت پس از تکمیل سفارش
	CallbackUrl string `json:"callback_url"`
	Merchant Merchant `json:"merchant"`
	Items []OrderItemCreate `json:"items"`
	SourceAddress interface{} `json:"source_address"`
	DestinationAddress interface{} `json:"destination_address"`
	SelectedShippingMethod ShippingMethod `json:"selected_shipping_method"`
	ShippingSelectedAt NullableTime `json:"shipping_selected_at"`
	AddressSelectedAt NullableTime `json:"address_selected_at"`
	// هزینه روش بسته‌بندی انتخاب‌شده (به تومان)
	PackingAmount int32 `json:"packing_amount"`
	PackingSelectedAt NullableTime `json:"packing_selected_at"`
	SelectedPacking Packing `json:"selected_packing"`
	CanSelectPacking bool `json:"can_select_packing"`
	CanSelectShipping bool `json:"can_select_shipping"`
	CanSelectAddress bool `json:"can_select_address"`
	CanProceedToPayment bool `json:"can_proceed_to_payment"`
	IsPaid bool `json:"is_paid"`
	User OrderUser `json:"user"`
	Payment PaymentOrder `json:"payment"`
	// Preparation time for the order (in days)
	PreparationTime int32 `json:"preparation_time"`
	// Total weight of the order (in grams)
	Weight float64 `json:"weight"`
	SelectedShippingData map[string]interface{} `json:"selected_shipping_data"`
	// کد مرجع یکتا برای پیگیری سفارش مشتری (قالب: BD-XXXXXXXX)
	ReferenceCode string `json:"reference_code"`
	PromotionDiscountAmount float64 `json:"promotion_discount_amount"`
	PromotionData map[string]interface{} `json:"promotion_data"`
	// Markup amount for the order (in Tomans)
	DigipayMarkupAmount int32 `json:"digipay_markup_amount"`
	// Markup commission percentage for the order (in percent)
	MarkupCommissionPercentage int32 `json:"markup_commission_percentage"`
	PreviousStatus NullableOrderStatusEnum `json:"previous_status"`
	PreviousStatusDisplay string `json:"previous_status_display"`
}

type _OrderDetail OrderDetail

// NewOrderDetail instantiates a new OrderDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetail(id int32, createdAt time.Time, orderUuid string, reservationExpiredAt NullableInt32, merchantOrderId string, status OrderStatusEnum, statusDisplay string, mainAmount int32, finalAmount int32, totalPaidAmount int32, discountAmount int32, taxAmount int32, shippingAmount int32, loyaltyAmount int32, callbackUrl string, merchant Merchant, items []OrderItemCreate, sourceAddress interface{}, destinationAddress interface{}, selectedShippingMethod ShippingMethod, shippingSelectedAt NullableTime, addressSelectedAt NullableTime, packingAmount int32, packingSelectedAt NullableTime, selectedPacking Packing, canSelectPacking bool, canSelectShipping bool, canSelectAddress bool, canProceedToPayment bool, isPaid bool, user OrderUser, payment PaymentOrder, preparationTime int32, weight float64, selectedShippingData map[string]interface{}, referenceCode string, promotionDiscountAmount float64, promotionData map[string]interface{}, digipayMarkupAmount int32, markupCommissionPercentage int32, previousStatus NullableOrderStatusEnum, previousStatusDisplay string) *OrderDetail {
	this := OrderDetail{}
	this.Id = id
	this.CreatedAt = createdAt
	this.OrderUuid = orderUuid
	this.ReservationExpiredAt = reservationExpiredAt
	this.MerchantOrderId = merchantOrderId
	this.Status = status
	this.StatusDisplay = statusDisplay
	this.MainAmount = mainAmount
	this.FinalAmount = finalAmount
	this.TotalPaidAmount = totalPaidAmount
	this.DiscountAmount = discountAmount
	this.TaxAmount = taxAmount
	this.ShippingAmount = shippingAmount
	this.LoyaltyAmount = loyaltyAmount
	this.CallbackUrl = callbackUrl
	this.Merchant = merchant
	this.Items = items
	this.SourceAddress = sourceAddress
	this.DestinationAddress = destinationAddress
	this.SelectedShippingMethod = selectedShippingMethod
	this.ShippingSelectedAt = shippingSelectedAt
	this.AddressSelectedAt = addressSelectedAt
	this.PackingAmount = packingAmount
	this.PackingSelectedAt = packingSelectedAt
	this.SelectedPacking = selectedPacking
	this.CanSelectPacking = canSelectPacking
	this.CanSelectShipping = canSelectShipping
	this.CanSelectAddress = canSelectAddress
	this.CanProceedToPayment = canProceedToPayment
	this.IsPaid = isPaid
	this.User = user
	this.Payment = payment
	this.PreparationTime = preparationTime
	this.Weight = weight
	this.SelectedShippingData = selectedShippingData
	this.ReferenceCode = referenceCode
	this.PromotionDiscountAmount = promotionDiscountAmount
	this.PromotionData = promotionData
	this.DigipayMarkupAmount = digipayMarkupAmount
	this.MarkupCommissionPercentage = markupCommissionPercentage
	this.PreviousStatus = previousStatus
	this.PreviousStatusDisplay = previousStatusDisplay
	return &this
}

// NewOrderDetailWithDefaults instantiates a new OrderDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailWithDefaults() *OrderDetail {
	this := OrderDetail{}
	return &this
}

// GetId returns the Id field value
func (o *OrderDetail) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrderDetail) SetId(v int32) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrderDetail) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrderDetail) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetOrderUuid returns the OrderUuid field value
func (o *OrderDetail) GetOrderUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderUuid
}

// GetOrderUuidOk returns a tuple with the OrderUuid field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetOrderUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderUuid, true
}

// SetOrderUuid sets field value
func (o *OrderDetail) SetOrderUuid(v string) {
	o.OrderUuid = v
}

// GetReservationExpiredAt returns the ReservationExpiredAt field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *OrderDetail) GetReservationExpiredAt() int32 {
	if o == nil || o.ReservationExpiredAt.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ReservationExpiredAt.Get()
}

// GetReservationExpiredAtOk returns a tuple with the ReservationExpiredAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderDetail) GetReservationExpiredAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReservationExpiredAt.Get(), o.ReservationExpiredAt.IsSet()
}

// SetReservationExpiredAt sets field value
func (o *OrderDetail) SetReservationExpiredAt(v int32) {
	o.ReservationExpiredAt.Set(&v)
}

// GetMerchantOrderId returns the MerchantOrderId field value
func (o *OrderDetail) GetMerchantOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantOrderId
}

// GetMerchantOrderIdOk returns a tuple with the MerchantOrderId field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetMerchantOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantOrderId, true
}

// SetMerchantOrderId sets field value
func (o *OrderDetail) SetMerchantOrderId(v string) {
	o.MerchantOrderId = v
}

// GetStatus returns the Status field value
func (o *OrderDetail) GetStatus() OrderStatusEnum {
	if o == nil {
		var ret OrderStatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetStatusOk() (*OrderStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OrderDetail) SetStatus(v OrderStatusEnum) {
	o.Status = v
}

// GetStatusDisplay returns the StatusDisplay field value
func (o *OrderDetail) GetStatusDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusDisplay
}

// GetStatusDisplayOk returns a tuple with the StatusDisplay field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetStatusDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusDisplay, true
}

// SetStatusDisplay sets field value
func (o *OrderDetail) SetStatusDisplay(v string) {
	o.StatusDisplay = v
}

// GetMainAmount returns the MainAmount field value
func (o *OrderDetail) GetMainAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MainAmount
}

// GetMainAmountOk returns a tuple with the MainAmount field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetMainAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MainAmount, true
}

// SetMainAmount sets field value
func (o *OrderDetail) SetMainAmount(v int32) {
	o.MainAmount = v
}

// GetFinalAmount returns the FinalAmount field value
func (o *OrderDetail) GetFinalAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FinalAmount
}

// GetFinalAmountOk returns a tuple with the FinalAmount field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetFinalAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinalAmount, true
}

// SetFinalAmount sets field value
func (o *OrderDetail) SetFinalAmount(v int32) {
	o.FinalAmount = v
}

// GetTotalPaidAmount returns the TotalPaidAmount field value
func (o *OrderDetail) GetTotalPaidAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalPaidAmount
}

// GetTotalPaidAmountOk returns a tuple with the TotalPaidAmount field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetTotalPaidAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPaidAmount, true
}

// SetTotalPaidAmount sets field value
func (o *OrderDetail) SetTotalPaidAmount(v int32) {
	o.TotalPaidAmount = v
}

// GetDiscountAmount returns the DiscountAmount field value
func (o *OrderDetail) GetDiscountAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetDiscountAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountAmount, true
}

// SetDiscountAmount sets field value
func (o *OrderDetail) SetDiscountAmount(v int32) {
	o.DiscountAmount = v
}

// GetTaxAmount returns the TaxAmount field value
func (o *OrderDetail) GetTaxAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetTaxAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxAmount, true
}

// SetTaxAmount sets field value
func (o *OrderDetail) SetTaxAmount(v int32) {
	o.TaxAmount = v
}

// GetShippingAmount returns the ShippingAmount field value
func (o *OrderDetail) GetShippingAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ShippingAmount
}

// GetShippingAmountOk returns a tuple with the ShippingAmount field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetShippingAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingAmount, true
}

// SetShippingAmount sets field value
func (o *OrderDetail) SetShippingAmount(v int32) {
	o.ShippingAmount = v
}

// GetLoyaltyAmount returns the LoyaltyAmount field value
func (o *OrderDetail) GetLoyaltyAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LoyaltyAmount
}

// GetLoyaltyAmountOk returns a tuple with the LoyaltyAmount field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetLoyaltyAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoyaltyAmount, true
}

// SetLoyaltyAmount sets field value
func (o *OrderDetail) SetLoyaltyAmount(v int32) {
	o.LoyaltyAmount = v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *OrderDetail) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetCallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *OrderDetail) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetMerchant returns the Merchant field value
func (o *OrderDetail) GetMerchant() Merchant {
	if o == nil {
		var ret Merchant
		return ret
	}

	return o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetMerchantOk() (*Merchant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Merchant, true
}

// SetMerchant sets field value
func (o *OrderDetail) SetMerchant(v Merchant) {
	o.Merchant = v
}

// GetItems returns the Items field value
func (o *OrderDetail) GetItems() []OrderItemCreate {
	if o == nil {
		var ret []OrderItemCreate
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetItemsOk() ([]OrderItemCreate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *OrderDetail) SetItems(v []OrderItemCreate) {
	o.Items = v
}

// GetSourceAddress returns the SourceAddress field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *OrderDetail) GetSourceAddress() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.SourceAddress
}

// GetSourceAddressOk returns a tuple with the SourceAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderDetail) GetSourceAddressOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SourceAddress) {
		return nil, false
	}
	return &o.SourceAddress, true
}

// SetSourceAddress sets field value
func (o *OrderDetail) SetSourceAddress(v interface{}) {
	o.SourceAddress = v
}

// GetDestinationAddress returns the DestinationAddress field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *OrderDetail) GetDestinationAddress() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.DestinationAddress
}

// GetDestinationAddressOk returns a tuple with the DestinationAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderDetail) GetDestinationAddressOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DestinationAddress) {
		return nil, false
	}
	return &o.DestinationAddress, true
}

// SetDestinationAddress sets field value
func (o *OrderDetail) SetDestinationAddress(v interface{}) {
	o.DestinationAddress = v
}

// GetSelectedShippingMethod returns the SelectedShippingMethod field value
func (o *OrderDetail) GetSelectedShippingMethod() ShippingMethod {
	if o == nil {
		var ret ShippingMethod
		return ret
	}

	return o.SelectedShippingMethod
}

// GetSelectedShippingMethodOk returns a tuple with the SelectedShippingMethod field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetSelectedShippingMethodOk() (*ShippingMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectedShippingMethod, true
}

// SetSelectedShippingMethod sets field value
func (o *OrderDetail) SetSelectedShippingMethod(v ShippingMethod) {
	o.SelectedShippingMethod = v
}

// GetShippingSelectedAt returns the ShippingSelectedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *OrderDetail) GetShippingSelectedAt() time.Time {
	if o == nil || o.ShippingSelectedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ShippingSelectedAt.Get()
}

// GetShippingSelectedAtOk returns a tuple with the ShippingSelectedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderDetail) GetShippingSelectedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShippingSelectedAt.Get(), o.ShippingSelectedAt.IsSet()
}

// SetShippingSelectedAt sets field value
func (o *OrderDetail) SetShippingSelectedAt(v time.Time) {
	o.ShippingSelectedAt.Set(&v)
}

// GetAddressSelectedAt returns the AddressSelectedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *OrderDetail) GetAddressSelectedAt() time.Time {
	if o == nil || o.AddressSelectedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.AddressSelectedAt.Get()
}

// GetAddressSelectedAtOk returns a tuple with the AddressSelectedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderDetail) GetAddressSelectedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressSelectedAt.Get(), o.AddressSelectedAt.IsSet()
}

// SetAddressSelectedAt sets field value
func (o *OrderDetail) SetAddressSelectedAt(v time.Time) {
	o.AddressSelectedAt.Set(&v)
}

// GetPackingAmount returns the PackingAmount field value
func (o *OrderDetail) GetPackingAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PackingAmount
}

// GetPackingAmountOk returns a tuple with the PackingAmount field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetPackingAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackingAmount, true
}

// SetPackingAmount sets field value
func (o *OrderDetail) SetPackingAmount(v int32) {
	o.PackingAmount = v
}

// GetPackingSelectedAt returns the PackingSelectedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *OrderDetail) GetPackingSelectedAt() time.Time {
	if o == nil || o.PackingSelectedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.PackingSelectedAt.Get()
}

// GetPackingSelectedAtOk returns a tuple with the PackingSelectedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderDetail) GetPackingSelectedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackingSelectedAt.Get(), o.PackingSelectedAt.IsSet()
}

// SetPackingSelectedAt sets field value
func (o *OrderDetail) SetPackingSelectedAt(v time.Time) {
	o.PackingSelectedAt.Set(&v)
}

// GetSelectedPacking returns the SelectedPacking field value
func (o *OrderDetail) GetSelectedPacking() Packing {
	if o == nil {
		var ret Packing
		return ret
	}

	return o.SelectedPacking
}

// GetSelectedPackingOk returns a tuple with the SelectedPacking field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetSelectedPackingOk() (*Packing, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectedPacking, true
}

// SetSelectedPacking sets field value
func (o *OrderDetail) SetSelectedPacking(v Packing) {
	o.SelectedPacking = v
}

// GetCanSelectPacking returns the CanSelectPacking field value
func (o *OrderDetail) GetCanSelectPacking() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanSelectPacking
}

// GetCanSelectPackingOk returns a tuple with the CanSelectPacking field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetCanSelectPackingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanSelectPacking, true
}

// SetCanSelectPacking sets field value
func (o *OrderDetail) SetCanSelectPacking(v bool) {
	o.CanSelectPacking = v
}

// GetCanSelectShipping returns the CanSelectShipping field value
func (o *OrderDetail) GetCanSelectShipping() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanSelectShipping
}

// GetCanSelectShippingOk returns a tuple with the CanSelectShipping field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetCanSelectShippingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanSelectShipping, true
}

// SetCanSelectShipping sets field value
func (o *OrderDetail) SetCanSelectShipping(v bool) {
	o.CanSelectShipping = v
}

// GetCanSelectAddress returns the CanSelectAddress field value
func (o *OrderDetail) GetCanSelectAddress() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanSelectAddress
}

// GetCanSelectAddressOk returns a tuple with the CanSelectAddress field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetCanSelectAddressOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanSelectAddress, true
}

// SetCanSelectAddress sets field value
func (o *OrderDetail) SetCanSelectAddress(v bool) {
	o.CanSelectAddress = v
}

// GetCanProceedToPayment returns the CanProceedToPayment field value
func (o *OrderDetail) GetCanProceedToPayment() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanProceedToPayment
}

// GetCanProceedToPaymentOk returns a tuple with the CanProceedToPayment field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetCanProceedToPaymentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanProceedToPayment, true
}

// SetCanProceedToPayment sets field value
func (o *OrderDetail) SetCanProceedToPayment(v bool) {
	o.CanProceedToPayment = v
}

// GetIsPaid returns the IsPaid field value
func (o *OrderDetail) GetIsPaid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPaid
}

// GetIsPaidOk returns a tuple with the IsPaid field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetIsPaidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPaid, true
}

// SetIsPaid sets field value
func (o *OrderDetail) SetIsPaid(v bool) {
	o.IsPaid = v
}

// GetUser returns the User field value
func (o *OrderDetail) GetUser() OrderUser {
	if o == nil {
		var ret OrderUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetUserOk() (*OrderUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *OrderDetail) SetUser(v OrderUser) {
	o.User = v
}

// GetPayment returns the Payment field value
func (o *OrderDetail) GetPayment() PaymentOrder {
	if o == nil {
		var ret PaymentOrder
		return ret
	}

	return o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetPaymentOk() (*PaymentOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payment, true
}

// SetPayment sets field value
func (o *OrderDetail) SetPayment(v PaymentOrder) {
	o.Payment = v
}

// GetPreparationTime returns the PreparationTime field value
func (o *OrderDetail) GetPreparationTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PreparationTime
}

// GetPreparationTimeOk returns a tuple with the PreparationTime field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetPreparationTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreparationTime, true
}

// SetPreparationTime sets field value
func (o *OrderDetail) SetPreparationTime(v int32) {
	o.PreparationTime = v
}

// GetWeight returns the Weight field value
func (o *OrderDetail) GetWeight() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *OrderDetail) SetWeight(v float64) {
	o.Weight = v
}

// GetSelectedShippingData returns the SelectedShippingData field value
func (o *OrderDetail) GetSelectedShippingData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.SelectedShippingData
}

// GetSelectedShippingDataOk returns a tuple with the SelectedShippingData field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetSelectedShippingDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.SelectedShippingData, true
}

// SetSelectedShippingData sets field value
func (o *OrderDetail) SetSelectedShippingData(v map[string]interface{}) {
	o.SelectedShippingData = v
}

// GetReferenceCode returns the ReferenceCode field value
func (o *OrderDetail) GetReferenceCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceCode
}

// GetReferenceCodeOk returns a tuple with the ReferenceCode field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetReferenceCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceCode, true
}

// SetReferenceCode sets field value
func (o *OrderDetail) SetReferenceCode(v string) {
	o.ReferenceCode = v
}

// GetPromotionDiscountAmount returns the PromotionDiscountAmount field value
func (o *OrderDetail) GetPromotionDiscountAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.PromotionDiscountAmount
}

// GetPromotionDiscountAmountOk returns a tuple with the PromotionDiscountAmount field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetPromotionDiscountAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromotionDiscountAmount, true
}

// SetPromotionDiscountAmount sets field value
func (o *OrderDetail) SetPromotionDiscountAmount(v float64) {
	o.PromotionDiscountAmount = v
}

// GetPromotionData returns the PromotionData field value
func (o *OrderDetail) GetPromotionData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.PromotionData
}

// GetPromotionDataOk returns a tuple with the PromotionData field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetPromotionDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.PromotionData, true
}

// SetPromotionData sets field value
func (o *OrderDetail) SetPromotionData(v map[string]interface{}) {
	o.PromotionData = v
}

// GetDigipayMarkupAmount returns the DigipayMarkupAmount field value
func (o *OrderDetail) GetDigipayMarkupAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DigipayMarkupAmount
}

// GetDigipayMarkupAmountOk returns a tuple with the DigipayMarkupAmount field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetDigipayMarkupAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DigipayMarkupAmount, true
}

// SetDigipayMarkupAmount sets field value
func (o *OrderDetail) SetDigipayMarkupAmount(v int32) {
	o.DigipayMarkupAmount = v
}

// GetMarkupCommissionPercentage returns the MarkupCommissionPercentage field value
func (o *OrderDetail) GetMarkupCommissionPercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MarkupCommissionPercentage
}

// GetMarkupCommissionPercentageOk returns a tuple with the MarkupCommissionPercentage field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetMarkupCommissionPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarkupCommissionPercentage, true
}

// SetMarkupCommissionPercentage sets field value
func (o *OrderDetail) SetMarkupCommissionPercentage(v int32) {
	o.MarkupCommissionPercentage = v
}

// GetPreviousStatus returns the PreviousStatus field value
// If the value is explicit nil, the zero value for OrderStatusEnum will be returned
func (o *OrderDetail) GetPreviousStatus() OrderStatusEnum {
	if o == nil || o.PreviousStatus.Get() == nil {
		var ret OrderStatusEnum
		return ret
	}

	return *o.PreviousStatus.Get()
}

// GetPreviousStatusOk returns a tuple with the PreviousStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderDetail) GetPreviousStatusOk() (*OrderStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousStatus.Get(), o.PreviousStatus.IsSet()
}

// SetPreviousStatus sets field value
func (o *OrderDetail) SetPreviousStatus(v OrderStatusEnum) {
	o.PreviousStatus.Set(&v)
}

// GetPreviousStatusDisplay returns the PreviousStatusDisplay field value
func (o *OrderDetail) GetPreviousStatusDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousStatusDisplay
}

// GetPreviousStatusDisplayOk returns a tuple with the PreviousStatusDisplay field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetPreviousStatusDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousStatusDisplay, true
}

// SetPreviousStatusDisplay sets field value
func (o *OrderDetail) SetPreviousStatusDisplay(v string) {
	o.PreviousStatusDisplay = v
}

func (o OrderDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["order_uuid"] = o.OrderUuid
	toSerialize["reservation_expired_at"] = o.ReservationExpiredAt.Get()
	toSerialize["merchant_order_id"] = o.MerchantOrderId
	toSerialize["status"] = o.Status
	toSerialize["status_display"] = o.StatusDisplay
	toSerialize["main_amount"] = o.MainAmount
	toSerialize["final_amount"] = o.FinalAmount
	toSerialize["total_paid_amount"] = o.TotalPaidAmount
	toSerialize["discount_amount"] = o.DiscountAmount
	toSerialize["tax_amount"] = o.TaxAmount
	toSerialize["shipping_amount"] = o.ShippingAmount
	toSerialize["loyalty_amount"] = o.LoyaltyAmount
	toSerialize["callback_url"] = o.CallbackUrl
	toSerialize["merchant"] = o.Merchant
	toSerialize["items"] = o.Items
	if o.SourceAddress != nil {
		toSerialize["source_address"] = o.SourceAddress
	}
	if o.DestinationAddress != nil {
		toSerialize["destination_address"] = o.DestinationAddress
	}
	toSerialize["selected_shipping_method"] = o.SelectedShippingMethod
	toSerialize["shipping_selected_at"] = o.ShippingSelectedAt.Get()
	toSerialize["address_selected_at"] = o.AddressSelectedAt.Get()
	toSerialize["packing_amount"] = o.PackingAmount
	toSerialize["packing_selected_at"] = o.PackingSelectedAt.Get()
	toSerialize["selected_packing"] = o.SelectedPacking
	toSerialize["can_select_packing"] = o.CanSelectPacking
	toSerialize["can_select_shipping"] = o.CanSelectShipping
	toSerialize["can_select_address"] = o.CanSelectAddress
	toSerialize["can_proceed_to_payment"] = o.CanProceedToPayment
	toSerialize["is_paid"] = o.IsPaid
	toSerialize["user"] = o.User
	toSerialize["payment"] = o.Payment
	toSerialize["preparation_time"] = o.PreparationTime
	toSerialize["weight"] = o.Weight
	toSerialize["selected_shipping_data"] = o.SelectedShippingData
	toSerialize["reference_code"] = o.ReferenceCode
	toSerialize["promotion_discount_amount"] = o.PromotionDiscountAmount
	toSerialize["promotion_data"] = o.PromotionData
	toSerialize["digipay_markup_amount"] = o.DigipayMarkupAmount
	toSerialize["markup_commission_percentage"] = o.MarkupCommissionPercentage
	toSerialize["previous_status"] = o.PreviousStatus.Get()
	toSerialize["previous_status_display"] = o.PreviousStatusDisplay
	return toSerialize, nil
}

func (o *OrderDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"order_uuid",
		"reservation_expired_at",
		"merchant_order_id",
		"status",
		"status_display",
		"main_amount",
		"final_amount",
		"total_paid_amount",
		"discount_amount",
		"tax_amount",
		"shipping_amount",
		"loyalty_amount",
		"callback_url",
		"merchant",
		"items",
		"source_address",
		"destination_address",
		"selected_shipping_method",
		"shipping_selected_at",
		"address_selected_at",
		"packing_amount",
		"packing_selected_at",
		"selected_packing",
		"can_select_packing",
		"can_select_shipping",
		"can_select_address",
		"can_proceed_to_payment",
		"is_paid",
		"user",
		"payment",
		"preparation_time",
		"weight",
		"selected_shipping_data",
		"reference_code",
		"promotion_discount_amount",
		"promotion_data",
		"digipay_markup_amount",
		"markup_commission_percentage",
		"previous_status",
		"previous_status_display",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrderDetail := _OrderDetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderDetail)

	if err != nil {
		return err
	}

	*o = OrderDetail(varOrderDetail)

	return err
}

type NullableOrderDetail struct {
	value *OrderDetail
	isSet bool
}

func (v NullableOrderDetail) Get() *OrderDetail {
	return v.value
}

func (v *NullableOrderDetail) Set(val *OrderDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetail(val *OrderDetail) *NullableOrderDetail {
	return &NullableOrderDetail{value: val, isSet: true}
}

func (v NullableOrderDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


