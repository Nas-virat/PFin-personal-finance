package utils

import "gorm.io/gorm"
import "github.com/Nas-virat/PFin-personal-finance/model"


func Migration(db *gorm.DB) {
	
	db.AutoMigrate(&model.Account{})
}