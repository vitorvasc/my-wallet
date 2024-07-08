package strategies

import (
	"errors"

	"transactions-service/internal/app/config"
	"transactions-service/internal/core/domain"
	"transactions-service/internal/core/services"
	"transactions-service/internal/core/utils"
	in "transactions-service/internal/ports/in/http"
	"transactions-service/internal/ports/out/repository"
)

type depositStrategy struct{}

func NewDepositStrategy() HandleTransactionStrategy {
	return &depositStrategy{}
}

func (s *depositStrategy) CanProcess(transactionType in.TransactionType) bool {
	return transactionType == in.Deposit
}

func (s *depositStrategy) Process(createTransaction in.CreateTransactionRequest) (*domain.Transaction, domain.ServiceError) {
	usersRestClient := config.Container.Get(config.UsersRestClient).(repository.UsersRepository)
	user, err := usersRestClient.GetUserByID(createTransaction.To.UserID)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, domain.ErrInvalidUsersInvolved
		}
		return nil, domain.ErrObtainingUserByID
	}

	transaction := mapToTransaction(createTransaction, user)
	defer persistTransaction(transaction)

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

func persistTransaction(transaction *domain.Transaction) {
	r := config.Container.Get(config.MongoRepository).(repository.TransactionRepository)
	_ = r.CreateTransaction(transaction)
}

func mapToTransaction(createTransaction in.CreateTransactionRequest, user *domain.User) *domain.Transaction {
	clock := config.Container.Get(config.Clock).(utils.Clock)
	return &domain.Transaction{
		UserID:      user.UserID,
		Type:        string(createTransaction.Type),
		Description: createTransaction.Description,
		From: domain.TransactionFrom{
			Amount:        createTransaction.From.Amount,
			PaymentMethod: string(createTransaction.From.PaymentMethod),
		},
		To: domain.TransactionTo{
			UserID: user.UserID,
			Amount: createTransaction.To.Amount,
		},
		CreatedAt: clock.Now(),
		UpdatedAt: clock.Now(),
	}
}
