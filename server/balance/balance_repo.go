package balance

import (
	"github.com/Nas-virat/PFin-personal-finance/account"
	"gorm.io/gorm"
)


type BalanceRepository interface {
	GetAllAccountBalances() ([]account.Account, error)
	GetAllDebtBalances() ([]Debt, error)
	//GetAllInvestmentBalances() ([]model.Investment, error)
	CreateDebt(debtInfo Debt)(*Debt, error)
}

type balanceRepository struct {
	db *gorm.DB
}


func NewBalanceRepositoryDB(db *gorm.DB) balanceRepository{
	return balanceRepository{db:db}
}


func (repo balanceRepository) GetAllAccountBalances() ([]account.Account, error){
	accounts := []account.Account{}

	err := repo.db.Find(&accounts).Error
	if err != nil{
		return nil, err
	}

	return accounts, nil
}

func (repo balanceRepository) GetAllDebtBalances() ([]Debt, error){

	debts := []Debt{}

	err := repo.db.Find(&debts).Error
	if err != nil{
		return nil,err
	}


	return debts,nil
}

func (repo balanceRepository) CreateDebt(debtInfo Debt)(*Debt, error){

	result := repo.db.Create(&debtInfo)
	if result.Error != nil{
		return nil,result.Error
	}

	return &debtInfo, nil
}