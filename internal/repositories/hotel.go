package repository

import (
	"errors"
	model "staycation/internal/models"
	database "staycation/pkg/databases"

	"gorm.io/gorm"
)

type HotelRepository interface {
	CreateHotel(hotel *model.Hotel) error
	FindByEmail(email string) (*model.Hotel, error)
	FindByPhone(phone string) (*model.Hotel, error)
}

type hotelRepo struct {
	db *gorm.DB
}

func NewHotelRepository() HotelRepository {
	return &hotelRepo{db: database.DB}
}

func (r *hotelRepo) CreateHotel(hotel *model.Hotel) error {
	result := r.db.Create(&hotel)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (r *hotelRepo) FindByEmail(email string) (*model.Hotel, error) {
	var hotel model.Hotel
	if err := r.db.Where("email = ?", email).First(&hotel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &hotel, nil
}

func (r *hotelRepo) FindByPhone(phone string) (*model.Hotel, error) {
	var hotel model.Hotel
	if err := r.db.Where("phone = ?", phone).First(&hotel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &hotel, nil
}
