package model

import (
	"time"
)

type Booking struct {
	// gorm.Model
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `gorm:"not null" json:"user_id"`
	HotelID      uint      `gorm:"not null" json:"hotel_id"`
	RoomID       uint      `gorm:"not null" json:"room_id"`
	CheckInDate  time.Time `gorm:"not null" json:"check_in_date" validate:"required"`
	CheckOutDate time.Time `gorm:"not null" json:"check_out_date" validate:"required"`
	TotalPrice   float64   `gorm:"not null" json:"total_price"`
	Status       string    `gorm:"type:booking_status_enum;default:'booked'"`

	User  User  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"-" validate:"-"`
	Hotel Hotel `gorm:"foreignKey:HotelID;constraint:OnDelete:CASCADE;" json:"-" validate:"-"`
	Room  Room  `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE;" json:"-" validate:"-"`
}

type Invoice struct {
	BookingID       uint    `gorm:"primaryKey" json:"booking_id" validate:"required"`
	XenditInvoiceID string  `gorm:"not null" json:"xendit_invoice_id" validate:"required"`
	InvoiceURL      string  `gorm:"type:text" json:"invoice_url"`
	Amount          float64 `gorm:"not null" json:"amount" validate:"required"`
	Status          string  `gorm:"type:invoice_status_enum;default:'PENDING'" json:"status"`

	Booking Booking `gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE;" json:"-" validate:"-"`
}

type Payment struct {
	// gorm.Model
	ID            uint      `gorm:"primaryKey" json:"id"`
	InvoiceID     uint      `gorm:"not null" json:"invoice_id" validate:"required"`
	PaymentMethod string    `gorm:"not null" json:"payment_method" validate:"required"`
	PaidAmount    float64   `gorm:"not null" json:"paid_amount" validate:"required"`
	PaidAt        time.Time `gorm:"not null" json:"paid_at"`

	Invoice Invoice `gorm:"foreignKey:InvoiceID;constraint:OnDelete:CASCADE;" json:"-" validate:"-"`
}

type BookingResponse struct {
	Booking *Booking `json:"booking"`
	Invoice *Invoice `json:"invoice"`
}
