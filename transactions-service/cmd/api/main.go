package main

import (
	"context"
	"log"
	"time"
	"transactions-service/internal/core/usecase"
	"transactions-service/internal/core/utils"

	in "transactions-service/internal/adapters/in/http"
	"transactions-service/internal/adapters/out/db"
	out "transactions-service/internal/adapters/out/http"
	"transactions-service/internal/adapters/out/kafka"
	"transactions-service/internal/adapters/out/metrics"
	"transactions-service/internal/app/config"
	"transactions-service/internal/core/services"
	"transactions-service/internal/core/strategies"

	"github.com/gin-gonic/gin"
)

func main() {
	clock := utils.NewClock()
	config.Container.Register(config.Clock, clock)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	database := db.InitDB(ctx)

	mongoRepository := db.NewMongoRepository(database)
	config.Container.Register(config.MongoRepository, mongoRepository)

	usersRestClient := out.NewUsersRestClient()
	config.Container.Register(config.UsersRestClient, usersRestClient)

	kafkaProducer := kafka.NewKafkaProducer()
	defer kafkaProducer.ProducerClose()

	accreditationService := services.NewAccreditationService(kafkaProducer)
	config.Container.Register(config.AccreditationService, accreditationService)

	transactionStrategies := []strategies.HandleTransactionStrategy{
		strategies.NewAccountTransferStrategy(),
		strategies.NewBillPaymentStrategy(),
		strategies.NewDepositStrategy(),
		strategies.NewWithdrawalStrategy(),
	}

	transactionService := usecase.NewTransactionUseCase(transactionStrategies)

	config.Container.Register(config.TransactionService, transactionService)

	r := gin.Default()

	in.MapRoutes(r)
	metrics.MapRoutes(r)

	if err := r.Run(":8081"); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
