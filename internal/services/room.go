package service

import (
	"errors"
	model "staycation/internal/models"
	repository "staycation/internal/repositories"
)

type RoomService interface {
	CreateRoom(userID float64, req *model.Room) (*model.Room, error)
}

type roomService struct {
	repoRoom     repository.RoomRepository
	repoRoomType repository.RoomTypeRepository
	repoHotel    repository.HotelRepository
}

func NewRoomService(repoRoom repository.RoomRepository, repoRoomType repository.RoomTypeRepository, repoHotel repository.HotelRepository) RoomService {
	return &roomService{repoRoom, repoRoomType, repoHotel}
}

func (s *roomService) CreateRoom(userID float64, req *model.Room) (*model.Room, error) {
	roomTypeExist, err := s.repoRoomType.FindByID(req.RoomTypeID)
	if err != nil {
		return nil, err
	} else if roomTypeExist == nil {
		return nil, errors.New("roomtype_not_found")
	}

	hotelExist, err := s.repoHotel.FindByID(roomTypeExist.HotelID)
	if err != nil {
		return nil, err
	} else if hotelExist == nil {
		return nil, errors.New("roomtype_not_found")
	}

	if roomTypeExist.ID != req.RoomTypeID {
		return nil, errors.New("roomtype_not_found")
	}

	if hotelExist.OwnerID != uint(userID) {
		return nil, errors.New("invalid_credentials")
	}

	room := &model.Room{
		RoomTypeID: req.RoomTypeID,
		RoomNumber: req.RoomNumber,
		Status:     model.RoomStatusEnum(req.Status),
	}

	if err := s.repoRoom.Create(room); err != nil {
		return nil, err
	}

	return room, nil
}
