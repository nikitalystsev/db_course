package intfRepo

import (
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
)

type IRatingRepo interface {
	Create(ctx context.Context, rating *models.RatingModel) error
	GetByUserAndSale(ctx context.Context, userID uuid.UUID, saleID uuid.UUID) (*models.RatingModel, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.RatingModel, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
}
