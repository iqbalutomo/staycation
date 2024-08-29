package service

import (
	"errors"
	model "staycation/internal/models"
	repository "staycation/internal/repositories"
)

type HotelService interface {
	CreateHotel(req *model.Hotel) (*model.Hotel, error)
	GetHotels(limit, offset int) ([]*model.Hotel, error)
	UpdateHotel(userID float64, hotel *model.Hotel) (*model.Hotel, error)
	DeleteHotel(hotelID int) error

	FindHotelByID(hotelID uint) (*model.Hotel, error)
}

type hotelService struct {
	repo repository.HotelRepository
}

func NewHotelService(repo repository.HotelRepository) HotelService {
	return &hotelService{repo}
}

func (s *hotelService) CreateHotel(req *model.Hotel) (*model.Hotel, error) {
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

	if err := s.repo.Create(req); err != nil {
		return nil, err
	}

	return req, nil
}

func (s *hotelService) GetHotels(limit, offset int) ([]*model.Hotel, error) {
	hotels, err := s.repo.FindAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return hotels, nil
}

func (s *hotelService) UpdateHotel(userID float64, hotel *model.Hotel) (*model.Hotel, error) {
	if hotel.OwnerID != uint(userID) {
		return nil, errors.New("invalid_credentials")
	}

	hotelExist, err := s.repo.FindByID(hotel.ID)
	if err != nil {
		return nil, err
	} else if hotelExist == nil {
		return nil, errors.New("hotel_not_found")
	}

	hotelExist.Name = hotel.Name
	hotelExist.Description = hotel.Description
	hotelExist.Address = hotel.Address
	hotelExist.City = hotel.City
	hotelExist.Zipcode = hotel.Zipcode
	hotelExist.Country = hotel.Country
	hotelExist.Phone = hotel.Phone
	hotelExist.Email = hotel.Email
	hotelExist.Star = hotel.Star

	if err := s.repo.Update(hotel); err != nil {
		return nil, err
	}

	return hotel, nil
}

func (s *hotelService) DeleteHotel(hotelID int) error {
	if err := s.repo.Delete(hotelID); err != nil {
		return err
	}

	return nil
}

func (s *hotelService) FindHotelByID(hotelID uint) (*model.Hotel, error) {
	hotel, err := s.repo.FindByID(hotelID)
	if err != nil {
		return nil, err
	} else if hotel == nil {
		return nil, errors.New("hotel_not_found")
	}

	return hotel, nil
}
