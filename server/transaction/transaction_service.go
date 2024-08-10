package transaction

import (
	"github.com/Nas-virat/PFin-personal-finance/errs"
	"time"
)

type TransactionService interface {
	CreateTransaction(transaction NewTransactionRequest) (*TransactionResponse, error)
	GetTransactionByID(id uint) (*TransactionResponse, error)
	GetTransactionInRangeMonthYear(month, year int) (*TransactionSummaryResponse, error)
	GetTransactionInRangeDayMonthYear(day, month, year int) (*TransactionSummaryResponse, error)
	GetSummaryRevenueExpenseYear() (*SummaryRevenueExpenseResponse, error)
	GetTransactions() ([]TransactionResponse, error)
	UpdateTransaction(id uint, newInfo NewTransactionRequest) (*TransactionResponse, error)
	DeleteTransaction(id uint) error
}

type transactionService struct {
	transactionRepo TransactionRepository
}

func NewTransactionService(transactionRepo TransactionRepository) TransactionService {
	return &transactionService{transactionRepo: transactionRepo}
}

func (s transactionService) CreateTransaction(transactionRequest NewTransactionRequest) (*TransactionResponse, error) {

	// Validate Transaction
	if transactionRequest.Amount <= 0 {
		return nil, errs.NewVaildationError("Amount must be greater than 0")
	}

	transaction := Transaction{
		TransactionType:  transactionRequest.TransactionType,
		Category:         transactionRequest.Category,
		Description:      transactionRequest.Description,
		Amount:           transactionRequest.Amount,
		TransactionDate:  transactionRequest.TransactionDate,
		TransactionMonth: transactionRequest.TransactionMonth,
		TransactionYear:  transactionRequest.TransactionYear,
	}

	transactionResult, err := s.transactionRepo.CreateTransaction(transaction)

	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	transactionResponse := TransactionResponse{
		CreateAt:         transactionResult.CreatedAt,
		TransactionType:  transactionResult.TransactionType,
		Category:         transactionResult.Category,
		Description:      transactionResult.Description,
		Amount:           transactionResult.Amount,
		TransactionDate:  transactionResult.TransactionDate,
		TransactionMonth: transactionResult.TransactionMonth,
		TransactionYear:  transactionResult.TransactionYear,
	}

	return &transactionResponse, nil
}

func (s transactionService) GetTransactionByID(id uint) (*TransactionResponse, error) {

	transaction, err := s.transactionRepo.GetTransactionByID(id)

	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	transactionResponse := TransactionResponse{
		CreateAt:         transaction.CreatedAt,
		TransactionType:  transaction.TransactionType,
		Category:         transaction.Category,
		Description:      transaction.Description,
		Amount:           transaction.Amount,
		TransactionDate:  transaction.TransactionDate,
		TransactionMonth: transaction.TransactionMonth,
		TransactionYear:  transaction.TransactionYear,
	}

	return &transactionResponse, nil
}

func (s transactionService) GetSummaryRevenueExpenseYear() (*SummaryRevenueExpenseResponse, error) {
	// Get current year
	year := time.Now().Year()

	TotalRevenue := [12]float64{}
	TotalExpense := [12]float64{}

	for month := 1; month <= 12; month++ {
		// Get all transaction in current month
		transactions, err := s.transactionRepo.GetTransactionInRangeMonthYear(month, year)
		if err != nil {
			return nil, errs.NewUnexpectedError()
		}
		for _, transaction := range transactions {
			if transaction.TransactionType == "income" {
				TotalRevenue[month-1] += transaction.Amount
			} else if transaction.TransactionType == "expense" {
				TotalExpense[month-1] += transaction.Amount
			}
		}
	}
	sumaryRevenueExpense := SummaryRevenueExpenseResponse{
		TotalRevenue: TotalRevenue[:],
		TotalExpense: TotalExpense[:],
	}

	return &sumaryRevenueExpense, nil
}

func (s transactionService) GetTransactions() ([]TransactionResponse, error) {

	transactions, err := s.transactionRepo.GetTransactions()

	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	transactionResponses := []TransactionResponse{}

	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses,
			TransactionResponse{
				CreateAt:         transaction.CreatedAt,
				TransactionType:  transaction.TransactionType,
				Category:         transaction.Category,
				Description:      transaction.Description,
				Amount:           transaction.Amount,
				TransactionDate:  transaction.TransactionDate,
				TransactionMonth: transaction.TransactionMonth,
				TransactionYear:  transaction.TransactionYear,
			},
		)
	}

	return transactionResponses, nil
}

func (s transactionService) GetTransactionInRangeMonthYear(month, year int) (*TransactionSummaryResponse, error) {

	// Validate month and year
	if month < 1 || month > 12 {
		return nil, errs.NewVaildationError("Month must be in range 1-12")
	}
	if year < 2000 {
		return nil, errs.NewVaildationError("Year must be greater than 2000")
	}

	transactions, err := s.transactionRepo.GetTransactionInRangeMonthYear(month, year)

	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	transactionCategoryResponse := []TransactionSummaryCategoryResponse{}

	//calculate Total Revenue, Total Expense, Total Credit, Total Remaining
	var totalRevenue, totalExpense, totalCredit float64

	categoryAmountMap := make(map[string]float64)
	categoryTypeMap := make(map[string]string)

	for _, transaction := range transactions {

		//check category is exist in map
		if _, ok := categoryAmountMap[transaction.Category]; !ok {
			categoryAmountMap[transaction.Category] = 0
			categoryTypeMap[transaction.Category] = transaction.TransactionType
		}

		//add amount to category
		categoryAmountMap[transaction.Category] += transaction.Amount

		if transaction.TransactionType == "income" {
			totalRevenue += transaction.Amount
		} else if transaction.TransactionType == "expense" {
			totalExpense += transaction.Amount
		} else {
			totalCredit += transaction.Amount
		}
	}

	for category, amount := range categoryAmountMap {
		transactionCategoryResponse = append(transactionCategoryResponse,
			TransactionSummaryCategoryResponse{
				Category:        category,
				TransactionType: categoryTypeMap[category],
				Amount:          amount,
			},
		)
	}

	transactionSummaryResponses := TransactionSummaryResponse{
		TotalRevenue:   totalRevenue,
		TotalExpense:   totalExpense,
		TotalCredit:    totalCredit,
		TotalRemaining: totalRevenue - totalExpense,
		Transactions:   transactionCategoryResponse,
	}

	return &transactionSummaryResponses, nil
}

func (s transactionService) GetTransactionInRangeDayMonthYear(day, month, year int) (*TransactionSummaryResponse, error) {
	// Validate month and year
	if month < 1 || month > 12 {
		return nil, errs.NewVaildationError("Month must be in range 1-12")
	}
	if year < 2000 {
		return nil, errs.NewVaildationError("Year must be greater than 2000")
	}

	transactions, err := s.transactionRepo.GetTransactionInRangeDayMonthYear(day, month, year)

	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	transactionCategoryResponse := []TransactionSummaryCategoryResponse{}

	//calculate Total Revenue, Total Expense, Total Credit, Total Remaining
	var totalRevenue, totalExpense, totalCredit float64

	categoryAmountMap := make(map[string]float64)
	categoryTypeMap := make(map[string]string)

	for _, transaction := range transactions {

		//check category is exist in map
		if _, ok := categoryAmountMap[transaction.Category]; !ok {
			categoryAmountMap[transaction.Category] = 0
			categoryTypeMap[transaction.Category] = transaction.TransactionType
		}

		//add amount to category
		categoryAmountMap[transaction.Category] += transaction.Amount

		if transaction.TransactionType == "income" {
			totalRevenue += transaction.Amount
		} else if transaction.TransactionType == "expense" {
			totalExpense += transaction.Amount
		} else {
			totalCredit += transaction.Amount
		}
	}

	for category, amount := range categoryAmountMap {
		transactionCategoryResponse = append(transactionCategoryResponse,
			TransactionSummaryCategoryResponse{
				Category:        category,
				TransactionType: categoryTypeMap[category],
				Amount:          amount,
			},
		)
	}

	transactionSummaryResponses := TransactionSummaryResponse{
		TotalRevenue:   totalRevenue,
		TotalExpense:   totalExpense,
		TotalCredit:    totalCredit,
		TotalRemaining: totalRevenue - totalExpense,
		Transactions:   transactionCategoryResponse,
	}

	return &transactionSummaryResponses, nil
}

func (s transactionService) UpdateTransaction(id uint, newInfo NewTransactionRequest) (*TransactionResponse, error) {

	// Validate newInfo transaction
	if newInfo.Amount <= 0 {
		return nil, errs.NewVaildationError("Amount need to more than 0")
	}

	newTransaction := Transaction{
		TransactionType:  newInfo.TransactionType,
		Category:         newInfo.Category,
		Description:      newInfo.Description,
		Amount:           newInfo.Amount,
		TransactionDate:  newInfo.TransactionDate,
		TransactionMonth: newInfo.TransactionMonth,
		TransactionYear:  newInfo.TransactionYear,
	}

	transaction, err := s.transactionRepo.UpdateTransaction(id, newTransaction)

	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	transactionResponse := TransactionResponse{
		CreateAt:         transaction.CreatedAt,
		TransactionType:  transaction.TransactionType,
		Category:         transaction.Category,
		Description:      transaction.Description,
		Amount:           transaction.Amount,
		TransactionDate:  transaction.TransactionDate,
		TransactionMonth: transaction.TransactionMonth,
		TransactionYear:  transaction.TransactionYear,
	}

	return &transactionResponse, nil
}

func (s transactionService) DeleteTransaction(id uint) error {
	err := s.transactionRepo.DeleteTransaction(id)

	if err != nil {
		return errs.NewUnexpectedError()
	}

	return nil
}
