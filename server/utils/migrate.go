package utils

import (
	"github.com/Nas-virat/PFin-personal-finance/model"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {

	db.AutoMigrate(
		&model.Account{},
		&model.Transaction{},
		&model.Debt{},
	)
}
