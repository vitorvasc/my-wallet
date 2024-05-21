package config

import (
	"account-balance-service/internal/adapters/db"
	"account-balance-service/internal/core/services"
)

var Container = NewContainer()

func InitializeDependencies() {
	dbConn := db.InitDB()
	defer dbConn.Close()
	
	accountBalanceRepository := db.NewPostgresRepository(dbConn)
	accountBalanceService := services.NewAccountBalanceService(accountBalanceRepository)

	Container.Register(AccountBalanceService, accountBalanceService)
}
