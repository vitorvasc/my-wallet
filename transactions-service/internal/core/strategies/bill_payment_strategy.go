package strategies

import (
	"transactions-service/internal/core/domain"
	in "transactions-service/internal/ports/in/http"
)

type billPaymentStrategy struct{}

func NewBillPaymentStrategy() HandleTransactionStrategy {
	return &billPaymentStrategy{}
}

func (s *billPaymentStrategy) CanProcess(transactionType in.TransactionType) bool {
	return transactionType == in.BillPayment
}

func (s *billPaymentStrategy) Process(createTransaction in.CreateTransactionRequest) (*domain.Transaction, domain.ServiceError) {
	return nil, nil
}

func (s *billPaymentStrategy) mapToTransaction(createTransaction in.CreateTransactionRequest, user *domain.User) *domain.Transaction {
	return nil
}

func (s *billPaymentStrategy) persistTransaction(transaction *domain.Transaction) {

}
