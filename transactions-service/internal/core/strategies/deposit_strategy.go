package strategies

import (
	"errors"
	"time"
	out "transactions-service/internal/adapters/out/http"
	"transactions-service/internal/app/config"
	"transactions-service/internal/core/domain"
	"transactions-service/internal/core/services"
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
	usersRestClient := config.Container.Get(config.UsersRestClient).(out.UsersRestClient)
	user, err := usersRestClient.GetUserByID(createTransaction.From.UserID)
	if err != nil {
		if errors.As(err, &domain.ErrUserNotFound) {
			return nil, domain.ErrInvalidUsersInvolved
		}
		return nil, domain.ErrObtainingUserByID
	}

	transaction := mapToTransaction(createTransaction, user)
	if createTransaction.From.Amount <= 0 {
		transaction.Status = domain.Rejected
		transaction.StatusDetail = domain.InvalidAmount
		return transaction, domain.ErrInvalidAmount
	}

	accreditationService := config.Container.Get(config.AccreditationService).(services.AccreditationService)
	err = accreditationService.CreditUserBalance(user.UserID, createTransaction.From.Amount)
	if err != nil {
		transaction.Status = domain.Rejected
		transaction.StatusDetail = domain.AccreditationProcessingError
		return transaction, domain.ErrProcessingAccreditation
	}

	transaction.Status = domain.Approved
	transaction.StatusDetail = domain.Processed
	return transaction, nil
}

func mapToTransaction(createTransaction in.CreateTransactionRequest, user *domain.User) *domain.Transaction {
	return &domain.Transaction{
		Type:        string(createTransaction.Type),
		Description: createTransaction.Description,
		From: domain.TransactionFrom{
			UserID:        user.UserID,
			Amount:        createTransaction.From.Amount,
			PaymentMethod: string(createTransaction.From.PaymentMethod),
		},
		To: domain.TransactionTo{
			UserID: user.UserID,
			Amount: createTransaction.To.Amount,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
