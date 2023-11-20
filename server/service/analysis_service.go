package service

import (
	"github.com/Nas-virat/PFin-personal-finance/model"
	"github.com/Nas-virat/PFin-personal-finance/repository"
)

type analysisService struct {
	transactionRepo repository.TransactionRepository
	balanceRepo     repository.BalanceRepository
}

func NewAnalysisService(transactionRepo repository.TransactionRepository, balanceRepo repository.BalanceRepository) AnalysisService {
	return analysisService{transactionRepo: transactionRepo, balanceRepo: balanceRepo}
}

// GetAnalysisWealth implements AnalysisService.
func (analysisService) GetAnalysisWealth() (*model.AnalysisWealth, error) {
	panic("unimplemented")
}

// GetFreeCashFlowByMonth implements AnalysisService.
func (analysisService) GetFreeCashFlowByMonth() (*model.FreeCashFlowByMonth, error) {
	panic("unimplemented")
}

// GetNetWorthByMonth implements AnalysisService.
func (analysisService) GetNetWorthByMonth() (*model.NetWorthByMonth, error) {
	panic("unimplemented")
}


