package strategies

import (
	"transactions-service/internal/core/domain"
	in "transactions-service/internal/ports/in/http"
)

type depositStrategy struct{}

func NewDepositStrategy() HandleTransactionStrategy {
	return &depositStrategy{}
}

func (s *depositStrategy) CanProcess(transactionType in.TransactionType) bool {
	return transactionType == in.Deposit
}

func (s *depositStrategy) Process(createTransaction in.CreateTransactionRequest) (*domain.Transaction, domain.ServiceError) {
	return nil, nil
}
