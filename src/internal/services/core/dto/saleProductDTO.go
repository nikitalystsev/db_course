package dto

import (
	"github.com/google/uuid"
	"time"
)

type SaleProductDTO struct {
	ShopTitle             string    `json:"shop_title" `
	ShopAddress           string    `json:"shop_address"`
	PromotionType         string    `json:"promotion_type"`
	PromotionDescription  string    `json:"promotion_description"`
	PromotionDiscountSize *int      `json:"promotion_discount_size"`
	Price                 float32   `json:"price"`
	Currency              string    `json:"currency"`
	SettingDate           time.Time `json:"setting_date"`
	AvgRating             *float32  `json:"avg_rating"`
}

type SaleProductShopDTO struct {
	ID                    uuid.UUID `json:"id"`
	ProductName           string    `json:"name"`
	ProductCategories     string    `json:"categories"`
	PromotionType         string    `json:"promotion_type"`
	PromotionDiscountSize *int      `json:"promotion_discount_size"`
	Price                 float32   `json:"price" db:"price"`
	Currency              string    `json:"currency" db:"currency"`
	SettingDate           time.Time `json:"setting_date" db:"setting_date"`
	AvgRating             *float32  `json:"avg_rating" db:"avg_rating"`
}
