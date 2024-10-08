package intfRepo

import (
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
)

type ISupplierRepo interface {
	CreateRetailer(ctx context.Context, retailer *models.SupplierModel) error
	CreateDistributor(ctx context.Context, distributor *models.SupplierModel) error
	CreateManufacturer(ctx context.Context, manufacturer *models.SupplierModel) error

	GetRetailerByID(ctx context.Context, retailerID uuid.UUID) (*models.SupplierModel, error)
	GetDistributorByID(ctx context.Context, distributorID uuid.UUID) (*models.SupplierModel, error)
	GetManufacturerByID(ctx context.Context, manufacturerID uuid.UUID) (*models.SupplierModel, error)

	DeleteRetailerByID(ctx context.Context, retailerID uuid.UUID) error
	DeleteDistributorByID(ctx context.Context, distributorID uuid.UUID) error
	DeleteManufacturerByID(ctx context.Context, manufacturerID uuid.UUID) error

	GetRetailerByAddress(ctx context.Context, retailerAddress string) (*models.SupplierModel, error)
	GetDistributorByAddress(ctx context.Context, distributorAddress string) (*models.SupplierModel, error)
	GetManufacturerByAddress(ctx context.Context, manufacturerAddress string) (*models.SupplierModel, error)

	IfExistsRetailerDistributor(ctx context.Context, retailerID, distributorID uuid.UUID) (bool, error)
	IfExistsDistributorManufacturer(ctx context.Context, distributorID, manufacturerID uuid.UUID) (bool, error)

	CreateRetailerDistributor(ctx context.Context, retailerID, distributorID uuid.UUID) error
	CreateDistributorManufacturer(ctx context.Context, distributorID, manufacturerID uuid.UUID) error
}
