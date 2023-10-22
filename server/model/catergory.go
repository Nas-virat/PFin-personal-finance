package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	TransactionType string `gorm:"not null"` // Revenue, Expense, Other 
	CategoryName string `gorm:"not null"`
}