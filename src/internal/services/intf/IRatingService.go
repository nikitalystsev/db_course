package intf

import (
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
)

type IRatingService interface {
	Create(ctx context.Context, rating *models.RatingModel) error
	GetByID(ctx context.Context, ID uuid.UUID) (*models.RatingModel, error)
	DeleteByID(ctx context.Context, ID uuid.UUID) error
}
