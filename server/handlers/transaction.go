package handlers

import (
	"github.com/Nas-virat/PFin-personal-finance/model"
	"github.com/Nas-virat/PFin-personal-finance/response"
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
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	transactionResponse, err := h.transactionSrv.CreateTransaction(request)
	if err != nil{
		return response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}
	return response.NewSuccessResponse(c, "insert successfully", fiber.StatusCreated, transactionResponse)
}

func (h transactionHandler) GetTransactionsHandler(c *fiber.Ctx) error {

	transactionResponses, err := h.transactionSrv.GetTransactions()

	if err != nil{
		return response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	return response.NewSuccessResponse(c, "Get all transaction", fiber.StatusAccepted, transactionResponses)
}

func (h transactionHandler) GetTransactionInRangeMonthYearHandler(c *fiber.Ctx) error{

	year, err := c.ParamsInt("year")
	if err != nil {
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	month, err := c.ParamsInt("month")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	transactionSummaryResponses, err := h.transactionSrv.GetTransactionInRangeMonthYear(
		month,
		year,
	)

	if err != nil{
		return response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	return response.NewSuccessResponse(c, "Get all transaction", fiber.StatusAccepted, transactionSummaryResponses)
}

func (h transactionHandler) GetTransactionInRangeDayMonthYearHandler(c *fiber.Ctx) error{

	year, err := c.ParamsInt("year")
	if err != nil {
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	month, err := c.ParamsInt("month")
	if err != nil {
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	day, err := c.ParamsInt("day")
	if err != nil {
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	transactionSummaryResponses, err := h.transactionSrv.GetTransactionInRangeDayMonthYear(
		day,
		month,
		year,
	)

	if err != nil{
		return response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	return response.NewSuccessResponse(c, "Get all transaction", fiber.StatusAccepted, transactionSummaryResponses)
}


func (h transactionHandler) GetTransactionByIDHandler(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil || id < 0{
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	transactionResponse, err := h.transactionSrv.GetTransactionByID(uint(id))

	if err != nil{
		return response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}
	return response.NewSuccessResponse(c, "Get Transaction with ID", fiber.StatusAccepted, transactionResponse)
}

func (h transactionHandler) UpdateTransactionHandler(c *fiber.Ctx) error {
	
	id, err := c.ParamsInt("id")

	if err != nil || id < 0{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail", 
			"message": err.Error(),
		})
	}

	request := model.NewTransactionRequest{}

	err = c.BodyParser(&request)

	if err != nil{
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	transactionResponse, err := h.transactionSrv.UpdateTransaction(uint(id), request)

	if err != nil{
		return response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	return response.NewSuccessResponse(c, "Update Transaction with ID", fiber.StatusAccepted, transactionResponse)
}
// TODO:Barchart home and summary by month and day
func (h transactionHandler) GetSummaryRevenueExpenseHandler(c *fiber.Ctx) error {

	//summaryRevenueExpenseResponse, err := h.transactionSrv.GetSummaryRevenueExpense()

	// if err != nil{
	// 	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
	// 		"status": "fail", 
	// 		"message": err.Error(),
	// 	})
	// }

	// return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
	// 	"status": "success", 
	// 	"message": "Get Summary Revenue Expense",
	// 	"data": summaryRevenueExpenseResponse,
	// })
	return nil
}

func (h transactionHandler) DeleteTransactionHandler(c *fiber.Ctx) error {
	
	id, err := c.ParamsInt("id")

	if err != nil || id < 0{
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	err = h.transactionSrv.DeleteTransaction(uint(id))

	if err != nil{
		return response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	return response.NewSuccessResponse(c, "Delete Transaction with ID", fiber.StatusOK, nil)
}

func (h transactionHandler) HealthCheck(c *fiber.Ctx) error {
	return response.NewSuccessResponse(c, "Health check", fiber.StatusOK, nil)
}