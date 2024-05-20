package main

import (
	"log"

	"account-balance-service/internal/adapters/db"
	"account-balance-service/internal/adapters/http"
	"account-balance-service/internal/adapters/kafka"
	"account-balance-service/internal/adapters/metrics"
	"account-balance-service/internal/core/services"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConn := db.InitDB()
	defer dbConn.Close()

	accountBalanceRepository := db.NewPostgresRepository(dbConn)

	accountBalanceService := services.NewAccountBalanceService(accountBalanceRepository)

	go kafka.Consume(accountBalanceService)

	r := gin.Default()

	http.MapRoutes(r)
	metrics.MapRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
