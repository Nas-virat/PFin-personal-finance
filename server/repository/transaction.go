package repository

import "github.com/Nas-virat/PFin-personal-finance/model"


type TransactionRepository interface{
	CreateTransaction(transaction model.Transaction)(*model.Transaction,error)
	GetTransactionByID(id uint) (*model.Transaction,error)
	GetTransactions() ([]model.Transaction,error)
	UpdateTransaction(id uint,newInfo model.Transaction) (*model.Transaction,error)
	DeleteTransaction(id uint) (*model.Transaction,error) 
}