package intfRepo

import (
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
)

type IShopRepo interface {
	Create(ctx context.Context, shop *models.ShopModel) error
	DeleteByID(ctx context.Context, ID uuid.UUID) error
	GetByID(ctx context.Context, ID uuid.UUID) (*models.ShopModel, error)
	GetByAddress(ctx context.Context, shopAddress string) (*models.ShopModel, error)
}
