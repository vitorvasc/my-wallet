package strategies

import (
	"transactions-service/internal/core/domain"
	in "transactions-service/internal/ports/in/http"
)

type accountTransferStrategy struct{}

func NewAccountTransferStrategy() HandleTransactionStrategy {
	return &accountTransferStrategy{}
}

func (s *accountTransferStrategy) CanProcess(transactionType in.TransactionType) bool {
	return transactionType == in.AccountTransfer
}

func (s *accountTransferStrategy) Process(createTransaction in.CreateTransactionRequest) (*domain.Transaction, domain.ServiceError) {
	return nil, nil
}

func (s *accountTransferStrategy) mapToTransaction(createTransaction in.CreateTransactionRequest, user *domain.User) *domain.Transaction {
	return nil
}

func (s *accountTransferStrategy) persistTransaction(transaction *domain.Transaction) {
	return
}
