package account


type AccountRepository interface {
	CreateAccount(account Account) (*Account, error)
	GetAccountById(id int) (*Account, error)
	GetAccounts() ([]Account, error)
	EditAccountInfo(account Account,id int) (*Account, error)
}