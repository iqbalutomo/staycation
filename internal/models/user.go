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

// swagger
type UserRegisterRequest struct {
	ID       uint         `json:"-"`
	Name     string       `json:"name" validate:"required"`
	Email    string       `json:"email" validate:"required,email"`
	Password string       `json:"password" validate:"required"`
	Phone    string       `json:"phone" validate:"required"`
	Role     UserRoleEnum `json:"role" validate:"required,oneof=customer hotel_owner"`
}

type RegisterSuccessResponse struct {
	Status string `json:"status"`
	Data   User   `json:"data"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

// FOR TESTTTTT
type UserTest struct {
	ID       uint         `gorm:"primaryKey" json:"id"`
	Name     string       `gorm:"size:100;not null" json:"name" validate:"required"`
	Email    string       `gorm:"size:100;unique;not null" json:"email" validate:"required,email"`
	Password string       `gorm:"size:60;not null" json:"-" validate:"required"`
	Phone    string       `gorm:"size:14;unique;not null" json:"phone" validate:"required"`
	Role     UserRoleEnum `gorm:"type:user_role_enum;not null" json:"role" validate:"required,oneof=customer hotel_owner"`
}
