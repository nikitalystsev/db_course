package models

import "github.com/google/uuid"

type SupplierModel struct {
	ID                uuid.UUID `json:"id" db:"id"`
	Name              string    `json:"name" db:"name"`
	Address           string    `json:"address" db:"address"`
	PhoneNumber       string    `json:"phone_number" db:"phone_number"`
	FioRepresentative string    `json:"fio_representative" db:"fio_representative"`
}
