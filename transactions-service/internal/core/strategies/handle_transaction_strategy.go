package strategies

import (
	"transactions-service/internal/core/domain"
	in "transactions-service/internal/ports/in/http"
)

type HandleTransactionStrategy interface {
	CanProcess(transactionType in.TransactionType) bool
	Process(createTransaction in.CreateTransactionRequest) (*domain.Transaction, domain.ServiceError)
	mapToTransaction(createTransaction in.CreateTransactionRequest, user *domain.User) *domain.Transaction
	persistTransaction(transaction *domain.Transaction)
}
