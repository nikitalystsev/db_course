package impl

import (
	"SmartShopper-services/core/dto"
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"SmartShopper-services/intf"
	"SmartShopper-services/intfRepo"
	"SmartShopper-services/pkg/transact"
	"context"
	"errors"
	"github.com/google/uuid"
)

type SaleProductService struct {
	saleProductRepo    intfRepo.ISaleProductRepo
	supplierRepo       intfRepo.ISupplierRepo
	productRepo        intfRepo.IProductRepo
	promotionRepo      intfRepo.IPromotionRepo
	shopRepo           intfRepo.IShopRepo
	transactionManager transact.ITransactionManager
}

func NewSaleProductService(
	saleProductRepo intfRepo.ISaleProductRepo,
	supplierRepo intfRepo.ISupplierRepo,
	productRepo intfRepo.IProductRepo,
	promotionRepo intfRepo.IPromotionRepo,
	shopRepo intfRepo.IShopRepo,
	transactionManager transact.ITransactionManager,
) intf.ISaleProductService {
	return &SaleProductService{
		saleProductRepo:    saleProductRepo,
		supplierRepo:       supplierRepo,
		productRepo:        productRepo,
		promotionRepo:      promotionRepo,
		shopRepo:           shopRepo,
		transactionManager: transactionManager,
	}
}

func (sps *SaleProductService) Create(ctx context.Context, saleProduct *dto.NewSaleProductDTO) error {
	return sps.transactionManager.Do(ctx, func(ctx context.Context) error {
		// логика
	})
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
