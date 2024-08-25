package models

import (
	"github.com/google/uuid"
	"time"
)

type PromotionModel struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Type         string    `json:"type" db:"type"`
	Description  string    `json:"description" db:"description"`
	DiscountSize int       `json:"discount_size" db:"discount_size"`
	StartDate    time.Time `json:"start_date" db:"start_date"`
	EndDate      time.Time `json:"end_date" db:"end_date"`
}
