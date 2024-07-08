package repository

import (
	"github.com/stretchr/testify/mock"
)

type AccountBalanceRepositoryMock struct {
	mock.Mock
}

func (m *AccountBalanceRepositoryMock) GetAccountBalance(userID uint64) (float64, error) {
	args := m.Called(userID)
	return args.Get(0).(float64), args.Error(1)
}

func (m *AccountBalanceRepositoryMock) Credit(userID uint64, amount float64) error {
	args := m.Called(userID, amount)
	return args.Error(0)
}

func (m *AccountBalanceRepositoryMock) Debit(userID uint64, amount float64) error {
	args := m.Called(userID, amount)
	return args.Error(0)
}
