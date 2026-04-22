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

// checks if the ShippingMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingMethod{}

// ShippingMethod Serializer for shipping method details.
type ShippingMethod struct {
	Id int32 `json:"id"`
	// نام روش ارسال
	Name string `json:"name"`
	// توضیحات روش ارسال و جزئیات تحویل آن
	Description *string `json:"description,omitempty"`
	// نوع روش ارسال: عادی یا دیجی اکسپرس  * `1` - سایر * `2` - دیجی اکسپرس
	ShippingType *ShippingTypeEnum `json:"shipping_type,omitempty"`
	GetShippingTypeDisplay string `json:"get_shipping_type_display"`
	ShippingTypeDisplay string `json:"shipping_type_display"`
	// هزینه ارسال برای منطقه اولیه (مثلاً تهران) به تومان
	Cost *int32 `json:"cost,omitempty"`
	// هزینه ارسال برای مناطق دیگر به تومان
	SecondaryCost *int32 `json:"secondary_cost,omitempty"`
	// حداقل تعداد روزها از تاریخ سفارش تا تحویل
	MinimumTimeSending *int32 `json:"minimum_time_sending,omitempty"`
	// حداکثر تعداد روزها از تاریخ سفارش تا تحویل
	MaximumTimeSending *int32 `json:"maximum_time_sending,omitempty"`
	DeliveryTimeDisplay string `json:"delivery_time_display"`
	DeliveryTimeRangeDisplay DeliveryTimeRangeDisplay `json:"delivery_time_range_display"`
	InventoryAddress BusinessAddress `json:"inventory_address"`
	// Whether the shipping method is pay at destination
	IsPayAtDestination *bool `json:"is_pay_at_destination,omitempty"`
}

type _ShippingMethod ShippingMethod

// NewShippingMethod instantiates a new ShippingMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethod(id int32, name string, getShippingTypeDisplay string, shippingTypeDisplay string, deliveryTimeDisplay string, deliveryTimeRangeDisplay DeliveryTimeRangeDisplay, inventoryAddress BusinessAddress) *ShippingMethod {
	this := ShippingMethod{}
	this.Id = id
	this.Name = name
	this.GetShippingTypeDisplay = getShippingTypeDisplay
	this.ShippingTypeDisplay = shippingTypeDisplay
	this.DeliveryTimeDisplay = deliveryTimeDisplay
	this.DeliveryTimeRangeDisplay = deliveryTimeRangeDisplay
	this.InventoryAddress = inventoryAddress
	return &this
}

// NewShippingMethodWithDefaults instantiates a new ShippingMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodWithDefaults() *ShippingMethod {
	this := ShippingMethod{}
	return &this
}

// GetId returns the Id field value
func (o *ShippingMethod) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ShippingMethod) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ShippingMethod) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ShippingMethod) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ShippingMethod) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ShippingMethod) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ShippingMethod) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethod) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ShippingMethod) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ShippingMethod) SetDescription(v string) {
	o.Description = &v
}

// GetShippingType returns the ShippingType field value if set, zero value otherwise.
func (o *ShippingMethod) GetShippingType() ShippingTypeEnum {
	if o == nil || IsNil(o.ShippingType) {
		var ret ShippingTypeEnum
		return ret
	}
	return *o.ShippingType
}

// GetShippingTypeOk returns a tuple with the ShippingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethod) GetShippingTypeOk() (*ShippingTypeEnum, bool) {
	if o == nil || IsNil(o.ShippingType) {
		return nil, false
	}
	return o.ShippingType, true
}

// HasShippingType returns a boolean if a field has been set.
func (o *ShippingMethod) HasShippingType() bool {
	if o != nil && !IsNil(o.ShippingType) {
		return true
	}

	return false
}

// SetShippingType gets a reference to the given ShippingTypeEnum and assigns it to the ShippingType field.
func (o *ShippingMethod) SetShippingType(v ShippingTypeEnum) {
	o.ShippingType = &v
}

// GetGetShippingTypeDisplay returns the GetShippingTypeDisplay field value
func (o *ShippingMethod) GetGetShippingTypeDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GetShippingTypeDisplay
}

// GetGetShippingTypeDisplayOk returns a tuple with the GetShippingTypeDisplay field value
// and a boolean to check if the value has been set.
func (o *ShippingMethod) GetGetShippingTypeDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GetShippingTypeDisplay, true
}

// SetGetShippingTypeDisplay sets field value
func (o *ShippingMethod) SetGetShippingTypeDisplay(v string) {
	o.GetShippingTypeDisplay = v
}

// GetShippingTypeDisplay returns the ShippingTypeDisplay field value
func (o *ShippingMethod) GetShippingTypeDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShippingTypeDisplay
}

// GetShippingTypeDisplayOk returns a tuple with the ShippingTypeDisplay field value
// and a boolean to check if the value has been set.
func (o *ShippingMethod) GetShippingTypeDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingTypeDisplay, true
}

// SetShippingTypeDisplay sets field value
func (o *ShippingMethod) SetShippingTypeDisplay(v string) {
	o.ShippingTypeDisplay = v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *ShippingMethod) GetCost() int32 {
	if o == nil || IsNil(o.Cost) {
		var ret int32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethod) GetCostOk() (*int32, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *ShippingMethod) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given int32 and assigns it to the Cost field.
func (o *ShippingMethod) SetCost(v int32) {
	o.Cost = &v
}

// GetSecondaryCost returns the SecondaryCost field value if set, zero value otherwise.
func (o *ShippingMethod) GetSecondaryCost() int32 {
	if o == nil || IsNil(o.SecondaryCost) {
		var ret int32
		return ret
	}
	return *o.SecondaryCost
}

// GetSecondaryCostOk returns a tuple with the SecondaryCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethod) GetSecondaryCostOk() (*int32, bool) {
	if o == nil || IsNil(o.SecondaryCost) {
		return nil, false
	}
	return o.SecondaryCost, true
}

// HasSecondaryCost returns a boolean if a field has been set.
func (o *ShippingMethod) HasSecondaryCost() bool {
	if o != nil && !IsNil(o.SecondaryCost) {
		return true
	}

	return false
}

// SetSecondaryCost gets a reference to the given int32 and assigns it to the SecondaryCost field.
func (o *ShippingMethod) SetSecondaryCost(v int32) {
	o.SecondaryCost = &v
}

// GetMinimumTimeSending returns the MinimumTimeSending field value if set, zero value otherwise.
func (o *ShippingMethod) GetMinimumTimeSending() int32 {
	if o == nil || IsNil(o.MinimumTimeSending) {
		var ret int32
		return ret
	}
	return *o.MinimumTimeSending
}

// GetMinimumTimeSendingOk returns a tuple with the MinimumTimeSending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethod) GetMinimumTimeSendingOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumTimeSending) {
		return nil, false
	}
	return o.MinimumTimeSending, true
}

// HasMinimumTimeSending returns a boolean if a field has been set.
func (o *ShippingMethod) HasMinimumTimeSending() bool {
	if o != nil && !IsNil(o.MinimumTimeSending) {
		return true
	}

	return false
}

// SetMinimumTimeSending gets a reference to the given int32 and assigns it to the MinimumTimeSending field.
func (o *ShippingMethod) SetMinimumTimeSending(v int32) {
	o.MinimumTimeSending = &v
}

// GetMaximumTimeSending returns the MaximumTimeSending field value if set, zero value otherwise.
func (o *ShippingMethod) GetMaximumTimeSending() int32 {
	if o == nil || IsNil(o.MaximumTimeSending) {
		var ret int32
		return ret
	}
	return *o.MaximumTimeSending
}

// GetMaximumTimeSendingOk returns a tuple with the MaximumTimeSending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethod) GetMaximumTimeSendingOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumTimeSending) {
		return nil, false
	}
	return o.MaximumTimeSending, true
}

// HasMaximumTimeSending returns a boolean if a field has been set.
func (o *ShippingMethod) HasMaximumTimeSending() bool {
	if o != nil && !IsNil(o.MaximumTimeSending) {
		return true
	}

	return false
}

// SetMaximumTimeSending gets a reference to the given int32 and assigns it to the MaximumTimeSending field.
func (o *ShippingMethod) SetMaximumTimeSending(v int32) {
	o.MaximumTimeSending = &v
}

// GetDeliveryTimeDisplay returns the DeliveryTimeDisplay field value
func (o *ShippingMethod) GetDeliveryTimeDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliveryTimeDisplay
}

// GetDeliveryTimeDisplayOk returns a tuple with the DeliveryTimeDisplay field value
// and a boolean to check if the value has been set.
func (o *ShippingMethod) GetDeliveryTimeDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryTimeDisplay, true
}

// SetDeliveryTimeDisplay sets field value
func (o *ShippingMethod) SetDeliveryTimeDisplay(v string) {
	o.DeliveryTimeDisplay = v
}

// GetDeliveryTimeRangeDisplay returns the DeliveryTimeRangeDisplay field value
func (o *ShippingMethod) GetDeliveryTimeRangeDisplay() DeliveryTimeRangeDisplay {
	if o == nil {
		var ret DeliveryTimeRangeDisplay
		return ret
	}

	return o.DeliveryTimeRangeDisplay
}

// GetDeliveryTimeRangeDisplayOk returns a tuple with the DeliveryTimeRangeDisplay field value
// and a boolean to check if the value has been set.
func (o *ShippingMethod) GetDeliveryTimeRangeDisplayOk() (*DeliveryTimeRangeDisplay, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryTimeRangeDisplay, true
}

// SetDeliveryTimeRangeDisplay sets field value
func (o *ShippingMethod) SetDeliveryTimeRangeDisplay(v DeliveryTimeRangeDisplay) {
	o.DeliveryTimeRangeDisplay = v
}

// GetInventoryAddress returns the InventoryAddress field value
func (o *ShippingMethod) GetInventoryAddress() BusinessAddress {
	if o == nil {
		var ret BusinessAddress
		return ret
	}

	return o.InventoryAddress
}

// GetInventoryAddressOk returns a tuple with the InventoryAddress field value
// and a boolean to check if the value has been set.
func (o *ShippingMethod) GetInventoryAddressOk() (*BusinessAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InventoryAddress, true
}

// SetInventoryAddress sets field value
func (o *ShippingMethod) SetInventoryAddress(v BusinessAddress) {
	o.InventoryAddress = v
}

// GetIsPayAtDestination returns the IsPayAtDestination field value if set, zero value otherwise.
func (o *ShippingMethod) GetIsPayAtDestination() bool {
	if o == nil || IsNil(o.IsPayAtDestination) {
		var ret bool
		return ret
	}
	return *o.IsPayAtDestination
}

// GetIsPayAtDestinationOk returns a tuple with the IsPayAtDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethod) GetIsPayAtDestinationOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPayAtDestination) {
		return nil, false
	}
	return o.IsPayAtDestination, true
}

// HasIsPayAtDestination returns a boolean if a field has been set.
func (o *ShippingMethod) HasIsPayAtDestination() bool {
	if o != nil && !IsNil(o.IsPayAtDestination) {
		return true
	}

	return false
}

// SetIsPayAtDestination gets a reference to the given bool and assigns it to the IsPayAtDestination field.
func (o *ShippingMethod) SetIsPayAtDestination(v bool) {
	o.IsPayAtDestination = &v
}

func (o ShippingMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ShippingType) {
		toSerialize["shipping_type"] = o.ShippingType
	}
	toSerialize["get_shipping_type_display"] = o.GetShippingTypeDisplay
	toSerialize["shipping_type_display"] = o.ShippingTypeDisplay
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.SecondaryCost) {
		toSerialize["secondary_cost"] = o.SecondaryCost
	}
	if !IsNil(o.MinimumTimeSending) {
		toSerialize["minimum_time_sending"] = o.MinimumTimeSending
	}
	if !IsNil(o.MaximumTimeSending) {
		toSerialize["maximum_time_sending"] = o.MaximumTimeSending
	}
	toSerialize["delivery_time_display"] = o.DeliveryTimeDisplay
	toSerialize["delivery_time_range_display"] = o.DeliveryTimeRangeDisplay
	toSerialize["inventory_address"] = o.InventoryAddress
	if !IsNil(o.IsPayAtDestination) {
		toSerialize["is_pay_at_destination"] = o.IsPayAtDestination
	}
	return toSerialize, nil
}

func (o *ShippingMethod) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"get_shipping_type_display",
		"shipping_type_display",
		"delivery_time_display",
		"delivery_time_range_display",
		"inventory_address",
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

	varShippingMethod := _ShippingMethod{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShippingMethod)

	if err != nil {
		return err
	}

	*o = ShippingMethod(varShippingMethod)

	return err
}

type NullableShippingMethod struct {
	value *ShippingMethod
	isSet bool
}

func (v NullableShippingMethod) Get() *ShippingMethod {
	return v.value
}

func (v *NullableShippingMethod) Set(val *ShippingMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethod(val *ShippingMethod) *NullableShippingMethod {
	return &NullableShippingMethod{value: val, isSet: true}
}

func (v NullableShippingMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


