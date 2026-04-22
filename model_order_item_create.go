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

// checks if the OrderItemCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderItemCreate{}

// OrderItemCreate struct for OrderItemCreate
type OrderItemCreate struct {
	// نام کامل محصول شامل تمام مشخصات
	Name string `json:"name"`
	// قیمت اولیه برای هر واحد بدون تخفیف (به تومان)
	PrimaryAmount *int32 `json:"primary_amount,omitempty"`
	// قیمت نهایی برای تمام واحدها بعد از تخفیف (به تومان)
	Amount *int32 `json:"amount,omitempty"`
	// تعداد واحدهای این کالا در سفارش
	Count int32 `json:"count"`
	// مبلغ کل تخفیف برای این کالا (به تومان)
	DiscountAmount *int32 `json:"discount_amount,omitempty"`
	// مبلغ کل مالیات برای این کالا (به تومان)
	TaxAmount *int32 `json:"tax_amount,omitempty"`
	// آدرس تصویر محصول
	ImageLink NullableString `json:"image_link,omitempty"`
	Options []Option `json:"options"`
	// Preparation time for the item (in days)
	PreparationTime *int32 `json:"preparation_time,omitempty"`
	// Weight of the item (in grams)
	Weight *float64 `json:"weight,omitempty"`
}

type _OrderItemCreate OrderItemCreate

// NewOrderItemCreate instantiates a new OrderItemCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderItemCreate(name string, count int32, options []Option) *OrderItemCreate {
	this := OrderItemCreate{}
	this.Name = name
	this.Count = count
	this.Options = options
	var preparationTime int32 = 2
	this.PreparationTime = &preparationTime
	return &this
}

// NewOrderItemCreateWithDefaults instantiates a new OrderItemCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderItemCreateWithDefaults() *OrderItemCreate {
	this := OrderItemCreate{}
	var preparationTime int32 = 2
	this.PreparationTime = &preparationTime
	return &this
}

// GetName returns the Name field value
func (o *OrderItemCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrderItemCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrderItemCreate) SetName(v string) {
	o.Name = v
}

// GetPrimaryAmount returns the PrimaryAmount field value if set, zero value otherwise.
func (o *OrderItemCreate) GetPrimaryAmount() int32 {
	if o == nil || IsNil(o.PrimaryAmount) {
		var ret int32
		return ret
	}
	return *o.PrimaryAmount
}

// GetPrimaryAmountOk returns a tuple with the PrimaryAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemCreate) GetPrimaryAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.PrimaryAmount) {
		return nil, false
	}
	return o.PrimaryAmount, true
}

// HasPrimaryAmount returns a boolean if a field has been set.
func (o *OrderItemCreate) HasPrimaryAmount() bool {
	if o != nil && !IsNil(o.PrimaryAmount) {
		return true
	}

	return false
}

// SetPrimaryAmount gets a reference to the given int32 and assigns it to the PrimaryAmount field.
func (o *OrderItemCreate) SetPrimaryAmount(v int32) {
	o.PrimaryAmount = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *OrderItemCreate) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemCreate) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *OrderItemCreate) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *OrderItemCreate) SetAmount(v int32) {
	o.Amount = &v
}

// GetCount returns the Count field value
func (o *OrderItemCreate) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *OrderItemCreate) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *OrderItemCreate) SetCount(v int32) {
	o.Count = v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *OrderItemCreate) GetDiscountAmount() int32 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret int32
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemCreate) GetDiscountAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *OrderItemCreate) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given int32 and assigns it to the DiscountAmount field.
func (o *OrderItemCreate) SetDiscountAmount(v int32) {
	o.DiscountAmount = &v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *OrderItemCreate) GetTaxAmount() int32 {
	if o == nil || IsNil(o.TaxAmount) {
		var ret int32
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemCreate) GetTaxAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.TaxAmount) {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *OrderItemCreate) HasTaxAmount() bool {
	if o != nil && !IsNil(o.TaxAmount) {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given int32 and assigns it to the TaxAmount field.
func (o *OrderItemCreate) SetTaxAmount(v int32) {
	o.TaxAmount = &v
}

// GetImageLink returns the ImageLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderItemCreate) GetImageLink() string {
	if o == nil || IsNil(o.ImageLink.Get()) {
		var ret string
		return ret
	}
	return *o.ImageLink.Get()
}

// GetImageLinkOk returns a tuple with the ImageLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderItemCreate) GetImageLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageLink.Get(), o.ImageLink.IsSet()
}

// HasImageLink returns a boolean if a field has been set.
func (o *OrderItemCreate) HasImageLink() bool {
	if o != nil && o.ImageLink.IsSet() {
		return true
	}

	return false
}

// SetImageLink gets a reference to the given NullableString and assigns it to the ImageLink field.
func (o *OrderItemCreate) SetImageLink(v string) {
	o.ImageLink.Set(&v)
}
// SetImageLinkNil sets the value for ImageLink to be an explicit nil
func (o *OrderItemCreate) SetImageLinkNil() {
	o.ImageLink.Set(nil)
}

// UnsetImageLink ensures that no value is present for ImageLink, not even an explicit nil
func (o *OrderItemCreate) UnsetImageLink() {
	o.ImageLink.Unset()
}

// GetOptions returns the Options field value
func (o *OrderItemCreate) GetOptions() []Option {
	if o == nil {
		var ret []Option
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *OrderItemCreate) GetOptionsOk() ([]Option, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *OrderItemCreate) SetOptions(v []Option) {
	o.Options = v
}

// GetPreparationTime returns the PreparationTime field value if set, zero value otherwise.
func (o *OrderItemCreate) GetPreparationTime() int32 {
	if o == nil || IsNil(o.PreparationTime) {
		var ret int32
		return ret
	}
	return *o.PreparationTime
}

// GetPreparationTimeOk returns a tuple with the PreparationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemCreate) GetPreparationTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.PreparationTime) {
		return nil, false
	}
	return o.PreparationTime, true
}

// HasPreparationTime returns a boolean if a field has been set.
func (o *OrderItemCreate) HasPreparationTime() bool {
	if o != nil && !IsNil(o.PreparationTime) {
		return true
	}

	return false
}

// SetPreparationTime gets a reference to the given int32 and assigns it to the PreparationTime field.
func (o *OrderItemCreate) SetPreparationTime(v int32) {
	o.PreparationTime = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *OrderItemCreate) GetWeight() float64 {
	if o == nil || IsNil(o.Weight) {
		var ret float64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemCreate) GetWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *OrderItemCreate) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float64 and assigns it to the Weight field.
func (o *OrderItemCreate) SetWeight(v float64) {
	o.Weight = &v
}

func (o OrderItemCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderItemCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.PrimaryAmount) {
		toSerialize["primary_amount"] = o.PrimaryAmount
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	toSerialize["count"] = o.Count
	if !IsNil(o.DiscountAmount) {
		toSerialize["discount_amount"] = o.DiscountAmount
	}
	if !IsNil(o.TaxAmount) {
		toSerialize["tax_amount"] = o.TaxAmount
	}
	if o.ImageLink.IsSet() {
		toSerialize["image_link"] = o.ImageLink.Get()
	}
	toSerialize["options"] = o.Options
	if !IsNil(o.PreparationTime) {
		toSerialize["preparation_time"] = o.PreparationTime
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

func (o *OrderItemCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"count",
		"options",
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

	varOrderItemCreate := _OrderItemCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderItemCreate)

	if err != nil {
		return err
	}

	*o = OrderItemCreate(varOrderItemCreate)

	return err
}

type NullableOrderItemCreate struct {
	value *OrderItemCreate
	isSet bool
}

func (v NullableOrderItemCreate) Get() *OrderItemCreate {
	return v.value
}

func (v *NullableOrderItemCreate) Set(val *OrderItemCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderItemCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderItemCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderItemCreate(val *OrderItemCreate) *NullableOrderItemCreate {
	return &NullableOrderItemCreate{value: val, isSet: true}
}

func (v NullableOrderItemCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderItemCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


