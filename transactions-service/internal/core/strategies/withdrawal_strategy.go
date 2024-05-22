package strategies

import (
	"transactions-service/internal/core/domain"
	in "transactions-service/internal/ports/in/http"
)

type withdrawalStrategy struct{}

func NewWithdrawalStrategy() HandleTransactionStrategy {
	return &withdrawalStrategy{}
}

func (s *withdrawalStrategy) CanProcess(transactionType in.TransactionType) bool {
	return transactionType == in.Withdrawal
}

func (s *withdrawalStrategy) Process(createTransaction in.CreateTransactionRequest) (*domain.Transaction, domain.ServiceError) {
	return nil, nil
}
