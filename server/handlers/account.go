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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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






