package models

import "github.com/google/uuid"

type Email struct {
	ID      uuid.UUID `json:"id"`
	To      string     `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func New() (e *Email) {
	e = &Email{}
	return e
}
