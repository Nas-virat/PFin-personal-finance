package account

import (
	"gorm.io/gorm"
)

type AccountRepository interface {
	CreateAccount(account Account) (*Account, error)
	GetAccountById(id int) (*Account, error)
	GetAccounts() ([]Account, error)
	EditAccountInfo(account Account,id int) (*Account, error)
}


type accountRepository struct {
	db *gorm.DB
}


// Constructor
func NewAccountRepositoryDB(db *gorm.DB) *accountRepository {
	return &accountRepository{db}
}

func (repo *accountRepository) CreateAccount(account Account) (*Account, error) {
	
	result := repo.db.Create(&account)

	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

func (repo *accountRepository) GetAccountById(id int) (*Account, error) {
	
	account := Account{}
	err := repo.db.Where("id = ?", id).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (repo *accountRepository) GetAccounts() ([]Account, error){

	accounts := []Account{}

	err := repo.db.Find(&accounts).Error
	if err != nil{
		return nil, err
	}

	return accounts, nil
}

func (repo *accountRepository) EditAccountInfo(account Account, id int) (*Account, error) {
	
	err := repo.db.Model(&account).Where("id = ?", id).Updates(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

