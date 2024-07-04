package repository

import "transactions-service/internal/core/domain"

type UsersRepository interface {
	GetUserByID(userID uint64) (*domain.User, error)
	GetAccountBalance(userID uint64) (*domain.AccountBalance, error)
}
