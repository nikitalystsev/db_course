package dto

import "github.com/google/uuid"

type RatingDTO struct {
	SaleProductID uuid.UUID `json:"sale_product_id" db:"sale_product_id"`
	Review        string    `json:"review" db:"review"`
	Rating        int       `json:"rating" db:"rating"`
}
