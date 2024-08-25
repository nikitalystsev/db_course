package intfRepo

import (
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
)

type IPromotionRepo interface {
	GetByID(ctx context.Context, ID uuid.UUID) (*models.PromotionModel, error)
}
