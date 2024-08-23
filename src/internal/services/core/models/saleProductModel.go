package models

import (
	"github.com/google/uuid"
	"time"
)

type SaleProductModel struct {
	ID          uuid.UUID `json:"id" db:"id"`
	ShopID      uuid.UUID `json:"shop_id" db:"shop_id"`
	ProductID   uuid.UUID `json:"product_id" db:"product_id"`
	PromotionID uuid.UUID `json:"promotion_id" db:"promotion_id"`
	Price       uuid.UUID `json:"price" db:"price"`
	Currency    string    `json:"currency" db:"currency"`
	SettingDate time.Time `json:"setting_date" db:"setting_date"`
	AvgRating   int       `json:"avg_rating" db:"avg_rating"`
}

type PromotionModel struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Type         string    `json:"type" db:"type"`
	Description  string    `json:"description" db:"description"`
	DiscountSize int       `json:"discount_size" db:"discount_size"`
	StartDate    time.Time `json:"start_date" db:"start_date"`
	EndDate      time.Time `json:"end_date" db:"end_date"`
}
