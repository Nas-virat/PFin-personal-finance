package service

import "github.com/Nas-virat/PFin-personal-finance/model"



type TransactionService interface {
	CreateTransaction(transaction model.NewTransactionRequest) (*model.TransactionResponse, error)
	GetTransactionByID(id uint) (*model.TransactionResponse, error)
	GetTransactionInRanageMonthYear(month, year int) (*model.TransactionSummaryResponse, error)
	GetTransactions() ([]model.TransactionResponse, error)
	UpdateTransaction(id uint, newInfo model.NewTransactionRequest) (*model.TransactionResponse, error)
	DeleteTransaction(id uint) error
}