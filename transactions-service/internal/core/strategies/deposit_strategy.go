package strategies

import (
	"transactions-service/internal/core/domain"
	in "transactions-service/internal/ports/in/http"
	"transactions-service/internal/ports/out/db"
)

type depositStrategy struct {
	repository db.TransactionRepository
}

func NewDepositStrategy(repository db.TransactionRepository) HandleTransactionStrategy {
	return &depositStrategy{repository: repository}
}

func (s *depositStrategy) CanProcess(transactionType in.TransactionType) bool {
	return transactionType == in.Deposit
}

func (s *depositStrategy) Process(createTransaction in.CreateTransactionRequest) (*domain.Transaction, domain.ServiceError) {
	return nil, nil
}
