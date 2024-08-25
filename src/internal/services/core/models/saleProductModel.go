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
	Price       float32   `json:"price" db:"price"`
	Currency    string    `json:"currency" db:"currency"`
	SettingDate time.Time `json:"setting_date" db:"setting_date"`
	AvgRating   *float32  `json:"avg_rating" db:"avg_rating"`
}
