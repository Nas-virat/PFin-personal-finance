package service

import (
	"github.com/Nas-virat/PFin-personal-finance/errs"
	"github.com/Nas-virat/PFin-personal-finance/model"
	"github.com/Nas-virat/PFin-personal-finance/repository"
)

type balanceService struct{
	balanceRepo repository.BalanceRepository
}


func NewBalanceService(balanceRepo repository.BalanceRepository) BalanceService{
	return balanceService{balanceRepo:balanceRepo}
}

func (s balanceService) GetSummaryBalance()(*model.SummaryBalanceResponse, error){
	
	// Get all account
	accounts, err := s.balanceRepo.GetAllAccountBalances()
	if err != nil{
		return nil, errs.NewUnexpectedError()
	}

	// Get all debt
	debts, err := s.balanceRepo.GetAllDebtBalances()
	if err != nil{
		return nil, errs.NewUnexpectedError()
	}

	// calculate total asset and total debts
	totalAsset 	:= 0.0
	totalDebt	:= 0.0

	for _, account := range accounts{
		totalAsset += float64(account.Amount)
	}

	for _, debt := range debts{
		totalDebt += float64(debt.Amount)
	}

	summaryBalance := model.SummaryBalanceResponse{
		TotalAsset: totalAsset,
		TotalDebt: totalDebt,
		Accounts: accounts,
		Debts: debts,
	}

	return &summaryBalance,nil
}