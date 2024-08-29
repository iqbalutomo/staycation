package repository

import (
	model "staycation/internal/models"
	database "staycation/pkg/databases"
	"time"

	"gorm.io/gorm"
)

type InvoiceRepository interface {
	CreateBooking(booking *model.Booking) error
	FindBookingByID(id uint) (*model.Booking, error)
	FindByRoomAndDate(roomID uint, checkInDate, checkOutDate time.Time) ([]model.Booking, error)

	CreateInvoice(invoice *model.Invoice) error
}

type invoiceRepo struct {
	db *gorm.DB
}

func NewInvoiceRepository() InvoiceRepository {
	return &invoiceRepo{db: database.DB}
}

func (r *invoiceRepo) CreateBooking(booking *model.Booking) error {
	return r.db.Create(booking).Error
}

func (r *invoiceRepo) FindBookingByID(id uint) (*model.Booking, error) {
	var booking model.Booking
	if err := r.db.First(&booking, id).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *invoiceRepo) FindByRoomAndDate(roomID uint, checkInDate, checkOutDate time.Time) ([]model.Booking, error) {
	var bookings []model.Booking
	if err := r.db.Where("room_id = ? AND ((check_in_date <= ? AND check_out_date >= ?) OR (check_in_date >= ? AND check_in_date < ?))", roomID, checkOutDate, checkInDate, checkInDate, checkOutDate).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *invoiceRepo) CreateInvoice(invoice *model.Invoice) error {
	return r.db.Create(invoice).Error
}
