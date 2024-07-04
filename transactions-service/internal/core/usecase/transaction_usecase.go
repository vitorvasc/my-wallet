package usecase

import (
	"transactions-service/internal/app/config"
	"transactions-service/internal/core/domain"
	"transactions-service/internal/core/strategies"
	in "transactions-service/internal/ports/in/http"
	out "transactions-service/internal/ports/out/repository"
)

type TransactionUseCase interface {
	CreateTransaction(transaction in.CreateTransactionRequest) (*domain.Transaction, domain.ServiceError)
	UpdateTransaction(transactionID uint64, transaction in.UpdateTransactionRequest) (*domain.Transaction, domain.ServiceError)
	GetTransactionsByUser(userID uint64) ([]*domain.Transaction, domain.ServiceError)
}

type transactionUseCase struct {
	strategies []strategies.HandleTransactionStrategy
}

func NewTransactionUseCase(
	strategies []strategies.HandleTransactionStrategy,
) TransactionUseCase {
	return &transactionUseCase{
		strategies: strategies,
	}
}

func (s *transactionUseCase) CreateTransaction(createTransaction in.CreateTransactionRequest) (*domain.Transaction, domain.ServiceError) {
	for _, strategy := range s.strategies {
		if strategy.CanProcess(createTransaction.Type) {
			return strategy.Process(createTransaction)
		}
	}

	return nil, domain.ErrProcessingTransactionStrategyNotFound
}

func (s *transactionUseCase) UpdateTransaction(transactionID uint64, updateTransaction in.UpdateTransactionRequest) (*domain.Transaction, domain.ServiceError) {
	repository := config.Container.Get(config.MongoRepository).(out.TransactionRepository)
	transaction, err := repository.GetTransactionByID(transactionID)
	if err != nil {
		return nil, domain.ErrObtainingTransaction
	}

	// TODO: map fields to update

	if err := repository.UpdateTransaction(transaction); err != nil {
		return nil, domain.ErrUpdatingTransaction
	}

	return transaction, nil
}

func (s *transactionUseCase) GetTransactionsByUser(userID uint64) ([]*domain.Transaction, domain.ServiceError) {
	repository := config.Container.Get(config.MongoRepository).(out.TransactionRepository)
	transactions, err := repository.GetTransactionsByUserID(userID)
	if err != nil {
		return nil, domain.ErrObtainingUserTransactions
	}

	return transactions, nil
}
