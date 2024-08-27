package intf

import (
	"SmartShopper-services/core/dto"
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
)

type IShopService interface {
	Create(ctx context.Context, shopDTO *dto.ShopDTO) error
	GetByID(ctx context.Context, ID uuid.UUID) (*models.ShopModel, error)
	DeleteByID(ctx context.Context, ID uuid.UUID) error
	GetByParams(ctx context.Context, params *dto.ShopParamsDTO) ([]*models.ShopModel, error)
}
