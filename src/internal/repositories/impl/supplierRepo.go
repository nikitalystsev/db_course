package impl

import (
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"SmartShopper-services/intfRepo"
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SupplierRepo struct {
	db *sqlx.DB
}

func NewSupplierRepo(db *sqlx.DB) intfRepo.ISupplierRepo {
	return &SupplierRepo{db: db}
}

func (sr *SupplierRepo) CreateRetailer(ctx context.Context, retailer *models.SupplierModel) error {
	query := `insert into ss.retailer values ($1, $2, $3, $4, $5)`

	result, err := sr.db.ExecContext(ctx, query, retailer.ID, retailer.Title, retailer.Address,
		retailer.PhoneNumber, retailer.FioRepresentative)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("supplierRepo.CreateRetailer expected 1 row affected")
	}

	return nil
}

func (sr *SupplierRepo) CreateDistributor(ctx context.Context, distributor *models.SupplierModel) error {
	query := `insert into ss.distributor values ($1, $2, $3, $4, $5)`

	result, err := sr.db.ExecContext(ctx, query, distributor.ID, distributor.Title, distributor.Address,
		distributor.PhoneNumber, distributor.FioRepresentative)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("supplierRepo.CreateDistributor expected 1 row affected")
	}

	return nil
}

func (sr *SupplierRepo) CreateManufacturer(ctx context.Context, manufacturer *models.SupplierModel) error {
	query := `insert into ss.manufacturer values ($1, $2, $3, $4, $5)`

	result, err := sr.db.ExecContext(ctx, query, manufacturer.ID, manufacturer.Title, manufacturer.Address,
		manufacturer.PhoneNumber, manufacturer.FioRepresentative)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("supplierRepo.CreateManufacturer expected 1 row affected")
	}

	return nil
}

func (sr *SupplierRepo) GetRetailerByID(ctx context.Context, retailerID uuid.UUID) (*models.SupplierModel, error) {
	query := `select id, title, address, phone_number, fio_representative from ss.retailer where id = $1`

	var retailer models.SupplierModel
	err := sr.db.GetContext(ctx, &retailer, query, retailerID.String())
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrRetailerDoesNotExists
	}

	return &retailer, nil
}

func (sr *SupplierRepo) GetDistributorByID(ctx context.Context, distributorID uuid.UUID) (*models.SupplierModel, error) {
	query := `select id, title, address, phone_number, fio_representative from ss.distributor where id = $1`

	var retailer models.SupplierModel
	err := sr.db.GetContext(ctx, &retailer, query, distributorID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrDistributorDoesNotExists
	}

	return &retailer, nil
}

func (sr *SupplierRepo) GetManufacturerByID(ctx context.Context, manufacturerID uuid.UUID) (*models.SupplierModel, error) {
	query := `select id, title, address, phone_number, fio_representative from ss.manufacturer where id = $1`

	var retailer models.SupplierModel
	err := sr.db.GetContext(ctx, &retailer, query, manufacturerID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrRetailerDoesNotExists
	}

	return &retailer, nil
}

func (sr *SupplierRepo) DeleteRetailerByID(ctx context.Context, retailerID uuid.UUID) error {
	query := `delete from ss.retailer where id = $1`

	result, err := sr.db.ExecContext(ctx, query, retailerID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("supplierRepo.DeleteRetailerByID expected 1 row affected")
	}

	return nil
}

func (sr *SupplierRepo) DeleteDistributorByID(ctx context.Context, distributorID uuid.UUID) error {
	query := `delete from ss.distributor where id = $1`

	result, err := sr.db.ExecContext(ctx, query, distributorID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("supplierRepo.DeleteDistributorByID expected 1 row affected")
	}

	return nil
}

func (sr *SupplierRepo) DeleteManufacturerByID(ctx context.Context, manufacturerID uuid.UUID) error {
	query := `delete from ss.manufacturer where id = $1`

	result, err := sr.db.ExecContext(ctx, query, manufacturerID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("supplierRepo.DeleteManufacturerByID expected 1 row affected")
	}

	return nil
}
