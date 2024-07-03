package http

import (
	"account-balance-service/internal/adapters/metrics"

	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine) {
	r.GET("/ping", ping())

	r.GET("/v1/users/:userId/account_balance", metrics.RecordMetrics(getAccountBalance()))
	r.POST("/v1/users/:userId/account_balance/debit", metrics.RecordMetrics(createDebit()))
	r.GET("/v1/users/:userId", metrics.RecordMetrics(getUser()))
}

func ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
