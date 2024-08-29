package service

import (
	"errors"
	model "staycation/internal/models"
	repository "staycation/internal/repositories"
)

type HotelService interface {
	NewHotel(req *model.Hotel) (*model.Hotel, error)
}

type hotelService struct {
	repo repository.HotelRepository
}

func NewHotelService(repo repository.HotelRepository) HotelService {
	return &hotelService{repo}
}

func (s *hotelService) NewHotel(req *model.Hotel) (*model.Hotel, error) {
	emailExist, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	} else if emailExist != nil {
		return nil, errors.New("email_exist")
	}

	phoneExist, err := s.repo.FindByPhone(req.Phone)
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
