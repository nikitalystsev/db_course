package models

import (
	"github.com/google/uuid"
	"time"
)

type SaleProductModel struct {
	ID          uuid.UUID `json:"id" db:"id"`
	ShopID      uuid.UUID `json:"shop_id" db:"shop_id"`
	ProductID   uuid.UUID `json:"product_id" db:"product_id"`
	PriceID     uuid.UUID `json:"price_id" db:"price_id"`
	PromotionID uuid.UUID `json:"promotion_id" db:"promotion_id"`
	Rating      int       `json:"rating" db:"rating"`
}

type PriceModel struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Price       uuid.UUID `json:"price" db:"price"`
	Currency    string    `json:"currency" db:"currency"`
	SettingDate time.Time `json:"setting_date" db:"setting_date"`
}

type PromotionModel struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Type         string    `json:"type" db:"type"`
	Description  string    `json:"description" db:"description"`
	DiscountSize int       `json:"discount_size" db:"discount_size"`
	StartDate    time.Time `json:"start_date" db:"start_date"`
	EndDate      time.Time `json:"end_date" db:"end_date"`
}

type RatingModel struct {
	ID            uuid.UUID `json:"id" db:"id"`
	UserID        uuid.UUID `json:"user_id" db:"user_id"`
	SaleProductID uuid.UUID `json:"sale_product_id" db:"sale_product_id"`
	Review        string    `json:"review" db:"review"`
	Rating        int       `json:"rating" db:"rating"`
}
