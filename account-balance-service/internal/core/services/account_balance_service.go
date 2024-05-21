package services

import (
	"account-balance-service/internal/core/domain"
	"account-balance-service/internal/ports/out/db"
)

type AccountBalance interface {
	GetBalance(userID uint64) (float64, *domain.ServiceError)
	AccreditValue(userID uint64, amount float64) *domain.ServiceError
	DebitValue(userID uint64, amount float64) *domain.ServiceError
}

type accountBalanceService struct {
	repository db.AccountBalanceRepository
}

func NewAccountBalanceService(repository db.AccountBalanceRepository) AccountBalance {
	return &accountBalanceService{repository}
}

func (s *accountBalanceService) GetBalance(userID uint64) (float64, *domain.ServiceError) {
	balance, err := s.repository.GetAccountBalance(userID)
	if err != nil {
		return 0, domain.ErrObtainingAccountBalance
	}
	return balance, nil
}

func (s *accountBalanceService) AccreditValue(userID uint64, amount float64) *domain.ServiceError {
	err := s.repository.Accredit(userID, amount)
	if err != nil {
		return domain.ErrAccreditingValue
	}
	return nil
}

func (s *accountBalanceService) DebitValue(userID uint64, amount float64) *domain.ServiceError {
	balance, err := s.repository.GetAccountBalance(userID)
	if err != nil {
		return domain.ErrObtainingAccountBalance
	}

	if balance < amount {
		return domain.ErrInsufficientFunds
	}

	err = s.repository.Debit(userID, amount)
	if err != nil {
		return domain.ErrDebitingValue
	}

	return nil
}
