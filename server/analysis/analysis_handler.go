package analysis

import (
	"github.com/Nas-virat/PFin-personal-finance/account"
	"github.com/gin-gonic/gin"
)

type analysisHandler struct {
	accSrv account.AccountService
}

func NewAnalysisHandler(accSrv account.AccountService) analysisHandler {
	return analysisHandler{accSrv: accSrv}
}

func (h analysisHandler) GetAnalysisWealth(c *gin.Context) error {
	return nil
}

func (h analysisHandler) GetNetWorthByMonth(c *gin.Context) error {
	return nil
}

func (h analysisHandler) GetFreeCashFlowByMonth(c *gin.Context) error {
	return nil
}
