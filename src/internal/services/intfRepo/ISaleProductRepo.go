package intfRepo

import (
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
)

type ISaleProductRepo interface {
	GetByProductID(ctx context.Context, productID uuid.UUID) ([]*models.SaleProductModel, error)
	GetByShopID(ctx context.Context, shopID uuid.UUID) ([]*models.SaleProductModel, error)
	GetByID(ctx context.Context, ID uuid.UUID) (*models.SaleProductModel, error)
	Update(ctx context.Context, saleProduct *models.SaleProductModel) error
	Create(ctx context.Context, saleProduct *models.SaleProductModel) error
}
