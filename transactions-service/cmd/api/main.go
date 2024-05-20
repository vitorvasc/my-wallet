package main

import (
	"log"
	"transactions-service/internal/adapters/db"
	"transactions-service/internal/adapters/http"
	"transactions-service/internal/adapters/metrics"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDB()

	_ = db.NewMongoRepository(database)

	r := gin.Default()

	http.MapRoutes(r)
	metrics.MapRoutes(r)

	if err := r.Run(":8081"); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
