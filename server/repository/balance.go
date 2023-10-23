package repository

import "github.com/Nas-virat/PFin-personal-finance/model"


type BalanceRepository interface {
	GetAllAccountBalances() ([]model.Account, error)
	GetAllDebtBalances() ([]model.Debt, error)
	//GetAllInvestmentBalances() ([]model.Investment, error)
	CreateDebt(debtInfo model.Debt)(*model.Debt, error)
}