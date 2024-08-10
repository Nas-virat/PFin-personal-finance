package transaction

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionType  string  `gorm:"not null"` // Revenue, Expense, Other
	Category         string  `gorm:"not null"`
	Description      string  // Food, Travel, Utility
	Amount           float64 `gorm:"not null"` // Amount
	TransactionDate  int     `gorm:"not null"` // Date
	TransactionMonth int     `gorm:"not null"` // Month
	TransactionYear  int     `gorm:"not null"` // Year
}

type NewTransactionRequest struct {
	TransactionType  string  `json:"transaction_type"`
	Category         string  `json:"category"`
	Description      string  `json:"description"`
	Amount           float64 `json:"amount"`
	TransactionDate  int     `json:"transaction_date"`  // Date
	TransactionMonth int     `json:"transaction_month"` // Month
	TransactionYear  int     `json:"transaction_year"`  // Year
}

type TransactionResponse struct {
	CreateAt         time.Time `json:"create_at"`
	TransactionType  string    `json:"transaction_type"`
	Category         string    `json:"category"`
	Description      string    `json:"description"`
	Amount           float64   `json:"amount"`
	TransactionDate  int       `json:"transaction_date"`  // Date
	TransactionMonth int       `json:"transaction_month"` // Month
	TransactionYear  int       `json:"transaction_year"`  // Year
}

type TransactionSummaryResponse struct {
	TotalRevenue   float64                              `json:"total_revenue"`
	TotalExpense   float64                              `json:"total_expense"`
	TotalCredit    float64                              `json:"total_credit"`
	TotalRemaining float64                              `json:"total_remaining"`
	Transactions   []TransactionSummaryCategoryResponse `json:"transactions"`
}

type TransactionSummaryCategoryResponse struct {
	TransactionType string  `json:"transaction_type"`
	Category        string  `json:"category"`
	Amount          float64 `json:"amount"`
}

type SummaryRevenueExpenseResponse struct {
	TotalRevenue []float64 `json:"total_revenue"`
	TotalExpense []float64 `json:"total_expense"`
}
