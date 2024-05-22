package db

import "transactions-service/internal/core/domain"

type TransactionRepository interface {
	CreateTransaction(transaction *domain.Transaction) error
	UpdateTransaction(transaction *domain.Transaction) error
	GetTransactionByID(transactionID uint64) (*domain.Transaction, error)
	GetTransactionsByUserID(userID uint64) ([]*domain.Transaction, error)
}
