package services

import (
	"account-balance-service/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (m *UserServiceMock) GetUserByID(userID uint64) (*domain.User, domain.ServiceError) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Get(1).(domain.ServiceError)
	}
	return args.Get(0).(*domain.User), nil
}
