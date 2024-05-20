package services

import "transactions-service/internal/ports/db"

type TransactionService struct {
	repository db.TransactionRepository
}

func NewTransactionService(repository db.TransactionRepository) *TransactionService {
	return &TransactionService{repository: repository}
}

func (s *TransactionService) CreateTransaction(transaction db.Transaction) error {
	return s.repository.CreateTransaction(transaction)
}

func (s *TransactionService) UpdateTransaction(transaction db.Transaction) error {
	return s.repository.UpdateTransaction(transaction)
}

func (s *TransactionService) GetTransactionsByUser(userId float64) ([]db.Transaction, error) {
	return s.repository.GetUserTransactions(userId)
}
