package config

import (
	"log"
	"os"
	database "staycation/pkg/databases"

	"github.com/joho/godotenv"
)

var (
	Port           string
	AppURL         string
	XenditAPIKey   string
	XenditAPIURL   string
	MailtrapAPIKey string
	MailtrapAPIURL string
	MailtrapSender string
	MailtrapName   string
)

func LoadEnv() error {
	return godotenv.Load()
}

func InitConfig() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "development" {
		if err := LoadEnv(); err != nil {
			log.Fatalf("failed to load .env file: %v", err)
		}
	}

	database.ConnectDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	Port = port
	AppURL = os.Getenv("APP_URL")
	XenditAPIKey = os.Getenv("XENDIT_API_KEY")
	XenditAPIURL = os.Getenv("XENDIT_API_URL")
	MailtrapAPIKey = os.Getenv("MAILTRAP_API_KEY")
	MailtrapAPIURL = os.Getenv("MAILTRAP_API_URL")
	MailtrapSender = os.Getenv("MAILTRAP_SENDER")
	MailtrapName = os.Getenv("MAILTRAP_NAME")
}
