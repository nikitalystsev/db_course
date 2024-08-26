package intf

import (
	"SmartShopper-services/core/dto"
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
)

type ISaleProductService interface {
	Create(ctx context.Context, saleProduct *dto.NewSaleProductDTO) error
	GetByProductID(ctx context.Context, productID uuid.UUID) ([]*models.SaleProductModel, error)
	GetByShopID(ctx context.Context, shopID uuid.UUID) ([]*models.SaleProductModel, error)
	GetByID(ctx context.Context, ID uuid.UUID) (*models.SaleProductModel, error)
	Update(ctx context.Context, saleProduct *models.SaleProductModel) error
}
