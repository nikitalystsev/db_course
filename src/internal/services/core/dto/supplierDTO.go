package dto

type SupplierDTO struct {
	Title             string `json:"title"`
	Address           string `json:"address"`
	PhoneNumber       string `json:"phone_number"`
	FioRepresentative string `json:"fio_representative"`
}
