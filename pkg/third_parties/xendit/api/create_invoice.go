package xendit

import (
	"encoding/json"
	"fmt"
	"staycation/config"
	helper "staycation/pkg/third_parties/helpers"
	model "staycation/pkg/third_parties/xendit/models"
	"time"
)

func CreateInvoice(amount float64, payerEmail, description, itemName string, quantity int, price float64) (*model.XenditInvoiceResponse, error) {
	xenditURL := config.XenditAPIURL
	apiKey := config.XenditAPIKey

	invoiceRequest := model.XenditInvoiceRequest{
		ExternalID:      fmt.Sprintf("invoice-%d", time.Now().Unix()),
		Amount:          amount,
		PayerEmail:      payerEmail,
		Description:     description,
		ShouldSendEmail: true,
		Items: []struct {
			Name        string  `json:"name"`
			Quantity    int     `json:"quantity"`
			Price       float64 `json:"price"`
			Category    string  `json:"category"`
			Description string  `json:"description"`
			URL         string  `json:"url"`
		}{
			{
				Name:        itemName,
				Quantity:    quantity,
				Price:       price,
				Category:    "Hotel",
				Description: "Thanks for booked:)",
				URL:         "wait to deploy...............",
			},
		},
	}

	jsonBody, err := json.Marshal(invoiceRequest)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Authorization": "Basic " + helper.BasicAuth(apiKey, ""),
		"Content-Type":  "application/json",
	}

	response, err := helper.FetchAPI(xenditURL, "POST", headers, jsonBody)
	if err != nil {
		return nil, err
	}

	var invoiceResponse model.XenditInvoiceResponse
	if err := json.Unmarshal([]byte(response), &invoiceResponse); err != nil {
		return nil, err
	}

	return &invoiceResponse, nil
}
