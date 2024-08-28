package service

import (
	"errors"
	repository "staycation/internal/repositories"
)

type BalanceService interface {
	Deposit(userID int, amount float64) error
}

type balanceService struct {
	repo repository.BalanceRepository
}

func NewBalanceService(repo repository.BalanceRepository) BalanceService {
	return &balanceService{repo}
}

func (s *balanceService) Deposit(userID int, amount float64) error {
	if amount < 100000 {
		return errors.New("min_topup")
	} else if amount > 10000000 {
		return errors.New("max_topup")
	}

	return s.repo.AddBalance(userID, amount)
}
