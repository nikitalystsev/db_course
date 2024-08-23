package models

import "github.com/google/uuid"

type RatingModel struct {
	ID            uuid.UUID `json:"id" db:"id"`
	UserID        uuid.UUID `json:"user_id" db:"user_id"`
	SaleProductID uuid.UUID `json:"sale_product_id" db:"sale_product_id"`
	Review        string    `json:"review" db:"review"`
	Rating        int       `json:"rating" db:"rating"`
}
