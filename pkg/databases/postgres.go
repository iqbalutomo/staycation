package database

import (
	"fmt"
	"log"
	"os"
	model "staycation/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := DB.AutoMigrate(&model.User{}, &model.Balance{}, &model.Hotel{}, &model.RoomType{}, &model.RoomBedType{}, &model.RoomFacilities{}, &model.Room{}, &model.Booking{}, &model.Invoice{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
