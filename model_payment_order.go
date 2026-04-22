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

// checks if the PaymentOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentOrder{}

// PaymentOrder struct for PaymentOrder
type PaymentOrder struct {
	FinalAmount int32 `json:"final_amount"`
	GatewayType GatewayTypeEnum `json:"gateway_type"`
	GatewayTypeDisplay string `json:"gateway_type_display"`
	PaidAt string `json:"paid_at"`
	GatewayName NullableString `json:"gateway_name"`
	InvoiceId NullableString `json:"invoice_id"`
	SettlementDate string `json:"settlement_date"`
	SettlementStatus int32 `json:"settlement_status"`
	SettlementStatusDisplay string `json:"settlement_status_display"`
}

type _PaymentOrder PaymentOrder

// NewPaymentOrder instantiates a new PaymentOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentOrder(finalAmount int32, gatewayType GatewayTypeEnum, gatewayTypeDisplay string, paidAt string, gatewayName NullableString, invoiceId NullableString, settlementDate string, settlementStatus int32, settlementStatusDisplay string) *PaymentOrder {
	this := PaymentOrder{}
	this.FinalAmount = finalAmount
	this.GatewayType = gatewayType
	this.GatewayTypeDisplay = gatewayTypeDisplay
	this.PaidAt = paidAt
	this.GatewayName = gatewayName
	this.InvoiceId = invoiceId
	this.SettlementDate = settlementDate
	this.SettlementStatus = settlementStatus
	this.SettlementStatusDisplay = settlementStatusDisplay
	return &this
}

// NewPaymentOrderWithDefaults instantiates a new PaymentOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentOrderWithDefaults() *PaymentOrder {
	this := PaymentOrder{}
	return &this
}

// GetFinalAmount returns the FinalAmount field value
func (o *PaymentOrder) GetFinalAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FinalAmount
}

// GetFinalAmountOk returns a tuple with the FinalAmount field value
// and a boolean to check if the value has been set.
func (o *PaymentOrder) GetFinalAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinalAmount, true
}

// SetFinalAmount sets field value
func (o *PaymentOrder) SetFinalAmount(v int32) {
	o.FinalAmount = v
}

// GetGatewayType returns the GatewayType field value
func (o *PaymentOrder) GetGatewayType() GatewayTypeEnum {
	if o == nil {
		var ret GatewayTypeEnum
		return ret
	}

	return o.GatewayType
}

// GetGatewayTypeOk returns a tuple with the GatewayType field value
// and a boolean to check if the value has been set.
func (o *PaymentOrder) GetGatewayTypeOk() (*GatewayTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayType, true
}

// SetGatewayType sets field value
func (o *PaymentOrder) SetGatewayType(v GatewayTypeEnum) {
	o.GatewayType = v
}

// GetGatewayTypeDisplay returns the GatewayTypeDisplay field value
func (o *PaymentOrder) GetGatewayTypeDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayTypeDisplay
}

// GetGatewayTypeDisplayOk returns a tuple with the GatewayTypeDisplay field value
// and a boolean to check if the value has been set.
func (o *PaymentOrder) GetGatewayTypeDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayTypeDisplay, true
}

// SetGatewayTypeDisplay sets field value
func (o *PaymentOrder) SetGatewayTypeDisplay(v string) {
	o.GatewayTypeDisplay = v
}

// GetPaidAt returns the PaidAt field value
func (o *PaymentOrder) GetPaidAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaidAt
}

// GetPaidAtOk returns a tuple with the PaidAt field value
// and a boolean to check if the value has been set.
func (o *PaymentOrder) GetPaidAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaidAt, true
}

// SetPaidAt sets field value
func (o *PaymentOrder) SetPaidAt(v string) {
	o.PaidAt = v
}

// GetGatewayName returns the GatewayName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentOrder) GetGatewayName() string {
	if o == nil || o.GatewayName.Get() == nil {
		var ret string
		return ret
	}

	return *o.GatewayName.Get()
}

// GetGatewayNameOk returns a tuple with the GatewayName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentOrder) GetGatewayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GatewayName.Get(), o.GatewayName.IsSet()
}

// SetGatewayName sets field value
func (o *PaymentOrder) SetGatewayName(v string) {
	o.GatewayName.Set(&v)
}

// GetInvoiceId returns the InvoiceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentOrder) GetInvoiceId() string {
	if o == nil || o.InvoiceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.InvoiceId.Get()
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentOrder) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceId.Get(), o.InvoiceId.IsSet()
}

// SetInvoiceId sets field value
func (o *PaymentOrder) SetInvoiceId(v string) {
	o.InvoiceId.Set(&v)
}

// GetSettlementDate returns the SettlementDate field value
func (o *PaymentOrder) GetSettlementDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SettlementDate
}

// GetSettlementDateOk returns a tuple with the SettlementDate field value
// and a boolean to check if the value has been set.
func (o *PaymentOrder) GetSettlementDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SettlementDate, true
}

// SetSettlementDate sets field value
func (o *PaymentOrder) SetSettlementDate(v string) {
	o.SettlementDate = v
}

// GetSettlementStatus returns the SettlementStatus field value
func (o *PaymentOrder) GetSettlementStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SettlementStatus
}

// GetSettlementStatusOk returns a tuple with the SettlementStatus field value
// and a boolean to check if the value has been set.
func (o *PaymentOrder) GetSettlementStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SettlementStatus, true
}

// SetSettlementStatus sets field value
func (o *PaymentOrder) SetSettlementStatus(v int32) {
	o.SettlementStatus = v
}

// GetSettlementStatusDisplay returns the SettlementStatusDisplay field value
func (o *PaymentOrder) GetSettlementStatusDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SettlementStatusDisplay
}

// GetSettlementStatusDisplayOk returns a tuple with the SettlementStatusDisplay field value
// and a boolean to check if the value has been set.
func (o *PaymentOrder) GetSettlementStatusDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SettlementStatusDisplay, true
}

// SetSettlementStatusDisplay sets field value
func (o *PaymentOrder) SetSettlementStatusDisplay(v string) {
	o.SettlementStatusDisplay = v
}

func (o PaymentOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["final_amount"] = o.FinalAmount
	toSerialize["gateway_type"] = o.GatewayType
	toSerialize["gateway_type_display"] = o.GatewayTypeDisplay
	toSerialize["paid_at"] = o.PaidAt
	toSerialize["gateway_name"] = o.GatewayName.Get()
	toSerialize["invoice_id"] = o.InvoiceId.Get()
	toSerialize["settlement_date"] = o.SettlementDate
	toSerialize["settlement_status"] = o.SettlementStatus
	toSerialize["settlement_status_display"] = o.SettlementStatusDisplay
	return toSerialize, nil
}

func (o *PaymentOrder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"final_amount",
		"gateway_type",
		"gateway_type_display",
		"paid_at",
		"gateway_name",
		"invoice_id",
		"settlement_date",
		"settlement_status",
		"settlement_status_display",
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

	varPaymentOrder := _PaymentOrder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymentOrder)

	if err != nil {
		return err
	}

	*o = PaymentOrder(varPaymentOrder)

	return err
}

type NullablePaymentOrder struct {
	value *PaymentOrder
	isSet bool
}

func (v NullablePaymentOrder) Get() *PaymentOrder {
	return v.value
}

func (v *NullablePaymentOrder) Set(val *PaymentOrder) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentOrder) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentOrder(val *PaymentOrder) *NullablePaymentOrder {
	return &NullablePaymentOrder{value: val, isSet: true}
}

func (v NullablePaymentOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


