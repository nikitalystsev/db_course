package intf

import (
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
)

type IPromotionService interface {
	GetByID(ctx context.Context, ID uuid.UUID) (*models.PromotionModel, error)
}
