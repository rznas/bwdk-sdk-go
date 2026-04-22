/*
BWDK API

<div dir=\"rtl\" style=\"text-align: right;\">  # مستندات فروشندگان در سرویس خرید با دیجی‌کالا  این پلتفرم برای فروشندگان (مرچنت‌ها) جهت یکپارچه‌سازی خدمات پرداخت و تجارت الکترونیکی با سیستم خرید با دیجی‌کالا. شامل مدیریت سفارشات، ارسال، و احراز هویت فروشندگان است.     ```mermaid flowchart TD     START([شروع]) --> INITIAL      INITIAL[\"1️⃣ INITIAL\\nسفارش ایجاد شد\"]     STARTED[\"2️⃣ STARTED\\nمشتری به BWDK هدایت شد\"]     PENDING[\"3️⃣ PENDING\\nمشتری وارد شد و سفارش در انتظار پرداخت\"]     WAITING_FOR_GATEWAY[\"4️⃣ WAITING_FOR_GATEWAY\\nمشتری به درگاه پرداخت هدایت شد\"]     PAID_BY_USER[\"7️⃣ PAID_BY_USER\\nپرداخت موفق\"]     VERIFIED_BY_MERCHANT[\"9️⃣ VERIFIED_BY_MERCHANT\\nتأیید شده توسط فروشنده\"]     SHIPPED[\"🚚 SHIPPED\\nارسال شد\"]     DELIVERED[\"✅ DELIVERED\\nتحویل داده شد\"]      EXPIRED[\"⏰ EXPIRED\\nمنقضی شد\"]     EXPIRATION_TIME_EXCEEDED[\"⏱️ EXPIRATION_TIME_EXCEEDED\\nزمان انقضا گذشت\"]     CANCELLED[\"❌ CANCELLED\\nلغو توسط مشتری\"]     FAILED_TO_PAY[\"💳 FAILED_TO_PAY\\nپرداخت ناموفق\"]     FAILED_TO_VERIFY_BY_MERCHANT[\"🔴 FAILED_TO_VERIFY_BY_MERCHANT\\nتأیید مرچنت ناموفق\"]     FAILED_BY_MERCHANT[\"🔴 FAILED_BY_MERCHANT\\nخطا از سمت مرچنت\"]     CANCELLED_BY_MERCHANT[\"🔴 CANCELLED_BY_MERCHANT\\nلغو توسط مرچنت\"]      R_CUSTOMER_REQUEST[\"1️⃣3️⃣ REQUEST_TO_REFUND\\nدرخواست استرداد توسط مشتری\"]     R_FAILED_VERIFY[\"1️⃣4️⃣ REQUEST_TO_REFUND\\nاسترداد پس از تأیید ناموفق مرچنت\"]     R_FAILED_MERCHANT[\"1️⃣5️⃣ REQUEST_TO_REFUND\\nاسترداد پس از خطای مرچنت\"]     R_CANCELLED_MERCHANT[\"1️⃣6️⃣ REQUEST_TO_REFUND\\nاسترداد پس از لغو مرچنت\"]     REFUND_COMPLETED[\"✅ REFUND_COMPLETED\\nاسترداد تکمیل شد\"]      INITIAL -->|\"مرچنت سفارش ایجاد کرد\"| STARTED     STARTED -->|\"مشتری وارد سیستم شد\"| PENDING     PENDING -->|\"مشتری سفارش را نهایی و ثبت کرد\"| WAITING_FOR_GATEWAY     WAITING_FOR_GATEWAY -->|\"پرداخت با موفقیت انجام شد\"| PAID_BY_USER     PAID_BY_USER -->|\"مرچنت سفارش را تأیید کرد\"| VERIFIED_BY_MERCHANT     VERIFIED_BY_MERCHANT -->|\"مرچنت وضعیت را به ارسال تغییر داد\"| SHIPPED     SHIPPED -->|\"مرچنت تحویل را تأیید کرد\"| DELIVERED      INITIAL -->|\"زمان رزرو به پایان رسید\"| EXPIRED     STARTED -->|\"زمان رزرو به پایان رسید\"| EXPIRED     PENDING -->|\"زمان رزرو به پایان رسید\"| EXPIRED     WAITING_FOR_GATEWAY -->|\"زمان رزرو به پایان رسید\"| EXPIRED      PENDING -->|\"زمان مجاز سفارش سپری شده بود\"| EXPIRATION_TIME_EXCEEDED     WAITING_FOR_GATEWAY -->|\"زمان مجاز سفارش سپری شده بود\"| EXPIRATION_TIME_EXCEEDED      PENDING -->|\"مشتری انصراف داد\"| CANCELLED     WAITING_FOR_GATEWAY -->|\"مشتری انصراف داد\"| CANCELLED      WAITING_FOR_GATEWAY -->|\"پرداخت ناموفق بود\"| FAILED_TO_PAY      PAID_BY_USER -->|\"مرچنت تأیید را رد کرد\"| FAILED_TO_VERIFY_BY_MERCHANT     PAID_BY_USER -->|\"مرچنت اعلام ناتوانی در انجام سفارش کرد\"| FAILED_BY_MERCHANT     PAID_BY_USER -->|\"مرچنت سفارش را لغو کرد\"| CANCELLED_BY_MERCHANT     VERIFIED_BY_MERCHANT -->|\"مرچنت سفارش را لغو کرد\"| CANCELLED_BY_MERCHANT      PAID_BY_USER -->|\"مرچنت درخواست استرداد داد\"| R_CUSTOMER_REQUEST     VERIFIED_BY_MERCHANT -->|\"مرچنت درخواست استرداد داد\"| R_CUSTOMER_REQUEST     FAILED_TO_VERIFY_BY_MERCHANT -->|\"سیستم استرداد را آغاز کرد\"| R_FAILED_VERIFY     FAILED_BY_MERCHANT -->|\"سیستم استرداد را آغاز کرد\"| R_FAILED_MERCHANT     CANCELLED_BY_MERCHANT -->|\"سیستم استرداد را آغاز کرد\"| R_CANCELLED_MERCHANT      R_CUSTOMER_REQUEST -->|\"استرداد توسط دیجی‌پی تأیید شد\"| REFUND_COMPLETED     R_FAILED_VERIFY -->|\"استرداد توسط دیجی‌پی تأیید شد\"| REFUND_COMPLETED     R_FAILED_MERCHANT -->|\"استرداد توسط دیجی‌پی تأیید شد\"| REFUND_COMPLETED     R_CANCELLED_MERCHANT -->|\"استرداد توسط دیجی‌پی تأیید شد\"| REFUND_COMPLETED      style INITIAL fill:#9e9e9e,color:#fff     style STARTED fill:#1565c0,color:#fff     style PENDING fill:#ef6c00,color:#fff     style WAITING_FOR_GATEWAY fill:#6a1b9a,color:#fff     style PAID_BY_USER fill:#2e7d32,color:#fff     style VERIFIED_BY_MERCHANT fill:#1b5e20,color:#fff     style SHIPPED fill:#0277bd,color:#fff     style DELIVERED fill:#1b5e20,color:#fff     style EXPIRED fill:#b71c1c,color:#fff     style EXPIRATION_TIME_EXCEEDED fill:#b71c1c,color:#fff     style CANCELLED fill:#7f0000,color:#fff     style FAILED_TO_PAY fill:#b71c1c,color:#fff     style FAILED_TO_VERIFY_BY_MERCHANT fill:#b71c1c,color:#fff     style FAILED_BY_MERCHANT fill:#b71c1c,color:#fff     style CANCELLED_BY_MERCHANT fill:#7f0000,color:#fff     style R_CUSTOMER_REQUEST fill:#e65100,color:#fff     style R_FAILED_VERIFY fill:#e65100,color:#fff     style R_FAILED_MERCHANT fill:#e65100,color:#fff     style R_CANCELLED_MERCHANT fill:#e65100,color:#fff     style REFUND_COMPLETED fill:#2e7d32,color:#fff ```  ---  <div dir=\"rtl\" style=\"text-align: right;\">  ## توضیح وضعیت‌های سفارش  ### ۱. INITIAL — ایجاد اولیه سفارش  **معنا:** سفارش توسط بک‌اند مرچنت ساخته شده ولی هنوز هیچ کاربری به آن اختصاص داده نشده است.  **چگونه اتفاق می‌افتد:** مرچنت با ارسال درخواست `POST /api/v1/orders/create` و ارائه اطلاعات کالاها، مبلغ و `callback_url`، یک سفارش جدید ایجاد می‌کند. BWDK یک `order_uuid` منحصربه‌فرد و لینک شروع سفارش (`order_start_url`) برمی‌گرداند.  **وابستگی‌ها:** نیازی به کاربر یا پرداخت ندارد. فقط اطلاعات کالا از سمت مرچنت کافی است.  ---  ### ۲. STARTED — آغاز جریان خرید  **معنا:** مشتری روی لینک شروع سفارش کلیک کرده و وارد محیط BWDK شده است، اما هنوز لاگین نکرده.  **چگونه اتفاق می‌افتد:** وقتی مشتری به `order_start_url` هدایت می‌شود، BWDK وضعیت سفارش را از `INITIAL` به `STARTED` تغییر می‌دهد. در این مرحله فرآیند احراز هویت (SSO) آغاز می‌شود.  **وابستگی‌ها:** مشتری باید به لینک شروع هدایت شده باشد.  ---  ### ۳. PENDING — انتظار برای تکمیل سفارش  **معنا:** مشتری با موفقیت وارد سیستم شده و سفارش به حساب او اختصاص یافته. مشتری در حال انتخاب آدرس، روش ارسال، بسته‌بندی یا تخفیف است.  **چگونه اتفاق می‌افتد:** پس از تکمیل ورود به سیستم (SSO)، BWDK سفارش را به کاربر وصل کرده و وضعیت را به `PENDING` تغییر می‌دهد.  **وابستگی‌ها:** ورود موفق کاربر به سیستم (SSO). در این مرحله مشتری می‌تواند آدرس، شیپینگ، پکینگ و تخفیف را انتخاب کند.  ---  ### ۴. WAITING_FOR_GATEWAY — انتظار برای پرداخت  **معنا:** مشتری اطلاعات سفارش را تأیید کرده و به درگاه پرداخت هدایت شده است.  **چگونه اتفاق می‌افتد:** مشتری دکمه «پرداخت» را می‌زند (`POST /api/v1/orders/submit`)، سیستم یک رکورد پرداخت ایجاد می‌کند و کاربر به درگاه Digipay هدایت می‌شود. وضعیت سفارش به `WAITING_FOR_GATEWAY` تغییر می‌کند.  **وابستگی‌ها:** انتخاب آدرس، روش ارسال و بسته‌بندی الزامی است. پرداخت باید ایجاد شده باشد.  ---  ### ۷. PAID_BY_USER — پرداخت موفق  **معنا:** تراکنش پرداخت با موفقیت انجام شده و وجه از حساب مشتری کسر شده است.  **چگونه اتفاق می‌افتد:** درگاه پرداخت نتیجه موفق را به BWDK اطلاع می‌دهد. سیستم پرداخت را تأیید و وضعیت سفارش را به `PAID_BY_USER` تغییر می‌دهد. در این لحظه مشتری به `callback_url` مرچنت هدایت می‌شود.  **وابستگی‌ها:** تأیید موفق تراکنش از سوی درگاه پرداخت (Digipay).  ---  ### ۹. VERIFIED_BY_MERCHANT — تأیید توسط مرچنت  **معنا:** مرچنت سفارش را بررسی کرده و موجودی کالا و صحت اطلاعات را تأیید نموده است. سفارش آماده ارسال است.  **چگونه اتفاق می‌افتد:** مرچنت با ارسال درخواست `POST /api/v1/orders/manager/{uuid}/verify` سفارش را تأیید می‌کند. این مرحله **اجباری** است و باید پس از پرداخت موفق انجام شود.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` باشد. مرچنت باید موجودی کالا را بررسی کند.  ---  ### ۲۰. SHIPPED — ارسال شد  **معنا:** سفارش از انبار خارج شده و در حال ارسال به مشتری است.  **چگونه اتفاق می‌افتد:** مرچنت پس از ارسال کالا، وضعیت سفارش را از طریق API به `SHIPPED` تغییر می‌دهد.  **وابستگی‌ها:** سفارش باید در وضعیت `VERIFIED_BY_MERCHANT` باشد.  ---  ### ۱۹. DELIVERED — تحویل داده شد  **معنا:** سفارش به دست مشتری رسیده و فرآیند خرید به پایان رسیده است.  **چگونه اتفاق می‌افتد:** مرچنت پس از تحویل موفق، وضعیت را به `DELIVERED` تغییر می‌دهد.  **وابستگی‌ها:** سفارش باید در وضعیت `SHIPPED` باشد.  ---  ### ۵. EXPIRED — منقضی شد  **معنا:** زمان رزرو سفارش به پایان رسیده و سفارش به صورت خودکار لغو شده است.  **چگونه اتفاق می‌افتد:** یک Task دوره‌ای به طور خودکار سفارش‌هایی که `reservation_expired_at` آن‌ها گذشته را پیدا کرده و وضعیتشان را به `EXPIRED` تغییر می‌دهد. این مکانیزم مانع بلوکه شدن موجودی کالا می‌شود.  **وابستگی‌ها:** سفارش باید در یکی از وضعیت‌های `INITIAL`، `STARTED`، `PENDING` یا `WAITING_FOR_GATEWAY` باشد و زمان رزرو آن گذشته باشد.  ---  ### ۱۸. EXPIRATION_TIME_EXCEEDED — زمان انقضا گذشت  **معنا:** در لحظه ثبت نهایی یا پرداخت، مشخص شد که زمان مجاز سفارش تمام شده است.  **چگونه اتفاق می‌افتد:** هنگام ارسال درخواست پرداخت (`submit_order`)، سیستم بررسی می‌کند که `expiration_time` سفارش هنوز معتبر است یا خیر. در صورت گذشتن زمان، وضعیت به `EXPIRATION_TIME_EXCEEDED` تغییر می‌کند.  **وابستگی‌ها:** سفارش در وضعیت `PENDING` یا `WAITING_FOR_GATEWAY` است و فیلد `expiration_time` سپری شده.  ---  ### ۶. CANCELLED — لغو توسط مشتری  **معنا:** مشتری در حین فرآیند خرید (قبل از پرداخت) سفارش را لغو کرده یا از صفحه خارج شده است.  **چگونه اتفاق می‌افتد:** مشتری در صفحه checkout دکمه «انصراف» را می‌زند یا پرداخت ناموفق بوده و سفارش به حالت لغو درمی‌آید.  **وابستگی‌ها:** سفارش باید در وضعیت `PENDING` یا `WAITING_FOR_GATEWAY` باشد. پرداختی انجام نشده است.  ---  ### ۸. FAILED_TO_PAY — پرداخت ناموفق  **معنا:** تراکنش پرداخت انجام نشد یا با خطا مواجه شد.  **چگونه اتفاق می‌افتد:** درگاه پرداخت نتیجه ناموفق برمی‌گرداند یا فرآیند بازگشت وجه در مرحله پرداخت با شکست مواجه می‌شود.  **وابستگی‌ها:** سفارش باید در وضعیت `WAITING_FOR_GATEWAY` بوده باشد.  ---  ### ۱۰. FAILED_TO_VERIFY_BY_MERCHANT — تأیید ناموفق توسط مرچنت  **معنا:** مرچنت سفارش را رد کرده است؛ معمولاً به دلیل ناموجود بودن کالا یا مغایرت اطلاعات.  **چگونه اتفاق می‌افتد:** مرچنت در پاسخ به درخواست verify، خطا برمی‌گرداند یا API آن وضعیت ناموفق تنظیم می‌کند. پس از این وضعیت، فرآیند استرداد وجه آغاز می‌شود.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` باشد.  ---  ### ۱۱. FAILED_BY_MERCHANT — خطا از سمت مرچنت  **معنا:** مرچنت پس از تأیید اولیه، اعلام می‌کند که قادر به انجام سفارش نیست (مثلاً به دلیل اتمام موجودی).  **چگونه اتفاق می‌افتد:** مرچنت وضعیت را به `FAILED_BY_MERCHANT` تغییر می‌دهد. وجه پرداختی مشتری مسترد خواهد شد.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` باشد.  ---  ### ۱۲. CANCELLED_BY_MERCHANT — لغو توسط مرچنت  **معنا:** مرچنت پس از پرداخت، سفارش را به هر دلیلی لغو کرده است.  **چگونه اتفاق می‌افتد:** مرچنت درخواست لغو سفارش را ارسال می‌کند. وجه پرداختی مشتری به او بازگردانده می‌شود.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` یا `VERIFIED_BY_MERCHANT` باشد.  ---  ### ۱۳. REQUEST_TO_REFUND — درخواست استرداد توسط مشتری  **معنا:** مشتری درخواست بازگشت وجه داده و سیستم در حال پردازش استرداد است.  **چگونه اتفاق می‌افتد:** مرچنت از طریق API درخواست استرداد را ثبت می‌کند (`POST /api/v1/orders/manager/{uuid}/refund`). سفارش وارد صف پردازش استرداد می‌شود.  **وابستگی‌ها:** سفارش باید در وضعیت `PAID_BY_USER` یا `VERIFIED_BY_MERCHANT` باشد.  ---  ### ۱۴، ۱۵، ۱۶. سایر وضعیت‌های درخواست استرداد  این وضعیت‌ها بر اساس دلیل استرداد از هم تفکیک می‌شوند:  - **۱۴ — REQUEST_TO_REFUND_TO_MERCHANT_AFTER_FAILED_TO_VERIFY:** استرداد پس از ناموفق بودن تأیید مرچنت؛ وجه به حساب مرچنت بازمی‌گردد. - **۱۵ — REQUEST_TO_REFUND_TO_CUSTOMER_AFTER_FAILED_BY_MERCHANT:** استرداد پس از خطای مرچنت؛ وجه به مشتری بازمی‌گردد. - **۱۶ — REQUEST_TO_REFUND_TO_MERCHANT_AFTER_CANCELLED_BY_MERCHANT:** استرداد پس از لغو توسط مرچنت؛ وجه به حساب مرچنت برمی‌گردد.  **چگونه اتفاق می‌افتد:** به صورت خودکار پس از رسیدن به وضعیت‌های ناموفق/لغو مربوطه توسط سیستم تنظیم می‌شود.  ---  ### ۱۷. REFUND_COMPLETED — استرداد تکمیل شد  **معنا:** وجه با موفقیت به صاحب آن (مشتری یا مرچنت بسته به نوع استرداد) بازگردانده شده است.  **چگونه اتفاق می‌افتد:** Task پردازش استرداد (`process_order_refund`) پس از تأیید موفق بازگشت وجه از سوی Digipay، وضعیت سفارش را به `REFUND_COMPLETED` تغییر می‌دهد.  **وابستگی‌ها:** یکی از وضعیت‌های درخواست استرداد (۱۳، ۱۴، ۱۵ یا ۱۶) باید فعال باشد و Digipay تراکنش استرداد را تأیید کرده باشد.  </div> 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bwdk_sdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiMerchantApiV1AuthStatusRetrieveRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
}

func (r ApiMerchantApiV1AuthStatusRetrieveRequest) Execute() (*AuthStatusResponse, *http.Response, error) {
	return r.ApiService.MerchantApiV1AuthStatusRetrieveExecute(r)
}

/*
MerchantApiV1AuthStatusRetrieve وضعیت لاگین بودن

<div dir="rtl" style="text-align: right;">

بررسی وضعیت احراز هویت فروشنده

## توضیحات

این endpoint برای بررسی اعتبار **API_KEY** فروشنده استفاده می‌شود. اگر کلید معتبر باشد، پاسخ `is_authenticated: true` برمی‌گردد. از این endpoint برای تأیید صحت کلید API قبل از شروع عملیات استفاده کنید.

نیاز به **API_KEY** فروشنده دارد (فقط Header لازم است، بدنه درخواست ندارد).

</div>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMerchantApiV1AuthStatusRetrieveRequest
*/
func (a *DefaultApiService) MerchantApiV1AuthStatusRetrieve(ctx context.Context) ApiMerchantApiV1AuthStatusRetrieveRequest {
	return ApiMerchantApiV1AuthStatusRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuthStatusResponse
func (a *DefaultApiService) MerchantApiV1AuthStatusRetrieveExecute(r ApiMerchantApiV1AuthStatusRetrieveRequest) (*AuthStatusResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthStatusResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.MerchantApiV1AuthStatusRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/api/v1/auth/status/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["MerchantAPIKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiOrderApiV1CreateOrderCreateRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orderCreate *OrderCreate
}

func (r ApiOrderApiV1CreateOrderCreateRequest) OrderCreate(orderCreate OrderCreate) ApiOrderApiV1CreateOrderCreateRequest {
	r.orderCreate = &orderCreate
	return r
}

func (r ApiOrderApiV1CreateOrderCreateRequest) Execute() (*OrderCreateResponse, *http.Response, error) {
	return r.ApiService.OrderApiV1CreateOrderCreateExecute(r)
}

/*
OrderApiV1CreateOrderCreate ساخت سفارش

<div dir="rtl" style="text-align: right;">

ساخت سفارش جدید در سیستم BWDK

## توضیحات

این endpoint برای ایجاد یک سفارش جدید در سیستم خرید با دیجی‌کالا استفاده می‌شود. در این درخواست باید اطلاعات سفارش، اقلام سبد خرید، و آدرس callback شامل شود.

برای شروع ارتباط با سرویس‌های **خرید با دیجی‌کالا** شما باید دارای **API_KEY** باشید که این مورد از سمت تیم دیجی‌فای در اختیار شما قرار خواهد گرفت.

### روند کاری

**مرحله ۱: درخواست ساخت سفارش**

کاربر پس از انتخاب کالاهای خود در سبد خرید، بر روی دکمه **خرید با دیجی‌کالا** کلیک می‌کند و بک‌اند مرچنت درخواستی برای ساخت سفارش BWDK دریافت می‌کند. در این مرحله اولین درخواست خود را به بک‌اند BWDK ارسال می‌نمایید:

BWDK یک سفارش جدید برای مرچنت با وضعیت **INITIAL** ایجاد می‌کند و **order_uuid** را جنریت می‌کند. آدرس **order_start_url** نقطه شروع مسیر تکمیل سفارش است (انتخاب آدرس، شیپینگ، پکینگ، پرداخت و غیره).
<br>
</div>

```mermaid
sequenceDiagram
    participant M as فروشنده
    participant API as BWDK API
    participant C as مشتری
    participant PG as درگاه پرداخت

    M->>API: POST /api/v1/orders/create
    Note over M,API: شناسه سفارش, کالاها, مبلغ, callback_url
    API-->>M: {لینک شروع سفارش, شناسه سفارش}

    M->>C: تغییر مسیر به لینک شروع

    C->>API: تکمیل درخواست خرید توسط مشتری
    API->>PG: شروع فرآیند پرداخت
    PG-->>C: نتیجه پرداخت
    PG->>API: تأیید پرداخت
    API-->>C: تغییر مسیر به callback_url

    M->>API: POST /api/v1/orders/manager/{uuid}/verify
    Note over M,API: {شناسه منحصربفرد فروشنده}
    API-->>M: سفارش تأیید شد و آماده ارسال است
```

</div>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOrderApiV1CreateOrderCreateRequest
*/
func (a *DefaultApiService) OrderApiV1CreateOrderCreate(ctx context.Context) ApiOrderApiV1CreateOrderCreateRequest {
	return ApiOrderApiV1CreateOrderCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OrderCreateResponse
func (a *DefaultApiService) OrderApiV1CreateOrderCreateExecute(r ApiOrderApiV1CreateOrderCreateRequest) (*OrderCreateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrderCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrderApiV1CreateOrderCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order/api/v1/create-order/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orderCreate == nil {
		return localVarReturnValue, nil, reportError("orderCreate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.orderCreate
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["MerchantAPIKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiOrderApiV1ManagerCancelShipmentCreateRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orderUuid string
}

func (r ApiOrderApiV1ManagerCancelShipmentCreateRequest) Execute() (*MerchantOrderCancelShipmentResponse, *http.Response, error) {
	return r.ApiService.OrderApiV1ManagerCancelShipmentCreateExecute(r)
}

/*
OrderApiV1ManagerCancelShipmentCreate لغو ارسال سفارش

<div dir="rtl" style="text-align: right;">

لغو مرسوله دیجی‌اکسپرس

## توضیحات

این endpoint برای لغو یک مرسوله ثبت‌شده در سرویس دیجی‌اکسپرس استفاده می‌شود. پس از لغو موفق، مرسوله از صف ارسال خارج می‌شود.

نیاز به **API_KEY** فروشنده دارد.

## شرایط لغو

* سفارش باید دارای روش ارسال **DigiExpress** باشد
* مرسوله باید در وضعیت **در انتظار تحویل به پیک** (Request for Pickup) باشد

</div>

```mermaid
sequenceDiagram
    participant M as فروشنده
    participant API as BWDK API
    participant DX as دیجی‌اکسپرس

    M->>API: POST /order/api/v1/manager/{order_uuid}/cancel-shipment/
    Note over M,API: Header: X-API-KEY (بدون بدنه)

    alt روش ارسال DigiExpress نیست
        API-->>M: 400 خطا
        Note over API,M: {error: "Selected shipping method is not DigiExpress"}
    else مرسوله قابل لغو نیست
        API-->>M: 400 خطا
        Note over API,M: {error: "..."}
    else لغو موفق
        API->>DX: لغو مرسوله
        DX-->>API: تأیید لغو
        API-->>M: 200 موفق
        Note over API,M: {message, order_uuid, status, status_display}
    end
```


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orderUuid
 @return ApiOrderApiV1ManagerCancelShipmentCreateRequest
*/
func (a *DefaultApiService) OrderApiV1ManagerCancelShipmentCreate(ctx context.Context, orderUuid string) ApiOrderApiV1ManagerCancelShipmentCreateRequest {
	return ApiOrderApiV1ManagerCancelShipmentCreateRequest{
		ApiService: a,
		ctx: ctx,
		orderUuid: orderUuid,
	}
}

// Execute executes the request
//  @return MerchantOrderCancelShipmentResponse
func (a *DefaultApiService) OrderApiV1ManagerCancelShipmentCreateExecute(r ApiOrderApiV1ManagerCancelShipmentCreateRequest) (*MerchantOrderCancelShipmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantOrderCancelShipmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrderApiV1ManagerCancelShipmentCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order/api/v1/manager/{order_uuid}/cancel-shipment/"
	localVarPath = strings.Replace(localVarPath, "{"+"order_uuid"+"}", url.PathEscape(parameterValueToString(r.orderUuid, "orderUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["MerchantAPIKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v OrderError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v OrderError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiOrderApiV1ManagerChangeShippingMethodUpdateRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orderUuid string
	orderDetail *OrderDetail
}

func (r ApiOrderApiV1ManagerChangeShippingMethodUpdateRequest) OrderDetail(orderDetail OrderDetail) ApiOrderApiV1ManagerChangeShippingMethodUpdateRequest {
	r.orderDetail = &orderDetail
	return r
}

func (r ApiOrderApiV1ManagerChangeShippingMethodUpdateRequest) Execute() (*OrderDetail, *http.Response, error) {
	return r.ApiService.OrderApiV1ManagerChangeShippingMethodUpdateExecute(r)
}

/*
OrderApiV1ManagerChangeShippingMethodUpdate تغییر روش ارسال

<div dir="rtl" style="text-align: right;">

تغییر روش ارسال سفارش

## توضیحات

این endpoint به فروشنده اجازه می‌دهد روش ارسال یک سفارش را تغییر دهد. این عملیات معمولاً زمانی استفاده می‌شود که فروشنده بخواهد از DigiExpress به روش ارسال پیش‌فرض (یا بالعکس) تغییر دهد.

نیاز به **API_KEY** فروشنده دارد.

## پارامترهای ورودی

* **updated_shipping**: شناسه روش ارسال جدید
* **preparation_time** (اختیاری): زمان آماده‌سازی (روز) برای DigiExpress

</div>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orderUuid
 @return ApiOrderApiV1ManagerChangeShippingMethodUpdateRequest
*/
func (a *DefaultApiService) OrderApiV1ManagerChangeShippingMethodUpdate(ctx context.Context, orderUuid string) ApiOrderApiV1ManagerChangeShippingMethodUpdateRequest {
	return ApiOrderApiV1ManagerChangeShippingMethodUpdateRequest{
		ApiService: a,
		ctx: ctx,
		orderUuid: orderUuid,
	}
}

// Execute executes the request
//  @return OrderDetail
func (a *DefaultApiService) OrderApiV1ManagerChangeShippingMethodUpdateExecute(r ApiOrderApiV1ManagerChangeShippingMethodUpdateRequest) (*OrderDetail, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrderDetail
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrderApiV1ManagerChangeShippingMethodUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order/api/v1/manager/{order_uuid}/change-shipping-method/"
	localVarPath = strings.Replace(localVarPath, "{"+"order_uuid"+"}", url.PathEscape(parameterValueToString(r.orderUuid, "orderUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orderDetail == nil {
		return localVarReturnValue, nil, reportError("orderDetail is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.orderDetail
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["MerchantAPIKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiOrderApiV1ManagerListRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	cities *string
	createdAt *string
	cursor *string
	orderIds *string
	ordering *string
	paymentTypes *string
	provinces *string
	referenceCode *string
	search *string
	shippingTypes *string
	status *int32
	statuses *string
	todayPickup *bool
}

func (r ApiOrderApiV1ManagerListRequest) Cities(cities string) ApiOrderApiV1ManagerListRequest {
	r.cities = &cities
	return r
}

func (r ApiOrderApiV1ManagerListRequest) CreatedAt(createdAt string) ApiOrderApiV1ManagerListRequest {
	r.createdAt = &createdAt
	return r
}

// مقدار نشانگر صفحه‌بندی.
func (r ApiOrderApiV1ManagerListRequest) Cursor(cursor string) ApiOrderApiV1ManagerListRequest {
	r.cursor = &cursor
	return r
}

func (r ApiOrderApiV1ManagerListRequest) OrderIds(orderIds string) ApiOrderApiV1ManagerListRequest {
	r.orderIds = &orderIds
	return r
}

// کدام فیلد باید هنگام مرتب‌سازی نتایج استفاده شود.
func (r ApiOrderApiV1ManagerListRequest) Ordering(ordering string) ApiOrderApiV1ManagerListRequest {
	r.ordering = &ordering
	return r
}

func (r ApiOrderApiV1ManagerListRequest) PaymentTypes(paymentTypes string) ApiOrderApiV1ManagerListRequest {
	r.paymentTypes = &paymentTypes
	return r
}

func (r ApiOrderApiV1ManagerListRequest) Provinces(provinces string) ApiOrderApiV1ManagerListRequest {
	r.provinces = &provinces
	return r
}

func (r ApiOrderApiV1ManagerListRequest) ReferenceCode(referenceCode string) ApiOrderApiV1ManagerListRequest {
	r.referenceCode = &referenceCode
	return r
}

// یک عبارت جستجو.
func (r ApiOrderApiV1ManagerListRequest) Search(search string) ApiOrderApiV1ManagerListRequest {
	r.search = &search
	return r
}

func (r ApiOrderApiV1ManagerListRequest) ShippingTypes(shippingTypes string) ApiOrderApiV1ManagerListRequest {
	r.shippingTypes = &shippingTypes
	return r
}

// * &#x60;1&#x60; - اولیه * &#x60;2&#x60; - شروع در * &#x60;3&#x60; - در انتظار * &#x60;4&#x60; - در انتظار درگاه * &#x60;5&#x60; - منقضی شده * &#x60;6&#x60; - لغو شده * &#x60;7&#x60; - ممنوع شده توسط ما * &#x60;8&#x60; - ناموفق در پرداخت * &#x60;9&#x60; - تأیید شده توسط فروشنده * &#x60;10&#x60; - ناموفق در تأیید توسط فروشنده * &#x60;11&#x60; - فروشگاه * &#x60;12&#x60; - لغو شده توسط فروشنده * &#x60;13&#x60; - درخواست بازگرداندن وجه به مشتری به دلیل درخواست مشتری * &#x60;14&#x60; - درخواست بازگرداندن وجه به فروشنده پس از ناموفقی در تأیید توسط فروشنده * &#x60;15&#x60; - درخواست بازگرداندن وجه به مشتری پس از ناموفقی توسط فروشنده * &#x60;16&#x60; - بازگردانده شده به فروشنده پس از لغو توسط فروشنده * &#x60;17&#x60; - بازگرداندن وجه تکمیل شد * &#x60;18&#x60; - زمان مجاز برای منقضی کردن گذشته است * &#x60;19&#x60; - تحویل نشده * &#x60;20&#x60; - مرسوله
func (r ApiOrderApiV1ManagerListRequest) Status(status int32) ApiOrderApiV1ManagerListRequest {
	r.status = &status
	return r
}

func (r ApiOrderApiV1ManagerListRequest) Statuses(statuses string) ApiOrderApiV1ManagerListRequest {
	r.statuses = &statuses
	return r
}

func (r ApiOrderApiV1ManagerListRequest) TodayPickup(todayPickup bool) ApiOrderApiV1ManagerListRequest {
	r.todayPickup = &todayPickup
	return r
}

func (r ApiOrderApiV1ManagerListRequest) Execute() (*PaginatedOrderDetailList, *http.Response, error) {
	return r.ApiService.OrderApiV1ManagerListExecute(r)
}

/*
OrderApiV1ManagerList لیست سفارشات

<div dir="rtl" style="text-align: right;">

لیست سفارشات فروشنده

## توضیحات

این endpoint لیست تمام سفارشات مرتبط با فروشنده را با امکان فیلتر، جستجو و مرتب‌سازی برمی‌گرداند. نتایج به صورت صفحه‌بندی‌شده (Cursor Pagination) ارسال می‌شوند و به ترتیب جدیدترین سفارش اول مرتب می‌شوند.

نیاز به **API_KEY** فروشنده دارد.

## پارامترهای فیلتر

* **status**: وضعیت سفارش (INITIAL, PENDING, PAID_BY_USER, VERIFIED_BY_MERCHANT, ...)
* **created_at__gte / created_at__lte**: فیلتر بر اساس تاریخ ایجاد
* **search**: جستجو در شماره تلفن مشتری، نام، کد مرجع، شناسه سفارش مرچنت
* **ordering**: مرتب‌سازی بر اساس created_at, final_amount, status, merchant_order_id

</div>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOrderApiV1ManagerListRequest
*/
func (a *DefaultApiService) OrderApiV1ManagerList(ctx context.Context) ApiOrderApiV1ManagerListRequest {
	return ApiOrderApiV1ManagerListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedOrderDetailList
func (a *DefaultApiService) OrderApiV1ManagerListExecute(r ApiOrderApiV1ManagerListRequest) (*PaginatedOrderDetailList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedOrderDetailList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrderApiV1ManagerList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order/api/v1/manager/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cities != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cities", r.cities, "form", "")
	}
	if r.createdAt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "created_at", r.createdAt, "form", "")
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
	if r.orderIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_ids", r.orderIds, "form", "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "form", "")
	}
	if r.paymentTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payment_types", r.paymentTypes, "form", "")
	}
	if r.provinces != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "provinces", r.provinces, "form", "")
	}
	if r.referenceCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reference_code", r.referenceCode, "form", "")
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "form", "")
	}
	if r.shippingTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "shipping_types", r.shippingTypes, "form", "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.statuses != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "statuses", r.statuses, "form", "")
	}
	if r.todayPickup != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "today_pickup", r.todayPickup, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["MerchantAPIKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiOrderApiV1ManagerPaidListRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	cities *string
	createdAt *string
	cursor *string
	orderIds *string
	ordering *string
	paymentTypes *string
	provinces *string
	referenceCode *string
	search *string
	shippingTypes *string
	status *int32
	statuses *string
	todayPickup *bool
}

func (r ApiOrderApiV1ManagerPaidListRequest) Cities(cities string) ApiOrderApiV1ManagerPaidListRequest {
	r.cities = &cities
	return r
}

func (r ApiOrderApiV1ManagerPaidListRequest) CreatedAt(createdAt string) ApiOrderApiV1ManagerPaidListRequest {
	r.createdAt = &createdAt
	return r
}

// مقدار نشانگر صفحه‌بندی.
func (r ApiOrderApiV1ManagerPaidListRequest) Cursor(cursor string) ApiOrderApiV1ManagerPaidListRequest {
	r.cursor = &cursor
	return r
}

func (r ApiOrderApiV1ManagerPaidListRequest) OrderIds(orderIds string) ApiOrderApiV1ManagerPaidListRequest {
	r.orderIds = &orderIds
	return r
}

// کدام فیلد باید هنگام مرتب‌سازی نتایج استفاده شود.
func (r ApiOrderApiV1ManagerPaidListRequest) Ordering(ordering string) ApiOrderApiV1ManagerPaidListRequest {
	r.ordering = &ordering
	return r
}

func (r ApiOrderApiV1ManagerPaidListRequest) PaymentTypes(paymentTypes string) ApiOrderApiV1ManagerPaidListRequest {
	r.paymentTypes = &paymentTypes
	return r
}

func (r ApiOrderApiV1ManagerPaidListRequest) Provinces(provinces string) ApiOrderApiV1ManagerPaidListRequest {
	r.provinces = &provinces
	return r
}

func (r ApiOrderApiV1ManagerPaidListRequest) ReferenceCode(referenceCode string) ApiOrderApiV1ManagerPaidListRequest {
	r.referenceCode = &referenceCode
	return r
}

// یک عبارت جستجو.
func (r ApiOrderApiV1ManagerPaidListRequest) Search(search string) ApiOrderApiV1ManagerPaidListRequest {
	r.search = &search
	return r
}

func (r ApiOrderApiV1ManagerPaidListRequest) ShippingTypes(shippingTypes string) ApiOrderApiV1ManagerPaidListRequest {
	r.shippingTypes = &shippingTypes
	return r
}

// * &#x60;1&#x60; - اولیه * &#x60;2&#x60; - شروع در * &#x60;3&#x60; - در انتظار * &#x60;4&#x60; - در انتظار درگاه * &#x60;5&#x60; - منقضی شده * &#x60;6&#x60; - لغو شده * &#x60;7&#x60; - ممنوع شده توسط ما * &#x60;8&#x60; - ناموفق در پرداخت * &#x60;9&#x60; - تأیید شده توسط فروشنده * &#x60;10&#x60; - ناموفق در تأیید توسط فروشنده * &#x60;11&#x60; - فروشگاه * &#x60;12&#x60; - لغو شده توسط فروشنده * &#x60;13&#x60; - درخواست بازگرداندن وجه به مشتری به دلیل درخواست مشتری * &#x60;14&#x60; - درخواست بازگرداندن وجه به فروشنده پس از ناموفقی در تأیید توسط فروشنده * &#x60;15&#x60; - درخواست بازگرداندن وجه به مشتری پس از ناموفقی توسط فروشنده * &#x60;16&#x60; - بازگردانده شده به فروشنده پس از لغو توسط فروشنده * &#x60;17&#x60; - بازگرداندن وجه تکمیل شد * &#x60;18&#x60; - زمان مجاز برای منقضی کردن گذشته است * &#x60;19&#x60; - تحویل نشده * &#x60;20&#x60; - مرسوله
func (r ApiOrderApiV1ManagerPaidListRequest) Status(status int32) ApiOrderApiV1ManagerPaidListRequest {
	r.status = &status
	return r
}

func (r ApiOrderApiV1ManagerPaidListRequest) Statuses(statuses string) ApiOrderApiV1ManagerPaidListRequest {
	r.statuses = &statuses
	return r
}

func (r ApiOrderApiV1ManagerPaidListRequest) TodayPickup(todayPickup bool) ApiOrderApiV1ManagerPaidListRequest {
	r.todayPickup = &todayPickup
	return r
}

func (r ApiOrderApiV1ManagerPaidListRequest) Execute() (*PaginatedMerchantPaidOrderListList, *http.Response, error) {
	return r.ApiService.OrderApiV1ManagerPaidListExecute(r)
}

/*
OrderApiV1ManagerPaidList سفارش پرداخت‌شده و تایید‌نشده

لیست تمامی سفارشاتی که توسط کاربر پرداخت شده اند ولی توسط فروشنده تایید نشده اند.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOrderApiV1ManagerPaidListRequest
*/
func (a *DefaultApiService) OrderApiV1ManagerPaidList(ctx context.Context) ApiOrderApiV1ManagerPaidListRequest {
	return ApiOrderApiV1ManagerPaidListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedMerchantPaidOrderListList
func (a *DefaultApiService) OrderApiV1ManagerPaidListExecute(r ApiOrderApiV1ManagerPaidListRequest) (*PaginatedMerchantPaidOrderListList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedMerchantPaidOrderListList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrderApiV1ManagerPaidList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order/api/v1/manager/paid/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cities != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cities", r.cities, "form", "")
	}
	if r.createdAt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "created_at", r.createdAt, "form", "")
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
	if r.orderIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_ids", r.orderIds, "form", "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "form", "")
	}
	if r.paymentTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payment_types", r.paymentTypes, "form", "")
	}
	if r.provinces != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "provinces", r.provinces, "form", "")
	}
	if r.referenceCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reference_code", r.referenceCode, "form", "")
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "form", "")
	}
	if r.shippingTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "shipping_types", r.shippingTypes, "form", "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.statuses != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "statuses", r.statuses, "form", "")
	}
	if r.todayPickup != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "today_pickup", r.todayPickup, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["MerchantAPIKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiOrderApiV1ManagerRefundCreateRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orderUuid string
	refundOrder *RefundOrder
}

func (r ApiOrderApiV1ManagerRefundCreateRequest) RefundOrder(refundOrder RefundOrder) ApiOrderApiV1ManagerRefundCreateRequest {
	r.refundOrder = &refundOrder
	return r
}

func (r ApiOrderApiV1ManagerRefundCreateRequest) Execute() (*MerchantOrderRefundResponse, *http.Response, error) {
	return r.ApiService.OrderApiV1ManagerRefundCreateExecute(r)
}

/*
OrderApiV1ManagerRefundCreate بازگشت سفارش

<div dir="rtl" style="text-align: right;">
بازگشت و لغو سفارش

## توضیحات

این endpoint برای بازگشت یا لغو سفارشی استفاده می‌شود که قبلاً پرداخت شده یا تایید شده است. در این مرحله مبلغ سفارش به کاربر بازگشت داده می‌شود و وضعیت سفارش به **REFUNDED** تغییر می‌یابد.


## شرایط بازگشت

سفارش باید در یکی از وضعیت‌های زیر باشد تا بازگشت امکان‌پذیر باشد:
* **PAID_BY_USER**: سفارش پرداخت شده است اما هنوز تایید نشده
* **VERIFIED_BY_MERCHANT**: سفارش تایید شده است

سفارش نباید قبلاً بازگشت داده شده باشد.

**پاسخ خطا** - اگر سفارش در وضعیت مناسب نباشد یا قبلاً بازگشت داده شده باشد
</div>

```mermaid
sequenceDiagram
    participant M as فروشنده
    participant API as BWDK API

    M->>API: POST /api/v1/orders/manager/{uuid}/refund
    Note over M,API: {reason: "درخواست مشتری"}

    alt سفارش قابل بازگشت نیست
        API-->>M: 500 خطا
        Note over API,M: {error: "شروع بازگشت ناموفق بود.<br/>لطفاً دوباره تلاش کنید."}
    else سفارش معتبر است
        API-->>M: 200 موفق
        Note over API,M: {message: "درخواست بازگشت با<br/>موفقیت شروع شد",<br/>order_uuid, status: 13,<br/>status_display,<br/>refund_reason}

        M->>API: GET /api/v1/orders/manager/{uuid}
        Note over M,API: بررسی وضعیت سفارش<br/>(endpoint جداگانه /refund/status وجود ندارد)

        alt status: 17 (بازگشت تکمیل شد)
            API-->>M: 200 موفق
            Note over API,M: {order_uuid, status: 17,<br/>status_display: "بازگشت تکمیل شد",<br/>metadata.refund_tracking_code,<br/>metadata.refund_completed_at}

        else status: 13 (در حال پردازش)
            API-->>M: 200 موفق
            Note over API,M: {order_uuid, status: 13,<br/>status_display: "درخواست بازگشت به مشتری<br/>به دلیل درخواست<br/>مشتری",<br/>metadata.refund_reason}

        else status: قبلی (بازگشت ناموفق)
            API-->>M: 200 موفق
            Note over API,M: {order_uuid, status: (قبلی),<br/>metadata.refund_error,<br/>metadata.refund_failed_at}
        end
    end
```


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orderUuid
 @return ApiOrderApiV1ManagerRefundCreateRequest
*/
func (a *DefaultApiService) OrderApiV1ManagerRefundCreate(ctx context.Context, orderUuid string) ApiOrderApiV1ManagerRefundCreateRequest {
	return ApiOrderApiV1ManagerRefundCreateRequest{
		ApiService: a,
		ctx: ctx,
		orderUuid: orderUuid,
	}
}

// Execute executes the request
//  @return MerchantOrderRefundResponse
func (a *DefaultApiService) OrderApiV1ManagerRefundCreateExecute(r ApiOrderApiV1ManagerRefundCreateRequest) (*MerchantOrderRefundResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantOrderRefundResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrderApiV1ManagerRefundCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order/api/v1/manager/{order_uuid}/refund/"
	localVarPath = strings.Replace(localVarPath, "{"+"order_uuid"+"}", url.PathEscape(parameterValueToString(r.orderUuid, "orderUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.refundOrder
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["MerchantAPIKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v OrderError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiOrderApiV1ManagerRetrieveRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orderUuid string
}

func (r ApiOrderApiV1ManagerRetrieveRequest) Execute() (*OrderDetail, *http.Response, error) {
	return r.ApiService.OrderApiV1ManagerRetrieveExecute(r)
}

/*
OrderApiV1ManagerRetrieve دریافت سفارش

<div dir="rtl" style="text-align: right;">

# مدیریت سفارشات

## توضیحات

این endpoint برای مدیریت و مشاهده سفارشات فروشنده استفاده می‌شود.

</div>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orderUuid
 @return ApiOrderApiV1ManagerRetrieveRequest
*/
func (a *DefaultApiService) OrderApiV1ManagerRetrieve(ctx context.Context, orderUuid string) ApiOrderApiV1ManagerRetrieveRequest {
	return ApiOrderApiV1ManagerRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		orderUuid: orderUuid,
	}
}

// Execute executes the request
//  @return OrderDetail
func (a *DefaultApiService) OrderApiV1ManagerRetrieveExecute(r ApiOrderApiV1ManagerRetrieveRequest) (*OrderDetail, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrderDetail
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrderApiV1ManagerRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order/api/v1/manager/{order_uuid}/"
	localVarPath = strings.Replace(localVarPath, "{"+"order_uuid"+"}", url.PathEscape(parameterValueToString(r.orderUuid, "orderUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["MerchantAPIKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiOrderApiV1ManagerReviveShipmentCreateRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orderUuid string
	reviveShipment *ReviveShipment
}

func (r ApiOrderApiV1ManagerReviveShipmentCreateRequest) ReviveShipment(reviveShipment ReviveShipment) ApiOrderApiV1ManagerReviveShipmentCreateRequest {
	r.reviveShipment = &reviveShipment
	return r
}

func (r ApiOrderApiV1ManagerReviveShipmentCreateRequest) Execute() (*MerchantOrderReviveShipmentResponse, *http.Response, error) {
	return r.ApiService.OrderApiV1ManagerReviveShipmentCreateExecute(r)
}

/*
OrderApiV1ManagerReviveShipmentCreate احیای ارسال سفارش

<div dir="rtl" style="text-align: right;">

احیای مرسوله دیجی‌اکسپرس

## توضیحات

این endpoint برای احیای (reactivate) یک مرسوله دیجی‌اکسپرس که قبلاً لغو شده یا در وضعیت غیرفعال است استفاده می‌شود. با ارسال `preparation_time` (زمان آماده‌سازی بر حسب روز)، زمان جدید آماده بودن بار تنظیم می‌شود.

نیاز به **API_KEY** فروشنده دارد.

## پارامترهای ورودی

* **preparation_time** (اختیاری، پیش‌فرض: ۲): تعداد روز تا آماده‌شدن بار برای تحویل به پیک

## شرایط

* سفارش باید دارای روش ارسال **DigiExpress** باشد
* مرسوله باید در وضعیت قابل احیا باشد

</div>

```mermaid
sequenceDiagram
    participant M as فروشنده
    participant API as BWDK API
    participant DX as دیجی‌اکسپرس

    M->>API: POST /order/api/v1/manager/{order_uuid}/revive-shipment/
    Note over M,API: Header: X-API-KEY<br/>{preparation_time: 2}

    alt روش ارسال DigiExpress نیست
        API-->>M: 400 خطا
        Note over API,M: {error: "Selected shipping method is not DigiExpress"}
    else احیا موفق
        API->>DX: احیای مرسوله با زمان جدید
        DX-->>API: تأیید احیا
        API-->>M: 200 موفق
        Note over API,M: {message, order_uuid, status, status_display}
    end
```


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orderUuid
 @return ApiOrderApiV1ManagerReviveShipmentCreateRequest
*/
func (a *DefaultApiService) OrderApiV1ManagerReviveShipmentCreate(ctx context.Context, orderUuid string) ApiOrderApiV1ManagerReviveShipmentCreateRequest {
	return ApiOrderApiV1ManagerReviveShipmentCreateRequest{
		ApiService: a,
		ctx: ctx,
		orderUuid: orderUuid,
	}
}

// Execute executes the request
//  @return MerchantOrderReviveShipmentResponse
func (a *DefaultApiService) OrderApiV1ManagerReviveShipmentCreateExecute(r ApiOrderApiV1ManagerReviveShipmentCreateRequest) (*MerchantOrderReviveShipmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantOrderReviveShipmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrderApiV1ManagerReviveShipmentCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order/api/v1/manager/{order_uuid}/revive-shipment/"
	localVarPath = strings.Replace(localVarPath, "{"+"order_uuid"+"}", url.PathEscape(parameterValueToString(r.orderUuid, "orderUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.reviveShipment
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["MerchantAPIKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v OrderError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v OrderError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiOrderApiV1ManagerUpdateStatusUpdateRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orderUuid string
	updateOrderStatus *UpdateOrderStatus
}

func (r ApiOrderApiV1ManagerUpdateStatusUpdateRequest) UpdateOrderStatus(updateOrderStatus UpdateOrderStatus) ApiOrderApiV1ManagerUpdateStatusUpdateRequest {
	r.updateOrderStatus = &updateOrderStatus
	return r
}

func (r ApiOrderApiV1ManagerUpdateStatusUpdateRequest) Execute() (*OrderDetail, *http.Response, error) {
	return r.ApiService.OrderApiV1ManagerUpdateStatusUpdateExecute(r)
}

/*
OrderApiV1ManagerUpdateStatusUpdate به‌روزرسانی وضعیت سفارش

<div dir="rtl" style="text-align: right;">

بروزرسانی وضعیت سفارش

## توضیحات

این endpoint به فروشنده امکان می‌دهد وضعیت یک سفارش را به‌صورت مستقیم تغییر دهد. این endpoint برای مواردی مانند علامت‌گذاری سفارش به‌عنوان «ارسال شده» یا «تحویل داده شده» توسط فروشنده استفاده می‌شود.

نیاز به **API_KEY** فروشنده دارد.

## وضعیت‌های مجاز

* **DELIVERED**: تحویل شد
* **DISPATCHED**: ارسال شد
* سایر وضعیت‌های تعریف‌شده در سیستم

</div>

```mermaid
sequenceDiagram
    participant M as فروشنده
    participant API as BWDK API

    M->>API: PUT /order/api/v1/manager/{order_uuid}/update-status/
    Note over M,API: Header: X-API-KEY<br/>{status: "DELIVERED"}

    alt وضعیت معتبر است
        API-->>M: 200 موفق
        Note over API,M: اطلاعات کامل سفارش با وضعیت جدید
    else وضعیت نامعتبر است
        API-->>M: 400 خطا
        Note over API,M: {error: "invalid status"}
    end
```


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orderUuid
 @return ApiOrderApiV1ManagerUpdateStatusUpdateRequest
*/
func (a *DefaultApiService) OrderApiV1ManagerUpdateStatusUpdate(ctx context.Context, orderUuid string) ApiOrderApiV1ManagerUpdateStatusUpdateRequest {
	return ApiOrderApiV1ManagerUpdateStatusUpdateRequest{
		ApiService: a,
		ctx: ctx,
		orderUuid: orderUuid,
	}
}

// Execute executes the request
//  @return OrderDetail
func (a *DefaultApiService) OrderApiV1ManagerUpdateStatusUpdateExecute(r ApiOrderApiV1ManagerUpdateStatusUpdateRequest) (*OrderDetail, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrderDetail
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrderApiV1ManagerUpdateStatusUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order/api/v1/manager/{order_uuid}/update-status/"
	localVarPath = strings.Replace(localVarPath, "{"+"order_uuid"+"}", url.PathEscape(parameterValueToString(r.orderUuid, "orderUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateOrderStatus == nil {
		return localVarReturnValue, nil, reportError("updateOrderStatus is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateOrderStatus
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["MerchantAPIKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiOrderApiV1ManagerVerifyCreateRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orderUuid string
	verifyOrder *VerifyOrder
}

func (r ApiOrderApiV1ManagerVerifyCreateRequest) VerifyOrder(verifyOrder VerifyOrder) ApiOrderApiV1ManagerVerifyCreateRequest {
	r.verifyOrder = &verifyOrder
	return r
}

func (r ApiOrderApiV1ManagerVerifyCreateRequest) Execute() (*OrderDetail, *http.Response, error) {
	return r.ApiService.OrderApiV1ManagerVerifyCreateExecute(r)
}

/*
OrderApiV1ManagerVerifyCreate تایید سفارش

<div dir="rtl" style="text-align: right;">

تایید سفارش پس از پرداخت

## توضیحات

پس از اتمام فرایند پرداخت توسط کاربر، مرچنت باید سفارش را تایید کند. این endpoint برای تایید و نهایی کردن سفارش پس از پرداخت موفق توسط مشتری استفاده می‌شود. در این مرحله مرچنت سفارش را تایید می‌کند و وضعیت سفارش به **VERIFIED_BY_MERCHANT** تغییر می‌یابد.

## روند کاری

**مرحله ۲: بازگشت کاربر و دریافت نتیجه**

پس از تکمیل فرایند پرداخت (موفق یا ناموفق)، کاربر به آدرس callback_url که هنگام ساخت سفارش ارسال کرده بودید بازگردانده می‌شود. در این درخواست وضعیت سفارش به صورت query parameters ارسال می‌شود:


**Query Parameters:**
* **success**: متغیر منطقی نشان‌دهنده موفق یا ناموفق بودن سفارش
* **status**: وضعیت سفارش (PAID، FAILED، وغیره)
* **phone_number**: شماره تلفن مشتری
* **order_uuid**: شناسه یکتای سفارش
* **merchant_order_id**: شناسه سفارش در سیستم مرچنت

**مرحله ۳: تایید سفارش در بک‌اند**

سپس شما باید این endpoint را فراخوانی کنید تا سفارش را تایید کنید و اطمینان حاصل کنید که سفارش موفقیت‌آمیز ثبت شده است:
</div>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orderUuid
 @return ApiOrderApiV1ManagerVerifyCreateRequest
*/
func (a *DefaultApiService) OrderApiV1ManagerVerifyCreate(ctx context.Context, orderUuid string) ApiOrderApiV1ManagerVerifyCreateRequest {
	return ApiOrderApiV1ManagerVerifyCreateRequest{
		ApiService: a,
		ctx: ctx,
		orderUuid: orderUuid,
	}
}

// Execute executes the request
//  @return OrderDetail
func (a *DefaultApiService) OrderApiV1ManagerVerifyCreateExecute(r ApiOrderApiV1ManagerVerifyCreateRequest) (*OrderDetail, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrderDetail
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrderApiV1ManagerVerifyCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order/api/v1/manager/{order_uuid}/verify/"
	localVarPath = strings.Replace(localVarPath, "{"+"order_uuid"+"}", url.PathEscape(parameterValueToString(r.orderUuid, "orderUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.verifyOrder == nil {
		return localVarReturnValue, nil, reportError("verifyOrder is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.verifyOrder
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["MerchantAPIKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiWalletsApiV1WalletBalanceRetrieveRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
}

func (r ApiWalletsApiV1WalletBalanceRetrieveRequest) Execute() (*WalletBalance, *http.Response, error) {
	return r.ApiService.WalletsApiV1WalletBalanceRetrieveExecute(r)
}

/*
WalletsApiV1WalletBalanceRetrieve دریافت موجودی کیف پول

<div dir="rtl" style="text-align: right;">

موجودی کیف پول فروشنده

## توضیحات

این endpoint موجودی کیف پول فروشنده را برمی‌گرداند. کیف پول برای پرداخت هزینه ارسال دیجی‌اکسپرس استفاده می‌شود. هنگام ثبت مرسوله دیجی‌اکسپرس، هزینه ارسال به‌صورت خودکار از کیف پول کسر می‌شود.

نیاز به **API_KEY** فروشنده دارد.

</div>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWalletsApiV1WalletBalanceRetrieveRequest
*/
func (a *DefaultApiService) WalletsApiV1WalletBalanceRetrieve(ctx context.Context) ApiWalletsApiV1WalletBalanceRetrieveRequest {
	return ApiWalletsApiV1WalletBalanceRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WalletBalance
func (a *DefaultApiService) WalletsApiV1WalletBalanceRetrieveExecute(r ApiWalletsApiV1WalletBalanceRetrieveRequest) (*WalletBalance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WalletBalance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.WalletsApiV1WalletBalanceRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wallets/api/v1/wallet-balance/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["MerchantAPIKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
