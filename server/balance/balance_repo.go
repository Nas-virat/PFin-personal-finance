package balance

import (
	"github.com/Nas-virat/PFin-personal-finance/account"
)


type BalanceRepository interface {
	GetAllAccountBalances() ([]account.Account, error)
	GetAllDebtBalances() ([]Debt, error)
	//GetAllInvestmentBalances() ([]model.Investment, error)
	CreateDebt(debtInfo Debt)(*Debt, error)
}