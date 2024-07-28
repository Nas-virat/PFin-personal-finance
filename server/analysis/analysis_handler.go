package analysis

import (
	"github.com/Nas-virat/PFin-personal-finance/account"
	"github.com/gofiber/fiber/v2"
)

type analysisHandler struct {
	accSrv account.AccountService
}

func NewAnalysisHandler(accSrv account.AccountService) analysisHandler {
	return analysisHandler{accSrv: accSrv}
}

func (h analysisHandler) GetAnalysisWealth(c *fiber.Ctx) error {
	return nil
}

func (h analysisHandler) GetNetWorthByMonth(c *fiber.Ctx) error {
	return nil
}

func (h analysisHandler) GetFreeCashFlowByMonth(c *fiber.Ctx) error {
	return nil
}
