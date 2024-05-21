package http

import (
	"account-balance-service/internal/core/services"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	pathParamUserId = "userId"
)

func getAccountBalance(service services.AccountBalance) gin.HandlerFunc {
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

		balance, err := service.GetBalance(userId)
		if err != nil {
			log.Printf("Error obtaining account balance: %v", err)
			c.JSON(500, err)
			return
		}

		c.JSON(200, gin.H{
			"user_id": userId,
			"balance": balance,
		})
	}
}

func createDebit() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
