package router

import (
	"github.com/Nas-virat/PFin-personal-finance/transaction"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func SetupTransactionRoute(app *gin.RouterGroup, db *gorm.DB){
	transactionRepositoryDB := transaction.NewTransactionRepositoryDB(db)
	transactionService := transaction.NewTransactionService(transactionRepositoryDB)
	transactionHandler := transaction.NewTransactionHandler(transactionService)

 	v1 := app.Group("/transaction")

	v1.GET("/health",transactionHandler.HealthCheck)
	v1.POST("/",transactionHandler.CreateTransactionHandler)
	v1.GET("/",transactionHandler.GetTransactionsHandler)
	v1.GET("/:id",transactionHandler.GetTransactionByIDHandler)
	v1.PUT("/:id",transactionHandler.UpdateTransactionByIdHandler)
	v1.DELETE("/:id",transactionHandler.DeleteTransactionHandler)
	v1.GET("/summary-year",transactionHandler.GetSummaryRevenueExpenseHandler)
	v1.GET("/month/:month/year/:year",transactionHandler.GetTransactionInRangeMonthYearHandler)
	v1.GET("/day/:day/month/:month/year/:year",transactionHandler.GetTransactionInRangeDayMonthYearHandler)
	
}