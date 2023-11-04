package repository

import "github.com/Nas-virat/PFin-personal-finance/model"



type AccountRepository interface {
	CreateAccount(account model.Account) (*model.Account, error)
	GetAccountById(id int) (*model.Account, error)
	GetAccounts() ([]model.Account, error)
	EditAccountInfo(account model.Account,id int) (*model.Account, error)
}