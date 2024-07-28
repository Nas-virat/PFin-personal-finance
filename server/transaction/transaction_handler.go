package transaction

import (
	"net/http"
	"strconv"

	"github.com/Nas-virat/PFin-personal-finance/response"
	"github.com/gin-gonic/gin"
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
func (h transactionHandler) CreateTransactionHandler(c *gin.Context) {

	request := NewTransactionRequest{}

	err := c.ShouldBindJSON(&request)

	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	transactionResponse, err := h.transactionSrv.CreateTransaction(request)
	if err != nil {
		response.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}
	response.NewSuccessResponse(c, "insert successfully", http.StatusCreated, transactionResponse)
}

//  GetTransaction		godoc
//	@Summary			Get All Transaction
//	@Description		Get All Transaction
//	@Tags				Transcation
//	@Accept				json
//	@Produce			json
//	@Success			200
//	@Router				/api/transaction [get]
func (h transactionHandler) GetTransactionsHandler(c *gin.Context) {

	transactionResponses, err := h.transactionSrv.GetTransactions()

	if err != nil {
		response.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	response.NewSuccessResponse(c, "Get all transaction", http.StatusOK, transactionResponses)
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
func (h transactionHandler) GetTransactionInRangeMonthYearHandler(c *gin.Context) {

	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	transactionSummaryResponses, err := h.transactionSrv.GetTransactionInRangeMonthYear(
		month,
		year,
	)

	if err != nil {
		response.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	response.NewSuccessResponse(c, "Get all transaction", http.StatusOK, transactionSummaryResponses)
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
func (h transactionHandler) GetTransactionInRangeDayMonthYearHandler(c *gin.Context){

	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	day, err := strconv.Atoi(c.Param("day"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	transactionSummaryResponses, err := h.transactionSrv.GetTransactionInRangeDayMonthYear(
		day,
		month,
		year,
	)

	if err != nil {
		response.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	response.NewSuccessResponse(c, "Get all transaction", http.StatusOK, transactionSummaryResponses)
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
func (h transactionHandler) GetTransactionByIDHandler(c *gin.Context) {

	id := c.Param("id")

	transactionId, err := strconv.Atoi(id)

	if err != nil || transactionId < 0 {
		response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	transactionResponse, err := h.transactionSrv.GetTransactionByID(uint(transactionId))

	if err != nil {
		response.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}
	response.NewSuccessResponse(c, "Get Transaction with ID", http.StatusOK, transactionResponse)
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
func (h transactionHandler) UpdateTransactionByIdHandler(c *gin.Context) {

	id := c.Param("id")

	transactionId, err := strconv.Atoi(id)
	if err != nil || transactionId < 0 {
		response.NewErrorResponse(c,http.StatusBadRequest,err)
	}

	request := NewTransactionRequest{}

	err = c.ShouldBindJSON(&request)

	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	transactionResponse, err := h.transactionSrv.UpdateTransaction(uint(transactionId), request)

	if err != nil {
		response.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	response.NewSuccessResponse(c, "Update Transaction with ID", http.StatusOK, transactionResponse)
}

//  UpdateTransaction	godoc
//	@Summary			Get Summary Revenue Expense 	
//	@Description		Get Summary Revenue Expense	
//	@Tags				Transcation
//	@Accept				json
//	@Produce			json
//	@Success			200
//	@Router				/api/transaction	[put]
func (h transactionHandler) GetSummaryRevenueExpenseHandler(c *gin.Context) {

	summaryRevenueExpenseResponse, err := h.transactionSrv.GetSummaryRevenueExpenseYear()

	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	response.NewSuccessResponse(c, "Get summary revenue expense", http.StatusOK, summaryRevenueExpenseResponse)
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
func (h transactionHandler) DeleteTransactionHandler(c *gin.Context){

	id := c.Param("id")

	transactionId, err := strconv.Atoi(id)


	if err != nil || transactionId < 0 {
		response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err = h.transactionSrv.DeleteTransaction(uint(transactionId))

	if err != nil {
		response.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	response.NewSuccessResponse(c, "Delete Transaction with ID", http.StatusOK, nil)
}

func (h transactionHandler) HealthCheck(c *gin.Context) {
	response.NewSuccessResponse(c, "Health check", http.StatusOK, nil)
}
