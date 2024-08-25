package dto

import (
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
