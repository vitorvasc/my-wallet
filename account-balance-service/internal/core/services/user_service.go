package services

import (
	"errors"

	"account-balance-service/internal/core/domain"
	"account-balance-service/internal/ports/out/repository"
)

type UserService interface {
	GetUserByID(userID uint64) (*domain.User, domain.ServiceError)
}

type userService struct {
	repository repository.UsersRepository
}

func NewUserService(repository repository.UsersRepository) UserService {
	return &userService{repository}
}

func (s *userService) GetUserByID(userID uint64) (*domain.User, domain.ServiceError) {
	user, err := s.repository.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, domain.ErrObtainingUser
	}
	return user, nil
}
