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

type ProductService struct {
	productRepo intfRepo.IProductRepo
}

func NewProductService(productRepo intfRepo.IProductRepo) intf.IProductService {
	return &ProductService{productRepo: productRepo}
}

func (ps *ProductService) Create(ctx context.Context, product *models.ProductModel) error {
	existingProduct, err := ps.productRepo.GetByID(ctx, product.ID)
	if err != nil && !errors.Is(err, errs.ErrProductDoesNotExists) {
		return err
	}

	if existingProduct != nil {
		return errs.ErrProductAlreadyExist
	}

	err = ps.productRepo.Create(ctx, product)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProductService) GetByID(ctx context.Context, ID uuid.UUID) (*models.ProductModel, error) {
	existingProduct, err := ps.productRepo.GetByID(ctx, ID)
	if err != nil && !errors.Is(err, errs.ErrProductDoesNotExists) {
		return nil, err
	}

	if existingProduct == nil {
		return nil, errs.ErrProductDoesNotExists
	}

	return existingProduct, nil
}

func (ps *ProductService) DeleteByID(ctx context.Context, ID uuid.UUID) error {
	existingProduct, err := ps.productRepo.GetByID(ctx, ID)
	if err != nil && !errors.Is(err, errs.ErrProductDoesNotExists) {
		return err
	}

	if existingProduct == nil {
		return errs.ErrProductDoesNotExists
	}

	err = ps.productRepo.DeleteByID(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProductService) GetPage(ctx context.Context, limit, offset int) ([]*models.ProductModel, error) {
	products, err := ps.productRepo.GetPage(ctx, limit, offset)
	if err != nil && !errors.Is(err, errs.ErrProductDoesNotExists) {
		return nil, err
	}
	if products == nil {
		return nil, errs.ErrProductDoesNotExists
	}

	return products, nil
}
