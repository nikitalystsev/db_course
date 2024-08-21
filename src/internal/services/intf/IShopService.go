package intf

import (
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
)

type IShopService interface {
	Create(ctx context.Context, shop *models.ShopModel) error
	GetByID(ctx context.Context, ID uuid.UUID) (*models.ShopModel, error)
	DeleteByID(ctx context.Context, ID uuid.UUID) error
}
