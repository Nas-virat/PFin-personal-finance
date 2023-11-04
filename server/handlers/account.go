package handlers

import (
	"github.com/Nas-virat/PFin-personal-finance/model"
	"github.com/Nas-virat/PFin-personal-finance/service"
	"github.com/gofiber/fiber/v2"
)


type accountHandler struct {
	accSrv service.AccountService
}

func NewAccountHandler(accSrv service.AccountService) accountHandler{
	return accountHandler{accSrv:accSrv}
}

func (h accountHandler) CreateAccountHandler(c *fiber.Ctx) error {

	request := model.NewAccountRequest{}

	err := c.BodyParser(&request)

	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	accountResponse ,err := h.accSrv.CreateAccount(request)

	if err != nil{
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status": "fail", 
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
			"status":"success",
			"message": "insert successfully",
			"data": accountResponse,
	})
}

func (h accountHandler) GetAccountsHandler(c *fiber.Ctx) error {

	accounts, err := h.accSrv.GetAccounts()

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "fail", 
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "success",
		"message": "get all accounts",
		"data": accounts,
	})
}

func (h accountHandler) EditAccountInfoHandler(c *fiber.Ctx) error {
	
	request := model.NewAccountRequest{}

	id, err := c.ParamsInt("id")

	if err != nil || id < 0{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail", 
			"message": err.Error(),
		})
	}

	err = c.BodyParser(&request)

	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	accountResponse, err := h.accSrv.EditAccountInfo(request,id)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "succuess",
		"message":"update sucuessfully",
		"data": accountResponse,
	})
}
 
func (h accountHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "success",
		"message": "health check",
	})
}







