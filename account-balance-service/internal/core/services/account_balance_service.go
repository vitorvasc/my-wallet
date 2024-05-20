package services

import (
	"account-balance-service/internal/core/domain"
	"account-balance-service/internal/ports/db"
)

type AccreditValueRequest struct {
	UserID uint64  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type AccountBalanceService struct {
	repository db.AccountBalanceRepository
}

func NewAccountBalanceService(repository db.AccountBalanceRepository) *AccountBalanceService {
	return &AccountBalanceService{repository}
}

func (s *AccountBalanceService) GetBalance(userID uint64) (float64, error) {
	return s.repository.GetAccountBalance(userID)
}

func (s *AccountBalanceService) AccreditValue(userID uint64, amount float64) error {
	balance, err := s.repository.GetAccountBalance(userID)
	if err != nil {
		return err
	}

	return s.repository.UpdateAccountBalance(userID, balance+amount)
}

func (s *AccountBalanceService) DebitValue(userID uint64, amount float64) error {
	balance, err := s.repository.GetAccountBalance(userID)
	if err != nil {
		return err
	}

	if balance < amount {
		return domain.ErrInsufficientFunds
	}

	return s.repository.UpdateAccountBalance(userID, balance-amount)
}
