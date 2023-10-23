package service

import(
	"github.com/Nas-virat/PFin-personal-finance/model"
)

type BalanceService interface{
	GetSummaryBalance()(*model.SummaryBalanceResponse, error)
	CreateDebt(debt model.NewDebtRequest)(*model.DebtResponse,error)
}