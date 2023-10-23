package service

import (
	"github.com/Nas-virat/PFin-personal-finance/errs"
	"github.com/Nas-virat/PFin-personal-finance/model"
	"github.com/Nas-virat/PFin-personal-finance/repository"
)

type transactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{transactionRepo: transactionRepo}
}

func (s transactionService) CreateTransaction(transactionRequest model.NewTransactionRequest) (*model.TransactionResponse, error) {

	// Validate Transaction
	if transactionRequest.Amount <= 0 {
		return nil, errs.NewVaildationError("Amount must be greater than 0")
	}

	transaction := model.Transaction{
		TransactionType: transactionRequest.TransactionType,
		Category:        transactionRequest.Category,
		Description:     transactionRequest.Description,
		Amount:          transactionRequest.Amount,
	}

	transactionResult, err := s.transactionRepo.CreateTransaction(transaction)

	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	transactionResponse := model.TransactionResponse{
		CreateAt:        transactionResult.CreatedAt,
		TransactionType: transactionResult.TransactionType,
		Category:       transactionResult.Category,
		Description:     transactionResult.Description,
		Amount:          transactionResult.Amount,
	}

	return &transactionResponse, nil
}

func (s transactionService) GetTransactionByID(id uint) (*model.TransactionResponse, error) {

	transaction, err := s.transactionRepo.GetTransactionByID(id)

	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	transactionResponse := model.TransactionResponse{
		CreateAt:        transaction.CreatedAt,
		TransactionType: transaction.TransactionType,
		Category:       transaction.Category,
		Description:     transaction.Description,
		Amount:          transaction.Amount,
	}

	return &transactionResponse, nil
}

func (s transactionService) GetTransactions() ([]model.TransactionResponse, error) {

	transactions, err := s.transactionRepo.GetTransactions()

	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	transactionResponses := []model.TransactionResponse{}

	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses,
			model.TransactionResponse{
				CreateAt:        transaction.CreatedAt,
				TransactionType: transaction.TransactionType,
				Category:       transaction.Category,
				Description:     transaction.Description,
				Amount:          transaction.Amount,
			},
		)
	}

	return transactionResponses, nil
}

func (s transactionService) GetTransactionInRangeMonthYear(month, year int) (*model.TransactionSummaryResponse, error) {

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

	transactionResponses := []model.TransactionResponse{}

	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses,
			model.TransactionResponse{
				CreateAt:        transaction.CreatedAt,
				TransactionType: transaction.TransactionType,
				Category:        transaction.Category,
				Description:     transaction.Description,
				Amount:          transaction.Amount,
			},
		)
	}

	//calculate Total Revenue, Total Expense, Total Credit, Total Remaining
	var totalRevenue, totalExpense, totalCredit float64

	for _, transaction := range transactionResponses {
		if transaction.TransactionType == "income" {
			totalRevenue += transaction.Amount
		} else if transaction.TransactionType == "expense" {
			totalExpense += transaction.Amount
		} else {
			totalCredit += transaction.Amount
		}
	}

	transactionSummaryResponses := model.TransactionSummaryResponse{
		TotalRevenue:   totalRevenue,
		TotalExpense:   totalExpense,
		TotalCredit:    totalCredit,
		TotalRemaining: totalRevenue - totalExpense,
		Transactions:   transactionResponses,
	}

	return &transactionSummaryResponses, nil
}

func (s transactionService) GetTransactionInRangeDayMonthYear(day, month, year int) (*model.TransactionSummaryResponse, error){
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

	transactionResponses := []model.TransactionResponse{}

	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses,
			model.TransactionResponse{
				CreateAt:        transaction.CreatedAt,
				TransactionType: transaction.TransactionType,
				Category:        transaction.Category,
				Description:     transaction.Description,
				Amount:          transaction.Amount,
			},
		)
	}

	//calculate Total Revenue, Total Expense, Total Credit, Total Remaining
	var totalRevenue, totalExpense, totalCredit float64

	for _, transaction := range transactionResponses {
		if transaction.TransactionType == "income" {
			totalRevenue += transaction.Amount
		} else if transaction.TransactionType == "expense" {
			totalExpense += transaction.Amount
		} else {
			totalCredit += transaction.Amount
		}
	}

	transactionSummaryResponses := model.TransactionSummaryResponse{
		TotalRevenue:   totalRevenue,
		TotalExpense:   totalExpense,
		TotalCredit:    totalCredit,
		TotalRemaining: totalRevenue - totalExpense,
		Transactions:   transactionResponses,
	}

	return &transactionSummaryResponses, nil
}

func (s transactionService) UpdateTransaction(id uint, newInfo model.NewTransactionRequest) (*model.TransactionResponse, error) {

	// Validate newInfo transaction
	if newInfo.Amount <= 0 {
		return nil, errs.NewVaildationError("Amount need to more than 0")
	}

	newTransaction := model.Transaction{
		TransactionType: newInfo.TransactionType,
		Category:        newInfo.Category,
		Description:     newInfo.Description,
		Amount:          newInfo.Amount,
	}

	transaction, err := s.transactionRepo.UpdateTransaction(id, newTransaction)

	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	transactionResponse := model.TransactionResponse{
		CreateAt:        transaction.CreatedAt,
		TransactionType: transaction.TransactionType,
		Category:        transaction.Category,
		Description:     transaction.Description,
		Amount:          transaction.Amount,
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
