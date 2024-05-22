package main

import (
	"context"
	"log"
	"time"

	"transactions-service/internal/adapters/config"
	"transactions-service/internal/adapters/db"
	"transactions-service/internal/adapters/http"
	"transactions-service/internal/adapters/metrics"
	"transactions-service/internal/core/services"
	"transactions-service/internal/core/strategies"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	database := db.InitDB(ctx)

	mongoRepository := db.NewMongoRepository(database)

	transactionStrategies := []strategies.HandleTransactionStrategy{
		strategies.NewAccountTransferStrategy(mongoRepository),
		strategies.NewBillPaymentStrategy(mongoRepository),
		strategies.NewDepositStrategy(mongoRepository),
		strategies.NewWithdrawalStrategy(mongoRepository),
	}

	transactionService := services.NewTransactionService(mongoRepository, transactionStrategies)

	config.Container.Register(config.TransactionService, transactionService)

	r := gin.Default()

	http.MapRoutes(r)
	metrics.MapRoutes(r)

	if err := r.Run(":8081"); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
