package repository

import (
	"github.com/Nas-virat/PFin-personal-finance/model"
	"gorm.io/gorm"
)

type transactionRepositoryDB struct {
	db *gorm.DB
}

// Constructor
func NewTransactionRepositoryDB(db *gorm.DB) TransactionRepository {
	return &transactionRepositoryDB{db: db}
}

func (r transactionRepositoryDB) CreateTransaction(transaction model.Transaction) (*model.Transaction, error) {

	result := r.db.Create(&transaction)

	if result.Error != nil {
		return nil, result.Error
	}

	return &transaction, nil
}

func (r transactionRepositoryDB) GetTransactionByID(id uint) (*model.Transaction, error) {

	transaction := model.Transaction{}

	result := r.db.Find(&transaction, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &transaction, nil
}

func (r transactionRepositoryDB) GetTransactions() ([]model.Transaction, error) {

	transactions := []model.Transaction{}

	err := r.db.Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r transactionRepositoryDB) GetTransactionInRangeMonthYear(month, year int) ([]model.Transaction, error) {

	transactions := []model.Transaction{}

	err := r.db.Where("extract(month from created_at) = ? AND extract(year from created_at) = ?", month, year).Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r transactionRepositoryDB) GetTransactionInRangeDayMonthYear(day,month,year int) ([]model.Transaction,error){
	transactions := []model.Transaction{}

	err := r.db.Where("extract(day from created_at) = ? AND extract(month from created_at) = ? AND extract(year from created_at) = ?", day,month,year).Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions,nil
}

func (r transactionRepositoryDB) UpdateTransaction(id uint, newInfo model.Transaction) (*model.Transaction, error) {
	transaction := model.Transaction{}

	// find transaction
	result := r.db.Find(&transaction, id)

	if result.Error != nil {
		return nil, result.Error
	}

	// update the transaction
	result = r.db.Model(&transaction).Updates(model.Transaction{
		TransactionType: newInfo.TransactionType,
		Category:        newInfo.Category,
		Description:     newInfo.Description,
		Amount:          newInfo.Amount,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return &transaction, nil

}

func (r transactionRepositoryDB) DeleteTransaction(id uint) error {

	transaction := model.Transaction{}

	result := r.db.Find(&transaction, id)

	if result.Error != nil {
		return result.Error
	}

	result = r.db.Delete(&transaction, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
