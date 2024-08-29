package repository

import (
	"errors"
	model "staycation/internal/models"
	database "staycation/pkg/databases"

	"gorm.io/gorm"
)

type HotelRepository interface {
	Create(hotel *model.Hotel) error
	FindAll(limit, offset int) ([]*model.Hotel, error)
	Update(hotel *model.Hotel) error
	Delete(hotelID int) error

	FindByEmail(email string) (*model.Hotel, error)
	FindByPhone(phone string) (*model.Hotel, error)
	FindByID(hotelID uint) (*model.Hotel, error)
}

type hotelRepo struct {
	db *gorm.DB
}

func NewHotelRepository() HotelRepository {
	return &hotelRepo{db: database.DB}
}

func (r *hotelRepo) Create(hotel *model.Hotel) error {
	result := r.db.Create(&hotel)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (r *hotelRepo) FindAll(limit, offset int) ([]*model.Hotel, error) {
	var hotels []*model.Hotel
	if err := r.db.Limit(limit).Offset(offset).Find(&hotels).Error; err != nil {
		return nil, err
	}

	return hotels, nil
}

func (r *hotelRepo) Update(hotel *model.Hotel) error {
	if err := r.db.Save(&hotel).Error; err != nil {
		return err
	}

	return nil
}

func (r *hotelRepo) Delete(hotelID int) error {
	var hotel model.Hotel
	if err := r.db.Where("id = ?", hotelID).Delete(&hotel).Error; err != nil {
		return err
	}

	return nil
}

func (r *hotelRepo) FindByID(hotelID uint) (*model.Hotel, error) {
	var hotel model.Hotel
	if err := r.db.Where("id = ?", hotelID).First(&hotel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &hotel, nil
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
