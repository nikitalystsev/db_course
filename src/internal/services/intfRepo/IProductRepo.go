package intfRepo

import (
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
)

type IProductRepo interface {
	Create(ctx context.Context, product *models.ProductModel) error
	GetByID(ctx context.Context, ID uuid.UUID) (*models.ProductModel, error)
	DeleteByID(ctx context.Context, ID uuid.UUID) error
}
