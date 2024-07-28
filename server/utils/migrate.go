package utils

import (
	"github.com/Nas-virat/PFin-personal-finance/account"
	"github.com/Nas-virat/PFin-personal-finance/balance"
	"github.com/Nas-virat/PFin-personal-finance/transaction"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {

	db.AutoMigrate(
		&account.Account{},
		&transaction.Transaction{},
		&balance.Debt{},
	)
}
