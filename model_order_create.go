/*
BWDK API

<div dir=\"rtl\" style=\"text-align: right;\">  # مستندات فروشندگان در سرویس خرید با دیجی‌کالا  این پلتفرم برای فروشندگان (مرچنت‌ها) جهت یکپارچه‌سازی خدمات پرداخت و تجارت الکترونیکی با سیستم خرید با دیجی‌کالا. شامل مدیریت سفارشات، ارسال، و احراز هویت فروشندگان است.   <div dir=\"rtl\" style=\"text-align: right;\">  <!-- ## توضیح وضعیت‌های سفارش  ### ۱. INITIAL — ایجاد اولیه سفارش  **معنا:** سفارش توسط بک‌اند مرچنت ساخته شده ولی هنوز هیچ کاربری به آن اختصاص داده نشده است.  **چگونه اتفاق می‌افتد:** مرچنت با ارسال درخواست `POST /api/v1/orders/create` و ارائه اطلاعات کالاها، مبلغ و `callback_url`، یک سفارش جدید ایجاد می‌کند. BWDK یک `order_uuid` منحصربه‌فرد و لینک شروع سفارش (`order_start_url`) برمی‌گرداند.  **وابستگی‌ها:** نیازی به کاربر یا پرداخت ندارد. فقط اطلاعات کالا از سمت مرچنت کافی است.  ---  ### ۲. STARTED — آغاز جریان خرید  **معنا:** مشتری روی لینک شروع سفارش کلیک کرده و وارد محیط BWDK شده است، اما هنوز لاگین نکرده.  **چگونه اتفاق می‌افتد:** وقتی مشتری به `order_start_url` هدایت می‌شود، BWDK وضعیت سفارش را از `INITIAL` به `STARTED` تغییر می‌دهد. در این مرحله فرآیند احراز هویت (SSO) آغاز می‌شود.  **وابستگی‌ها:** مشتری باید به لینک شروع هدایت شده باشد.  ---  ### ۳. PENDING — انتظار برای تکمیل سفارش  **معنا:** مشتری با موفقیت وارد سیستم شده و سفارش به حساب او اختصاص یافته. مشتری در حال انتخاب آدرس، روش ارسال، بسته‌بندی یا تخفیف است.  **چگونه اتفاق می‌افتد:** پس از تکمیل ورود به سیستم (SSO)، BWDK سفارش را به کاربر وصل کرده و وضعیت را به `PENDING` تغییر می‌دهد.  **وابستگی‌ها:** ورود موفق کاربر به سیستم (SSO). در این مرحله مشتری می‌تواند آدرس، شیپینگ، پکینگ و تخفیف را انتخاب کند.  ---  ### ۴. WAITING_FOR_GATEWAY — انتظار برای پرداخت  **معنا:** مشتری اطلاعات سفارش را تأیید کرده و به درگاه پرداخت هدایت شده است.  **چگونه اتفاق می‌افتد:** مشتری دکمه «پرداخت» را می‌زند (`POST /api/v1/orders/submit`)، سیستم یک رکورد پرداخت ایجاد می‌کند و کاربر به درگاه Digipay هدایت می‌شود. وضعیت سفارش به `WAITING_FOR_GATEWAY` تغییر می‌کند.  **وابستگی‌ها:** انتخاب آدرس، روش ارسال و بسته‌بندی الزامی است. پرداخت باید ایجاد شده باشد.  ---  ### ۷. PAID_BY_USER — پرداخت موفق  **معنا:** تراکنش پرداخت با موفقیت انجام شده و وجه از حساب مشتری کسر شده است.  **چگونه اتفاق می‌افتد:** درگاه پرداخت نتیجه موفق را به BWDK اطلاع می‌دهد. سیستم پرداخت را تأیید و وضعیت سفارش را به `PAID_BY_USER` تغییر می‌دهد. در این لحظه مشتری به `callback_url` مرچنت هدایت می‌شود.  **وابستگی‌ها:** تأیید موفق تراکنش از سوی درگاه پرداخت (Digipay).  ---  ### ۹. VERIFIED_BY_MERCHANT — تأیید توسط مرچنت  **معنا:** مرچنت سفارش را بررسی کرده و موجودی کالا و صحت اطلاعات را تأیید نموده است. سفارش آماده ارسال است.  **چگونه اتفاق می‌افتد:** مرچنت با ارسال درخواست `POST /api/v1/orders/manager/{uuid}/verify` سفارش را تأیید می‌کند. این مرحله **اجباری** است و باید پس از پرداخت موفق انجام شود.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` باشد. مرچنت باید موجودی کالا را بررسی کند.  ---  ### ۲۰. SHIPPED — ارسال شد  **معنا:** سفارش از انبار خارج شده و در حال ارسال به مشتری است.  **چگونه اتفاق می‌افتد:** مرچنت پس از ارسال کالا، وضعیت سفارش را از طریق API به `SHIPPED` تغییر می‌دهد.  **وابستگی‌ها:** سفارش باید در وضعیت `VERIFIED_BY_MERCHANT` باشد.  ---  ### ۱۹. DELIVERED — تحویل داده شد  **معنا:** سفارش به دست مشتری رسیده و فرآیند خرید به پایان رسیده است.  **چگونه اتفاق می‌افتد:** مرچنت پس از تحویل موفق، وضعیت را به `DELIVERED` تغییر می‌دهد.  **وابستگی‌ها:** سفارش باید در وضعیت `SHIPPED` باشد.  ---  ### ۵. EXPIRED — منقضی شد  **معنا:** زمان رزرو سفارش به پایان رسیده و سفارش به صورت خودکار لغو شده است.  **چگونه اتفاق می‌افتد:** یک Task دوره‌ای به طور خودکار سفارش‌هایی که `reservation_expired_at` آن‌ها گذشته را پیدا کرده و وضعیتشان را به `EXPIRED` تغییر می‌دهد. این مکانیزم مانع بلوکه شدن موجودی کالا می‌شود.  **وابستگی‌ها:** سفارش باید در یکی از وضعیت‌های `INITIAL`، `STARTED`، `PENDING` یا `WAITING_FOR_GATEWAY` باشد و زمان رزرو آن گذشته باشد.  ---  ### ۱۸. EXPIRATION_TIME_EXCEEDED — زمان انقضا گذشت  **معنا:** در لحظه ثبت نهایی یا پرداخت، مشخص شد که زمان مجاز سفارش تمام شده است.  **چگونه اتفاق می‌افتد:** هنگام ارسال درخواست پرداخت (`submit_order`)، سیستم بررسی می‌کند که `expiration_time` سفارش هنوز معتبر است یا خیر. در صورت گذشتن زمان، وضعیت به `EXPIRATION_TIME_EXCEEDED` تغییر می‌کند.  **وابستگی‌ها:** سفارش در وضعیت `PENDING` یا `WAITING_FOR_GATEWAY` است و فیلد `expiration_time` سپری شده.  ---  ### ۶. CANCELLED — لغو توسط مشتری  **معنا:** مشتری در حین فرآیند خرید (قبل از پرداخت) سفارش را لغو کرده یا از صفحه خارج شده است.  **چگونه اتفاق می‌افتد:** مشتری در صفحه checkout دکمه «انصراف» را می‌زند یا پرداخت ناموفق بوده و سفارش به حالت لغو درمی‌آید.  **وابستگی‌ها:** سفارش باید در وضعیت `PENDING` یا `WAITING_FOR_GATEWAY` باشد. پرداختی انجام نشده است.  ---  ### ۸. FAILED_TO_PAY — پرداخت ناموفق  **معنا:** تراکنش پرداخت انجام نشد یا با خطا مواجه شد.  **چگونه اتفاق می‌افتد:** درگاه پرداخت نتیجه ناموفق برمی‌گرداند یا فرآیند بازگشت وجه در مرحله پرداخت با شکست مواجه می‌شود.  **وابستگی‌ها:** سفارش باید در وضعیت `WAITING_FOR_GATEWAY` بوده باشد.  ---  ### ۱۰. FAILED_TO_VERIFY_BY_MERCHANT — تأیید ناموفق توسط مرچنت  **معنا:** مرچنت سفارش را رد کرده است؛ معمولاً به دلیل ناموجود بودن کالا یا مغایرت اطلاعات.  **چگونه اتفاق می‌افتد:** مرچنت در پاسخ به درخواست verify، خطا برمی‌گرداند یا API آن وضعیت ناموفق تنظیم می‌کند. پس از این وضعیت، فرآیند استرداد وجه آغاز می‌شود.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` باشد.  ---  ### ۱۱. FAILED_BY_MERCHANT — خطا از سمت مرچنت  **معنا:** مرچنت پس از تأیید اولیه، اعلام می‌کند که قادر به انجام سفارش نیست (مثلاً به دلیل اتمام موجودی).  **چگونه اتفاق می‌افتد:** مرچنت وضعیت را به `FAILED_BY_MERCHANT` تغییر می‌دهد. وجه پرداختی مشتری مسترد خواهد شد.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` باشد.  ---  ### ۱۲. CANCELLED_BY_MERCHANT — لغو توسط مرچنت  **معنا:** مرچنت پس از پرداخت، سفارش را به هر دلیلی لغو کرده است.  **چگونه اتفاق می‌افتد:** مرچنت درخواست لغو سفارش را ارسال می‌کند. وجه پرداختی مشتری به او بازگردانده می‌شود.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` یا `VERIFIED_BY_MERCHANT` باشد.  ---  ### ۱۳. REQUEST_TO_REFUND — درخواست استرداد توسط مشتری  **معنا:** مشتری درخواست بازگشت وجه داده و سیستم در حال پردازش استرداد است.  **چگونه اتفاق می‌افتد:** مرچنت از طریق API درخواست استرداد را ثبت می‌کند (`POST /api/v1/orders/manager/{uuid}/refund`). سفارش وارد صف پردازش استرداد می‌شود.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` یا `VERIFIED_BY_MERCHANT` باشد.  ---  ### ۱۴، ۱۵، ۱۶. سایر وضعیت‌های درخواست استرداد  این وضعیت‌ها بر اساس دلیل استرداد از هم تفکیک می‌شوند:  - **۱۴ — REQUEST_TO_REFUND_TO_MERCHANT_AFTER_FAILED_TO_VERIFY:** استرداد پس از ناموفق بودن تأیید مرچنت؛ وجه به حساب مرچنت بازمی‌گردد. - **۱۵ — REQUEST_TO_REFUND_TO_CUSTOMER_AFTER_FAILED_BY_MERCHANT:** استرداد پس از خطای مرچنت؛ وجه به مشتری بازمی‌گردد. - **۱۶ — REQUEST_TO_REFUND_TO_MERCHANT_AFTER_CANCELLED_BY_MERCHANT:** استرداد پس از لغو توسط مرچنت؛ وجه به حساب مرچنت برمی‌گردد.  **چگونه اتفاق می‌افتد:** به صورت خودکار پس از رسیدن به وضعیت‌های ناموفق/لغو مربوطه توسط سیستم تنظیم می‌شود.  ---  ### ۱۷. REFUND_COMPLETED — استرداد تکمیل شد  **معنا:** وجه با موفقیت به صاحب آن (مشتری یا مرچنت بسته به نوع استرداد) بازگردانده شده است.  **چگونه اتفاق می‌افتد:** Task پردازش استرداد (`process_order_refund`) پس از تأیید موفق بازگشت وجه از سوی Digipay، وضعیت سفارش را به `REFUND_COMPLETED` تغییر می‌دهد.  **وابستگی‌ها:** یکی از وضعیت‌های درخواست استرداد (۱۳، ۱۴، ۱۵ یا ۱۶) باید فعال باشد و Digipay تراکنش استرداد را تأیید کرده باشد.  --> </div> 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bwdk_sdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OrderCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreate{}

// OrderCreate struct for OrderCreate
type OrderCreate struct {
	// شناسه منحصر به فرد این سفارش در سیستم فروشنده
	MerchantOrderId string `json:"merchant_order_id"`
	// شناسه منحصر به فرد برای تأیید اصالت سفارش
	MerchantUniqueId string `json:"merchant_unique_id"`
	// مجموع قیمت‌های اولیه تمام کالاها بدون تخفیف (به تومان)
	MainAmount *int32 `json:"main_amount,omitempty"`
	// مبلغ نهایی: مبلغ_اصلی - مبلغ_تخفیف + مبلغ_مالیات (به تومان)
	FinalAmount *int32 `json:"final_amount,omitempty"`
	// مبلغ کل پرداخت شده توسط کاربر: مبلغ_نهایی + هزینه_ارسال (به تومان)
	TotalPaidAmount *int32 `json:"total_paid_amount,omitempty"`
	// مبلغ کل تخفیف برای تمام سفارش (به تومان)
	DiscountAmount *int32 `json:"discount_amount,omitempty"`
	// مبلغ کل مالیات برای تمام سفارش (به تومان)
	TaxAmount *int32 `json:"tax_amount,omitempty"`
	// هزینه ارسال برای سفارش (به تومان)
	ShippingAmount *int32 `json:"shipping_amount,omitempty"`
	// مبلغ تخفیف باشگاه مشتریان/پاداش (به تومان)
	LoyaltyAmount *int32 `json:"loyalty_amount,omitempty"`
	// آدرس وب‌هوک برای دریافت اطلاع رسانی وضعیت پرداخت
	CallbackUrl string `json:"callback_url"`
	DestinationAddress interface{} `json:"destination_address"`
	Items []OrderItemCreate `json:"items"`
	// مقدار توسط سیستم جایگذاری می شود
	Merchant *int32 `json:"merchant,omitempty"`
	// مقدار توسط سیستم جایگذاری می شود
	SourceAddress interface{} `json:"source_address,omitempty"`
	User NullableInt32 `json:"user"`
	// مهلت پرداخت (به عنوان Unix timestamp) قبل از اتمام سفارش
	ReservationExpiredAt NullableInt32 `json:"reservation_expired_at,omitempty"`
	// کد مرجع یکتا برای پیگیری سفارش مشتری (قالب: BD-XXXXXXXX)
	ReferenceCode string `json:"reference_code"`
	// Preparation time for the order (in days)
	PreparationTime *int32 `json:"preparation_time,omitempty"`
	// Total Weight of the order (in grams)
	Weight *float64 `json:"weight,omitempty"`
}

type _OrderCreate OrderCreate

// NewOrderCreate instantiates a new OrderCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreate(merchantOrderId string, merchantUniqueId string, callbackUrl string, destinationAddress interface{}, items []OrderItemCreate, user NullableInt32, referenceCode string) *OrderCreate {
	this := OrderCreate{}
	this.MerchantOrderId = merchantOrderId
	this.MerchantUniqueId = merchantUniqueId
	this.CallbackUrl = callbackUrl
	this.DestinationAddress = destinationAddress
	this.Items = items
	this.User = user
	this.ReferenceCode = referenceCode
	var preparationTime int32 = 2
	this.PreparationTime = &preparationTime
	return &this
}

// NewOrderCreateWithDefaults instantiates a new OrderCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateWithDefaults() *OrderCreate {
	this := OrderCreate{}
	var preparationTime int32 = 2
	this.PreparationTime = &preparationTime
	return &this
}

// GetMerchantOrderId returns the MerchantOrderId field value
func (o *OrderCreate) GetMerchantOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantOrderId
}

// GetMerchantOrderIdOk returns a tuple with the MerchantOrderId field value
// and a boolean to check if the value has been set.
func (o *OrderCreate) GetMerchantOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantOrderId, true
}

// SetMerchantOrderId sets field value
func (o *OrderCreate) SetMerchantOrderId(v string) {
	o.MerchantOrderId = v
}

// GetMerchantUniqueId returns the MerchantUniqueId field value
func (o *OrderCreate) GetMerchantUniqueId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantUniqueId
}

// GetMerchantUniqueIdOk returns a tuple with the MerchantUniqueId field value
// and a boolean to check if the value has been set.
func (o *OrderCreate) GetMerchantUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantUniqueId, true
}

// SetMerchantUniqueId sets field value
func (o *OrderCreate) SetMerchantUniqueId(v string) {
	o.MerchantUniqueId = v
}

// GetMainAmount returns the MainAmount field value if set, zero value otherwise.
func (o *OrderCreate) GetMainAmount() int32 {
	if o == nil || IsNil(o.MainAmount) {
		var ret int32
		return ret
	}
	return *o.MainAmount
}

// GetMainAmountOk returns a tuple with the MainAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreate) GetMainAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.MainAmount) {
		return nil, false
	}
	return o.MainAmount, true
}

// HasMainAmount returns a boolean if a field has been set.
func (o *OrderCreate) HasMainAmount() bool {
	if o != nil && !IsNil(o.MainAmount) {
		return true
	}

	return false
}

// SetMainAmount gets a reference to the given int32 and assigns it to the MainAmount field.
func (o *OrderCreate) SetMainAmount(v int32) {
	o.MainAmount = &v
}

// GetFinalAmount returns the FinalAmount field value if set, zero value otherwise.
func (o *OrderCreate) GetFinalAmount() int32 {
	if o == nil || IsNil(o.FinalAmount) {
		var ret int32
		return ret
	}
	return *o.FinalAmount
}

// GetFinalAmountOk returns a tuple with the FinalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreate) GetFinalAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.FinalAmount) {
		return nil, false
	}
	return o.FinalAmount, true
}

// HasFinalAmount returns a boolean if a field has been set.
func (o *OrderCreate) HasFinalAmount() bool {
	if o != nil && !IsNil(o.FinalAmount) {
		return true
	}

	return false
}

// SetFinalAmount gets a reference to the given int32 and assigns it to the FinalAmount field.
func (o *OrderCreate) SetFinalAmount(v int32) {
	o.FinalAmount = &v
}

// GetTotalPaidAmount returns the TotalPaidAmount field value if set, zero value otherwise.
func (o *OrderCreate) GetTotalPaidAmount() int32 {
	if o == nil || IsNil(o.TotalPaidAmount) {
		var ret int32
		return ret
	}
	return *o.TotalPaidAmount
}

// GetTotalPaidAmountOk returns a tuple with the TotalPaidAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreate) GetTotalPaidAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPaidAmount) {
		return nil, false
	}
	return o.TotalPaidAmount, true
}

// HasTotalPaidAmount returns a boolean if a field has been set.
func (o *OrderCreate) HasTotalPaidAmount() bool {
	if o != nil && !IsNil(o.TotalPaidAmount) {
		return true
	}

	return false
}

// SetTotalPaidAmount gets a reference to the given int32 and assigns it to the TotalPaidAmount field.
func (o *OrderCreate) SetTotalPaidAmount(v int32) {
	o.TotalPaidAmount = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *OrderCreate) GetDiscountAmount() int32 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret int32
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreate) GetDiscountAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *OrderCreate) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given int32 and assigns it to the DiscountAmount field.
func (o *OrderCreate) SetDiscountAmount(v int32) {
	o.DiscountAmount = &v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *OrderCreate) GetTaxAmount() int32 {
	if o == nil || IsNil(o.TaxAmount) {
		var ret int32
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreate) GetTaxAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.TaxAmount) {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *OrderCreate) HasTaxAmount() bool {
	if o != nil && !IsNil(o.TaxAmount) {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given int32 and assigns it to the TaxAmount field.
func (o *OrderCreate) SetTaxAmount(v int32) {
	o.TaxAmount = &v
}

// GetShippingAmount returns the ShippingAmount field value if set, zero value otherwise.
func (o *OrderCreate) GetShippingAmount() int32 {
	if o == nil || IsNil(o.ShippingAmount) {
		var ret int32
		return ret
	}
	return *o.ShippingAmount
}

// GetShippingAmountOk returns a tuple with the ShippingAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreate) GetShippingAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.ShippingAmount) {
		return nil, false
	}
	return o.ShippingAmount, true
}

// HasShippingAmount returns a boolean if a field has been set.
func (o *OrderCreate) HasShippingAmount() bool {
	if o != nil && !IsNil(o.ShippingAmount) {
		return true
	}

	return false
}

// SetShippingAmount gets a reference to the given int32 and assigns it to the ShippingAmount field.
func (o *OrderCreate) SetShippingAmount(v int32) {
	o.ShippingAmount = &v
}

// GetLoyaltyAmount returns the LoyaltyAmount field value if set, zero value otherwise.
func (o *OrderCreate) GetLoyaltyAmount() int32 {
	if o == nil || IsNil(o.LoyaltyAmount) {
		var ret int32
		return ret
	}
	return *o.LoyaltyAmount
}

// GetLoyaltyAmountOk returns a tuple with the LoyaltyAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreate) GetLoyaltyAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.LoyaltyAmount) {
		return nil, false
	}
	return o.LoyaltyAmount, true
}

// HasLoyaltyAmount returns a boolean if a field has been set.
func (o *OrderCreate) HasLoyaltyAmount() bool {
	if o != nil && !IsNil(o.LoyaltyAmount) {
		return true
	}

	return false
}

// SetLoyaltyAmount gets a reference to the given int32 and assigns it to the LoyaltyAmount field.
func (o *OrderCreate) SetLoyaltyAmount(v int32) {
	o.LoyaltyAmount = &v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *OrderCreate) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *OrderCreate) GetCallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *OrderCreate) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetDestinationAddress returns the DestinationAddress field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *OrderCreate) GetDestinationAddress() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.DestinationAddress
}

// GetDestinationAddressOk returns a tuple with the DestinationAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderCreate) GetDestinationAddressOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DestinationAddress) {
		return nil, false
	}
	return &o.DestinationAddress, true
}

// SetDestinationAddress sets field value
func (o *OrderCreate) SetDestinationAddress(v interface{}) {
	o.DestinationAddress = v
}

// GetItems returns the Items field value
func (o *OrderCreate) GetItems() []OrderItemCreate {
	if o == nil {
		var ret []OrderItemCreate
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *OrderCreate) GetItemsOk() ([]OrderItemCreate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *OrderCreate) SetItems(v []OrderItemCreate) {
	o.Items = v
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *OrderCreate) GetMerchant() int32 {
	if o == nil || IsNil(o.Merchant) {
		var ret int32
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreate) GetMerchantOk() (*int32, bool) {
	if o == nil || IsNil(o.Merchant) {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *OrderCreate) HasMerchant() bool {
	if o != nil && !IsNil(o.Merchant) {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given int32 and assigns it to the Merchant field.
func (o *OrderCreate) SetMerchant(v int32) {
	o.Merchant = &v
}

// GetSourceAddress returns the SourceAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderCreate) GetSourceAddress() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SourceAddress
}

// GetSourceAddressOk returns a tuple with the SourceAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderCreate) GetSourceAddressOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SourceAddress) {
		return nil, false
	}
	return &o.SourceAddress, true
}

// HasSourceAddress returns a boolean if a field has been set.
func (o *OrderCreate) HasSourceAddress() bool {
	if o != nil && !IsNil(o.SourceAddress) {
		return true
	}

	return false
}

// SetSourceAddress gets a reference to the given interface{} and assigns it to the SourceAddress field.
func (o *OrderCreate) SetSourceAddress(v interface{}) {
	o.SourceAddress = v
}

// GetUser returns the User field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *OrderCreate) GetUser() int32 {
	if o == nil || o.User.Get() == nil {
		var ret int32
		return ret
	}

	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderCreate) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// SetUser sets field value
func (o *OrderCreate) SetUser(v int32) {
	o.User.Set(&v)
}

// GetReservationExpiredAt returns the ReservationExpiredAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderCreate) GetReservationExpiredAt() int32 {
	if o == nil || IsNil(o.ReservationExpiredAt.Get()) {
		var ret int32
		return ret
	}
	return *o.ReservationExpiredAt.Get()
}

// GetReservationExpiredAtOk returns a tuple with the ReservationExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderCreate) GetReservationExpiredAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReservationExpiredAt.Get(), o.ReservationExpiredAt.IsSet()
}

// HasReservationExpiredAt returns a boolean if a field has been set.
func (o *OrderCreate) HasReservationExpiredAt() bool {
	if o != nil && o.ReservationExpiredAt.IsSet() {
		return true
	}

	return false
}

// SetReservationExpiredAt gets a reference to the given NullableInt32 and assigns it to the ReservationExpiredAt field.
func (o *OrderCreate) SetReservationExpiredAt(v int32) {
	o.ReservationExpiredAt.Set(&v)
}
// SetReservationExpiredAtNil sets the value for ReservationExpiredAt to be an explicit nil
func (o *OrderCreate) SetReservationExpiredAtNil() {
	o.ReservationExpiredAt.Set(nil)
}

// UnsetReservationExpiredAt ensures that no value is present for ReservationExpiredAt, not even an explicit nil
func (o *OrderCreate) UnsetReservationExpiredAt() {
	o.ReservationExpiredAt.Unset()
}

// GetReferenceCode returns the ReferenceCode field value
func (o *OrderCreate) GetReferenceCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceCode
}

// GetReferenceCodeOk returns a tuple with the ReferenceCode field value
// and a boolean to check if the value has been set.
func (o *OrderCreate) GetReferenceCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceCode, true
}

// SetReferenceCode sets field value
func (o *OrderCreate) SetReferenceCode(v string) {
	o.ReferenceCode = v
}

// GetPreparationTime returns the PreparationTime field value if set, zero value otherwise.
func (o *OrderCreate) GetPreparationTime() int32 {
	if o == nil || IsNil(o.PreparationTime) {
		var ret int32
		return ret
	}
	return *o.PreparationTime
}

// GetPreparationTimeOk returns a tuple with the PreparationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreate) GetPreparationTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.PreparationTime) {
		return nil, false
	}
	return o.PreparationTime, true
}

// HasPreparationTime returns a boolean if a field has been set.
func (o *OrderCreate) HasPreparationTime() bool {
	if o != nil && !IsNil(o.PreparationTime) {
		return true
	}

	return false
}

// SetPreparationTime gets a reference to the given int32 and assigns it to the PreparationTime field.
func (o *OrderCreate) SetPreparationTime(v int32) {
	o.PreparationTime = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *OrderCreate) GetWeight() float64 {
	if o == nil || IsNil(o.Weight) {
		var ret float64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreate) GetWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *OrderCreate) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float64 and assigns it to the Weight field.
func (o *OrderCreate) SetWeight(v float64) {
	o.Weight = &v
}

func (o OrderCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchant_order_id"] = o.MerchantOrderId
	toSerialize["merchant_unique_id"] = o.MerchantUniqueId
	if !IsNil(o.MainAmount) {
		toSerialize["main_amount"] = o.MainAmount
	}
	if !IsNil(o.FinalAmount) {
		toSerialize["final_amount"] = o.FinalAmount
	}
	if !IsNil(o.TotalPaidAmount) {
		toSerialize["total_paid_amount"] = o.TotalPaidAmount
	}
	if !IsNil(o.DiscountAmount) {
		toSerialize["discount_amount"] = o.DiscountAmount
	}
	if !IsNil(o.TaxAmount) {
		toSerialize["tax_amount"] = o.TaxAmount
	}
	if !IsNil(o.ShippingAmount) {
		toSerialize["shipping_amount"] = o.ShippingAmount
	}
	if !IsNil(o.LoyaltyAmount) {
		toSerialize["loyalty_amount"] = o.LoyaltyAmount
	}
	toSerialize["callback_url"] = o.CallbackUrl
	if o.DestinationAddress != nil {
		toSerialize["destination_address"] = o.DestinationAddress
	}
	toSerialize["items"] = o.Items
	if !IsNil(o.Merchant) {
		toSerialize["merchant"] = o.Merchant
	}
	if o.SourceAddress != nil {
		toSerialize["source_address"] = o.SourceAddress
	}
	toSerialize["user"] = o.User.Get()
	if o.ReservationExpiredAt.IsSet() {
		toSerialize["reservation_expired_at"] = o.ReservationExpiredAt.Get()
	}
	toSerialize["reference_code"] = o.ReferenceCode
	if !IsNil(o.PreparationTime) {
		toSerialize["preparation_time"] = o.PreparationTime
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

func (o *OrderCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"merchant_order_id",
		"merchant_unique_id",
		"callback_url",
		"destination_address",
		"items",
		"user",
		"reference_code",
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

	varOrderCreate := _OrderCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderCreate)

	if err != nil {
		return err
	}

	*o = OrderCreate(varOrderCreate)

	return err
}

type NullableOrderCreate struct {
	value *OrderCreate
	isSet bool
}

func (v NullableOrderCreate) Get() *OrderCreate {
	return v.value
}

func (v *NullableOrderCreate) Set(val *OrderCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreate(val *OrderCreate) *NullableOrderCreate {
	return &NullableOrderCreate{value: val, isSet: true}
}

func (v NullableOrderCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


