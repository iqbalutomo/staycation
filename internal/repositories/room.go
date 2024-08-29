package repository

import (
	model "staycation/internal/models"
	database "staycation/pkg/databases"
	"time"

	"gorm.io/gorm"
)

type RoomRepository interface {
	Create(room *model.Room) error
	FindByID(id uint) (*model.Room, error)
	FindRoomsToUpdate(checkOutDate time.Time) ([]model.Room, error)
	UpdateRoom(room *model.Room) error
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

func (r *roomRepo) FindByID(id uint) (*model.Room, error) {
	var room model.Room
	if err := r.db.First(&room, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &room, nil
}

func (r *roomRepo) FindRoomsToUpdate(checkOutDate time.Time) ([]model.Room, error) {
	var rooms []model.Room
	err := r.db.Where("check_out_date < ? AND status != ?", checkOutDate, "available").Find(&rooms).Error
	return rooms, err
}

func (r *roomRepo) UpdateRoom(room *model.Room) error {
	return r.db.Save(room).Error
}
