package service

import "github.com/Nas-virat/PFin-personal-finance/model"

type TransactionService interface {
	CreateTransaction(transaction model.NewTransactionRequest) (*model.TransactionResponse, error)
	GetTransactionByID(id uint) (*model.TransactionResponse, error)
	GetTransactionInRangeMonthYear(month, year int) (*model.TransactionSummaryResponse, error)
	GetTransactionInRangeDayMonthYear(day, month, year int) (*model.TransactionSummaryResponse, error)
	GetSummaryRevenueExpenseYear() (*model.SummaryRevenueExpenseResponse, error)
	GetTransactions() ([]model.TransactionResponse, error)
	UpdateTransaction(id uint, newInfo model.NewTransactionRequest) (*model.TransactionResponse, error)
	DeleteTransaction(id uint) error
}
