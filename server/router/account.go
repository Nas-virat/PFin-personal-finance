package router

import (
	"github.com/Nas-virat/PFin-personal-finance/handlers"
	"github.com/Nas-virat/PFin-personal-finance/repository"
	"github.com/Nas-virat/PFin-personal-finance/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes func
func SetupAccountRoutes(app *fiber.App, db *gorm.DB) {
 // grouping

	accountRepositoryDB := repository.NewAccountRepositoryDB(db)
	accountService := service.NewAccountService(accountRepositoryDB)
	accountHandler := handlers.NewAccountHandler(accountService)

	api := app.Group("/api")
	v1 := api.Group("/account")
	// routes
	v1.Get("/id/:id", accountHandler.GetAccountByIdHandler)
	v1.Get("/health", accountHandler.HealthCheck)
	v1.Post("/create", accountHandler.CreateAccountHandler)
	v1.Get("/getaccounts", accountHandler.GetAccountsHandler)
	v1.Put("/edit/:id", accountHandler.EditAccountInfoHandler)
}