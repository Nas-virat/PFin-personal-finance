package balance

import (
	"github.com/Nas-virat/PFin-personal-finance/account"
	"gorm.io/gorm"
)

type balanceRepositoryDB struct {
	db *gorm.DB
}


func NewBalanceRepositoryDB(db *gorm.DB) BalanceRepository{
	return balanceRepositoryDB{db:db}
}


func (repo balanceRepositoryDB) GetAllAccountBalances() ([]account.Account, error){
	accounts := []account.Account{}

	err := repo.db.Find(&accounts).Error
	if err != nil{
		return nil, err
	}

	return accounts, nil
}

func (repo balanceRepositoryDB) GetAllDebtBalances() ([]Debt, error){

	debts := []Debt{}

	err := repo.db.Find(&debts).Error
	if err != nil{
		return nil,err
	}


	return debts,nil
}

func (repo balanceRepositoryDB) CreateDebt(debtInfo Debt)(*Debt, error){

	result := repo.db.Create(&debtInfo)
	if result.Error != nil{
		return nil,result.Error
	}

	return &debtInfo, nil
}