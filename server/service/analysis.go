package service

import "github.com/Nas-virat/PFin-personal-finance/model"


type AnalysisService interface{
	GetAnalysisWealth() (*model.AnalysisWealth,error)
	GetNetWorthByMonth() (*model.NetWorthByMonth,error)
	GetFreeCashFlowByMonth() (*model.FreeCashFlowByMonth,error)
}