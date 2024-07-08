package strategies

import (
	"errors"
	"transactions-service/internal/app/config"
	"transactions-service/internal/core/domain"
	"transactions-service/internal/core/utils"
	in "transactions-service/internal/ports/in/http"
	"transactions-service/internal/ports/out/repository"
)

type withdrawalStrategy struct{}

func NewWithdrawalStrategy() HandleTransactionStrategy {
	return &withdrawalStrategy{}
}

func (s *withdrawalStrategy) CanProcess(transactionType in.TransactionType) bool {
	return transactionType == in.Withdrawal
}

func (s *withdrawalStrategy) Process(createTransaction in.CreateTransactionRequest) (*domain.Transaction, domain.ServiceError) {
	usersRestClient := config.Container.Get(config.UsersRestClient).(repository.UsersRepository)
	user, err := usersRestClient.GetUserByID(createTransaction.From.UserID)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, domain.ErrInvalidUsersInvolved
		}
		return nil, domain.ErrObtainingUserByID
	}

	userBalance, err := usersRestClient.GetAccountBalance(user.UserID)
	if err != nil {
		return nil, domain.ErrObtainingAccountBalance
	}

	transaction := s.mapToTransaction(createTransaction, user)
	defer s.persistTransaction(transaction)

	if userBalance.Balance < createTransaction.From.Amount {
		return transaction, domain.ErrInsufficientFunds
	}

	// TODO: Create debit

	transaction.Status = domain.Approved
	transaction.StatusDetail = domain.Processed

	return transaction, nil
}

func (s *withdrawalStrategy) mapToTransaction(createTransaction in.CreateTransactionRequest, user *domain.User) *domain.Transaction {
	clock := config.Container.Get(config.Clock).(utils.Clock)
	return &domain.Transaction{
		UserID:      user.UserID,
		Type:        string(createTransaction.Type),
		Description: createTransaction.Description,
		From: domain.TransactionFrom{
			UserID:        user.UserID,
			Amount:        createTransaction.From.Amount,
			PaymentMethod: string(createTransaction.From.PaymentMethod),
		},
		To: domain.TransactionTo{
			Amount: createTransaction.To.Amount,
		},
		CreatedAt: clock.Now(),
		UpdatedAt: clock.Now(),
	}
}

func (s *withdrawalStrategy) persistTransaction(transaction *domain.Transaction) {
	return
}
