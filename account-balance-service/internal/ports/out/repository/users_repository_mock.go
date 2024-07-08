package repository

import (
	"account-balance-service/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type UsersRepositoryMock struct {
	mock.Mock
}

func (m *UsersRepositoryMock) GetUserByID(userID uint64) (*domain.User, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), nil
}
