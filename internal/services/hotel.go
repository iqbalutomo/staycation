package service

import (
	"errors"
	model "staycation/internal/models"
	repository "staycation/internal/repositories"
)

type HotelService interface {
	NewHotel(req *model.Hotel) (*model.Hotel, error)
	NewRoomType(userID float64, roomType *model.RoomType, bedType *model.RoomBedType, facilities *model.RoomFacilities) (*model.RoomTypeRequest, error)
}

type hotelService struct {
	repo repository.HotelRepository
}

func NewHotelService(repo repository.HotelRepository) HotelService {
	return &hotelService{repo}
}

func (s *hotelService) NewHotel(req *model.Hotel) (*model.Hotel, error) {
	emailExist, err := s.repo.FindHotelByEmail(req.Email)
	if err != nil {
		return nil, err
	} else if emailExist != nil {
		return nil, errors.New("email_exist")
	}

	phoneExist, err := s.repo.FindHotelByPhone(req.Phone)
	if err != nil {
		return nil, err
	} else if phoneExist != nil {
		return nil, errors.New("phone_exist")
	}

	if err := s.repo.CreateHotel(req); err != nil {
		return nil, err
	}

	return req, nil
}

func (s *hotelService) NewRoomType(userID float64, roomType *model.RoomType, bedType *model.RoomBedType, facilities *model.RoomFacilities) (*model.RoomTypeRequest, error) {
	hotelExist, err := s.repo.FindHotelByID(roomType.HotelID)
	if err != nil {
		return nil, err
	} else if hotelExist == nil {
		return nil, errors.New("hotel not found")
	}

	if hotelExist.ID != roomType.HotelID {
		return nil, errors.New("hotel not found")
	}

	if hotelExist.OwnerID != uint(userID) {
		return nil, errors.New("invalid credentials")
	}

	return s.repo.CreateRoomType(roomType, bedType, facilities)
}
