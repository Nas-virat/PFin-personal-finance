package analysis

import (
	"github.com/Nas-virat/PFin-personal-finance/balance"
	"github.com/Nas-virat/PFin-personal-finance/transaction"
)


type AnalysisService interface{
	GetAnalysisWealth() (*AnalysisWealth,error)
	GetNetWorthByMonth() (*NetWorthByMonth,error)
	GetFreeCashFlowByMonth() (*FreeCashFlowByMonth,error)
}

type analysisService struct {
	transactionRepo transaction.TransactionRepository
	balanceRepo     balance.BalanceRepository
}

func NewAnalysisService(transactionRepo transaction.TransactionRepository, balanceRepo balance.BalanceRepository) AnalysisService {
	return analysisService{transactionRepo: transactionRepo, balanceRepo: balanceRepo}
}

// GetAnalysisWealth implements AnalysisService.
func (analysisService) GetAnalysisWealth() (*AnalysisWealth, error) {
	panic("unimplemented")
}

// GetFreeCashFlowByMonth implements AnalysisService.
func (analysisService) GetFreeCashFlowByMonth() (*FreeCashFlowByMonth, error) {
	panic("unimplemented")
}

// GetNetWorthByMonth implements AnalysisService.
func (analysisService) GetNetWorthByMonth() (*NetWorthByMonth, error) {
	panic("unimplemented")
}


