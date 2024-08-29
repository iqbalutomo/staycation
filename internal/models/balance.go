package model

import (
	"time"

	"gorm.io/gorm"
)

type Balance struct {
	UserID    uint    `gorm:"primaryKey;not null" json:"user_id"`
	Balance   float64 `gorm:"type:decimal(10, 2); not null" json:"balance"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
