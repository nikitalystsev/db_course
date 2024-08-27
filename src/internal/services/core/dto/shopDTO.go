package dto

type ShopDTO struct {
	Retailer   SupplierDTO
	ShopParams ShopParamsDTO
}

type ShopParamsDTO struct {
	Title       string `json:"title"`
	Address     string `json:"address" `
	PhoneNumber string `json:"phone_number"`
	FioDirector string `json:"fio_director"`
	Limit       int    `json:"limit"`
	Offset      int    `json:"offset"`
}
