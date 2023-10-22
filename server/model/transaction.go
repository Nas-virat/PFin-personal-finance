package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionType string  `gorm:"not null"` // Revenue, Expense, Other
	Category        string  `gorm:"not null"`
	Description     string  // Food, Travel, Utility
	Amount          float64 `gorm:"not null"` // Amount
}

type NewTransactionRequest struct {
	TransactionType string  `json:"transaction_type"`
	Category       string  `json:"category"`
	Description     string  `json:"description"`
	Amount          float64 `json:"amount"`
}

type TransactionResponse struct {
	CreateAt        time.Time `json:"create_at"`
	TransactionType string    `json:"transaction_type"`
	Category        string    `json:"category"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
}

type TransactionSummaryResponse struct {
	TotalRevenue   float64               `json:"total_revenue"`
	TotalExpense   float64               `json:"total_expense"`
	TotalCredit    float64               `json:"total_credit"`
	TotalRemaining float64               `json:"total_remaining"`
	Transactions   []TransactionResponse `json:"transactions"`
}
