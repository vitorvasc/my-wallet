package main

import (
	"log"

	"account-balance-service/internal/adapters/config"
	"account-balance-service/internal/adapters/db/postgres"
	"account-balance-service/internal/adapters/http"
	"account-balance-service/internal/adapters/kafka"
	"account-balance-service/internal/adapters/metrics"
	"account-balance-service/internal/core/services"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConn := postgres.InitDB()
	defer dbConn.Close()

	pgRepository := postgres.NewPostgresRepository(dbConn)

	userService := services.NewUserService(pgRepository)
	accountBalanceService := services.NewAccountBalanceService(pgRepository, userService)

	config.Container.Register(config.AccountBalanceService, accountBalanceService)
	config.Container.Register(config.UsersService, userService)

	r := gin.Default()

	http.MapRoutes(r)
	metrics.MapRoutes(r)

	go kafka.Consume()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
