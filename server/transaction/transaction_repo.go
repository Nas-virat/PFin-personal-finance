package transaction



type TransactionRepository interface{
	CreateTransaction(transaction Transaction)(*Transaction,error)
	GetTransactionByID(id uint) (*Transaction,error)
	GetTransactions() ([]Transaction,error)
	GetTransactionInYear(year int) ([] Transaction,error)
	GetTransactionInRangeMonthYear(month,year int) ([]Transaction,error)
	GetTransactionInRangeDayMonthYear(day,month,year int) ([]Transaction,error)
	UpdateTransaction(id uint,newInfo  Transaction) (*Transaction,error)
	DeleteTransaction(id uint) error 
}