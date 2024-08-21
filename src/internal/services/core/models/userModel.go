package models

import (
	"github.com/google/uuid"
	"time"
)

type UserModel struct {
	ID               uuid.UUID `json:"id" db:"id"`
	Fio              string    `json:"fio" db:"fio"`
	PhoneNumber      string    `json:"phone_number" db:"phone_number"`
	Password         string    `json:"password" db:"password"`
	RegistrationDate time.Time `json:"registration_date" db:"registration_date"`
}
