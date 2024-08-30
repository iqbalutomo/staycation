package email_mailtrap

import (
	"encoding/json"
	"fmt"
	"staycation/config"
	helper "staycation/pkg/third_parties/helpers"
	model_mailtrap "staycation/pkg/third_parties/mailtrap/models"
)

func SendEmailRegister(toEmail, name string) error {
	url := config.MailtrapAPIURL + "/api/send"

	payload := model_mailtrap.EmailPayload{
		From: model_mailtrap.EmailAddress{
			Email: config.MailtrapSender,
			Name:  config.MailtrapName,
		},
		To: []model_mailtrap.EmailAddress{
			{
				Email: toEmail,
			},
		},
		Subject:  fmt.Sprintf("Welcome to Staycation, %s!", name),
		Text:     fmt.Sprintf("Hey %s,\n\n🎉 Welcome to Staycation 🎉 Your registration was successful, and lets go booking hotel to best your enjoyed holiday 🏨🏖️😎\n\nCheers,\nThe Staycation Team", name),
		Category: "Registration Success",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	headers := map[string]string{
		"Authorization": "Bearer " + config.MailtrapAPIKey,
		"Content-Type":  "application/json",
	}

	response, err := helper.FetchAPI(url, "POST", headers, payloadBytes)
	if err != nil {
		return err
	}

	fmt.Println("Email sent successfully:", response)
	return nil
}
