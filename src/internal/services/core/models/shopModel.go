package models

import "github.com/google/uuid"

type ShopModel struct {
	ID          uuid.UUID `json:"id" db:"id"`
	RetailerID  uuid.UUID `json:"retailer_id" db:"retailer_id"`
	Title       string    `json:"title" db:"title"`
	Address     string    `json:"address" db:"address"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	FioDirector string    `json:"fio_director" db:"fio_director"`
}
