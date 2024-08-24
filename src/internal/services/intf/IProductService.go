package intf

import (
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
)

type IProductService interface {
	Create(ctx context.Context, product *models.ProductModel) error
	GetByID(ctx context.Context, ID uuid.UUID) (*models.ProductModel, error)
	DeleteByID(ctx context.Context, ID uuid.UUID) error
	GetPage(ctx context.Context, limit, offset int) ([]*models.ProductModel, error)
}
