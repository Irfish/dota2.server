package gin

const(
	ORDER_RESPONSE_STATUS_CREATED = "CREATED"
	ORDER_RESPONSE_STATUS_SAVED = "SAVED"
	ORDER_RESPONSE_STATUS_APPROVED = "APPROVED"
	ORDER_RESPONSE_STATUS_VOIDED = "VOIDED"
	ORDER_RESPONSE_STATUS_COMPLETED = "COMPLETED"
	INTENT_CAPTURE ="CAPTURE"
	INTENT_AUTHORIZE ="AUTHORIZE"
)


type PayPalAccess struct{
	Scope string `json:"scope"`
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	AppID string `json:"app_id"`
	ExpiresIn int64 `json:"expires_in"`
	Nonce string `json:"nonce"`
}

type  AmountData struct {
	CurrencyCode string `json:"currency_code"`
	Value float32 `json:"value"`
}

type  PurchaseUnits struct {
	Amount  AmountData `json:"amount"`
	ReferenceId string `json:"reference_id"`
	Shipping Shipping `json:"shipping"`
	Payments  CapturePayment `json:"payments"`
}

type  OrderReq struct {
	Intent string `json:"intent"`
	PurchaseUnits []PurchaseUnits `json:"purchase_units"`
}

type link struct {
	Href string `json:"href"`
	Rel string `json:"rel"`
	Method string `json:"method"`
}

type OrderResp struct {
	Id string `json:"id"`
	Status string `json:"status"`
	Links []link `json:"links"`
}

type PayerName struct {
	GivenName string `json:"given_name"`
	Surname string `json:"surname"`
}

type Payer struct {
	Name PayerName `json:"name"`
	EmailAddress string `json:"email_address"`
	PayerId string `json:"payer_id"`
}

type ShippingAddress struct {
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	AdminArea2 string `json:"admin_area_2"`
	AdminArea1 string `json:"admin_area_1"`
	PostalCode string `json:"postal_code"`
	CountryCode string `json:"country_code"`
}

type Shipping struct {
	Address  ShippingAddress `json:"address"`
}

type SellerProtection struct {
	Status string `json:"status"`
	DisputeCategories []string `json:"dispute_categories"`
}

type SellerReceivableBreakdown struct {
	GrossAmount AmountData `json:"gross_amount"`
	PaypalFee AmountData `json:"paypal_fee"`
	NetAmount AmountData `json:"net_amount"`
}

type Capture struct {
	Id string `json:"id"`
	Status string `json:"status"`
	Amount AmountData `json:"amount"`
	SellerProtection SellerProtection `json:"seller_protection"`
	FinalCapture bool `json:"final_capture"`
	DisbursementMode string `json:"disbursement_mode"`
	SellerReceivableBreakdown SellerReceivableBreakdown `json:"seller_receivable_breakdown"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	Links []link `json:"links"`
}

type  CapturePayment struct {
	Captures []Capture 	`json:"captures"`
}

type CaptureOrderResp struct {
	Id string `json:"id"`
	Status string `json:"status"`
	Payer Payer `json:"payer"`
	PurchaseUnits []PurchaseUnits `json:"purchase_units"`
	Links []link `json:"links"`
}
