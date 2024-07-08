package strategies

import (
	"testing"
	"time"
	"transactions-service/internal/app/config"
	"transactions-service/internal/core/domain"
	"transactions-service/internal/core/utils"
	in "transactions-service/internal/ports/in/http"
	"transactions-service/internal/ports/out/repository"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WithdrawalStrategyTestSuite struct {
	suite.Suite
}

func TestWithdrawalStrategy(t *testing.T) {
	suite.Run(t, new(WithdrawalStrategyTestSuite))
}

func (suite *WithdrawalStrategyTestSuite) TestCanProcess() {
	strategy := NewWithdrawalStrategy()

	suite.True(strategy.CanProcess(in.Withdrawal))
}

func (suite *WithdrawalStrategyTestSuite) TestProcess() {
	tests := []struct {
		name     string
		given    in.CreateTransactionRequest
		mocks    func()
		expected func() (*domain.Transaction, domain.ServiceError)
	}{
		{
			name: "Should return transaction with status approved and status detail processed",
			given: in.CreateTransactionRequest{
				Type: in.Withdrawal,
				From: in.TransactionFrom{
					UserID:        1,
					Amount:        100,
					PaymentMethod: in.AccountBalance,
				},
				To: in.TransactionTo{
					Amount: 100,
				},
			},
			mocks: func() {
				clockMock := new(utils.ClockMock)
				clockMock.On("Now").Return(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))
				config.Container.Register(config.Clock, clockMock)

				usersRepositoryMock := new(repository.UsersRepositoryMock)
				usersRepositoryMock.On("GetUserByID", mock.Anything).Return(&domain.User{
					UserID: 1,
				}, nil)
				usersRepositoryMock.On("GetAccountBalance", mock.Anything).Return(&domain.AccountBalance{
					UserID:  1,
					Balance: 100,
				}, nil)
				config.Container.Register(config.UsersRestClient, usersRepositoryMock)

				transactionRepositoryMock := new(repository.TransactionRepositoryMock)
				transactionRepositoryMock.On("CreateTransaction", mock.Anything).Return(nil)
				config.Container.Register(config.MongoRepository, transactionRepositoryMock)
			},
			expected: func() (*domain.Transaction, domain.ServiceError) {
				return &domain.Transaction{
					ID:           primitive.ObjectID{},
					UserID:       1,
					Type:         "withdrawal",
					Status:       "approved",
					StatusDetail: "processed",
					From: domain.TransactionFrom{
						UserID:        1,
						Amount:        100,
						PaymentMethod: "account_balance",
					},
					To: domain.TransactionTo{
						Amount: 100,
					},
					CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				}, nil
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			test.mocks()
			strategy := NewWithdrawalStrategy()

			transaction, err := strategy.Process(test.given)

			expectedTransaction, expectedErr := test.expected()

			suite.Equal(expectedTransaction, transaction)
			suite.Equal(expectedErr, err)
		})
	}
}
