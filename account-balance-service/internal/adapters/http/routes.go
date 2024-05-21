package http

import (
	"account-balance-service/internal/adapters/metrics"

	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine) {
	r.GET("/v1/users/:userId/account_balance", metrics.RecordMetrics(getAccountBalance()))
	r.POST("/v1/users/:userId/account_balance/debit", metrics.RecordMetrics(createAccountDebit()))
}
