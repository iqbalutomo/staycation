package repository

import (
	"errors"
	model "staycation/internal/models"
	database "staycation/pkg/databases"

	"gorm.io/gorm"
)

type HotelRepository interface {
	// HOTEL
	CreateHotel(hotel *model.Hotel) error
	FindHotelByEmail(email string) (*model.Hotel, error)
	FindHotelByPhone(phone string) (*model.Hotel, error)
	FindHotelByID(hotelID uint) (*model.Hotel, error)

	// ROOM TYPE
	CreateRoomType(roomType *model.RoomType, bedType *model.RoomBedType, facilities *model.RoomFacilities) (*model.RoomTypeRequest, error)
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

func (r *hotelRepo) FindHotelByEmail(email string) (*model.Hotel, error) {
	var hotel model.Hotel
	if err := r.db.Where("email = ?", email).First(&hotel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &hotel, nil
}

func (r *hotelRepo) FindHotelByPhone(phone string) (*model.Hotel, error) {
	var hotel model.Hotel
	if err := r.db.Where("phone = ?", phone).First(&hotel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &hotel, nil
}

func (r *hotelRepo) FindHotelByID(hotelID uint) (*model.Hotel, error) {
	var hotel model.Hotel
	if err := r.db.Where("id = ?", hotelID).First(&hotel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &hotel, nil
}

func (r *hotelRepo) CreateRoomType(roomType *model.RoomType, bedType *model.RoomBedType, facilities *model.RoomFacilities) (*model.RoomTypeRequest, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&roomType).Error; err != nil {
			return err
		}

		bedType.RoomTypeID = roomType.ID
		if err := tx.Create(&bedType).Error; err != nil {
			return err
		}

		facilities.RoomTypeID = roomType.ID
		if err := tx.Create(facilities).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	respData := model.RoomTypeRequest{
		RoomType:       *roomType,
		RoomBedType:    *bedType,
		RoomFacilities: *facilities,
	}

	return &respData, nil
}
