package repository

import (
	model "staycation/internal/models"
	database "staycation/pkg/databases"

	"gorm.io/gorm"
)

type BalanceRepository interface {
	AddBalance(userID int, amount float64) error
}

type balanceRepo struct {
	db *gorm.DB
}

func NewBalanceRepository() BalanceRepository {
	return &balanceRepo{db: database.DB}
}

func (r *balanceRepo) AddBalance(userID int, amount float64) error {
	var balance model.Balance
	if err := r.db.Where("user_id = ?", userID).First(&balance).Error; err != nil {
		return err
	}

	newBalance := balance.Balance + amount

	if err := r.db.Model(&model.Balance{}).Where("user_id = ?", userID).Update("balance", newBalance).Error; err != nil {
		return err
	}

	return nil
}
