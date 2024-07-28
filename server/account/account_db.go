package account

import (
	"gorm.io/gorm"
)

type accountRepositoryDB struct {
	db *gorm.DB
}


// Constructor
func NewAccountRepositoryDB(db *gorm.DB) AccountRepository {
	return &accountRepositoryDB{db}
}

func (repo *accountRepositoryDB) CreateAccount(account Account) (*Account, error) {
	
	result := repo.db.Create(&account)

	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

func (repo *accountRepositoryDB) GetAccountById(id int) (*Account, error) {
	
	account := Account{}
	err := repo.db.Where("id = ?", id).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (repo *accountRepositoryDB) GetAccounts() ([]Account, error){

	accounts := []Account{}

	err := repo.db.Find(&accounts).Error
	if err != nil{
		return nil, err
	}

	return accounts, nil
}

func (repo *accountRepositoryDB) EditAccountInfo(account Account, id int) (*Account, error) {
	
	err := repo.db.Model(&account).Where("id = ?", id).Updates(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

