package config

import (
	"log"
	"os"
	database "staycation/pkg/databases"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	XenditAPIKey string
}

func LoadEnv() error {
	return godotenv.Load()
}

func InitConfig() *Config {
	if err := LoadEnv(); err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}

	database.ConnectDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Port:         port,
		XenditAPIKey: os.Getenv("XENDIT_API_KEY"),
	}
}
