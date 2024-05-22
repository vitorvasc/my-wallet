package main

import (
	"context"
	"log"
	"time"
	"transactions-service/internal/adapters/kafka"

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

	kafkaProducer := kafka.NewKafkaProducer()
	defer kafkaProducer.ProducerClose()

	transactionStrategies := []strategies.HandleTransactionStrategy{
		strategies.NewAccountTransferStrategy(),
		strategies.NewBillPaymentStrategy(),
		strategies.NewDepositStrategy(),
		strategies.NewWithdrawalStrategy(),
	}

	accreditationService := services.NewAccreditationService(kafkaProducer)
	transactionService := services.NewTransactionService(
		mongoRepository,
		transactionStrategies,
		accreditationService,
	)

	config.Container.Register(config.TransactionService, transactionService)

	r := gin.Default()

	http.MapRoutes(r)
	metrics.MapRoutes(r)

	if err := r.Run(":8081"); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
