package services

import out "transactions-service/internal/ports/out/repository"

type AccreditationService interface {
	CreditUserBalance(userID uint64, amount float64) error
}

type accreditationService struct {
	repository out.AccreditationRepository
}

func NewAccreditationService(repository out.AccreditationRepository) AccreditationService {
	return &accreditationService{repository: repository}
}

func (s *accreditationService) CreditUserBalance(userID uint64, amount float64) error {
	return s.repository.CreateUserBalanceCredit(out.UserBalanceCredit{
		UserID: userID,
		Amount: amount,
	})
}
