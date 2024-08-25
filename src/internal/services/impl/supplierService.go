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

type SupplierService struct {
	supplierRepo intfRepo.ISupplierRepo
}

func NewSupplierService(supplierRepo intfRepo.ISupplierRepo) intf.ISupplierService {
	return &SupplierService{supplierRepo: supplierRepo}
}

func (ss *SupplierService) CreateRetailer(ctx context.Context, retailer *models.SupplierModel) error {
	existingRetailer, err := ss.supplierRepo.GetRetailerByAddress(ctx, retailer.Address)
	if err != nil && !errors.Is(err, errs.ErrRetailerDoesNotExists) {
		return err
	}

	if existingRetailer != nil {
		return errs.ErrRetailerAlreadyExist
	}

	err = ss.supplierRepo.CreateRetailer(ctx, retailer)
	if err != nil {
		return err
	}

	return nil
}

func (ss *SupplierService) CreateDistributor(ctx context.Context, retailer *models.SupplierModel) error {
	existingRetailer, err := ss.supplierRepo.GetDistributorByID(ctx, retailer.ID)
	if err != nil && !errors.Is(err, errs.ErrDistributorDoesNotExists) {
		return err
	}

	if existingRetailer != nil {
		return errs.ErrDistributorAlreadyExist
	}

	err = ss.supplierRepo.CreateDistributor(ctx, retailer)
	if err != nil {
		return err
	}

	return nil
}

func (ss *SupplierService) CreateManufacturer(ctx context.Context, retailer *models.SupplierModel) error {
	existingRetailer, err := ss.supplierRepo.GetManufacturerByID(ctx, retailer.ID)
	if err != nil && !errors.Is(err, errs.ErrManufacturerDoesNotExists) {
		return err
	}

	if existingRetailer != nil {
		return errs.ErrManufacturerAlreadyExist
	}

	err = ss.supplierRepo.CreateManufacturer(ctx, retailer)
	if err != nil {
		return err
	}

	return nil
}

func (ss *SupplierService) GetRetailerByID(ctx context.Context, retailerID uuid.UUID) (*models.SupplierModel, error) {
	existingRetailer, err := ss.supplierRepo.GetRetailerByID(ctx, retailerID)
	if err != nil && !errors.Is(err, errs.ErrRetailerDoesNotExists) {
		return nil, err
	}

	if existingRetailer == nil {
		return nil, errs.ErrRetailerDoesNotExists
	}

	return existingRetailer, nil
}

func (ss *SupplierService) GetDistributorByID(ctx context.Context, distributorID uuid.UUID) (*models.SupplierModel, error) {
	existingRetailer, err := ss.supplierRepo.GetDistributorByID(ctx, distributorID)
	if err != nil && !errors.Is(err, errs.ErrDistributorDoesNotExists) {
		return nil, err
	}

	if existingRetailer == nil {
		return nil, errs.ErrDistributorDoesNotExists
	}

	return existingRetailer, nil
}

func (ss *SupplierService) GetManufacturerByID(ctx context.Context, retailerID uuid.UUID) (*models.SupplierModel, error) {
	existingRetailer, err := ss.supplierRepo.GetManufacturerByID(ctx, retailerID)
	if err != nil && !errors.Is(err, errs.ErrManufacturerDoesNotExists) {
		return nil, err
	}

	if existingRetailer == nil {
		return nil, errs.ErrManufacturerDoesNotExists
	}

	return existingRetailer, nil
}

func (ss *SupplierService) DeleteRetailerByID(ctx context.Context, retailerID uuid.UUID) error {
	existingRetailer, err := ss.supplierRepo.GetRetailerByID(ctx, retailerID)
	if err != nil && !errors.Is(err, errs.ErrRetailerDoesNotExists) {
		return err
	}

	if existingRetailer == nil {
		return errs.ErrRetailerDoesNotExists
	}

	err = ss.supplierRepo.DeleteRetailerByID(ctx, retailerID)
	if err != nil {
		return err
	}

	return nil
}

func (ss *SupplierService) DeleteDistributorByID(ctx context.Context, distributorID uuid.UUID) error {
	existingRetailer, err := ss.supplierRepo.GetDistributorByID(ctx, distributorID)
	if err != nil && !errors.Is(err, errs.ErrDistributorDoesNotExists) {
		return err
	}

	if existingRetailer == nil {
		return errs.ErrDistributorDoesNotExists
	}

	err = ss.supplierRepo.DeleteDistributorByID(ctx, distributorID)
	if err != nil {
		return err
	}

	return nil
}

func (ss *SupplierService) DeleteManufacturerByID(ctx context.Context, manufacturerID uuid.UUID) error {
	existingRetailer, err := ss.supplierRepo.GetManufacturerByID(ctx, manufacturerID)
	if err != nil && !errors.Is(err, errs.ErrManufacturerDoesNotExists) {
		return err
	}

	if existingRetailer == nil {
		return errs.ErrManufacturerDoesNotExists
	}

	err = ss.supplierRepo.DeleteManufacturerByID(ctx, manufacturerID)
	if err != nil {
		return err
	}

	return nil
}

func (ss *SupplierService) GetRetailerByAddress(ctx context.Context, retailerAddress string) (*models.SupplierModel, error) {
	existingRetailer, err := ss.supplierRepo.GetRetailerByAddress(ctx, retailerAddress)
	if err != nil && !errors.Is(err, errs.ErrRetailerDoesNotExists) {
		return nil, err
	}

	if existingRetailer == nil {
		return nil, errs.ErrRetailerDoesNotExists
	}

	return existingRetailer, nil
}
