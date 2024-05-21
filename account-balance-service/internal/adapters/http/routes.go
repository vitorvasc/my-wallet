package http

import (
	"account-balance-service/internal/adapters/metrics"
	"account-balance-service/internal/core/services"

	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine, service services.AccountBalance) {
	r.GET("/v1/users/:userId/account_balance", metrics.RecordMetrics(getAccountBalance(service)))
	r.POST("/v1/users/:userId/account_balance/debit", metrics.RecordMetrics(createDebit()))
}
