package transaction

import (
	"github.com/Nas-virat/PFin-personal-finance/response"
	"github.com/gofiber/fiber/v2"
)

type transactionHandler struct {
	transactionSrv TransactionService
}

func NewTransactionHandler(transactionSrv TransactionService) transactionHandler {
	return transactionHandler{transactionSrv: transactionSrv}
}

//  CreateTransaction	godoc
//	@Summary			Create Transaction
//	@Description		Create Transaction
//	@Tags				Transcation
//	@Accept				json
//	@Produce			json
//	@Success			201	
//	@Router				/api/transaction [post]
func (h transactionHandler) CreateTransactionHandler(c *fiber.Ctx) error {

	request := NewTransactionRequest{}

	err := c.BodyParser(&request)

	if err != nil {
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	transactionResponse, err := h.transactionSrv.CreateTransaction(request)
	if err != nil {
		return response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}
	return response.NewSuccessResponse(c, "insert successfully", fiber.StatusCreated, transactionResponse)
}

//  GetTransaction		godoc
//	@Summary			Get All Transaction
//	@Description		Get All Transaction
//	@Tags				Transcation
//	@Accept				json
//	@Produce			json
//	@Success			200
//	@Router				/api/transaction [get]
func (h transactionHandler) GetTransactionsHandler(c *fiber.Ctx) error {

	transactionResponses, err := h.transactionSrv.GetTransactions()

	if err != nil {
		return response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	return response.NewSuccessResponse(c, "Get all transaction", fiber.StatusOK, transactionResponses)
}

//  GetTransactionInRangeMonthYear	godoc
//	@Summary						Get All Transaction in Range Month and Year
//	@Description					Get All Transaction in Range Month and Year
//	@Tags							Transcation
//	@Accept							json
//	@Produce						json
//  @Param        					month   path     int  true  "month"
//  @Param        					year   	path     int  true  "year"
//	@Success						200
//	@Router							/api/transaction/month/{month}/year/{year} [get]
func (h transactionHandler) GetTransactionInRangeMonthYearHandler(c *fiber.Ctx) error {

	year, err := c.ParamsInt("year")
	if err != nil {
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	month, err := c.ParamsInt("month")
	if err != nil {
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	transactionSummaryResponses, err := h.transactionSrv.GetTransactionInRangeMonthYear(
		month,
		year,
	)

	if err != nil {
		return response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	return response.NewSuccessResponse(c, "Get all transaction", fiber.StatusOK, transactionSummaryResponses)
}

//  GGetTransactionInRangeDayMonthYear	godoc
//	@Summary							Get All Transaction in Range Day Month and Year
//	@Description						Get All Transaction in Range Day Month and Year
//	@Tags								Transcation
//	@Accept								json
//	@Produce							json
//  @Param        						day   	path     int  true  "day"
//  @Param        						month   path     int  true  "month"
//  @Param        						year   	path     int  true  "year"
//	@Success							200
//	@Router								/api/transaction/day/{day}/month/{month}/year/{year} [get]
func (h transactionHandler) GetTransactionInRangeDayMonthYearHandler(c *fiber.Ctx) error {

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

	if err != nil {
		return response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	return response.NewSuccessResponse(c, "Get all transaction", fiber.StatusOK, transactionSummaryResponses)
}

//  GetTransactionByID	godoc
//	@Summary			Get Transaction By Id	
//	@Description		Get Transaction By Id	
//	@Tags				Transcation
//	@Accept				json
//	@Produce			json
//  @Param        		id   	path     int  true  "Transaction ID"
//	@Success			200
//	@Router				/api/transaction/{id}	[get]
func (h transactionHandler) GetTransactionByIDHandler(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil || id < 0 {
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	transactionResponse, err := h.transactionSrv.GetTransactionByID(uint(id))

	if err != nil {
		return response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}
	return response.NewSuccessResponse(c, "Get Transaction with ID", fiber.StatusOK, transactionResponse)
}


//  UpdateTransaction	godoc
//	@Summary			Update Transaction By Id	
//	@Description		Update Transaction By Id	
//	@Tags				Transcation
//	@Accept				json
//	@Produce			json
//  @Param        		id   	path     int  true  "Transaction ID"
//	@Success			200
//	@Router				/api/transaction/{id}	[put]
func (h transactionHandler) UpdateTransactionByIdHandler(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil || id < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	request := NewTransactionRequest{}

	err = c.BodyParser(&request)

	if err != nil {
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	transactionResponse, err := h.transactionSrv.UpdateTransaction(uint(id), request)

	if err != nil {
		return response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	return response.NewSuccessResponse(c, "Update Transaction with ID", fiber.StatusOK, transactionResponse)
}

//  UpdateTransaction	godoc
//	@Summary			Get Summary Revenue Expense 	
//	@Description		Get Summary Revenue Expense	
//	@Tags				Transcation
//	@Accept				json
//	@Produce			json
//	@Success			200
//	@Router				/api/transaction	[put]
func (h transactionHandler) GetSummaryRevenueExpenseHandler(c *fiber.Ctx) error {

	summaryRevenueExpenseResponse, err := h.transactionSrv.GetSummaryRevenueExpenseYear()

	if err != nil {
		return response.NewErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return response.NewSuccessResponse(c, "Get summary revenue expense", fiber.StatusOK, summaryRevenueExpenseResponse)
}


//  DeleteTransaction	godoc
//	@Summary			Delete Transaction 	
//	@Description		Delete Transaction	
//	@Tags				Transcation
//	@Accept				json
//	@Produce			json
//  @Param        		id   	path     int  true  "Transaction ID"
//	@Success			200
//	@Router				/api/transaction/summary-year	[get]
func (h transactionHandler) DeleteTransactionHandler(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil || id < 0 {
		return response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}

	err = h.transactionSrv.DeleteTransaction(uint(id))

	if err != nil {
		return response.NewErrorResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	return response.NewSuccessResponse(c, "Delete Transaction with ID", fiber.StatusOK, nil)
}

func (h transactionHandler) HealthCheck(c *fiber.Ctx) error {
	return response.NewSuccessResponse(c, "Health check", fiber.StatusOK, nil)
}
