package model

import (
	"time"

	"gorm.io/gorm"
)


type Transaction struct {
	gorm.Model
	TransactionType string `gorm:"not null"` // Revenue, Expense, Other 
	Catergory		string `gorm:"not null"`
	Description 	string					 // Food, Travel, Utility 
	Amount			float64 `gorm:"not null"`// Amount
}

type NewTransactionRequest struct{
	TransactionType string 		`json:"transaction_type"`
	Catergory		string 		`json:"catergory"`
	Description		string 		`json:"description"`
	Amount			float64		`json:"amount"`
}

type TransactionResponse struct {
	CreateAt		time.Time	`json:"create_at"`
	TransactionType string 		`json:"transaction_type"`
	Catergory		string 		`json:"catergory"`
	Description		string 		`json:"description"`
	Amount			float64		`json:"amount"`
}
