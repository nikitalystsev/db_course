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

type ShopService struct {
	shopRepo           intfRepo.IShopRepo
	supplierRepo       intfRepo.ISupplierRepo
	transactionManager transact.ITransactionManager
}

func NewShopService(
	shopRepo intfRepo.IShopRepo,
	supplierRepo intfRepo.ISupplierRepo,
	transactionManager transact.ITransactionManager,
) intf.IShopService {
	return &ShopService{
		shopRepo:           shopRepo,
		supplierRepo:       supplierRepo,
		transactionManager: transactionManager,
	}
}

func (ss *ShopService) Create(ctx context.Context, shopDTO *dto.ShopDTO) error {
	return ss.transactionManager.Do(ctx, func(ctx context.Context) error {
		retailerID, err := ss.addRetailerIfNotExists(ctx, &shopDTO.Retailer)
		if err != nil {
			return err
		}

		err = ss.addShopIfNotExists(ctx, retailerID, &shopDTO.ShopParams)
		if err != nil {
			return err
		}

		return nil
	})
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

func (ss *ShopService) GetByParams(ctx context.Context, params *dto.ShopParamsDTO) ([]*models.ShopModel, error) {
	shops, err := ss.shopRepo.GetByParams(ctx, params)
	if err != nil && !errors.Is(err, errs.ErrShopDoesNotExists) {
		return nil, err
	}

	if errors.Is(err, errs.ErrShopDoesNotExists) {
		return nil, errs.ErrShopDoesNotExists
	}

	return shops, nil
}

func (ss *ShopService) addRetailerIfNotExists(ctx context.Context, retailerDTO *dto.SupplierDTO) (uuid.UUID, error) {
	existingRetailer, err := ss.supplierRepo.GetRetailerByAddress(ctx, retailerDTO.Address)
	if err != nil && !errors.Is(err, errs.ErrRetailerDoesNotExists) {
		return uuid.Nil, err
	}

	if existingRetailer != nil {
		return existingRetailer.ID, nil
	}

	retailer := &models.SupplierModel{
		ID:                uuid.New(),
		Title:             retailerDTO.Title,
		Address:           retailerDTO.Address,
		PhoneNumber:       retailerDTO.PhoneNumber,
		FioRepresentative: retailerDTO.FioRepresentative,
	}

	if err = ss.supplierRepo.CreateRetailer(ctx, retailer); err != nil {
		return uuid.Nil, err
	}

	return retailer.ID, nil
}

func (ss *ShopService) addShopIfNotExists(ctx context.Context, retailerID uuid.UUID, shopParams *dto.ShopParamsDTO) error {
	existingShop, err := ss.shopRepo.GetByAddress(ctx, shopParams.Address)
	if err != nil && !errors.Is(err, errs.ErrShopDoesNotExists) {
		return err
	}

	if existingShop != nil {
		return errs.ErrShopAlreadyExist
	}

	shop := &models.ShopModel{
		ID:          uuid.New(),
		RetailerID:  retailerID,
		Title:       shopParams.Title,
		Address:     shopParams.Address,
		PhoneNumber: shopParams.PhoneNumber,
		FioDirector: shopParams.FioDirector,
	}

	err = ss.shopRepo.Create(ctx, shop)
	if err != nil {
		return err
	}

	return nil
}
