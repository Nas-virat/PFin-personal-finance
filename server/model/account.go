package model

import (
	"time"

	"gorm.io/gorm"
)


type Account struct {
	gorm.Model
	AccountName string 	`gorm:"unique;not null"`
	Type 		string 	`gorm:"not null"`
	Amount 		float64 `gorm:"not null"`
	Description string
	Currency 	string 	`gorm:"not null"`
}

type NewAccountRequest struct{
	Name		string	`json:"account_name"`
	Type 		string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Currency 	string 	`json:"currency"`
}

type AccountResponse struct{
	AccountID 	int 		`json:"account_id"`
	Opendate 	time.Time 	`json:"opendate"`
	Type 		string  	`json:"account_type"`
	Amount      float64 	`json:"amount"`
	Status bool
}