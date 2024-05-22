package strategies

import (
	"transactions-service/internal/core/domain"
	in "transactions-service/internal/ports/in/http"
	"transactions-service/internal/ports/out/db"
)

type withdrawalStrategy struct {
	repository db.TransactionRepository
}

func NewWithdrawalStrategy(repository db.TransactionRepository) HandleTransactionStrategy {
	return &withdrawalStrategy{repository: repository}
}

func (s *withdrawalStrategy) CanProcess(transactionType in.TransactionType) bool {
	return transactionType == in.Withdrawal
}

func (s *withdrawalStrategy) Process(createTransaction in.CreateTransactionRequest) (*domain.Transaction, domain.ServiceError) {
	return nil, nil
}
