package handlers

import (
	"github.com/Nas-virat/PFin-personal-finance/service"
	"github.com/gofiber/fiber/v2"
)


type balanceHandler struct{
	balanceSrv service.BalanceService
}

func NewBalanceHandler(balanceSrv service.BalanceService) balanceHandler {
	return balanceHandler{balanceSrv: balanceSrv}
}

func (h balanceHandler) GetSummaryBalanceHandler(c *fiber.Ctx) error {
	summaryBalance, err := h.balanceSrv.GetSummaryBalance()
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "success",
		"message": "get all accounts",
		"data": summaryBalance,
	})
}

func (h balanceHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "success",
		"message": "health check",
	})
}


