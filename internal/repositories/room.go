package repository

import (
	model "staycation/internal/models"
	database "staycation/pkg/databases"

	"gorm.io/gorm"
)

type RoomRepository interface {
	Create(room *model.Room) error
}

type roomRepo struct {
	db *gorm.DB
}

func NewRoomRepository() RoomRepository {
	return &roomRepo{db: database.DB}
}

func (r *roomRepo) Create(room *model.Room) error {
	result := r.db.Create(&room)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}
