package service

import (
	"errors"
	"fmt"
	model "staycation/internal/models"
	repository "staycation/internal/repositories"
	xendit "staycation/pkg/third_parties/xendit/api"
)

type InvoiceService interface {
	BookRoom(userID uint, email string, booking *model.Booking) (*model.BookingResponse, error)
}

type invoiceService struct {
	repoInvoice  repository.InvoiceRepository
	repoRoom     repository.RoomRepository
	roomTypeRepo repository.RoomTypeRepository
	repoBalance  repository.BalanceRepository
}

func NewInvoiceService(repoInvoice repository.InvoiceRepository, repoRoom repository.RoomRepository, roomTypeRepo repository.RoomTypeRepository, repoBalance repository.BalanceRepository) InvoiceService {
	return &invoiceService{repoInvoice, repoRoom, roomTypeRepo, repoBalance}
}

// invoice *model.Invoice
func (s *invoiceService) BookRoom(userID uint, email string, booking *model.Booking) (*model.BookingResponse, error) {
	room, err := s.repoRoom.FindByID(booking.RoomID)
	if err != nil {
		return nil, err
	} else if room == nil {
		return nil, errors.New("room_not_found")
	}

	// check room is already
	bookings, err := s.repoInvoice.FindByRoomAndDate(booking.RoomID, booking.CheckInDate, booking.CheckOutDate)
	if err != nil {
		return nil, err
	}
	for _, b := range bookings {
		if b.CheckInDate.Before(booking.CheckOutDate) && b.CheckOutDate.After(booking.CheckInDate) {
			return nil, errors.New("already-booked")
		}
	}

	// validate check-in and check-out time
	if booking.CheckInDate.Hour() != 14 || booking.CheckOutDate.Hour() != 12 {
		return nil, errors.New("validate-time")
	}

	duration := booking.CheckOutDate.Sub(booking.CheckInDate)
	days := int(duration.Hours() / 24)
	if duration.Hours() > float64(days*24) {
		days++ // roundup if additional time
	}

	if days <= 0 || duration.Hours() < 22 {
		return nil, errors.New("booking-duration")
	}

	roomType, err := s.roomTypeRepo.FindByID(room.RoomTypeID)
	if err != nil {
		return nil, err
	}
	if roomType == nil {
		return nil, errors.New("room type not found")
	}

	hotelID := roomType.HotelID
	totalPrice := roomType.Price * float64(days)

	booking.UserID = userID
	booking.HotelID = hotelID
	booking.TotalPrice = totalPrice

	if err := s.repoInvoice.CreateBooking(booking); err != nil {
		return nil, err
	}

	invoiceResponse, err := xendit.CreateInvoice(totalPrice, email, fmt.Sprintf("Invoice Staycation #%d", booking.ID), fmt.Sprintf("%s - %d", roomType.Name, room.RoomNumber), days, roomType.Price)
	if err != nil {
		return nil, err
	}

	invoice := &model.Invoice{
		BookingID:       booking.ID,
		XenditInvoiceID: invoiceResponse.ID,
		InvoiceURL:      invoiceResponse.InvoiceURL,
		Amount:          totalPrice,
		Status:          invoiceResponse.Status,
	}

	if err := s.repoInvoice.CreateInvoice(invoice); err != nil {
		return nil, err
	}

	respData := &model.BookingResponse{
		Booking: booking,
		Invoice: invoice,
	}

	return respData, nil
}
