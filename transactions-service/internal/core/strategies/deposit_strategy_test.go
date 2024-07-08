package strategies

import (
	"errors"
	"testing"
	"time"
	"transactions-service/internal/app/config"
	"transactions-service/internal/core/services"
	"transactions-service/internal/core/utils"
	"transactions-service/internal/ports/out/repository"

	"transactions-service/internal/core/domain"
	in "transactions-service/internal/ports/in/http"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DepositStrategyTestSuite struct {
	suite.Suite
}

func TestDepositStrategy(t *testing.T) {
	suite.Run(t, new(DepositStrategyTestSuite))
}

func (suite *DepositStrategyTestSuite) TestCanProcess() {
	strategy := NewDepositStrategy()

	suite.True(strategy.CanProcess(in.Deposit))
}

func (suite *DepositStrategyTestSuite) TestProcess() {
	tests := []struct {
		name     string
		given    in.CreateTransactionRequest
		mocks    func()
		expected func() (*domain.Transaction, domain.ServiceError)
	}{
		{
			name: "Should return transaction with status approved and status detail processed",
			given: in.CreateTransactionRequest{
				Type: in.Deposit,
				From: in.TransactionFrom{
					Amount:        100,
					PaymentMethod: in.PureCash,
				},
				To: in.TransactionTo{
					UserID: 1,
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
				config.Container.Register(config.UsersRestClient, usersRepositoryMock)

				accreditationRepositoryMock := new(repository.AccreditationRepositoryMock)
				accreditationRepositoryMock.On("CreateUserBalanceCredit", mock.Anything).Return(nil)
				accreditationService := services.NewAccreditationService(accreditationRepositoryMock)
				config.Container.Register(config.AccreditationService, accreditationService)

				transactionRepositoryMock := new(repository.TransactionRepositoryMock)
				transactionRepositoryMock.On("CreateTransaction", mock.Anything).Return(nil)
				config.Container.Register(config.MongoRepository, transactionRepositoryMock)
			},
			expected: func() (*domain.Transaction, domain.ServiceError) {
				return &domain.Transaction{
					ID:           primitive.ObjectID{},
					UserID:       1,
					Type:         "deposit",
					Status:       "approved",
					StatusDetail: "processed",
					From: domain.TransactionFrom{
						Amount:        100,
						PaymentMethod: "pure_cash",
					},
					To: domain.TransactionTo{
						UserID: 1,
						Amount: 100,
					},
					CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				}, nil
			},
		},
		{
			name: "Should return transaction with status rejected and status detail invalid amount",
			given: in.CreateTransactionRequest{
				Type: in.Deposit,
				From: in.TransactionFrom{
					Amount:        0,
					PaymentMethod: in.PureCash,
				},
				To: in.TransactionTo{
					UserID: 1,
					Amount: 0,
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
				config.Container.Register(config.UsersRestClient, usersRepositoryMock)

				accreditationRepositoryMock := new(repository.AccreditationRepositoryMock)
				accreditationRepositoryMock.On("CreateUserBalanceCredit", mock.Anything).Return(nil)
				accreditationService := services.NewAccreditationService(accreditationRepositoryMock)
				config.Container.Register(config.AccreditationService, accreditationService)

				transactionRepositoryMock := new(repository.TransactionRepositoryMock)
				transactionRepositoryMock.On("CreateTransaction", mock.Anything).Return(nil)
				config.Container.Register(config.MongoRepository, transactionRepositoryMock)
			},
			expected: func() (*domain.Transaction, domain.ServiceError) {
				return &domain.Transaction{
					ID:           primitive.ObjectID{},
					UserID:       1,
					Type:         "deposit",
					Status:       "rejected",
					StatusDetail: "invalid_amount",
					From: domain.TransactionFrom{
						Amount:        0,
						PaymentMethod: "pure_cash",
					},
					To: domain.TransactionTo{
						UserID: 1,
						Amount: 0,
					},
					CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				}, domain.ErrInvalidAmount
			},
		},
		{
			name: "Should return error invalid users involved given an non-existent user id",
			given: in.CreateTransactionRequest{
				Type: in.Deposit,
				From: in.TransactionFrom{
					Amount:        0,
					PaymentMethod: in.PureCash,
				},
				To: in.TransactionTo{
					UserID: 0,
					Amount: 0,
				},
			},
			mocks: func() {
				usersRepositoryMock := new(repository.UsersRepositoryMock)
				usersRepositoryMock.On("GetUserByID", mock.Anything).Return(nil, domain.ErrUserNotFound)
				config.Container.Register(config.UsersRestClient, usersRepositoryMock)
			},
			expected: func() (*domain.Transaction, domain.ServiceError) {
				return nil, domain.ErrInvalidUsersInvolved
			},
		},
		{
			name: "Should return error invalid obtaining user by id given any other error type than user not found",
			given: in.CreateTransactionRequest{
				Type: in.Deposit,
				From: in.TransactionFrom{
					Amount:        0,
					PaymentMethod: in.PureCash,
				},
				To: in.TransactionTo{
					UserID: 0,
					Amount: 0,
				},
			},
			mocks: func() {
				usersRepositoryMock := new(repository.UsersRepositoryMock)
				usersRepositoryMock.On("GetUserByID", mock.Anything).Return(nil, domain.ErrParsingUserResponse)
				config.Container.Register(config.UsersRestClient, usersRepositoryMock)
			},
			expected: func() (*domain.Transaction, domain.ServiceError) {
				return nil, domain.ErrObtainingUserByID
			},
		},
		{
			name: "Given failure on credit user balance, should return rejected transaction with status detail accreditation processing error",
			given: in.CreateTransactionRequest{
				Type: in.Deposit,
				From: in.TransactionFrom{
					Amount:        100,
					PaymentMethod: in.PureCash,
				},
				To: in.TransactionTo{
					UserID: 1,
					Amount: 100,
				},
			},
			mocks: func() {
				usersRepositoryMock := new(repository.UsersRepositoryMock)
				usersRepositoryMock.On("GetUserByID", mock.Anything).Return(&domain.User{
					UserID: 1,
				}, nil)
				config.Container.Register(config.UsersRestClient, usersRepositoryMock)

				accreditationRepositoryMock := new(repository.AccreditationRepositoryMock)
				accreditationRepositoryMock.On("CreateUserBalanceCredit", mock.Anything).Return(errors.New("error"))
				accreditationService := services.NewAccreditationService(accreditationRepositoryMock)
				config.Container.Register(config.AccreditationService, accreditationService)
			},
			expected: func() (*domain.Transaction, domain.ServiceError) {
				return &domain.Transaction{
					ID:           primitive.ObjectID{},
					UserID:       1,
					Type:         "deposit",
					Status:       "rejected",
					StatusDetail: "accreditation_processing_error",
					From: domain.TransactionFrom{
						Amount:        100,
						PaymentMethod: "pure_cash",
					},
					To: domain.TransactionTo{
						UserID: 1,
						Amount: 100,
					},
					CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				}, domain.ErrProcessingAccreditation
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {

			test.mocks()

			strategy := NewDepositStrategy()
			transaction, err := strategy.Process(test.given)

			expectedTransaction, expectedErr := test.expected()

			suite.Equal(expectedTransaction, transaction)
			suite.Equal(expectedErr, err)
		})
	}
}
