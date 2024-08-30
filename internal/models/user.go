package model

type UserRoleEnum string

const (
	Customer   UserRoleEnum = "customer"
	HotelOwner UserRoleEnum = "hotel_owner"
)

type User struct {
	// gorm.Model
	ID       uint         `gorm:"primaryKey" json:"id"`
	Name     string       `gorm:"size:100;not null" json:"name" validate:"required"`
	Email    string       `gorm:"size:100;unique;not null" json:"email" validate:"required,email"`
	Password string       `gorm:"size:60;not null" json:"-" validate:"required"`
	Phone    string       `gorm:"size:14;unique;not null" json:"phone" validate:"required"`
	Role     UserRoleEnum `gorm:"type:user_role_enum;not null" json:"role" validate:"required,oneof=customer hotel_owner"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
