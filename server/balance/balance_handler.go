package balance

import (
	"github.com/Nas-virat/PFin-personal-finance/response"
	"github.com/gofiber/fiber/v2"
)


type balanceHandler struct{
	balanceSrv BalanceService
}

func NewBalanceHandler(balanceSrv BalanceService) balanceHandler {
	return balanceHandler{balanceSrv: balanceSrv}
}

func (h balanceHandler) HealthCheck(c *fiber.Ctx) error {
	return response.NewSuccessResponse(c, "Health check", fiber.StatusOK, nil)
}

//  GetSummaryBalance	godoc
//	@Summary			Get Summary Balance
//	@Description		Get Summary Balance
//	@Tags				Balance	
//	@Accept				json
//	@Produce			json
//	@Success			200		
//	@Router				/api/balance/summmary [get]
func (h balanceHandler) GetSummaryBalanceHandler(c *fiber.Ctx) error {
	summaryBalance, err := h.balanceSrv.GetSummaryBalance()
	if err != nil{
		return response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	return response.NewSuccessResponse(c, "Get summary balance", fiber.StatusOK, summaryBalance)
}

//  CreateDebt			godoc
//	@Summary			Create Debt
//	@Description		Create Debt
//	@Tags				Balance	
//	@Accept				json
//	@Produce			json
//	@Success			201	
//	@Router				/api/balance/debt [post]
func (h balanceHandler) CreateDebtHandler(c *fiber.Ctx) error {

	request := NewDebtRequest{}

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


