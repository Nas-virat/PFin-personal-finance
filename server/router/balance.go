package router

import (
	"github.com/Nas-virat/PFin-personal-finance/handlers"
	"github.com/Nas-virat/PFin-personal-finance/repository"
	"github.com/Nas-virat/PFin-personal-finance/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupBalanceRoutes(app *fiber.App, db *gorm.DB) {
	balanceRepositoryDB := repository.NewBalanceRepositoryDB(db)
	balanceService := service.NewBalanceService(balanceRepositoryDB)
	balanceHandler := handlers.NewBalanceHandler(balanceService)
	
	api := app.Group("/api")
 	v1 := api.Group("/balance")

	v1.Get("/health",balanceHandler.HealthCheck)
	v1.Get("/summary",balanceHandler.GetSummaryBalanceHandler)
}