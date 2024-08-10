package transaction

import (
	"gorm.io/gorm"
)

type transactionRepositoryDB struct {
	db *gorm.DB
}

// Constructor
func NewTransactionRepositoryDB(db *gorm.DB) TransactionRepository {
	return &transactionRepositoryDB{db: db}
}

func (r transactionRepositoryDB) CreateTransaction(transaction Transaction) (*Transaction, error) {

	result := r.db.Create(&transaction)

	if result.Error != nil {
		return nil, result.Error
	}

	return &transaction, nil
}

func (r transactionRepositoryDB) GetTransactionByID(id uint) (*Transaction, error) {

	transaction := Transaction{}

	result := r.db.Find(&transaction, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &transaction, nil
}

func (r transactionRepositoryDB) GetTransactions() ([]Transaction, error) {

	transactions := []Transaction{}

	err := r.db.Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r transactionRepositoryDB) GetTransactionInYear(year int) ([]Transaction, error) {

	transactions := []Transaction{}

	err := r.db.Where("transaction_year = ?", year).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r transactionRepositoryDB) GetTransactionInRangeMonthYear(month, year int) ([]Transaction, error) {

	transactions := []Transaction{}

	err := r.db.Where("transaction_month = ? AND transaction_year = ?", month, year).Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r transactionRepositoryDB) GetTransactionInRangeDayMonthYear(day, month, year int) ([]Transaction, error) {
	transactions := []Transaction{}

	err := r.db.Where("transaction_date = ? AND transaction_month = ? AND transaction_year = ?", day, month, year).Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r transactionRepositoryDB) UpdateTransaction(id uint, newInfo Transaction) (*Transaction, error) {
	transaction := Transaction{}

	// find transaction
	result := r.db.Find(&transaction, id)

	if result.Error != nil {
		return nil, result.Error
	}

	// update the transaction
	result = r.db.Model(&transaction).Updates(Transaction{
		TransactionType:  newInfo.TransactionType,
		Category:         newInfo.Category,
		Description:      newInfo.Description,
		Amount:           newInfo.Amount,
		TransactionDate:  newInfo.TransactionDate,
		TransactionMonth: newInfo.TransactionMonth,
		TransactionYear:  newInfo.TransactionYear,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return &transaction, nil

}

func (r transactionRepositoryDB) DeleteTransaction(id uint) error {

	transaction := Transaction{}

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
