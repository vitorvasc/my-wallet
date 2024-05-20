package http

import (
	"transactions-service/internal/adapters/metrics"

	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine) {
	r.POST("/v1/transactions", metrics.RecordMetrics(createTransaction()))
	r.PUT("/v1/transactions/:transactionId", metrics.RecordMetrics(updateTransaction()))
	r.GET("/v1/transactions/:userId/timeline", metrics.RecordMetrics(getUserTransactions()))
}
