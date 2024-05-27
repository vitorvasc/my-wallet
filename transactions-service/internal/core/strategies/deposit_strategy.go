package strategies

import (
	"time"
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
	transaction := mapToTransaction(createTransaction)
	if createTransaction.From.Amount <= 0 {
		transaction.Status = domain.Rejected
		transaction.StatusDetail = domain.InvalidAmount
		return transaction, domain.ErrInvalidAmount
	}

	transaction.Status = domain.Approved
	transaction.StatusDetail = domain.Processed
	return mapToTransaction(createTransaction), nil
}

func mapToTransaction(createTransaction in.CreateTransactionRequest) *domain.Transaction {
	return &domain.Transaction{
		Type:        string(createTransaction.Type),
		Description: createTransaction.Description,
		From: domain.TransactionFrom{
			UserID:        createTransaction.From.UserID,
			Amount:        createTransaction.From.Amount,
			PaymentMethod: string(createTransaction.From.PaymentMethod),
		},
		To: domain.TransactionTo{
			UserID: createTransaction.To.UserID,
			Amount: createTransaction.To.Amount,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
