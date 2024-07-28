package balance

import (
	"github.com/Nas-virat/PFin-personal-finance/account"
	"github.com/Nas-virat/PFin-personal-finance/errs"
)

type BalanceService interface{
	GetSummaryBalance()(*SummaryBalanceResponse, error)
	CreateDebt(debt NewDebtRequest)(*DebtResponse,error)
}

type balanceService struct{
	balanceRepo BalanceRepository
}


func NewBalanceService(balanceRepo BalanceRepository) BalanceService{
	return balanceService{balanceRepo:balanceRepo}
}

func (s balanceService) GetSummaryBalance()(*SummaryBalanceResponse, error){
	
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

	AccountListResponses := []account.AccountListResponse{}

	for _, accountItem := range accounts{
		totalAsset += float64(accountItem.Amount)

		AccountListResponses = append(AccountListResponses, account.AccountListResponse{
			AccountName: accountItem.AccountName,
			Type: accountItem.Type,
			Amount: accountItem.Amount,
		})
	}

	for _, debt := range debts{
		totalDebt += float64(debt.Amount)
	}


	summaryBalance := SummaryBalanceResponse{
		TotalAsset: totalAsset,
		TotalDebt: totalDebt,
		Accounts: AccountListResponses,
		Debts: debts,
	}

	return &summaryBalance,nil
}

func (s balanceService) CreateDebt(debt NewDebtRequest)(*DebtResponse,error){

	debtModel := Debt{
		Name: debt.Name,
		Amount: debt.Amount,
		InterestRate: debt.InterestRate,
		MinimumPayment: debt.MinimumPayment,
		AccountID: debt.AccountID,
	}
	
	result, err := s.balanceRepo.CreateDebt(debtModel)

	if err != nil{
		return nil, errs.NewUnexpectedError()
	}

	resultResponse := DebtResponse{
		ID: result.ID,
		Name: result.Name,
		Amount: result.Amount,
		InterestRate: result.InterestRate,
		MinimumPayment: result.MinimumPayment,
		AccountID: result.AccountID,
	}

	return &resultResponse, nil

}