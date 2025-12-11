package payment

import (
	"order-service/internal/primitive/enum"
	"encoding/json"
)

type TransactionStatus string

var (
	TransactionStatusCapture       TransactionStatus = "capture"
	TransactionStatusAuthorize     TransactionStatus = "authorize"
	TransactionStatusSettlement    TransactionStatus = "settlement"
	TransactionStatusPending       TransactionStatus = "pending"
	TransactionStatusExpire        TransactionStatus = "expire"
	TransactionStatusDeny          TransactionStatus = "deny"
	TransactionStatusCancel        TransactionStatus = "cancel"
	TransactionStatusRefund        TransactionStatus = "refund"
	TransactionStatusPartialRefund TransactionStatus = "partial_refund"
	TransactionStatusFailure       TransactionStatus = "failure"
	TransactionStatusChargeback    TransactionStatus = "chargeback"
)

type ChargeRequest struct {
	PaymentType       enum.PaymentType  `json:"payment_type" validate:"required"`
	TransactionDetail TransactionDetail `json:"transaction_details" validate:"required"`
	CustomExpiry      *CustomExpiry     `json:"custom_expiry,omitempty" validate:"omitempty"`
	ItemDetails       []ItemDetail      `json:"item_details,omitempty" validate:"omitempty,dive"`
	CustomerDetail    *CustomerDetail   `json:"customer_details,omitempty" validate:"omitempty"`
	SellerDetail      *SellerDetail     `json:"seller_details,omitempty" validate:"omitempty"`

	// payment
	CreditCard *CreditCard `json:"credit_card,omitempty" validate:"omitempty"`
	QRIS       *QRIS       `json:"qris,omitempty" validate:"omitempty"`
	ShopeePay  *ShopeePay  `json:"shopeepay,omitempty" validate:"omitempty"`
	GoPay      *GoPay      `json:"gopay,omitempty" validate:"omitempty"`

	MetaData map[string]string `json:"meta_data,omitempty"`
}

type ChargeOutput struct {
	ResponseSuccess MidtransResponse

	ChargeRequestBody json.RawMessage
}

type PaymentStatusOutput struct {
	OrderId           uint64
	Status            string
	StatusCode        string
	TransactionStatus string
	TransactionID     string
}

type MidtransStatusResponse struct {
	StatusCode        string `json:"status_code"`
	TransactionID     string `json:"transaction_id"`
	GrossAmount       string `json:"gross_amount"`
	Currency          string `json:"currency"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	TransactionStatus string `json:"transaction_status"`
	FraudStatus       string `json:"fraud_status"`
	StatusMessage     string `json:"status_message"`
	TransactionTime   string `json:"transaction_time"`
	ExpiryTime        string `json:"expiry_time"`
}

// GOPAY
// REF https://docs.midtrans.com/reference/gopay-object
type GoPay struct {
	EnableCallback bool   `json:"enable_callback,omitempty" validate:"omitempty"`
	CallbackURL    string `json:"callback_url,omitempty" validate:"omitempty,url"`

	// for gopay token
	AccountID          string   `json:"account_id,omitempty" validate:"omitempty"` // removed uuid4
	PaymentOptionToken string   `json:"payment_option_token,omitempty" validate:"omitempty"`
	PreAuth            bool     `json:"pre_auth,omitempty" validate:"omitempty"`
	Recurring          bool     `json:"recurring,omitempty" validate:"omitempty"`
	PromotionIDs       []string `json:"promotion_ids,omitempty" validate:"omitempty,dive,required"`
}

// CREDIT CARD
// REF https://docs.midtrans.com/reference/credit-card-object
type CreditCard struct {
	TokenID         string   `json:"token_id" validate:"required"`
	Bank            string   `json:"bank,omitempty" validate:"omitempty,oneof=mandiri bni cimb bca maybank bri"`
	InstallmentTerm int      `json:"installment_term,omitempty" validate:"omitempty,min=1"`
	Bins            []string `json:"bins,omitempty" validate:"omitempty,dive,numeric,min=6,max=8"`
	Type            string   `json:"type,omitempty" validate:"omitempty,oneof=authorize"`
	SaveTokenID     bool     `json:"save_token_id,omitempty" validate:"omitempty"`
	Channel         string   `json:"channel,omitempty" validate:"omitempty,oneof=dragon mti migs cybersource braintree mpgs"`
}

// SHOPEEPAY
// REF https://docs.midtrans.com/reference/shopeepay-object
type ShopeePay struct {
	CallbackURL string `json:"callback_url" validate:"required,url"`
}

// QRIS
// REF https://docs.midtrans.com/reference/qris-object
type QRIS struct {
	Acquirer string `json:"acquirer,omitempty" validate:"omitempty,oneof=airpay shopee gopay"`
}

// VIRTUAL ACCOUNT
// REF https://docs.midtrans.com/reference/bank-transfer-object
type BankTransfer struct {
	Bank     string          `json:"bank" validate:"required,oneof=permata bni bri bca cimb"`
	VANumber string          `json:"va_number,omitempty" validate:"omitempty,min=6,max=18"`
	FreeText *FreeTextObject `json:"free_text,omitempty"`
	BCA      *BCAOptions     `json:"bca,omitempty" validate:"required_if=Bank bca"`
	Permata  *PermataOptions `json:"permata,omitempty" validate:"required_if=Bank permata"`
}
type FreeTextObject struct {
	Inquiry []InquiryPayment `json:"inquiry,omitempty" validate:"omitempty,dive,max=10"`
	Payment []InquiryPayment `json:"payment,omitempty" validate:"omitempty,dive,max=10"`
}
type InquiryPayment struct {
	ID string `json:"id" validate:"required,max=50"` // Free text message in Bahasa Indonesia
	EN string `json:"en" validate:"required,max=50"` // Free text message in English
}
type BCAOptions struct {
	SubCompanyCode string `json:"sub_company_code,omitempty" validate:"max=5"` // Default is 00000
}
type PermataOptions struct {
	RecipientName string `json:"recipient_name,omitempty" validate:"max=20"` // Uppercase string
}

// ==== GLOBAL RESPONSE ENVELOPE ====
// MidtransResponse: global Midtrans response for all payment types.
// This struct embeds common fields (always present) and
// optional fields that only apply to specific payment methods.
type MidtransResponse struct {
	MidtransCommonResponse

	// Only for Virtual Account (Bank Transfer).
	// Midtrans returns a list of VA numbers for each bank.
	VANumbers []BankTransferResponse `json:"va_numbers,omitempty"`

	// Only for QRIS.
	// Indicates the acquirer being used, e.g. "gopay", "shopee", "airpay".
	Acquirer string `json:"acquirer,omitempty"`

	// Typically present in GoPay responses (and sometimes ShopeePay).
	// Used to validate the integrity of the response via hashing.
	// Note: in Midtrans callbacks (notifications), this field is ALWAYS included.
	SignatureKey string `json:"signature_key,omitempty"`

	// Commonly appears in GoPay and ShopeePay responses.
	// Contains the original channel/acquirer (bank/e-wallet) response code & message.
	// Useful for debugging why a transaction failed or was denied.
	ChannelResponseCode    string `json:"channel_response_code,omitempty"`
	ChannelResponseMessage string `json:"channel_response_message,omitempty"`

	// Refund.
	// Contains refund details if the transaction has been refunded.
	Refunds      []Refund `json:"refunds,omitempty"`
	RefundAmount string   `json:"refund_amount,omitempty"`
	// ApprovalCode is the authorization code returned by the card issuer
	// for a successful credit card transaction. It uniquely identifies the
	// authorization and can be used as a reference during refunds or disputes.
	ApprovalCode string `json:"approval_code,omitempty"`

	// Credit Card.
	// Additional details for card transactions.
	PaymentOptionType string `json:"payment_option_type,omitempty"`
	CardType          string `json:"card_type,omitempty"`

	// if return error
	ValidationMessages []string `json:"validation_messages,omitempty"`
}

type MidtransCommonResponse struct {
	// Midtrans status code (e.g. "200", "201").
	StatusCode    string `json:"status_code,omitempty"`
	StatusMessage string `json:"status_message,omitempty"`

	// Transaction identity.
	TransactionID     string            `json:"transaction_id,omitempty"`   // Midtrans-generated transaction ID
	TransactionTime   string            `json:"transaction_time,omitempty"` // Transaction timestamp (YYYY-MM-DD HH:mm:ss)
	TransactionStatus TransactionStatus `json:"transaction_status"`         // capture | settlement | pending | expire | cancel | deny | refund | partial_refund | failure | chargeback
	OrderID           string            `json:"order_id"`                   // Merchant’s order ID
	GrossAmount       string            `json:"gross_amount"`               // Total transaction amount (always returned as string by Midtrans)
	PaymentType       enum.PaymentType  `json:"payment_type"`               // qris | gopay | shopeepay | credit_card | bank_transfer | echannel
	Currency          string            `json:"currency"`                   // Currency code (usually "IDR")

	// Fraud Detection System (FDS) result.
	// Only relevant for credit card and some risky payments.
	// Possible values: "accept", "challenge", "deny".
	FraudStatus string `json:"fraud_status,omitempty"`

	// Merchant identity (always present in Midtrans response).
	MerchantID string `json:"merchant_id"`

	// Action URLs provided by Midtrans (e.g. QR code link, deeplink, cancel URL).
	Actions []Action `json:"actions,omitempty"`

	// Optional metadata (custom key-value pairs).
	Metadata map[string]string `json:"metadata,omitempty"`
}

// CUSTOMER DETAIL
// REF https://docs.midtrans.com/reference/customer-details-object
type CustomerDetail struct {
	FirstName       string                   `json:"first_name,omitempty" validate:"omitempty,max=30"`
	LastName        string                   `json:"last_name,omitempty" validate:"omitempty,max=30"`
	Email           string                   `json:"email,omitempty" validate:"omitempty,email"`
	Phone           string                   `json:"phone,omitempty" validate:"omitempty,max=255"`
	BillingAddress  *CustomerBillingAddress  `json:"billing_address,omitempty" validate:"omitempty"`
	ShippingAddress *CustomerShippingAddress `json:"shipping_address,omitempty" validate:"omitempty"`
}
type CustomerBillingAddress struct {
	FirstName   string `json:"first_name,omitempty" validate:"omitempty,max=255"`
	LastName    string `json:"last_name,omitempty" validate:"omitempty,max=255"`
	Phone       string `json:"phone,omitempty" validate:"omitempty,max=255"`
	Address     string `json:"address,omitempty" validate:"omitempty,max=255"`
	City        string `json:"city,omitempty" validate:"omitempty,max=255"`
	PostalCode  string `json:"postal_code,omitempty" validate:"omitempty,max=255,alphanumunicode"`
	CountryCode string `json:"country_code,omitempty" validate:"omitempty,eq=IDN,len=3"`
}
type CustomerShippingAddress struct {
	FirstName   string `json:"first_name,omitempty" validate:"omitempty,max=255"`
	LastName    string `json:"last_name,omitempty" validate:"omitempty,max=255"`
	Phone       string `json:"phone,omitempty" validate:"omitempty,max=255"`
	Address     string `json:"address,omitempty" validate:"omitempty,max=255"`
	City        string `json:"city,omitempty" validate:"omitempty,max=255"`
	PostalCode  string `json:"postal_code,omitempty" validate:"omitempty,max=255,alphanumunicode"`
	CountryCode string `json:"country_code,omitempty" validate:"omitempty,eq=IDN,len=3"`
}

// ITEM DETAIL
// REF https://docs.midtrans.com/reference/item-details-object
type ItemDetail struct {
	ID           string `json:"id,omitempty" validate:"omitempty"`
	Name         string `json:"name" validate:"required"`
	Price        int64  `json:"price" validate:"required,gte=0"`
	Qty          int32  `json:"quantity" validate:"required,gt=0"`
	Brand        string `json:"brand,omitempty" validate:"omitempty"`
	Category     string `json:"category,omitempty" validate:"omitempty"`
	MerchantName string `json:"merchant_name,omitempty" validate:"omitempty"`
	Tenor        int    `json:"tenor,omitempty" validate:"omitempty,min=0,max=99"`
	CodePlan     int    `json:"code_plan,omitempty" validate:"omitempty,min=0,max=999"`
	MID          int    `json:"mid,omitempty" validate:"omitempty,min=0,max=999999999"`
	URL          string `json:"url,omitempty" validate:"omitempty,url"`
}

// TRANSACTION DETAIL
// REF https://docs.midtrans.com/reference/transaction-details-object
type TransactionDetail struct {
	OrderID     string `json:"order_id" validate:"required,max=50"`
	GrossAmount int64  `json:"gross_amount" validate:"required,gte=0"`
}

// CUSTOM EXPIRY
// REF https://docs.midtrans.com/reference/custom-expiry-object
type CustomExpiry struct {
	ExpiryDuration int32  `json:"expiry_duration,omitempty" validate:"omitempty,gt=0"`
	Unit           string `json:"unit,omitempty" validate:"omitempty,oneof=second minute hour"`
}

// SELLER DETAIL
// REF https://docs.midtrans.com/reference/seller-details-object
type SellerDetail struct {
	ID      string         `json:"id,omitempty"`
	Name    string         `json:"name,omitempty"`
	Email   string         `json:"email,omitempty"`
	URL     string         `json:"url,omitempty"`
	Address *SellerAddress `json:"address,omitempty"`
}
type SellerAddress struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Address     string `json:"address,omitempty"`
	City        string `json:"city,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

type Refund struct {
	RefundChargebackID int    `json:"refund_chargeback_id"`
	RefundAmount       string `json:"refund_amount"`
	CreatedAt          string `json:"created_at"`
	Reason             string `json:"reason"`
	RefundKey          string `json:"refund_key"`

	// bank confirmation
	RefundMethod    string `json:"refund_method,omitempty"`
	BankConfirmedAt string `json:"bank_confirmed_at"`
}

// VA number untuk bank transfer (kalau dipakai)
type BankTransferResponse struct {
	Bank     string `json:"bank"`
	VANumber string `json:"va_number"`
}

// Action Midtrans (disederhanakan)
type Action struct {
	Name   string            `json:"name"`
	Method string            `json:"method"`
	URL    string            `json:"url"`
	Fields map[string]string `json:"fields,omitempty"`
}
