package repository

import (
	"github.com/Nas-virat/PFin-personal-finance/model"
	"gorm.io/gorm"
)

type balanceRepositoryDB struct {
	db *gorm.DB
}


func NewBalanceRepositoryDB(db *gorm.DB) BalanceRepository{
	return balanceRepositoryDB{db:db}
}


func (repo balanceRepositoryDB) GetAllAccountBalances() ([]model.Account, error){
	accounts := []model.Account{}

	err := repo.db.Find(&accounts).Error
	if err != nil{
		return nil, err
	}

	return accounts, nil
}

func (repo balanceRepositoryDB) GetAllDebtBalances() ([]model.Debt, error){

	debts := []model.Debt{}

	err := repo.db.Find(&debts).Error
	if err != nil{
		return nil,err
	}


	return debts,nil
}

func (repo balanceRepositoryDB) CreateDebt(debtInfo model.Debt)(*model.Debt, error){

	result := repo.db.Create(&debtInfo)
	if result.Error != nil{
		return nil,result.Error
	}

	return &debtInfo, nil
}