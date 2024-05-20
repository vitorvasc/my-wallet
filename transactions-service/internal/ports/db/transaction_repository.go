package db

import "transactions-service/internal/core/domain"

type TransactionRepository interface {
	CreateTransaction(transaction domain.Transaction) error
	UpdateTransaction(transaction domain.Transaction) error
	GetUserTransactions(userId float64) ([]domain.Transaction, error)
}
