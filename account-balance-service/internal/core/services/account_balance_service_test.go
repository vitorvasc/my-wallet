package services

import (
	"testing"

	"account-balance-service/internal/core/domain"
	"account-balance-service/internal/ports/out/repository"

	"github.com/stretchr/testify/suite"
)

type AccountBalanceServiceTestSuite struct {
	suite.Suite
	repository  *repository.AccountBalanceRepositoryMock
	userService *UserServiceMock
}

func TestAccountBalanceService(t *testing.T) {
	suite.Run(t, new(AccountBalanceServiceTestSuite))
}

func (suite *AccountBalanceServiceTestSuite) setup() AccountBalanceService {
	suite.repository = new(repository.AccountBalanceRepositoryMock)
	suite.userService = new(UserServiceMock)
	return NewAccountBalanceService(suite.repository, suite.userService)
}

func (suite *AccountBalanceServiceTestSuite) TestGetAccountBalance() {
	service := suite.setup()

	userID := uint64(1)
	user := &domain.User{
		UserID: userID,
	}

	accountBalance := &domain.AccountBalance{
		UserID:  userID,
		Balance: float64(100),
	}

	suite.userService.On("GetUserByID", userID).Return(user, nil)
	suite.repository.On("GetAccountBalance", userID).Return(float64(100), nil)

	result, err := service.GetBalance(userID)

	suite.repository.AssertExpectations(suite.T())
	suite.Nil(err)
	suite.Equal(accountBalance, result)
}

func (suite *AccountBalanceServiceTestSuite) TestGetAccountBalanceUserIDNotFound() {
	service := suite.setup()

	accountBalanceID := uint64(1)

	suite.userService.On("GetUserByID", accountBalanceID).Return(nil, domain.ErrUserNotFound)

	result, err := service.GetBalance(accountBalanceID)

	suite.repository.AssertExpectations(suite.T())
	suite.Error(err)
	suite.Nil(result)
}

func (suite *AccountBalanceServiceTestSuite) TestGetAccountBalanceUserUnexpectedError() {
	service := suite.setup()

	accountBalanceID := uint64(1)

	suite.userService.On("GetUserByID", accountBalanceID).Return(nil, domain.ErrObtainingAccountBalance)

	result, err := service.GetBalance(accountBalanceID)

	suite.repository.AssertExpectations(suite.T())
	suite.Error(err)
	suite.Nil(result)
}

func (suite *AccountBalanceServiceTestSuite) TestGetAccountBalanceUnexpectedError() {
	service := suite.setup()

	accountBalanceID := uint64(1)
	user := &domain.User{
		UserID: accountBalanceID,
	}

	suite.userService.On("GetUserByID", accountBalanceID).Return(user, nil)
	suite.repository.On("GetAccountBalance", accountBalanceID).Return(float64(0), domain.ErrObtainingAccountBalance)

	result, err := service.GetBalance(accountBalanceID)

	suite.repository.AssertExpectations(suite.T())
	suite.Error(err)
	suite.Nil(result)
}

func (suite *AccountBalanceServiceTestSuite) TestAccountCredit() {
	service := suite.setup()

	userID := uint64(1)
	amount := float64(100)

	user := &domain.User{
		UserID: userID,
	}

	suite.userService.On("GetUserByID", userID).Return(user, nil)
	suite.repository.On("Credit", userID, amount).Return(nil)

	err := service.AccountCredit(userID, amount)

	suite.repository.AssertExpectations(suite.T())
	suite.Nil(err)
}

func (suite *AccountBalanceServiceTestSuite) TestAccountCreditUserIDNotFound() {
	service := suite.setup()

	accountBalanceID := uint64(1)
	amount := float64(100)

	suite.userService.On("GetUserByID", accountBalanceID).Return(nil, domain.ErrUserNotFound)

	err := service.AccountCredit(accountBalanceID, amount)

	suite.repository.AssertExpectations(suite.T())
	suite.Error(err)
}

func (suite *AccountBalanceServiceTestSuite) TestAccountCreditError() {
	service := suite.setup()

	accountBalanceID := uint64(1)
	amount := float64(100)

	user := &domain.User{
		UserID: accountBalanceID,
	}

	suite.userService.On("GetUserByID", accountBalanceID).Return(user, nil)
	suite.repository.On("Credit", accountBalanceID, amount).Return(domain.ErrAccreditingValue)

	err := service.AccountCredit(accountBalanceID, amount)

	suite.repository.AssertExpectations(suite.T())
	suite.Error(err)
}

func (suite *AccountBalanceServiceTestSuite) TestAccountDebit() {
	service := suite.setup()

	userID := uint64(1)
	amount := float64(100)

	user := &domain.User{
		UserID: userID,
	}

	suite.userService.On("GetUserByID", userID).Return(user, nil)
	suite.repository.On("GetAccountBalance", userID).Return(float64(200), nil)
	suite.repository.On("Debit", userID, amount).Return(nil)

	err := service.AccountDebit(userID, amount)

	suite.repository.AssertExpectations(suite.T())
	suite.Nil(err)
}

func (suite *AccountBalanceServiceTestSuite) TestAccountDebitUserIDNotFound() {
	service := suite.setup()

	accountBalanceID := uint64(1)
	amount := float64(100)

	suite.userService.On("GetUserByID", accountBalanceID).Return(nil, domain.ErrUserNotFound)

	err := service.AccountDebit(accountBalanceID, amount)

	suite.repository.AssertExpectations(suite.T())
	suite.Error(err)
}

func (suite *AccountBalanceServiceTestSuite) TestAccountDebitErrorObtainingAccountBalance() {
	service := suite.setup()

	accountBalanceID := uint64(1)
	amount := float64(100)

	user := &domain.User{
		UserID: accountBalanceID,
	}

	suite.userService.On("GetUserByID", accountBalanceID).Return(user, nil)
	suite.repository.On("GetAccountBalance", accountBalanceID).Return(float64(200), domain.ErrObtainingAccountBalance)

	err := service.AccountDebit(accountBalanceID, amount)

	suite.repository.AssertExpectations(suite.T())
	suite.Error(err)
}

func (suite *AccountBalanceServiceTestSuite) TestAccountDebitInsufficientFundsError() {
	service := suite.setup()

	accountBalanceID := uint64(1)
	amount := float64(100)

	user := &domain.User{
		UserID: accountBalanceID,
	}

	suite.userService.On("GetUserByID", accountBalanceID).Return(user, nil)
	suite.repository.On("GetAccountBalance", accountBalanceID).Return(float64(50), nil)

	err := service.AccountDebit(accountBalanceID, amount)

	suite.repository.AssertExpectations(suite.T())
	suite.Error(err)
}

func (suite *AccountBalanceServiceTestSuite) TestAccountDebitUnexpectedError() {
	service := suite.setup()

	accountBalanceID := uint64(1)
	amount := float64(100)

	user := &domain.User{
		UserID: accountBalanceID,
	}

	suite.userService.On("GetUserByID", accountBalanceID).Return(user, nil)
	suite.repository.On("GetAccountBalance", accountBalanceID).Return(float64(200), nil)
	suite.repository.On("Debit", accountBalanceID, amount).Return(domain.ErrDebitingValue)

	err := service.AccountDebit(accountBalanceID, amount)

	suite.repository.AssertExpectations(suite.T())
	suite.Error(err)
}
