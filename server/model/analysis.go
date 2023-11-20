package model

type AnalysisWealth struct {
	NetWorth                 float64 `json:"net_worth"`
	RevenuePerExpense        float64 `json:"revenue_per_expense"`
	FreeCashFlow             float64 `json:"free_cash_flow"`
	RevenuePassivePerExpense float64 `json:"revenue_passive_per_expense"`
}

type FreeCashFlowByMonth struct {
	FreeCashFlow []float64 `json:"free_cash_flow"`
}

type NetWorthByMonth struct {
	NetWorth []float64 `json:"net_worth"`
}
