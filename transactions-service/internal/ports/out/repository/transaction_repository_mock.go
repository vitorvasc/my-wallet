package repository

import (
	"transactions-service/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (m *TransactionRepositoryMock) CreateTransaction(transaction *domain.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) UpdateTransaction(transaction *domain.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) GetTransactionByID(transactionID uint64) (*domain.Transaction, error) {
	args := m.Called(transactionID)
	return args.Get(0).(*domain.Transaction), args.Error(1)
}

func (m *TransactionRepositoryMock) GetTransactionsByUserID(userID uint64) ([]*domain.Transaction, error) {
	args := m.Called(userID)
	return args.Get(0).([]*domain.Transaction), args.Error(1)
}
