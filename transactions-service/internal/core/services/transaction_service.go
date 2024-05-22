package services

import (
	"transactions-service/internal/core/domain"
	"transactions-service/internal/core/strategies"
	in "transactions-service/internal/ports/in/http"
	out "transactions-service/internal/ports/out/repository"
)

type TransactionService interface {
	CreateTransaction(transaction in.CreateTransactionRequest) (*domain.Transaction, domain.ServiceError)
	UpdateTransaction(transactionID uint64, transaction in.UpdateTransactionRequest) (*domain.Transaction, domain.ServiceError)
	GetTransactionsByUser(userID uint64) ([]*domain.Transaction, domain.ServiceError)
}

type transactionService struct {
	repository           out.TransactionRepository
	strategies           []strategies.HandleTransactionStrategy
	accreditationService AccreditationService
}

func NewTransactionService(
	repository out.TransactionRepository,
	strategies []strategies.HandleTransactionStrategy,
	accreditationService AccreditationService,
) TransactionService {
	return &transactionService{
		repository:           repository,
		strategies:           strategies,
		accreditationService: accreditationService,
	}
}

func (s *transactionService) CreateTransaction(createTransaction in.CreateTransactionRequest) (*domain.Transaction, domain.ServiceError) {

	for _, strategy := range s.strategies {
		if strategy.CanProcess(createTransaction.Type) {
			return strategy.Process(createTransaction)
		}
	}

	return nil, domain.ErrProcessingTransactionStrategyNotFound
}

func (s *transactionService) UpdateTransaction(transactionID uint64, updateTransaction in.UpdateTransactionRequest) (*domain.Transaction, domain.ServiceError) {
	transaction, err := s.repository.GetTransactionByID(transactionID)
	if err != nil {
		return nil, domain.ErrObtainingTransaction
	}

	// TODO: map fields to update

	if err := s.repository.UpdateTransaction(transaction); err != nil {
		return nil, domain.ErrUpdatingTransaction
	}

	return transaction, nil
}

func (s *transactionService) GetTransactionsByUser(userID uint64) ([]*domain.Transaction, domain.ServiceError) {
	transactions, err := s.repository.GetTransactionsByUserID(userID)
	if err != nil {
		return nil, domain.ErrObtainingUserTransactions
	}

	return transactions, nil
}
