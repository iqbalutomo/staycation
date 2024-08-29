package model

import "time"

type XenditInvoiceRequest struct {
	ExternalID      string  `json:"external_id"`
	Amount          float64 `json:"amount"`
	PayerEmail      string  `json:"payer_email"`
	Description     string  `json:"description"`
	ShouldSendEmail bool    `json:"should_send_email"`
	Items           []struct {
		Name        string  `json:"name"`
		Quantity    int     `json:"quantity"`
		Price       float64 `json:"price"`
		Category    string  `json:"category"`
		Description string  `json:"description"`
		URL         string  `json:"url"`
	} `json:"items"`
}

type XenditInvoiceResponse struct {
	ID          string    `json:"id"`
	ExternalID  string    `json:"external_id"`
	Status      string    `json:"status"`
	Amount      float64   `json:"amount"`
	PayerEmail  string    `json:"payer_email"`
	Description string    `json:"description"`
	ExpiryDate  time.Time `json:"expiry_date"`
	InvoiceURL  string    `json:"invoice_url"`
	Currency    string    `json:"currency"`
	Items       []struct {
		Name     string  `json:"name"`
		Quantity int     `json:"quantity"`
		Price    float64 `json:"price"`
		Category string  `json:"category"`
		URL      string  `json:"url"`
	} `json:"items"`
	Customer struct {
		Email string `json:"email"`
	} `json:"customer"`
}

type XenditInvoiceWebhookPayload struct {
	ID                 string  `json:"id"`
	ExternalID         string  `json:"external_id"`
	UserID             string  `json:"user_id"`
	PaymentMethod      string  `json:"payment_method"`
	Status             string  `json:"status"`
	Amount             float64 `json:"amount"`
	PaidAmount         float64 `json:"paid_amount"`
	BankCode           string  `json:"bank_code"`
	PaidAt             string  `json:"paid_at"`
	PayerEmail         string  `json:"payer_email"`
	Description        string  `json:"description"`
	Created            string  `json:"created"`
	Updated            string  `json:"updated"`
	Currency           string  `json:"currency"`
	PaymentChannel     string  `json:"payment_channel"`
	PaymentDestination string  `json:"payment_destination"`
}
