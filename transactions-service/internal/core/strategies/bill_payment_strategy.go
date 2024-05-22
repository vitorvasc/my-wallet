package strategies

import (
	"transactions-service/internal/core/domain"
	in "transactions-service/internal/ports/in/http"
	"transactions-service/internal/ports/out/db"
)

type billPaymentStrattegy struct {
	repository db.TransactionRepository
}

func NewBillPaymentStrategy(repository db.TransactionRepository) HandleTransactionStrategy {
	return &billPaymentStrattegy{repository: repository}
}

func (s *billPaymentStrattegy) CanProcess(transactionType in.TransactionType) bool {
	return transactionType == in.BillPayment
}

func (s *billPaymentStrattegy) Process(createTransaction in.CreateTransactionRequest) (*domain.Transaction, domain.ServiceError) {
	return nil, nil
}
