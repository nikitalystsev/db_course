package models

import "github.com/google/uuid"

type SupplierModel struct {
	ID                uuid.UUID `json:"id" db:"id"`
	Title             string    `json:"title" db:"title"`
	Address           string    `json:"address" db:"address"`
	PhoneNumber       string    `json:"phone_number" db:"phone_number"`
	FioRepresentative string    `json:"fio_representative" db:"fio_representative"`
}

type RetailerDistributorModel struct {
	RetailerID    uuid.UUID `json:"retailer_id" db:"retailer_id"`
	DistributorID uuid.UUID `json:"distributor_id" db:"distributor_id"`
}

type DistributorManufacturerModel struct {
	DistributorID  uuid.UUID `json:"distributor_id" db:"distributor_id"`
	ManufacturerID uuid.UUID `json:"manufacturer_id" db:"manufacturer_id"`
}
