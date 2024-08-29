package config

import (
	"log"
	"os"
	database "staycation/pkg/databases"

	"github.com/joho/godotenv"
)

var (
	Port         string
	XenditAPIKey string
	XenditAPIURL string
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
	XenditAPIKey = os.Getenv("XENDIT_API_KEY")
	XenditAPIURL = os.Getenv("XENDIT_API_URL")
}
