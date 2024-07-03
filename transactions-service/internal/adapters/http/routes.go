package http

import (
	"transactions-service/internal/adapters/metrics"

	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine) {
	r.GET("/ping", ping())

	r.POST("/v1/transactions", metrics.RecordMetrics(createTransaction()))
	r.PUT("/v1/transactions/:transactionId", metrics.RecordMetrics(updateTransaction()))
	r.GET("/v1/transactions/:userId/timeline", metrics.RecordMetrics(getUserTransactions()))
}

func ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
