package repository

import "account-balance-service/internal/core/domain"

type UsersRepository interface {
	GetUserByID(userID uint64) (*domain.User, error)
}
