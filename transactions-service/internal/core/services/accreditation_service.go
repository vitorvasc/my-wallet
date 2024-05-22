package services

import out "transactions-service/internal/ports/out/repository"

type AccreditationService interface {
	AccreditateUserBalance(userID uint64, amount float64) error
}

type accreditationService struct {
	repository out.AccreditationRepository
}

func NewAccreditationService(repository out.AccreditationRepository) AccreditationService {
	return &accreditationService{repository: repository}
}

func (s *accreditationService) AccreditateUserBalance(userID uint64, amount float64) error {
	return s.repository.AccreditateUserBalance(out.UserBalanceAccreditation{
		UserID: userID,
		Amount: amount,
	})
}
