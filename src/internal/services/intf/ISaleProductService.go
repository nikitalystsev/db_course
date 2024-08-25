package intf

import (
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
)

type ISaleProductService interface {
	GetByProductID(ctx context.Context, productID uuid.UUID) ([]*models.SaleProductModel, error)
}
