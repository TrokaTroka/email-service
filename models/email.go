package models

import (
	"time"

	"github.com/google/uuid"
)

type Email struct {
	ID      uuid.UUID `json:"id" gorm:"primarykey"`
	To      string     `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Created time.Time `json:"created" gorm:"autoCreateTime"`
}

func New() (e *Email) {
	e = &Email{}
	return e
}
