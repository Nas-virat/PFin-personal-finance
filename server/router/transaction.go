package router

import (
	"github.com/Nas-virat/PFin-personal-finance/transaction"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func SetupTransactionRoute(app *fiber.App, db *gorm.DB){
	transactionRepositoryDB := transaction.NewTransactionRepositoryDB(db)
	transactionService := transaction.NewTransactionService(transactionRepositoryDB)
	transactionHandler := transaction.NewTransactionHandler(transactionService)

	api := app.Group("/api")
 	v1 := api.Group("/transaction")

	v1.Get("/health",transactionHandler.HealthCheck)
	v1.Post("/create",transactionHandler.CreateTransactionHandler)
	v1.Get("/all",transactionHandler.GetTransactionsHandler)
	v1.Get("/summary-year",transactionHandler.GetSummaryRevenueExpenseHandler)
	v1.Get("/month/:month/year/:year",transactionHandler.GetTransactionInRangeMonthYearHandler)
	v1.Get("/day/:day/month/:month/year/:year",transactionHandler.GetTransactionInRangeDayMonthYearHandler)
	v1.Get("/:id",transactionHandler.GetTransactionByIDHandler)
	v1.Put("/:id",transactionHandler.UpdateTransactionHandler)
	v1.Delete("/:id",transactionHandler.DeleteTransactionHandler)
}