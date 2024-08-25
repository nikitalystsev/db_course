package impl

import (
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"SmartShopper-services/intfRepo"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type ShopService struct {
	shopRepo intfRepo.IShopRepo
}

func NewShopService(shopRepo intfRepo.IShopRepo) *ShopService {
	return &ShopService{shopRepo: shopRepo}
}

func (ss *ShopService) Create(ctx context.Context, shop *models.ShopModel) error {
	fmt.Println("call shopRepo.Create")

	existingShop, err := ss.shopRepo.GetByAddress(ctx, shop.Address)
	if err != nil && !errors.Is(err, errs.ErrShopDoesNotExists) {
		return err
	}

	if existingShop != nil {
		return errs.ErrShopAlreadyExist
	}

	fmt.Println("Все тип топ. Магазина еще не создавали")

	err = ss.shopRepo.Create(ctx, shop)
	if err != nil {
		return err
	}

	return nil
}

func (ss *ShopService) GetByID(ctx context.Context, ID uuid.UUID) (*models.ShopModel, error) {
	existingShop, err := ss.shopRepo.GetByID(ctx, ID)
	if err != nil && !errors.Is(err, errs.ErrShopDoesNotExists) {
		return nil, err
	}

	if existingShop == nil {
		return nil, errs.ErrShopDoesNotExists
	}

	return existingShop, nil
}

func (ss *ShopService) DeleteByID(ctx context.Context, ID uuid.UUID) error {
	existingShop, err := ss.shopRepo.GetByID(ctx, ID)
	if err != nil && !errors.Is(err, errs.ErrShopDoesNotExists) {
		return err
	}

	if existingShop == nil {
		return errs.ErrShopDoesNotExists
	}

	err = ss.shopRepo.DeleteByID(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
