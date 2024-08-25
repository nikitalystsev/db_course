package intfRepo

import (
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
)

type ISaleProductRepo interface {
	GetByProductID(ctx context.Context, productID uuid.UUID) ([]*models.SaleProductModel, error)
}
