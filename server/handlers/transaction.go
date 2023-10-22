package handlers

import (
	"github.com/Nas-virat/PFin-personal-finance/model"
	"github.com/Nas-virat/PFin-personal-finance/service"
	"github.com/gofiber/fiber/v2"
)


type transactionHandler struct{
	transactionSrv service.TransactionService
}


func NewTransactionHandler(transactionSrv service.TransactionService) transactionHandler{
	return transactionHandler{transactionSrv:transactionSrv}
}

func (h transactionHandler) CreateTransactionHandler(c *fiber.Ctx) error {

	request := model.NewTransactionRequest{}

	err := c.BodyParser(&request)

	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	transactionResponse, err := h.transactionSrv.CreateTransaction(request)
	if err != nil{
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status": "fail", 
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status" : "success",
		"message": "insert successfully",
		"data"	 : transactionResponse,
	})
}

func (h transactionHandler) GetTransactionsHandler(c *fiber.Ctx) error {

	transactionResponses, err := h.transactionSrv.GetTransactions()

	if err != nil{
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status": "fail", 
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status" : "success",
		"message": "Get all transaction",
		"data"	 : transactionResponses,
	})
}

func (h transactionHandler) GetTransactionByIDHandler(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil || id < 0{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail", 
			"message": err.Error(),
		})
	}

	transactionResponse, err := h.transactionSrv.GetTransactionByID(uint(id))

	if err != nil{
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status": "fail", 
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "success", 
		"message": "Get Transaction with ID",
		"data": transactionResponse,
	})
}