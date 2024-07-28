package router

import (
	"github.com/Nas-virat/PFin-personal-finance/account"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes func
func SetupAccountRoutes(app *fiber.App, db *gorm.DB) {
 // grouping

	accountRepositoryDB := account.NewAccountRepositoryDB(db)
	accountService := account.NewAccountService(accountRepositoryDB)
	accountHandler := account.NewAccountHandler(accountService)

	api := app.Group("/api")
	v1 := api.Group("/account")
	// routes
	v1.Get("/", accountHandler.GetAccountsHandler)
	v1.Get("/:id", accountHandler.GetAccountByIdHandler)
	v1.Put("/:id", accountHandler.EditAccountInfoHandler)
	v1.Get("/health", accountHandler.HealthCheck)
	v1.Post("/create", accountHandler.CreateAccountHandler)
}