package model

import (
	"time"

	"gorm.io/gorm"
)

type RoomStatusEnum string

const (
	Available   RoomStatusEnum = "available"
	Booked      RoomStatusEnum = "booked"
	Maintenance RoomStatusEnum = "maintenance"
)

type Hotel struct {
	gorm.Model
	// ID       uint         `gorm:"primaryKey" json:"id"`
	OwnerID     uint   `gorm:"not null" json:"owner_id" validate:"required"`
	Name        string `gorm:"size:100;not null" json:"name" validate:"required"`
	Description string `gorm:"type:text" json:"description"`
	Address     string `gorm:"size:255;not null" json:"address" validate:"required"`
	City        string `gorm:"size:100;not null" json:"city" validate:"required"`
	Zipcode     string `gorm:"size:20;not null" json:"zipcode" validate:"required"`
	Country     string `gorm:"size:100;not null" json:"country" validate:"required"`
	Phone       string `gorm:"size:14;unique;not null" json:"phone" validate:"required"`
	Email       string `gorm:"size:100;unique;not null" json:"email" validate:"required,email"`
	Star        int    `gorm:"default:1" json:"star"`
}

type RoomType struct {
	gorm.Model
	HotelID     uint    `gorm:"not null" json:"hotel_id" validate:"required"`
	Name        string  `gorm:"size:100;not null" json:"name" validate:"required,min=1,max=100"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"type:decimal(10,2);check:price >= 0;not null" json:"price" validate:"required,gt=0"`
	RoomSize    float64 `gorm:"type:decimal(4,1);not null" json:"room_size" validate:"required,gt=0"`
	Guest       int     `gorm:"check:guest >= 1;not null" json:"guest" validate:"required,gt=0"`

	Hotel Hotel `gorm:"foreignKey:HotelID;constraint:OnDelete:CASCADE;" json:"-"`
}

type RoomBedType struct {
	RoomTypeID uint `gorm:"primaryKey" json:"-"`
	DoubleBed  int  `gorm:"default:0" json:"double_bed"`
	SingleBed  int  `gorm:"default:0" json:"single_bed"`
	KingBed    int  `gorm:"default:0" json:"king_bed"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	RoomType RoomType `gorm:"foreignKey:RoomTypeID;constraint:OnDelete:CASCADE;" json:"-"`
}

type RoomFacilities struct {
	RoomTypeID      uint `gorm:"primaryKey"`
	HasShower       bool `gorm:"default:false"`
	HasRefrigerator bool `gorm:"default:false"`
	SeatingArea     bool `gorm:"default:false"`
	AirConditioning bool `gorm:"default:false"`
	HasBreakfast    bool `gorm:"default:false"`
	HasWifi         bool `gorm:"default:false"`
	SmokingAllowed  bool `gorm:"default:false"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`

	RoomType RoomType `gorm:"foreignKey:RoomTypeID;constraint:OnDelete:CASCADE;" json:"-"`
}

type Room struct {
	gorm.Model
	RoomTypeID uint           `gorm:"not null"`
	RoomNumber int            `gorm:"not null"`
	Status     RoomStatusEnum `gorm:"type:room_status_enum;default:'available'"`

	RoomType RoomType `gorm:"foreignKey:RoomTypeID;constraint:OnDelete:CASCADE;"`
}
