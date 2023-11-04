package repository

import (
	"github.com/Nas-virat/PFin-personal-finance/model"
	"gorm.io/gorm"
)

type accountRepositoryDB struct {
	db *gorm.DB
}


// Constructor
func NewAccountRepositoryDB(db *gorm.DB) AccountRepository {
	return &accountRepositoryDB{db}
}

func (repo *accountRepositoryDB) CreateAccount(account model.Account) (*model.Account, error) {
	
	result := repo.db.Create(&account)

	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

func (repo *accountRepositoryDB) GetAccountById(id int) (*model.Account, error) {
	
	account := model.Account{}
	err := repo.db.Where("id = ?", id).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (repo *accountRepositoryDB) GetAccounts() ([]model.Account, error){

	accounts := []model.Account{}

	err := repo.db.Find(&accounts).Error
	if err != nil{
		return nil, err
	}

	return accounts, nil
}

func (repo *accountRepositoryDB) EditAccountInfo(account model.Account, id int) (*model.Account, error) {
	
	err := repo.db.Model(&account).Where("id = ?", id).Updates(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

