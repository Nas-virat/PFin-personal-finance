package balance

import (
	"net/http"

	"github.com/Nas-virat/PFin-personal-finance/response"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)


type balanceHandler struct{
	balanceSrv BalanceService
}

func NewBalanceHandler(balanceSrv BalanceService) balanceHandler {
	return balanceHandler{balanceSrv: balanceSrv}
}

func (h balanceHandler) HealthCheck(c *gin.Context){
	response.NewSuccessResponse(c, "Health check", http.StatusOK, nil)
}

//  GetSummaryBalance	godoc
//	@Summary			Get Summary Balance
//	@Description		Get Summary Balance
//	@Tags				Balance	
//	@Accept				json
//	@Produce			json
//	@Success			200		
//	@Router				/api/balance/summmary [get]
func (h balanceHandler) GetSummaryBalanceHandler(c *gin.Context){
	summaryBalance, err := h.balanceSrv.GetSummaryBalance()
	if err != nil{
		response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	response.NewSuccessResponse(c, "Get summary balance", http.StatusOK, summaryBalance)
}

//  CreateDebt			godoc
//	@Summary			Create Debt
//	@Description		Create Debt
//	@Tags				Balance	
//	@Accept				json
//	@Produce			json
//	@Success			201	
//	@Router				/api/balance/debt [post]
func (h balanceHandler) CreateDebtHandler(c *gin.Context) {

	request := NewDebtRequest{}

	err := c.ShouldBindJSON(&request)

	if err != nil{
		response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	debtResponse, err := h.balanceSrv.CreateDebt(request)
	if err != nil{
		response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	response.NewSuccessResponse(c, "insert successfully", http.StatusCreated, debtResponse)
}


