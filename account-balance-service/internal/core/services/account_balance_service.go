package services

import (
	"account-balance-service/internal/core/domain"
	"account-balance-service/internal/ports/out/repository"
)

type AccountBalanceService interface {
	GetBalance(userID uint64) (*domain.AccountBalance, domain.ServiceError)
	AccountCredit(userID uint64, amount float64) domain.ServiceError
	AccountDebit(userID uint64, amount float64) domain.ServiceError
}

type accountBalanceService struct {
	repository  repository.AccountBalanceRepository
	userService UserService
}

func NewAccountBalanceService(repository repository.AccountBalanceRepository, userService UserService) AccountBalanceService {
	return &accountBalanceService{
		repository:  repository,
		userService: userService,
	}
}

func (s *accountBalanceService) GetBalance(userID uint64) (*domain.AccountBalance, domain.ServiceError) {
	_, userErr := s.userService.GetUserByID(userID)
	if userErr != nil {
		return nil, userErr
	}

	balance, err := s.repository.GetAccountBalance(userID)
	if err != nil {
		return nil, domain.ErrObtainingAccountBalance
	}

	return &domain.AccountBalance{
		UserID:  userID,
		Balance: balance,
	}, nil
}

func (s *accountBalanceService) AccountCredit(userID uint64, amount float64) domain.ServiceError {
	_, userErr := s.userService.GetUserByID(userID)
	if userErr != nil {
		return userErr
	}

	err := s.repository.Credit(userID, amount)
	if err != nil {
		return domain.ErrAccreditingValue
	}
	return nil
}

func (s *accountBalanceService) AccountDebit(userID uint64, amount float64) domain.ServiceError {
	_, userErr := s.userService.GetUserByID(userID)
	if userErr != nil {
		return userErr
	}

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
