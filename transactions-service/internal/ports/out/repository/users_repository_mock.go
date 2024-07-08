package repository

import (
	"transactions-service/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type UsersRepositoryMock struct {
	mock.Mock
}

func (m *UsersRepositoryMock) GetUserByID(userID uint64) (*domain.User, error) {
	args := m.Called(userID)
	if err := args.Error(1); err != nil {
		return nil, err
	}
	return args.Get(0).(*domain.User), nil
}

func (m *UsersRepositoryMock) GetAccountBalance(userID uint64) (*domain.AccountBalance, error) {
	args := m.Called(userID)
	if err := args.Error(1); err != nil {
		return nil, err
	}
	return args.Get(0).(*domain.AccountBalance), nil
}
