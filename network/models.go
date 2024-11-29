package network

import validation "github.com/go-ozzo/ozzo-validation"

type PayerInfo struct {
	PayerID string `json:"payerID"`
}

type CreateRequestPayout struct {
	ClientOrderID string    `json:"clientOrderID"`
	Sum           int       `json:"sum"`
	Ttl           int       `json:"ttl"`
	Message       string    `json:"message"`
	Type          string    `json:"type"`
	WalletID      int       `json:"walletID"`
	WebhookUrl    string    `json:"webhookUrl"`
	CardNumber    string    `json:"cardNumber"`
	PayerInfo     PayerInfo `json:"payerInfo"`
}

type CreateInvoiceP2P struct {
	ClientOrderID string `json:"clientOrderId"`
	ClientIP      string `json:"clientIp"`
	PayerID       string `json:"payerId"`
	Sum           int    `json:"sum"`
	WalletID      int    `json:"walletId"`
	Bank          string `json:"bank"`
	Message       string `json:"message"`
	WebhookUrl    string `json:"webhookUrl"`
	ExpireAt      int    `json:"expireAt"`
}

type CreateInvoiceFPS struct {
	ClientOrderID string `json:"clientOrderID"`
	ClientIP      string `json:"clientIP"`
	PayerID       string `json:"payerID"`
	CurrencyID    int    `json:"currencyID"`
	Amount        int    `json:"amount"`
	Comment       string `json:"comment"`
	WebhookURL    string `json:"webhookURL"`
	ExpireAt      int    `json:"expireAt"`
}

type RequestPostmanInvoice struct {
	Type string `json:"type"`
	Sum  int    `json:"sum"`
}

type RequestPostmanPayout struct {
	Type string `json:"type"`
	Sum  int    `json:"sum"`
	Card string `json:"card"`
}

func (r RequestPostmanInvoice) Validate() error {
	err := validation.ValidateStruct(&r,
		validation.Field(&r.Type, validation.Required),
		validation.Field(&r.Sum, validation.Required))
	return err
}

func (r RequestPostmanPayout) Validate() error {
	err := validation.ValidateStruct(&r,
		validation.Field(&r.Type, validation.Required),
		validation.Field(&r.Sum, validation.Required),
		validation.Field(&r.Card, validation.Required))
	return err
}

type P2PInvoiceResponse struct {
	Success        bool   `json:"success"`
	Id             int    `json:"id"`
	Uuid           string `json:"uuid"`
	Sum            int    `json:"sum"`
	InitSum        int    `json:"initSum"`
	BankName       string `json:"bank_Name"`
	Card           string `json:"card"`
	CardHolderName string `json:"cardHolderName"`
}

type FPSInvoiceResponse struct {
	Success        bool   `json:"success"`
	ExternalID     string `json:"externalID"`
	PhoneNumber    string `json:"phoneNumber"`
	BankName       string `json:"bankName"`
	CardHolderName string `json:"cardHolderName"`
	Amount         int    `json:"amount"`
	InitAmount     int    `json:"initAmount"`
}

type PayoutResponse struct {
	Id      int    `json:"id"`
	Status  string `json:"status"`
	Success bool   `json:"success"`
	Uuid    string `json:"uuid"`
}

type FinalResponsePayout struct {
	Id      int    `json:"id"`
	Uuid    string `json:"uuid"`
	Success bool   `json:"success"`
}

type FinalResponseInvoice struct {
	Uuid           string `json:"uuid"`
	Card           string `json:"card"`
	Sum            int    `json:"sum"`
	CardHolderName string `json:"cardHolderName"`
	BankName       string `json:"bankName"`
}

type FinalResponseError struct {
	Error string `json:"error"`
}
