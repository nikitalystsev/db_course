package impl

import (
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"SmartShopper-services/intf"
	"SmartShopper-services/intfRepo"
	"context"
	"errors"
	"github.com/google/uuid"
)

type SaleProductService struct {
	saleProductRepo intfRepo.ISaleProductRepo
}

func NewSaleProductService(saleProductRepo intfRepo.ISaleProductRepo) intf.ISaleProductService {
	return &SaleProductService{saleProductRepo: saleProductRepo}
}

func (sps *SaleProductService) GetByProductID(ctx context.Context, productID uuid.UUID) ([]*models.SaleProductModel, error) {
	sales, err := sps.saleProductRepo.GetByProductID(ctx, productID)
	if err != nil && !errors.Is(err, errs.ErrSaleProductDoesNotExists) {
		return nil, err
	}

	if errors.Is(err, errs.ErrSaleProductDoesNotExists) {
		return nil, errs.ErrSaleProductDoesNotExists
	}

	return sales, nil
}

func (sps *SaleProductService) GetByShopID(ctx context.Context, shopID uuid.UUID) ([]*models.SaleProductModel, error) {
	sales, err := sps.saleProductRepo.GetByShopID(ctx, shopID)
	if err != nil && !errors.Is(err, errs.ErrSaleProductDoesNotExists) {
		return nil, err
	}

	if errors.Is(err, errs.ErrSaleProductDoesNotExists) {
		return nil, errs.ErrSaleProductDoesNotExists
	}

	return sales, nil
}

func (sps *SaleProductService) GetByID(ctx context.Context, ID uuid.UUID) (*models.SaleProductModel, error) {
	sale, err := sps.saleProductRepo.GetByID(ctx, ID)
	if err != nil && !errors.Is(err, errs.ErrSaleProductDoesNotExists) {
		return nil, err
	}

	if errors.Is(err, errs.ErrSaleProductDoesNotExists) {
		return nil, errs.ErrSaleProductDoesNotExists
	}

	return sale, nil
}

func (sps *SaleProductService) Update(ctx context.Context, saleProduct *models.SaleProductModel) error {
	err := sps.saleProductRepo.Update(ctx, saleProduct)
	if err != nil {
		return err
	}

	return nil
}
