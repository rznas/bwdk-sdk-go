/*
BWDK API

<div dir=\"rtl\" style=\"text-align: right;\">  # مستندات فروشندگان در سرویس خرید با دیجی‌کالا  این پلتفرم برای فروشندگان (مرچنت‌ها) جهت یکپارچه‌سازی خدمات پرداخت و تجارت الکترونیکی با سیستم خرید با دیجی‌کالا. شامل مدیریت سفارشات، ارسال، و احراز هویت فروشندگان است.     ```mermaid flowchart TD     START([شروع]) --> INITIAL      INITIAL[\"1️⃣ INITIAL\\nسفارش ایجاد شد\"]     STARTED[\"2️⃣ STARTED\\nمشتری به BWDK هدایت شد\"]     PENDING[\"3️⃣ PENDING\\nمشتری وارد شد و سفارش در انتظار پرداخت\"]     WAITING_FOR_GATEWAY[\"4️⃣ WAITING_FOR_GATEWAY\\nمشتری به درگاه پرداخت هدایت شد\"]     PAID_BY_USER[\"7️⃣ PAID_BY_USER\\nپرداخت موفق\"]     VERIFIED_BY_MERCHANT[\"9️⃣ VERIFIED_BY_MERCHANT\\nتأیید شده توسط فروشنده\"]     SHIPPED[\"🚚 SHIPPED\\nارسال شد\"]     DELIVERED[\"✅ DELIVERED\\nتحویل داده شد\"]      EXPIRED[\"⏰ EXPIRED\\nمنقضی شد\"]     EXPIRATION_TIME_EXCEEDED[\"⏱️ EXPIRATION_TIME_EXCEEDED\\nزمان انقضا گذشت\"]     CANCELLED[\"❌ CANCELLED\\nلغو توسط مشتری\"]     FAILED_TO_PAY[\"💳 FAILED_TO_PAY\\nپرداخت ناموفق\"]     FAILED_TO_VERIFY_BY_MERCHANT[\"🔴 FAILED_TO_VERIFY_BY_MERCHANT\\nتأیید مرچنت ناموفق\"]     FAILED_BY_MERCHANT[\"🔴 FAILED_BY_MERCHANT\\nخطا از سمت مرچنت\"]     CANCELLED_BY_MERCHANT[\"🔴 CANCELLED_BY_MERCHANT\\nلغو توسط مرچنت\"]      R_CUSTOMER_REQUEST[\"1️⃣3️⃣ REQUEST_TO_REFUND\\nدرخواست استرداد توسط مشتری\"]     R_FAILED_VERIFY[\"1️⃣4️⃣ REQUEST_TO_REFUND\\nاسترداد پس از تأیید ناموفق مرچنت\"]     R_FAILED_MERCHANT[\"1️⃣5️⃣ REQUEST_TO_REFUND\\nاسترداد پس از خطای مرچنت\"]     R_CANCELLED_MERCHANT[\"1️⃣6️⃣ REQUEST_TO_REFUND\\nاسترداد پس از لغو مرچنت\"]     REFUND_COMPLETED[\"✅ REFUND_COMPLETED\\nاسترداد تکمیل شد\"]      INITIAL -->|\"مرچنت سفارش ایجاد کرد\"| STARTED     STARTED -->|\"مشتری وارد سیستم شد\"| PENDING     PENDING -->|\"مشتری سفارش را نهایی و ثبت کرد\"| WAITING_FOR_GATEWAY     WAITING_FOR_GATEWAY -->|\"پرداخت با موفقیت انجام شد\"| PAID_BY_USER     PAID_BY_USER -->|\"مرچنت سفارش را تأیید کرد\"| VERIFIED_BY_MERCHANT     VERIFIED_BY_MERCHANT -->|\"مرچنت وضعیت را به ارسال تغییر داد\"| SHIPPED     SHIPPED -->|\"مرچنت تحویل را تأیید کرد\"| DELIVERED      INITIAL -->|\"زمان رزرو به پایان رسید\"| EXPIRED     STARTED -->|\"زمان رزرو به پایان رسید\"| EXPIRED     PENDING -->|\"زمان رزرو به پایان رسید\"| EXPIRED     WAITING_FOR_GATEWAY -->|\"زمان رزرو به پایان رسید\"| EXPIRED      PENDING -->|\"زمان مجاز سفارش سپری شده بود\"| EXPIRATION_TIME_EXCEEDED     WAITING_FOR_GATEWAY -->|\"زمان مجاز سفارش سپری شده بود\"| EXPIRATION_TIME_EXCEEDED      PENDING -->|\"مشتری انصراف داد\"| CANCELLED     WAITING_FOR_GATEWAY -->|\"مشتری انصراف داد\"| CANCELLED      WAITING_FOR_GATEWAY -->|\"پرداخت ناموفق بود\"| FAILED_TO_PAY      PAID_BY_USER -->|\"مرچنت تأیید را رد کرد\"| FAILED_TO_VERIFY_BY_MERCHANT     PAID_BY_USER -->|\"مرچنت اعلام ناتوانی در انجام سفارش کرد\"| FAILED_BY_MERCHANT     PAID_BY_USER -->|\"مرچنت سفارش را لغو کرد\"| CANCELLED_BY_MERCHANT     VERIFIED_BY_MERCHANT -->|\"مرچنت سفارش را لغو کرد\"| CANCELLED_BY_MERCHANT      PAID_BY_USER -->|\"مرچنت درخواست استرداد داد\"| R_CUSTOMER_REQUEST     VERIFIED_BY_MERCHANT -->|\"مرچنت درخواست استرداد داد\"| R_CUSTOMER_REQUEST     FAILED_TO_VERIFY_BY_MERCHANT -->|\"سیستم استرداد را آغاز کرد\"| R_FAILED_VERIFY     FAILED_BY_MERCHANT -->|\"سیستم استرداد را آغاز کرد\"| R_FAILED_MERCHANT     CANCELLED_BY_MERCHANT -->|\"سیستم استرداد را آغاز کرد\"| R_CANCELLED_MERCHANT      R_CUSTOMER_REQUEST -->|\"استرداد توسط دیجی‌پی تأیید شد\"| REFUND_COMPLETED     R_FAILED_VERIFY -->|\"استرداد توسط دیجی‌پی تأیید شد\"| REFUND_COMPLETED     R_FAILED_MERCHANT -->|\"استرداد توسط دیجی‌پی تأیید شد\"| REFUND_COMPLETED     R_CANCELLED_MERCHANT -->|\"استرداد توسط دیجی‌پی تأیید شد\"| REFUND_COMPLETED      style INITIAL fill:#9e9e9e,color:#fff     style STARTED fill:#1565c0,color:#fff     style PENDING fill:#ef6c00,color:#fff     style WAITING_FOR_GATEWAY fill:#6a1b9a,color:#fff     style PAID_BY_USER fill:#2e7d32,color:#fff     style VERIFIED_BY_MERCHANT fill:#1b5e20,color:#fff     style SHIPPED fill:#0277bd,color:#fff     style DELIVERED fill:#1b5e20,color:#fff     style EXPIRED fill:#b71c1c,color:#fff     style EXPIRATION_TIME_EXCEEDED fill:#b71c1c,color:#fff     style CANCELLED fill:#7f0000,color:#fff     style FAILED_TO_PAY fill:#b71c1c,color:#fff     style FAILED_TO_VERIFY_BY_MERCHANT fill:#b71c1c,color:#fff     style FAILED_BY_MERCHANT fill:#b71c1c,color:#fff     style CANCELLED_BY_MERCHANT fill:#7f0000,color:#fff     style R_CUSTOMER_REQUEST fill:#e65100,color:#fff     style R_FAILED_VERIFY fill:#e65100,color:#fff     style R_FAILED_MERCHANT fill:#e65100,color:#fff     style R_CANCELLED_MERCHANT fill:#e65100,color:#fff     style REFUND_COMPLETED fill:#2e7d32,color:#fff ```  ---  <div dir=\"rtl\" style=\"text-align: right;\">  ## توضیح وضعیت‌های سفارش  ### ۱. INITIAL — ایجاد اولیه سفارش  **معنا:** سفارش توسط بک‌اند مرچنت ساخته شده ولی هنوز هیچ کاربری به آن اختصاص داده نشده است.  **چگونه اتفاق می‌افتد:** مرچنت با ارسال درخواست `POST /api/v1/orders/create` و ارائه اطلاعات کالاها، مبلغ و `callback_url`، یک سفارش جدید ایجاد می‌کند. BWDK یک `order_uuid` منحصربه‌فرد و لینک شروع سفارش (`order_start_url`) برمی‌گرداند.  **وابستگی‌ها:** نیازی به کاربر یا پرداخت ندارد. فقط اطلاعات کالا از سمت مرچنت کافی است.  ---  ### ۲. STARTED — آغاز جریان خرید  **معنا:** مشتری روی لینک شروع سفارش کلیک کرده و وارد محیط BWDK شده است، اما هنوز لاگین نکرده.  **چگونه اتفاق می‌افتد:** وقتی مشتری به `order_start_url` هدایت می‌شود، BWDK وضعیت سفارش را از `INITIAL` به `STARTED` تغییر می‌دهد. در این مرحله فرآیند احراز هویت (SSO) آغاز می‌شود.  **وابستگی‌ها:** مشتری باید به لینک شروع هدایت شده باشد.  ---  ### ۳. PENDING — انتظار برای تکمیل سفارش  **معنا:** مشتری با موفقیت وارد سیستم شده و سفارش به حساب او اختصاص یافته. مشتری در حال انتخاب آدرس، روش ارسال، بسته‌بندی یا تخفیف است.  **چگونه اتفاق می‌افتد:** پس از تکمیل ورود به سیستم (SSO)، BWDK سفارش را به کاربر وصل کرده و وضعیت را به `PENDING` تغییر می‌دهد.  **وابستگی‌ها:** ورود موفق کاربر به سیستم (SSO). در این مرحله مشتری می‌تواند آدرس، شیپینگ، پکینگ و تخفیف را انتخاب کند.  ---  ### ۴. WAITING_FOR_GATEWAY — انتظار برای پرداخت  **معنا:** مشتری اطلاعات سفارش را تأیید کرده و به درگاه پرداخت هدایت شده است.  **چگونه اتفاق می‌افتد:** مشتری دکمه «پرداخت» را می‌زند (`POST /api/v1/orders/submit`)، سیستم یک رکورد پرداخت ایجاد می‌کند و کاربر به درگاه Digipay هدایت می‌شود. وضعیت سفارش به `WAITING_FOR_GATEWAY` تغییر می‌کند.  **وابستگی‌ها:** انتخاب آدرس، روش ارسال و بسته‌بندی الزامی است. پرداخت باید ایجاد شده باشد.  ---  ### ۷. PAID_BY_USER — پرداخت موفق  **معنا:** تراکنش پرداخت با موفقیت انجام شده و وجه از حساب مشتری کسر شده است.  **چگونه اتفاق می‌افتد:** درگاه پرداخت نتیجه موفق را به BWDK اطلاع می‌دهد. سیستم پرداخت را تأیید و وضعیت سفارش را به `PAID_BY_USER` تغییر می‌دهد. در این لحظه مشتری به `callback_url` مرچنت هدایت می‌شود.  **وابستگی‌ها:** تأیید موفق تراکنش از سوی درگاه پرداخت (Digipay).  ---  ### ۹. VERIFIED_BY_MERCHANT — تأیید توسط مرچنت  **معنا:** مرچنت سفارش را بررسی کرده و موجودی کالا و صحت اطلاعات را تأیید نموده است. سفارش آماده ارسال است.  **چگونه اتفاق می‌افتد:** مرچنت با ارسال درخواست `POST /api/v1/orders/manager/{uuid}/verify` سفارش را تأیید می‌کند. این مرحله **اجباری** است و باید پس از پرداخت موفق انجام شود.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` باشد. مرچنت باید موجودی کالا را بررسی کند.  ---  ### ۲۰. SHIPPED — ارسال شد  **معنا:** سفارش از انبار خارج شده و در حال ارسال به مشتری است.  **چگونه اتفاق می‌افتد:** مرچنت پس از ارسال کالا، وضعیت سفارش را از طریق API به `SHIPPED` تغییر می‌دهد.  **وابستگی‌ها:** سفارش باید در وضعیت `VERIFIED_BY_MERCHANT` باشد.  ---  ### ۱۹. DELIVERED — تحویل داده شد  **معنا:** سفارش به دست مشتری رسیده و فرآیند خرید به پایان رسیده است.  **چگونه اتفاق می‌افتد:** مرچنت پس از تحویل موفق، وضعیت را به `DELIVERED` تغییر می‌دهد.  **وابستگی‌ها:** سفارش باید در وضعیت `SHIPPED` باشد.  ---  ### ۵. EXPIRED — منقضی شد  **معنا:** زمان رزرو سفارش به پایان رسیده و سفارش به صورت خودکار لغو شده است.  **چگونه اتفاق می‌افتد:** یک Task دوره‌ای به طور خودکار سفارش‌هایی که `reservation_expired_at` آن‌ها گذشته را پیدا کرده و وضعیتشان را به `EXPIRED` تغییر می‌دهد. این مکانیزم مانع بلوکه شدن موجودی کالا می‌شود.  **وابستگی‌ها:** سفارش باید در یکی از وضعیت‌های `INITIAL`، `STARTED`، `PENDING` یا `WAITING_FOR_GATEWAY` باشد و زمان رزرو آن گذشته باشد.  ---  ### ۱۸. EXPIRATION_TIME_EXCEEDED — زمان انقضا گذشت  **معنا:** در لحظه ثبت نهایی یا پرداخت، مشخص شد که زمان مجاز سفارش تمام شده است.  **چگونه اتفاق می‌افتد:** هنگام ارسال درخواست پرداخت (`submit_order`)، سیستم بررسی می‌کند که `expiration_time` سفارش هنوز معتبر است یا خیر. در صورت گذشتن زمان، وضعیت به `EXPIRATION_TIME_EXCEEDED` تغییر می‌کند.  **وابستگی‌ها:** سفارش در وضعیت `PENDING` یا `WAITING_FOR_GATEWAY` است و فیلد `expiration_time` سپری شده.  ---  ### ۶. CANCELLED — لغو توسط مشتری  **معنا:** مشتری در حین فرآیند خرید (قبل از پرداخت) سفارش را لغو کرده یا از صفحه خارج شده است.  **چگونه اتفاق می‌افتد:** مشتری در صفحه checkout دکمه «انصراف» را می‌زند یا پرداخت ناموفق بوده و سفارش به حالت لغو درمی‌آید.  **وابستگی‌ها:** سفارش باید در وضعیت `PENDING` یا `WAITING_FOR_GATEWAY` باشد. پرداختی انجام نشده است.  ---  ### ۸. FAILED_TO_PAY — پرداخت ناموفق  **معنا:** تراکنش پرداخت انجام نشد یا با خطا مواجه شد.  **چگونه اتفاق می‌افتد:** درگاه پرداخت نتیجه ناموفق برمی‌گرداند یا فرآیند بازگشت وجه در مرحله پرداخت با شکست مواجه می‌شود.  **وابستگی‌ها:** سفارش باید در وضعیت `WAITING_FOR_GATEWAY` بوده باشد.  ---  ### ۱۰. FAILED_TO_VERIFY_BY_MERCHANT — تأیید ناموفق توسط مرچنت  **معنا:** مرچنت سفارش را رد کرده است؛ معمولاً به دلیل ناموجود بودن کالا یا مغایرت اطلاعات.  **چگونه اتفاق می‌افتد:** مرچنت در پاسخ به درخواست verify، خطا برمی‌گرداند یا API آن وضعیت ناموفق تنظیم می‌کند. پس از این وضعیت، فرآیند استرداد وجه آغاز می‌شود.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` باشد.  ---  ### ۱۱. FAILED_BY_MERCHANT — خطا از سمت مرچنت  **معنا:** مرچنت پس از تأیید اولیه، اعلام می‌کند که قادر به انجام سفارش نیست (مثلاً به دلیل اتمام موجودی).  **چگونه اتفاق می‌افتد:** مرچنت وضعیت را به `FAILED_BY_MERCHANT` تغییر می‌دهد. وجه پرداختی مشتری مسترد خواهد شد.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` باشد.  ---  ### ۱۲. CANCELLED_BY_MERCHANT — لغو توسط مرچنت  **معنا:** مرچنت پس از پرداخت، سفارش را به هر دلیلی لغو کرده است.  **چگونه اتفاق می‌افتد:** مرچنت درخواست لغو سفارش را ارسال می‌کند. وجه پرداختی مشتری به او بازگردانده می‌شود.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` یا `VERIFIED_BY_MERCHANT` باشد.  ---  ### ۱۳. REQUEST_TO_REFUND — درخواست استرداد توسط مشتری  **معنا:** مشتری درخواست بازگشت وجه داده و سیستم در حال پردازش استرداد است.  **چگونه اتفاق می‌افتد:** مرچنت از طریق API درخواست استرداد را ثبت می‌کند (`POST /api/v1/orders/manager/{uuid}/refund`). سفارش وارد صف پردازش استرداد می‌شود.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` یا `VERIFIED_BY_MERCHANT` باشد.  ---  ### ۱۴، ۱۵، ۱۶. سایر وضعیت‌های درخواست استرداد  این وضعیت‌ها بر اساس دلیل استرداد از هم تفکیک می‌شوند:  - **۱۴ — REQUEST_TO_REFUND_TO_MERCHANT_AFTER_FAILED_TO_VERIFY:** استرداد پس از ناموفق بودن تأیید مرچنت؛ وجه به حساب مرچنت بازمی‌گردد. - **۱۵ — REQUEST_TO_REFUND_TO_CUSTOMER_AFTER_FAILED_BY_MERCHANT:** استرداد پس از خطای مرچنت؛ وجه به مشتری بازمی‌گردد. - **۱۶ — REQUEST_TO_REFUND_TO_MERCHANT_AFTER_CANCELLED_BY_MERCHANT:** استرداد پس از لغو توسط مرچنت؛ وجه به حساب مرچنت برمی‌گردد.  **چگونه اتفاق می‌افتد:** به صورت خودکار پس از رسیدن به وضعیت‌های ناموفق/لغو مربوطه توسط سیستم تنظیم می‌شود.  ---  ### ۱۷. REFUND_COMPLETED — استرداد تکمیل شد  **معنا:** وجه با موفقیت به صاحب آن (مشتری یا مرچنت بسته به نوع استرداد) بازگردانده شده است.  **چگونه اتفاق می‌افتد:** Task پردازش استرداد (`process_order_refund`) پس از تأیید موفق بازگشت وجه از سوی Digipay، وضعیت سفارش را به `REFUND_COMPLETED` تغییر می‌دهد.  **وابستگی‌ها:** یکی از وضعیت‌های درخواست استرداد (۱۳، ۱۴، ۱۵ یا ۱۶) باید فعال باشد و Digipay تراکنش استرداد را تأیید کرده باشد.  </div> 

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

// checks if the BusinessAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BusinessAddress{}

// BusinessAddress Serializer for BusinessAddress model. Used for merchant and shipping addresses.
type BusinessAddress struct {
	Id int32 `json:"id"`
	Address string `json:"address"`
	PostalCode NullableString `json:"postal_code,omitempty"`
	CityName string `json:"city_name"`
	StateName string `json:"state_name"`
	DistrictId NullableInt32 `json:"district_id,omitempty"`
	DistrictName NullableString `json:"district_name,omitempty"`
	Longitude NullableFloat64 `json:"longitude,omitempty" validate:"regexp=^-?\\\\d{0,6}(?:\\\\.\\\\d{0,16})?$"`
	Latitude NullableFloat64 `json:"latitude,omitempty" validate:"regexp=^-?\\\\d{0,6}(?:\\\\.\\\\d{0,16})?$"`
	BuildingNumber NullableString `json:"building_number,omitempty"`
	Unit NullableString `json:"unit,omitempty"`
	ReceiverName NullableString `json:"receiver_name,omitempty"`
	ReceiverPhone NullableString `json:"receiver_phone,omitempty"`
	IsAccurate *bool `json:"is_accurate,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type _BusinessAddress BusinessAddress

// NewBusinessAddress instantiates a new BusinessAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessAddress(id int32, address string, cityName string, stateName string, createdAt time.Time, modifiedAt time.Time) *BusinessAddress {
	this := BusinessAddress{}
	this.Id = id
	this.Address = address
	this.CityName = cityName
	this.StateName = stateName
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewBusinessAddressWithDefaults instantiates a new BusinessAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessAddressWithDefaults() *BusinessAddress {
	this := BusinessAddress{}
	return &this
}

// GetId returns the Id field value
func (o *BusinessAddress) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BusinessAddress) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BusinessAddress) SetId(v int32) {
	o.Id = v
}

// GetAddress returns the Address field value
func (o *BusinessAddress) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *BusinessAddress) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *BusinessAddress) SetAddress(v string) {
	o.Address = v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BusinessAddress) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode.Get()) {
		var ret string
		return ret
	}
	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BusinessAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// HasPostalCode returns a boolean if a field has been set.
func (o *BusinessAddress) HasPostalCode() bool {
	if o != nil && o.PostalCode.IsSet() {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given NullableString and assigns it to the PostalCode field.
func (o *BusinessAddress) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}
// SetPostalCodeNil sets the value for PostalCode to be an explicit nil
func (o *BusinessAddress) SetPostalCodeNil() {
	o.PostalCode.Set(nil)
}

// UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
func (o *BusinessAddress) UnsetPostalCode() {
	o.PostalCode.Unset()
}

// GetCityName returns the CityName field value
func (o *BusinessAddress) GetCityName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CityName
}

// GetCityNameOk returns a tuple with the CityName field value
// and a boolean to check if the value has been set.
func (o *BusinessAddress) GetCityNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CityName, true
}

// SetCityName sets field value
func (o *BusinessAddress) SetCityName(v string) {
	o.CityName = v
}

// GetStateName returns the StateName field value
func (o *BusinessAddress) GetStateName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateName
}

// GetStateNameOk returns a tuple with the StateName field value
// and a boolean to check if the value has been set.
func (o *BusinessAddress) GetStateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateName, true
}

// SetStateName sets field value
func (o *BusinessAddress) SetStateName(v string) {
	o.StateName = v
}

// GetDistrictId returns the DistrictId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BusinessAddress) GetDistrictId() int32 {
	if o == nil || IsNil(o.DistrictId.Get()) {
		var ret int32
		return ret
	}
	return *o.DistrictId.Get()
}

// GetDistrictIdOk returns a tuple with the DistrictId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BusinessAddress) GetDistrictIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DistrictId.Get(), o.DistrictId.IsSet()
}

// HasDistrictId returns a boolean if a field has been set.
func (o *BusinessAddress) HasDistrictId() bool {
	if o != nil && o.DistrictId.IsSet() {
		return true
	}

	return false
}

// SetDistrictId gets a reference to the given NullableInt32 and assigns it to the DistrictId field.
func (o *BusinessAddress) SetDistrictId(v int32) {
	o.DistrictId.Set(&v)
}
// SetDistrictIdNil sets the value for DistrictId to be an explicit nil
func (o *BusinessAddress) SetDistrictIdNil() {
	o.DistrictId.Set(nil)
}

// UnsetDistrictId ensures that no value is present for DistrictId, not even an explicit nil
func (o *BusinessAddress) UnsetDistrictId() {
	o.DistrictId.Unset()
}

// GetDistrictName returns the DistrictName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BusinessAddress) GetDistrictName() string {
	if o == nil || IsNil(o.DistrictName.Get()) {
		var ret string
		return ret
	}
	return *o.DistrictName.Get()
}

// GetDistrictNameOk returns a tuple with the DistrictName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BusinessAddress) GetDistrictNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DistrictName.Get(), o.DistrictName.IsSet()
}

// HasDistrictName returns a boolean if a field has been set.
func (o *BusinessAddress) HasDistrictName() bool {
	if o != nil && o.DistrictName.IsSet() {
		return true
	}

	return false
}

// SetDistrictName gets a reference to the given NullableString and assigns it to the DistrictName field.
func (o *BusinessAddress) SetDistrictName(v string) {
	o.DistrictName.Set(&v)
}
// SetDistrictNameNil sets the value for DistrictName to be an explicit nil
func (o *BusinessAddress) SetDistrictNameNil() {
	o.DistrictName.Set(nil)
}

// UnsetDistrictName ensures that no value is present for DistrictName, not even an explicit nil
func (o *BusinessAddress) UnsetDistrictName() {
	o.DistrictName.Unset()
}

// GetLongitude returns the Longitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BusinessAddress) GetLongitude() float64 {
	if o == nil || IsNil(o.Longitude.Get()) {
		var ret float64
		return ret
	}
	return *o.Longitude.Get()
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BusinessAddress) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Longitude.Get(), o.Longitude.IsSet()
}

// HasLongitude returns a boolean if a field has been set.
func (o *BusinessAddress) HasLongitude() bool {
	if o != nil && o.Longitude.IsSet() {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given NullableFloat64 and assigns it to the Longitude field.
func (o *BusinessAddress) SetLongitude(v float64) {
	o.Longitude.Set(&v)
}
// SetLongitudeNil sets the value for Longitude to be an explicit nil
func (o *BusinessAddress) SetLongitudeNil() {
	o.Longitude.Set(nil)
}

// UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
func (o *BusinessAddress) UnsetLongitude() {
	o.Longitude.Unset()
}

// GetLatitude returns the Latitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BusinessAddress) GetLatitude() float64 {
	if o == nil || IsNil(o.Latitude.Get()) {
		var ret float64
		return ret
	}
	return *o.Latitude.Get()
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BusinessAddress) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Latitude.Get(), o.Latitude.IsSet()
}

// HasLatitude returns a boolean if a field has been set.
func (o *BusinessAddress) HasLatitude() bool {
	if o != nil && o.Latitude.IsSet() {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given NullableFloat64 and assigns it to the Latitude field.
func (o *BusinessAddress) SetLatitude(v float64) {
	o.Latitude.Set(&v)
}
// SetLatitudeNil sets the value for Latitude to be an explicit nil
func (o *BusinessAddress) SetLatitudeNil() {
	o.Latitude.Set(nil)
}

// UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
func (o *BusinessAddress) UnsetLatitude() {
	o.Latitude.Unset()
}

// GetBuildingNumber returns the BuildingNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BusinessAddress) GetBuildingNumber() string {
	if o == nil || IsNil(o.BuildingNumber.Get()) {
		var ret string
		return ret
	}
	return *o.BuildingNumber.Get()
}

// GetBuildingNumberOk returns a tuple with the BuildingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BusinessAddress) GetBuildingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildingNumber.Get(), o.BuildingNumber.IsSet()
}

// HasBuildingNumber returns a boolean if a field has been set.
func (o *BusinessAddress) HasBuildingNumber() bool {
	if o != nil && o.BuildingNumber.IsSet() {
		return true
	}

	return false
}

// SetBuildingNumber gets a reference to the given NullableString and assigns it to the BuildingNumber field.
func (o *BusinessAddress) SetBuildingNumber(v string) {
	o.BuildingNumber.Set(&v)
}
// SetBuildingNumberNil sets the value for BuildingNumber to be an explicit nil
func (o *BusinessAddress) SetBuildingNumberNil() {
	o.BuildingNumber.Set(nil)
}

// UnsetBuildingNumber ensures that no value is present for BuildingNumber, not even an explicit nil
func (o *BusinessAddress) UnsetBuildingNumber() {
	o.BuildingNumber.Unset()
}

// GetUnit returns the Unit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BusinessAddress) GetUnit() string {
	if o == nil || IsNil(o.Unit.Get()) {
		var ret string
		return ret
	}
	return *o.Unit.Get()
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BusinessAddress) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unit.Get(), o.Unit.IsSet()
}

// HasUnit returns a boolean if a field has been set.
func (o *BusinessAddress) HasUnit() bool {
	if o != nil && o.Unit.IsSet() {
		return true
	}

	return false
}

// SetUnit gets a reference to the given NullableString and assigns it to the Unit field.
func (o *BusinessAddress) SetUnit(v string) {
	o.Unit.Set(&v)
}
// SetUnitNil sets the value for Unit to be an explicit nil
func (o *BusinessAddress) SetUnitNil() {
	o.Unit.Set(nil)
}

// UnsetUnit ensures that no value is present for Unit, not even an explicit nil
func (o *BusinessAddress) UnsetUnit() {
	o.Unit.Unset()
}

// GetReceiverName returns the ReceiverName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BusinessAddress) GetReceiverName() string {
	if o == nil || IsNil(o.ReceiverName.Get()) {
		var ret string
		return ret
	}
	return *o.ReceiverName.Get()
}

// GetReceiverNameOk returns a tuple with the ReceiverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BusinessAddress) GetReceiverNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceiverName.Get(), o.ReceiverName.IsSet()
}

// HasReceiverName returns a boolean if a field has been set.
func (o *BusinessAddress) HasReceiverName() bool {
	if o != nil && o.ReceiverName.IsSet() {
		return true
	}

	return false
}

// SetReceiverName gets a reference to the given NullableString and assigns it to the ReceiverName field.
func (o *BusinessAddress) SetReceiverName(v string) {
	o.ReceiverName.Set(&v)
}
// SetReceiverNameNil sets the value for ReceiverName to be an explicit nil
func (o *BusinessAddress) SetReceiverNameNil() {
	o.ReceiverName.Set(nil)
}

// UnsetReceiverName ensures that no value is present for ReceiverName, not even an explicit nil
func (o *BusinessAddress) UnsetReceiverName() {
	o.ReceiverName.Unset()
}

// GetReceiverPhone returns the ReceiverPhone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BusinessAddress) GetReceiverPhone() string {
	if o == nil || IsNil(o.ReceiverPhone.Get()) {
		var ret string
		return ret
	}
	return *o.ReceiverPhone.Get()
}

// GetReceiverPhoneOk returns a tuple with the ReceiverPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BusinessAddress) GetReceiverPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceiverPhone.Get(), o.ReceiverPhone.IsSet()
}

// HasReceiverPhone returns a boolean if a field has been set.
func (o *BusinessAddress) HasReceiverPhone() bool {
	if o != nil && o.ReceiverPhone.IsSet() {
		return true
	}

	return false
}

// SetReceiverPhone gets a reference to the given NullableString and assigns it to the ReceiverPhone field.
func (o *BusinessAddress) SetReceiverPhone(v string) {
	o.ReceiverPhone.Set(&v)
}
// SetReceiverPhoneNil sets the value for ReceiverPhone to be an explicit nil
func (o *BusinessAddress) SetReceiverPhoneNil() {
	o.ReceiverPhone.Set(nil)
}

// UnsetReceiverPhone ensures that no value is present for ReceiverPhone, not even an explicit nil
func (o *BusinessAddress) UnsetReceiverPhone() {
	o.ReceiverPhone.Unset()
}

// GetIsAccurate returns the IsAccurate field value if set, zero value otherwise.
func (o *BusinessAddress) GetIsAccurate() bool {
	if o == nil || IsNil(o.IsAccurate) {
		var ret bool
		return ret
	}
	return *o.IsAccurate
}

// GetIsAccurateOk returns a tuple with the IsAccurate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessAddress) GetIsAccurateOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAccurate) {
		return nil, false
	}
	return o.IsAccurate, true
}

// HasIsAccurate returns a boolean if a field has been set.
func (o *BusinessAddress) HasIsAccurate() bool {
	if o != nil && !IsNil(o.IsAccurate) {
		return true
	}

	return false
}

// SetIsAccurate gets a reference to the given bool and assigns it to the IsAccurate field.
func (o *BusinessAddress) SetIsAccurate(v bool) {
	o.IsAccurate = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *BusinessAddress) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessAddress) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *BusinessAddress) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *BusinessAddress) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *BusinessAddress) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BusinessAddress) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BusinessAddress) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *BusinessAddress) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *BusinessAddress) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *BusinessAddress) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

func (o BusinessAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["address"] = o.Address
	if o.PostalCode.IsSet() {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	toSerialize["city_name"] = o.CityName
	toSerialize["state_name"] = o.StateName
	if o.DistrictId.IsSet() {
		toSerialize["district_id"] = o.DistrictId.Get()
	}
	if o.DistrictName.IsSet() {
		toSerialize["district_name"] = o.DistrictName.Get()
	}
	if o.Longitude.IsSet() {
		toSerialize["longitude"] = o.Longitude.Get()
	}
	if o.Latitude.IsSet() {
		toSerialize["latitude"] = o.Latitude.Get()
	}
	if o.BuildingNumber.IsSet() {
		toSerialize["building_number"] = o.BuildingNumber.Get()
	}
	if o.Unit.IsSet() {
		toSerialize["unit"] = o.Unit.Get()
	}
	if o.ReceiverName.IsSet() {
		toSerialize["receiver_name"] = o.ReceiverName.Get()
	}
	if o.ReceiverPhone.IsSet() {
		toSerialize["receiver_phone"] = o.ReceiverPhone.Get()
	}
	if !IsNil(o.IsAccurate) {
		toSerialize["is_accurate"] = o.IsAccurate
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["modified_at"] = o.ModifiedAt
	return toSerialize, nil
}

func (o *BusinessAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"address",
		"city_name",
		"state_name",
		"created_at",
		"modified_at",
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

	varBusinessAddress := _BusinessAddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBusinessAddress)

	if err != nil {
		return err
	}

	*o = BusinessAddress(varBusinessAddress)

	return err
}

type NullableBusinessAddress struct {
	value *BusinessAddress
	isSet bool
}

func (v NullableBusinessAddress) Get() *BusinessAddress {
	return v.value
}

func (v *NullableBusinessAddress) Set(val *BusinessAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessAddress(val *BusinessAddress) *NullableBusinessAddress {
	return &NullableBusinessAddress{value: val, isSet: true}
}

func (v NullableBusinessAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


