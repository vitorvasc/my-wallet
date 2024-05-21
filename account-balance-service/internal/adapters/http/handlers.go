package http

import (
	"account-balance-service/internal/adapters/config"
	"account-balance-service/internal/core/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	pathParamUserId = "userId"
)

func getAccountBalance() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Param(pathParamUserId) == "" {
			c.JSON(400, gin.H{
				"message": "userId is required",
			})
			return
		}

		userId, err := strconv.ParseUint(c.Param(pathParamUserId), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "userId must be valid",
			})
			return
		}

		accountBalanceService := config.Container.Get(config.AccountBalanceService).(services.AccountBalance)
		balance, err := accountBalanceService.GetBalance(userId)
		if err != nil {
			return
		}

		c.JSON(200, gin.H{
			"user_id": userId,
			"balance": balance,
		})
	}
}

func createAccountDebit() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
