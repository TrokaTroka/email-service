package migrations

import (
	"mail-sender/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.Email{})
}