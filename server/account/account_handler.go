package account

import (
	"net/http"
	"strconv"

	"github.com/Nas-virat/PFin-personal-finance/response"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)


type accountHandler struct {
	accSrv AccountService
}

func NewAccountHandler(accSrv AccountService) accountHandler{
	return accountHandler{accSrv:accSrv}
}

// CreateAccountHandler godoc
//	@Summary		CreateAccount
//	@Description	Create Account
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	NewAccountRequest
//	@Router			/api/account/create [post]
func (h accountHandler) CreateAccountHandler(c *gin.Context) {

	request := NewAccountRequest{}

	err := c.ShouldBindJSON(&request)

	if err != nil{
		response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	accountResponse ,err := h.accSrv.CreateAccount(request)

	if err != nil{
		response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	response.NewSuccessResponse(c, "insert successfully", http.StatusCreated, accountResponse)
}

//  GetAccountById 	godoc
//	@Summary		GetAccountById
//	@Description	Get account by id
//	@Tags			account	
//	@Accept			json
//	@Produce		json
//  @Param        	id   	path     int  true  "Account ID"
//	@Success		200	{object}	AccountResponse
//	@Router			/api/account/{id} [get]
func (h accountHandler) GetAccountByIdHandler(c *gin.Context) {

	id := c.Param("id")
	
	accountId, err := strconv.Atoi(id)
	if err != nil {
		response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}
	
	account, err := h.accSrv.GetAccountById(accountId)

	if err != nil{
		response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	response.NewSuccessResponse(c, "get account by id", fiber.StatusOK, account)
}

//  GetAccounts		godoc
//	@Summary		GetAccounts
//	@Description	Get AllAccounts
//	@Tags			account	
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	AccountResponse	
//	@Router			/api/account [get]
func (h accountHandler) GetAccountsHandler(c *gin.Context) {

	accounts, err := h.accSrv.GetAccounts()

	if err != nil{
		response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	response.NewSuccessResponse(c, "get all accounts", fiber.StatusOK, accounts)
}

//  EditAccountInfo	godoc
//	@Summary		Edit Account Info
//	@Description	Edit Account Info
//	@Tags			account	
//	@Accept			json
//	@Produce		json
//  @Param        	id   	path     int  true  "Account ID"
//	@Success		200	{object}	AccountResponse	
//	@Router			/api/account/{id} [put]
func (h accountHandler) EditAccountInfoHandler(c *gin.Context) {
	
	request := NewAccountRequest{}

	id := c.Param("id")

	accountId, err := strconv.Atoi(id)

	if err != nil || accountId < 0{
		response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err = c.ShouldBindJSON(&request)

	if err != nil{
		response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	accountResponse, err := h.accSrv.EditAccountInfo(request,accountId)

	if err != nil{
		response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}
	
	response.NewSuccessResponse(c, "edit account info", fiber.StatusOK, accountResponse)
}
 
func (h accountHandler) HealthCheck(c *gin.Context){
	response.NewSuccessResponse(c, "health check", fiber.StatusOK, nil)
}







