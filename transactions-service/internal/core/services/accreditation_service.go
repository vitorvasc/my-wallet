package services

import (
	"transactions-service/internal/ports/out/dto"
	"transactions-service/internal/ports/out/repository"
)

type AccreditationService interface {
	CreditUserBalance(userID uint64, amount float64) error
}

type accreditationService struct {
	repository repository.AccreditationRepository
}

func NewAccreditationService(repository repository.AccreditationRepository) AccreditationService {
	return &accreditationService{repository: repository}
}

func (s *accreditationService) CreditUserBalance(userID uint64, amount float64) error {
	return s.repository.CreateUserBalanceCredit(dto.AccountCreditCreation{
		UserID: userID,
		Amount: amount,
	})
}
