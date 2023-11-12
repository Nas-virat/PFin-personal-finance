package handlers

import (
	"github.com/Nas-virat/PFin-personal-finance/model"
	"github.com/Nas-virat/PFin-personal-finance/response"
	"github.com/Nas-virat/PFin-personal-finance/service"
	"github.com/gofiber/fiber/v2"
)


type balanceHandler struct{
	balanceSrv service.BalanceService
}

func NewBalanceHandler(balanceSrv service.BalanceService) balanceHandler {
	return balanceHandler{balanceSrv: balanceSrv}
}

func (h balanceHandler) HealthCheck(c *fiber.Ctx) error {
	return response.NewSuccessResponse(c, "Health check", fiber.StatusOK, nil)
}

func (h balanceHandler) GetSummaryBalanceHandler(c *fiber.Ctx) error {
	summaryBalance, err := h.balanceSrv.GetSummaryBalance()
	if err != nil{
		return response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	return response.NewSuccessResponse(c, "Get summary balance", fiber.StatusOK, summaryBalance)
}

func (h balanceHandler) CreateDebtHandler(c *fiber.Ctx) error {

	request := model.NewDebtRequest{}

	err := c.BodyParser(&request)

	if err != nil{
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	debtResponse, err := h.balanceSrv.CreateDebt(request)
	if err != nil{
		return response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	return response.NewSuccessResponse(c, "insert successfully", fiber.StatusCreated, debtResponse)
}


