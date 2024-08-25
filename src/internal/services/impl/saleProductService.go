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
