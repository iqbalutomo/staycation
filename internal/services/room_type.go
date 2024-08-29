package service

import (
	"errors"
	model "staycation/internal/models"
	repository "staycation/internal/repositories"
)

type RoomTypeService interface {
	CreateRoomType(userID float64, roomType *model.RoomType, bedType *model.RoomBedType, facilities *model.RoomFacilities) (*model.RoomTypeRequest, error)
}

type roomTypeService struct {
	repoRoomType repository.RoomTypeRepository
	repoHotel    repository.HotelRepository
}

func NewRoomTypeService(repoRoomType repository.RoomTypeRepository, repoHotel repository.HotelRepository) RoomTypeService {
	return &roomTypeService{repoRoomType, repoHotel}
}

func (s *roomTypeService) CreateRoomType(userID float64, roomType *model.RoomType, bedType *model.RoomBedType, facilities *model.RoomFacilities) (*model.RoomTypeRequest, error) {
	hotelExist, err := s.repoHotel.FindByID(roomType.HotelID)
	if err != nil {
		return nil, err
	} else if hotelExist == nil {
		return nil, errors.New("hotel_not_found")
	}

	if hotelExist.ID != roomType.HotelID {
		return nil, errors.New("hotel_not_found")
	}

	if hotelExist.OwnerID != uint(userID) {
		return nil, errors.New("invalid_credentials")
	}

	return s.repoRoomType.Create(roomType, bedType, facilities)
}
