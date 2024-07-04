package http

import (
	"net/http"
	"strconv"

	"transactions-service/internal/app/config"
	"transactions-service/internal/core/usecase"
	in "transactions-service/internal/ports/in/http"
	out "transactions-service/internal/ports/out/dto"

	"github.com/gin-gonic/gin"
)

const (
	pathParamUserID        = "userId"
	pathParamTransactionID = "transactionId"
)

func createTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var createTransactionRequest in.CreateTransactionRequest
		if err := c.ShouldBindJSON(&createTransactionRequest); err != nil {
			c.JSON(http.StatusBadRequest, out.NewApiError(http.StatusBadRequest, err.Error()))
			return
		}

		service := config.Container.Get(config.TransactionService).(usecase.TransactionUseCase)
		transaction, err := service.CreateTransaction(createTransactionRequest)
		if err != nil {
			c.JSON(err.GetCode(), out.NewApiError(err.GetCode(), err.GetMessage()))
			return
		}

		c.JSON(http.StatusCreated, out.TransactionResponse{
			UserID: transaction.ID,
			Status: string(transaction.Status),
		})
	}
}

func updateTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Param(pathParamTransactionID) == "" {
			c.JSON(http.StatusBadRequest, out.NewRequiredFieldApiError(http.StatusBadRequest, pathParamTransactionID))
			return
		}

		transactionID, parsingError := strconv.ParseUint(c.Param(pathParamTransactionID), 10, 64)
		if parsingError != nil {
			c.JSON(http.StatusBadRequest, out.NewInvalidFieldApiError(http.StatusBadRequest, pathParamTransactionID))
			return
		}

		var updateTransactionRequest in.UpdateTransactionRequest
		if err := c.ShouldBindJSON(&updateTransactionRequest); err != nil {
			c.JSON(http.StatusBadRequest, out.NewApiError(http.StatusBadRequest, err.Error()))
			return
		}

		service := config.Container.Get(config.TransactionService).(usecase.TransactionUseCase)
		transaction, err := service.UpdateTransaction(transactionID, updateTransactionRequest)
		if err != nil {
			c.JSON(err.GetCode(), out.NewApiError(err.GetCode(), err.GetMessage()))
			return
		}

		c.JSON(http.StatusCreated, out.TransactionResponse{
			ID:     transaction.ID,
			Status: string(transaction.Status),
		})
	}
}

func getUserTransactions() gin.HandlerFunc {
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

		service := config.Container.Get(config.TransactionService).(usecase.TransactionUseCase)
		transactions, err := service.GetTransactionsByUser(userID)
		if err != nil {
			c.JSON(err.GetCode(), out.NewApiError(err.GetCode(), err.GetMessage()))
			return
		}

		c.JSON(http.StatusOK, transactions)
	}
}
