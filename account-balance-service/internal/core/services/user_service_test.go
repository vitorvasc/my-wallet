package services

import (
	"account-balance-service/internal/core/domain"
	"account-balance-service/internal/ports/out/db"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	repository *db.UsersRepositoryMock
}

func TestService(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

func (suite *UserServiceTestSuite) setup() UserService {
	suite.repository = new(db.UsersRepositoryMock)
	return NewUserService(suite.repository)
}

func (suite *UserServiceTestSuite) TestGetUserByID() {
	service := suite.setup()

	userID := uint64(1)
	user := &domain.User{
		UserID: userID,
	}

	suite.repository.On("GetUserByID", userID).Return(user, nil)

	result, err := service.GetUserByID(userID)

	suite.repository.AssertExpectations(suite.T())
	suite.Nil(err)
	suite.Equal(user, result)
}

func (suite *UserServiceTestSuite) TestGetUserByIDNotFound() {
	service := suite.setup()

	userID := uint64(1)

	suite.repository.On("GetUserByID", userID).Return(nil, domain.ErrUserNotFound)

	result, err := service.GetUserByID(userID)

	suite.repository.AssertExpectations(suite.T())
	suite.Error(err)
	suite.Nil(result)
}

func (suite *UserServiceTestSuite) TestGetUserByIDError() {
	service := suite.setup()

	userID := uint64(1)

	suite.repository.On("GetUserByID", userID).Return(nil, domain.ErrObtainingUser)

	result, err := service.GetUserByID(userID)

	suite.repository.AssertExpectations(suite.T())
	suite.Error(err)
	suite.Nil(result)
}
