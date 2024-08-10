package router

import (
	"github.com/Nas-virat/PFin-personal-finance/account"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes func
func SetupAccountRoutes(app *gin.RouterGroup, db *gorm.DB) {
	// grouping

	accountRepositoryDB := account.NewAccountRepositoryDB(db)
	accountService := account.NewAccountService(accountRepositoryDB)
	accountHandler := account.NewAccountHandler(accountService)

	// routes
	app.GET("/account", accountHandler.GetAccountsHandler)
	app.GET("/account/:id", accountHandler.GetAccountByIdHandler)
	app.PUT("/account/:id", accountHandler.EditAccountInfoHandler)
	app.GET("/account/health", accountHandler.HealthCheck)
	app.POST("/account/create", accountHandler.CreateAccountHandler)
}
