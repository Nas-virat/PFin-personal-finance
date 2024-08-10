package router

import (
	"github.com/Nas-virat/PFin-personal-finance/balance"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupBalanceRoutes(app *gin.RouterGroup, db *gorm.DB) {
	balanceRepositoryDB := balance.NewBalanceRepositoryDB(db)
	balanceService := balance.NewBalanceService(balanceRepositoryDB)
	balanceHandler := balance.NewBalanceHandler(balanceService)

	v1 := app.Group("/balance")

	v1.GET("/health", balanceHandler.HealthCheck)
	v1.GET("/summary", balanceHandler.GetSummaryBalanceHandler)
	v1.POST("/debt", balanceHandler.CreateDebtHandler)
}
