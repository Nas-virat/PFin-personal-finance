package account

import (
	"gorm.io/gorm"
)

type AccountRepository interface {
	CreateAccount(accountModel Account) (*Account, error)
	GetAccountById(id int) (*Account, error)
	GetAccounts() ([]Account, error)
	EditAccountInfo(accountModel Account, id int) (*Account, error)
}

type accountRepository struct {
	db *gorm.DB
}

// Constructor
func NewAccountRepositoryDB(db *gorm.DB) *accountRepository {
	return &accountRepository{db}
}

func (repo *accountRepository) CreateAccount(accountModel Account) (*Account, error) {

	result := repo.db.Create(&accountModel)

	if result.Error != nil {
		return nil, result.Error
	}
	return &accountModel, nil
}

func (repo *accountRepository) GetAccountById(id int) (*Account, error) {

	account := Account{}
	err := repo.db.Where("id = ?", id).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (repo *accountRepository) GetAccounts() ([]Account, error) {

	accounts := []Account{}

	err := repo.db.Find(&accounts).Error
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (repo *accountRepository) EditAccountInfo(accountModel Account, id int) (*Account, error) {

	err := repo.db.Model(&accountModel).Where("id = ?", id).Updates(&accountModel).Error
	if err != nil {
		return nil, err
	}
	return &accountModel, nil
}
