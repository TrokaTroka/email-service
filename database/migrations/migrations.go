package migrations

import (
	"gorm.io/gorm"
	"mail-sender/models"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Email{})
}