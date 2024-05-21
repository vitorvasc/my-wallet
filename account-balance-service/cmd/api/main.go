package main

import (
	"log"

	"account-balance-service/internal/adapters/config"
	"account-balance-service/internal/adapters/http"
	"account-balance-service/internal/adapters/kafka"
	"account-balance-service/internal/adapters/metrics"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.InitializeDependencies()

	http.MapRoutes(r)
	metrics.MapRoutes(r)

	go kafka.Consume()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
