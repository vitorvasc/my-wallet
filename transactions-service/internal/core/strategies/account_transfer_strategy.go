package strategies

import (
	"transactions-service/internal/core/domain"
	in "transactions-service/internal/ports/in/http"
	"transactions-service/internal/ports/out/db"
)

type accountTransferStrategy struct {
	repository db.TransactionRepository
}

func NewAccountTransferStrategy(repository db.TransactionRepository) HandleTransactionStrategy {
	return &accountTransferStrategy{repository: repository}
}

func (s *accountTransferStrategy) CanProcess(transactionType in.TransactionType) bool {
	return transactionType == in.AccountTransfer
}

func (s *accountTransferStrategy) Process(createTransaction in.CreateTransactionRequest) (*domain.Transaction, domain.ServiceError) {
	return nil, nil
}
