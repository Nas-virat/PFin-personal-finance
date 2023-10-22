package model

import "gorm.io/gorm"

type Catergory struct {
	gorm.Model
	TransactionType string `gorm:"not null"` // Revenue, Expense, Other 
	CatergoryName string `gorm:"not null"`
}