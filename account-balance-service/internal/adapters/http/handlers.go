package http

import (
	"net/http"
	"strconv"

	"account-balance-service/internal/adapters/config"
	"account-balance-service/internal/core/services"
	in "account-balance-service/internal/ports/in/http"
	out "account-balance-service/internal/ports/out/http"

	"github.com/gin-gonic/gin"
)

const (
	pathParamUserID = "userId"
)

func getAccountBalance() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Param(pathParamUserID) == "" {
			c.JSON(http.StatusBadRequest, out.NewRequiredFieldApiError(http.StatusBadRequest, pathParamUserID))
			return
		}

		userID, parsingError := strconv.ParseUint(c.Param(pathParamUserID), 10, 64)
		if parsingError != nil {
			c.JSON(http.StatusBadRequest, out.NewInvalidFieldApiError(http.StatusBadRequest, pathParamUserID))
			return
		}

		service := config.Container.Get(config.AccountBalanceService).(services.AccountBalanceService)
		balance, err := service.GetBalance(userID)
		if err != nil {
			c.JSON(err.GetCode(), out.NewApiError(err.GetCode(), err.GetMessage()))
			return
		}

		c.JSON(http.StatusOK, out.BalanceResponse{
			UserID:  balance.UserID,
			Balance: balance.Balance,
		})
	}
}

func createDebit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Param(pathParamUserID) == "" {
			c.JSON(http.StatusBadRequest, out.NewRequiredFieldApiError(http.StatusBadRequest, pathParamUserID))
			return
		}

		userID, parsingError := strconv.ParseUint(c.Param(pathParamUserID), 10, 64)
		if parsingError != nil {
			c.JSON(http.StatusBadRequest, out.NewInvalidFieldApiError(http.StatusBadRequest, pathParamUserID))
			return
		}

		var debitRequest in.DebitRequest
		if err := c.ShouldBindJSON(&debitRequest); err != nil {
			c.JSON(http.StatusBadRequest, out.NewApiError(http.StatusBadRequest, err.Error()))
			return
		}

		service := config.Container.Get(config.AccountBalanceService).(services.AccountBalanceService)
		err := service.AccountDebit(debitRequest.UserID, debitRequest.Amount)
		if err != nil {
			c.JSON(err.GetCode(), out.NewApiError(err.GetCode(), err.GetMessage()))
			return
		}

		c.JSON(http.StatusCreated, out.DebitResponse{
			UserID:  userID,
			Message: "debited successfully",
		})
	}
}

func getUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Param(pathParamUserID) == "" {
			c.JSON(http.StatusBadRequest, out.NewRequiredFieldApiError(http.StatusBadRequest, pathParamUserID))
			return
		}

		userID, parsingError := strconv.ParseUint(c.Param(pathParamUserID), 10, 64)
		if parsingError != nil {
			c.JSON(http.StatusBadRequest, out.NewInvalidFieldApiError(http.StatusBadRequest, pathParamUserID))
			return
		}

		service := config.Container.Get(config.UsersService).(services.UserService)
		user, err := service.GetUserByID(userID)
		if err != nil {
			c.JSON(err.GetCode(), out.NewApiError(err.GetCode(), err.GetMessage()))
			return
		}

		c.JSON(http.StatusOK, out.UserResponse{
			UserID: user.UserID,
		})
	}
}
