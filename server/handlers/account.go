package handlers

import (
	"strconv"

	"github.com/Nas-virat/PFin-personal-finance/model"
	"github.com/Nas-virat/PFin-personal-finance/response"
	"github.com/Nas-virat/PFin-personal-finance/service"
	"github.com/gofiber/fiber/v2"
)


type accountHandler struct {
	accSrv service.AccountService
}

func NewAccountHandler(accSrv service.AccountService) accountHandler{
	return accountHandler{accSrv:accSrv}
}

// CreateAccountHandler godoc
// @Summary      CreateAccount
// @Description  Create Account
// @Tags         account
// @Accept       json
// @Produce      json
// @Success      201  {object}  model.NewAccountRequest
// @Router /api/account/create [post]
func (h accountHandler) CreateAccountHandler(c *fiber.Ctx) error {

	request := model.NewAccountRequest{}

	err := c.BodyParser(&request)

	if err != nil{
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	accountResponse ,err := h.accSrv.CreateAccount(request)

	if err != nil{
		return response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	return response.NewSuccessResponse(c, "insert successfully", fiber.StatusCreated, accountResponse)
}

func (h accountHandler) GetAccountByIdHandler(c *fiber.Ctx) error{

	id := c.Params("id")
	
	accountId, err := strconv.Atoi(id)
	if err != nil {
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}
	
	account, err := h.accSrv.GetAccountById(accountId)

	if err != nil{
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, "get account by id", fiber.StatusOK, account)
}

func (h accountHandler) GetAccountsHandler(c *fiber.Ctx) error {

	accounts, err := h.accSrv.GetAccounts()

	if err != nil{
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, "get all accounts", fiber.StatusOK, accounts)
}

func (h accountHandler) EditAccountInfoHandler(c *fiber.Ctx) error {
	
	request := model.NewAccountRequest{}

	id, err := c.ParamsInt("id")

	if err != nil || id < 0{
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	err = c.BodyParser(&request)

	if err != nil{
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	accountResponse, err := h.accSrv.EditAccountInfo(request,id)

	if err != nil{
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}
	
	return response.NewSuccessResponse(c, "edit account info", fiber.StatusOK, accountResponse)
}
 
func (h accountHandler) HealthCheck(c *fiber.Ctx) error {
	return response.NewSuccessResponse(c, "health check", fiber.StatusOK, nil)
}







