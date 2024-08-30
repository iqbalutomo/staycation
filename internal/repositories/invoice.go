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
	FindInvoiceByID(id string) (*model.Invoice, error)
	UpdateInvoiceStatus(id uint, status string) error

	CreatePayment(payment *model.Payment) error
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
	if err := r.db.Where("id = ?", id).First(&booking).Error; err != nil {
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

func (r *invoiceRepo) FindInvoiceByID(id string) (*model.Invoice, error) {
	var invoice model.Invoice
	if err := r.db.Where("xendit_invoice_id = ?", id).First(&invoice).Error; err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (r *invoiceRepo) UpdateInvoiceStatus(id uint, status string) error {
	return r.db.Model(&model.Invoice{}).Where("booking_id = ?", id).Update("status", status).Error
}

func (r *invoiceRepo) CreatePayment(payment *model.Payment) error {
	return r.db.Create(payment).Error
}
