package email

import (
	"mail-sender/database"
	"mail-sender/models"

	"github.com/google/uuid"
)

func CreateEmail(email *models.Email) error {

	email.ID = uuid.New()

	db := database.GetDB()

	db.Create(&email)
	return nil
}