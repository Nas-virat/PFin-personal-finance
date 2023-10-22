package router

import (
	"github.com/Nas-virat/PFin-personal-finance/handlers"
	"github.com/Nas-virat/PFin-personal-finance/repository"
	"github.com/Nas-virat/PFin-personal-finance/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func SetupTransactionRoute(app *fiber.App, db *gorm.DB){
	transactionRepositoryDB := repository.NewTransactionRepositoryDB(db)
	transactionService := service.NewTransactionService(transactionRepositoryDB)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	api := app.Group("/api")
 	v1 := api.Group("/transaction")

	v1.Post("/create",transactionHandler.CreateTransactionHandler)
	v1.Get("/all",transactionHandler.GetTransactionsHandler)
	v1.Get("/month/:month/year/:year",transactionHandler.GetTransactionInRanageMonthYearHandler)
	v1.Get("/:id",transactionHandler.GetTransactionByIDHandler)
	v1.Put("/:id",transactionHandler.UpdateTransactionHandler)
	v1.Delete("/:id",transactionHandler.DeleteTransactionHandler)
}