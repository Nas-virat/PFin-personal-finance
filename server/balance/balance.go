package balance

import (
	"github.com/Nas-virat/PFin-personal-finance/account"
	"gorm.io/gorm"
)


type Debt struct{
	gorm.Model
	Name 			string 	`gorm:"not null"`
	Amount 			float64 `gorm:"not null"`
	InterestRate 	float64 `gorm:"not null"`
	MinimumPayment 	float64 
	AccountID 		int
	Account 		account.Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type NewDebtRequest struct {
	Name 			string 	`json:"name"`
	Amount 			float64 `json:"amount"`
	InterestRate 	float64 `json:"interest_rate"`
	MinimumPayment 	float64 `json:"minimum_payment"`
	AccountID 		int 	`json:"account_id"`
}

type DebtResponse struct {
	ID 				uint 	`json:"id"`
	Name 			string 	`json:"name"`
	Amount 			float64 `json:"amount"`
	InterestRate 	float64 `json:"interest_rate"`
	MinimumPayment 	float64 `json:"minimum_payment"`
	AccountID 		int 	`json:"account_id"`
}

type SummaryBalanceResponse struct {
	TotalAsset	float64		`json:"total_asset"`
	TotalDebt	float64		`json:"total_debt"`
	Accounts 	[]account.AccountListResponse	`json:"accounts"` // bank and investment account
	Debts    	[]Debt		`json:"debts"`	  	
}