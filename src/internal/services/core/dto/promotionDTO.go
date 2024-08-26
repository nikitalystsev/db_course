package dto

import (
	"time"
)

type PromotionDTO struct {
	Type         string    `json:"type"`
	Description  string    `json:"description"`
	DiscountSize int       `json:"discount_size"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
}
