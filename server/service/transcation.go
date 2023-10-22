package service

import "github.com/Nas-virat/PFin-personal-finance/model"



type TransactionService interface {
	CreateTransaction(transaction model.NewTransactionRequest) (*model.TransactionResponse, error)
	GetTransactionByID(id uint) (*model.TransactionResponse, error)
	GetTransactions() ([]model.TransactionResponse, error)
	UpdateTransaction(id uint, newInfo model.Transaction) (*model.TransactionResponse, error)
	DeleteTransaction(id uint) (*model.TransactionResponse, error)
}