package handlers

import (
	"github.com/Nas-virat/PFin-personal-finance/service"
	"github.com/gofiber/fiber/v2"
)

type analysisHandler struct {
	accSrv service.AccountService
}

func NewAnalysisHandler(accSrv service.AccountService) accountHandler {
	return accountHandler{accSrv: accSrv}
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
