package model

import "gorm.io/gorm"

type UserRoleEnum string

const (
	Customer   UserRoleEnum = "customer"
	HotelOwner UserRoleEnum = "hotel_owner"
)

type User struct {
	// ID       uint         `gorm:"primaryKey" json:"id"`
	gorm.Model
	Name     string       `gorm:"size:100;not null" json:"name" validate:"required"`
	Email    string       `gorm:"size:100;unique;not null" json:"email" validate:"required,email"`
	Password string       `gorm:"size:60;not null" json:"-" validate:"required"`
	Phone    string       `gorm:"size:14;unique;not null" json:"phone" validate:"required"`
	Role     UserRoleEnum `gorm:"type:user_role_enum;not null" json:"role" validate:"required,oneof=customer hotel_owner"`
}
