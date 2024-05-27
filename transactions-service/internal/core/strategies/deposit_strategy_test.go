package strategies

import (
	"testing"
	"time"

	"transactions-service/internal/core/domain"
	in "transactions-service/internal/ports/in/http"

	"github.com/stretchr/testify/suite"
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
			expected: func() (*domain.Transaction, domain.ServiceError) {
				return &domain.Transaction{
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
					CreatedAt: time.Now(), // TODO: Assert time
					UpdatedAt: time.Now(), // TODO: Assert time
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
			expected: func() (*domain.Transaction, domain.ServiceError) {
				return &domain.Transaction{
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
					CreatedAt: time.Now(), // TODO: Assert time
					UpdatedAt: time.Now(), // TODO: Assert time
				}, domain.ErrInvalidAmount
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			strategy := NewDepositStrategy()
			transaction, err := strategy.Process(test.given)

			expectedTransaction, expectedErr := test.expected()

			suite.Equal(expectedTransaction, transaction)
			suite.Equal(expectedErr, err)
		})
	}
}
