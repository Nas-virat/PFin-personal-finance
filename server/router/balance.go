package router

import (
	"github.com/Nas-virat/PFin-personal-finance/balance"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupBalanceRoutes(app *fiber.App, db *gorm.DB) {
	balanceRepositoryDB := balance.NewBalanceRepositoryDB(db)
	balanceService := balance.NewBalanceService(balanceRepositoryDB)
	balanceHandler := balance.NewBalanceHandler(balanceService)
	
	api := app.Group("/api")
 	v1 := api.Group("/balance")

	v1.Get("/health",balanceHandler.HealthCheck)
	v1.Get("/summary",balanceHandler.GetSummaryBalanceHandler)
	v1.Post("/debt",balanceHandler.CreateDebtHandler)
}