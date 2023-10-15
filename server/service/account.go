package service

import "github.com/Nas-virat/PFin-personal-finance/model"


type AccountService interface {
 	CreateAccount(account model.NewAccountRequest) (*model.NewAccountResponse, error)
	GetAccountById(accountID int) (*model.AccountResponse, error)
	GetAccounts() ([]model.AccountResponse, error)
}