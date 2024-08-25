package dto

import "github.com/google/uuid"

type ShopDTO struct {
	RetailerID  uuid.UUID `json:"retailer_id"`
	Title       string    `json:"title"`
	Address     string    `json:"address" `
	PhoneNumber string    `json:"phone_number"`
	FioDirector string    `json:"fio_director"`
}
