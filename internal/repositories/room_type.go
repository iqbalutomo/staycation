package repository

import (
	"errors"
	model "staycation/internal/models"
	database "staycation/pkg/databases"

	"gorm.io/gorm"
)

type RoomTypeRepository interface {
	Create(roomType *model.RoomType, bedType *model.RoomBedType, facilities *model.RoomFacilities) (*model.RoomTypeRequest, error)
	FindByID(roomTypeID uint) (*model.RoomType, error)
}

type roomtTypeRepo struct {
	db *gorm.DB
}

func NewRoomTypeRepository() RoomTypeRepository {
	return &roomtTypeRepo{db: database.DB}
}

func (r *roomtTypeRepo) Create(roomType *model.RoomType, bedType *model.RoomBedType, facilities *model.RoomFacilities) (*model.RoomTypeRequest, error) {
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

func (r *roomtTypeRepo) FindByID(roomTypeID uint) (*model.RoomType, error) {
	var roomtype model.RoomType
	if err := r.db.Where("id = ?", roomTypeID).First(&roomtype).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &roomtype, nil
}
