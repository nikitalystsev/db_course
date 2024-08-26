package impl

import (
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"SmartShopper-services/intfRepo"
	"context"
	"database/sql"
	"errors"
	"fmt"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SupplierRepo struct {
	db     *sqlx.DB
	getter *trmsqlx.CtxGetter
}

func NewSupplierRepo(db *sqlx.DB) intfRepo.ISupplierRepo {
	return &SupplierRepo{db: db, getter: trmsqlx.DefaultCtxGetter}
}

func (sr *SupplierRepo) CreateRetailer(ctx context.Context, retailer *models.SupplierModel) error {
	query := `insert into ss.retailer values ($1, $2, $3, $4, $5)`

	result, err := sr.getter.DefaultTrOrDB(ctx, sr.db).ExecContext(ctx, query, retailer.ID, retailer.Title, retailer.Address,
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

	result, err := sr.getter.DefaultTrOrDB(ctx, sr.db).ExecContext(ctx, query, distributor.ID, distributor.Title, distributor.Address,
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

	result, err := sr.getter.DefaultTrOrDB(ctx, sr.db).ExecContext(ctx, query, manufacturer.ID, manufacturer.Title, manufacturer.Address,
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
	err := sr.getter.DefaultTrOrDB(ctx, sr.db).GetContext(ctx, &retailer, query, retailerID.String())
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
	err := sr.getter.DefaultTrOrDB(ctx, sr.db).GetContext(ctx, &retailer, query, distributorID)
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
	err := sr.getter.DefaultTrOrDB(ctx, sr.db).GetContext(ctx, &retailer, query, manufacturerID)
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

	result, err := sr.getter.DefaultTrOrDB(ctx, sr.db).ExecContext(ctx, query, retailerID)
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

	result, err := sr.getter.DefaultTrOrDB(ctx, sr.db).ExecContext(ctx, query, distributorID)
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

	result, err := sr.getter.DefaultTrOrDB(ctx, sr.db).ExecContext(ctx, query, manufacturerID)
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

func (sr *SupplierRepo) GetRetailerByAddress(ctx context.Context, retailerAddress string) (*models.SupplierModel, error) {
	query := `select id, title, address, phone_number, fio_representative from ss.retailer where address = $1`

	var retailer models.SupplierModel
	err := sr.getter.DefaultTrOrDB(ctx, sr.db).GetContext(ctx, &retailer, query, retailerAddress)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrRetailerDoesNotExists
	}

	return &retailer, nil
}

func (sr *SupplierRepo) GetDistributorByAddress(ctx context.Context, distributorAddress string) (*models.SupplierModel, error) {
	query := `select id, title, address, phone_number, fio_representative from ss.distributor where address = $1`

	var distributor models.SupplierModel
	err := sr.getter.DefaultTrOrDB(ctx, sr.db).GetContext(ctx, &distributor, query, distributorAddress)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrRetailerDoesNotExists
	}

	return &distributor, nil
}

func (sr *SupplierRepo) GetManufacturerByAddress(ctx context.Context, manufacturerAddress string) (*models.SupplierModel, error) {
	query := `select id, title, address, phone_number, fio_representative from ss.manufacturer where address = $1`

	var manufacturer models.SupplierModel
	err := sr.getter.DefaultTrOrDB(ctx, sr.db).GetContext(ctx, &manufacturer, query, manufacturerAddress)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrRetailerDoesNotExists
	}

	return &manufacturer, nil
}

func (sr *SupplierRepo) IfExistsRetailerDistributor(ctx context.Context, retailerID, distributorID uuid.UUID) (bool, error) {
	fmt.Println("Дрочим на IfExistsRetailerDistributor")

	fmt.Println(retailerID.String(), " ", distributorID.String())
	query := `select retailer_id, distributor_id from ss.retailer_distributor where retailer_id = $1 and distributor_id = $2`

	var ids models.RetailerDistributorModel
	err := sr.getter.DefaultTrOrDB(ctx, sr.db).GetContext(ctx, &ids, query, retailerID, distributorID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		fmt.Println("ебаная женская писечка")
		return false, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("нету, бля")
		return false, nil
	}

	fmt.Println("есть")

	return true, nil
}

func (sr *SupplierRepo) IfExistsDistributorManufacturer(ctx context.Context, distributorID, manufacturerID uuid.UUID) (bool, error) {
	query := `select distributor_id, manufacturer_id from ss.distributor_manufacturer where distributor_id = $1 and manufacturer_id = $2`

	var ids models.DistributorManufacturerModel
	err := sr.getter.DefaultTrOrDB(ctx, sr.db).GetContext(ctx, &ids, query, distributorID, manufacturerID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	return true, nil
}

func (sr *SupplierRepo) CreateRetailerDistributor(ctx context.Context, retailerID, distributorID uuid.UUID) error {
	fmt.Println("call CreateRetailerDistributor")

	query := `insert into ss.retailer_distributor values ($1, $2)`

	result, err := sr.getter.DefaultTrOrDB(ctx, sr.db).ExecContext(ctx, query, retailerID, distributorID)
	if err != nil {
		fmt.Println("Ошибка, блять")
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Че сука еще то надо")
		return err
	}
	if rows != 1 {
		return errors.New("supplierRepo.CreateRetailerDistributor expected 1 row affected")
	}

	fmt.Println("все тип топ")

	return nil
}

func (sr *SupplierRepo) CreateDistributorManufacturer(ctx context.Context, distributorID, manufacturerID uuid.UUID) error {
	query := `insert into ss.distributor_manufacturer values ($1, $2)`

	result, err := sr.getter.DefaultTrOrDB(ctx, sr.db).ExecContext(ctx, query, distributorID, manufacturerID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("supplierRepo.CreateDistributorManufacturer expected 1 row affected")
	}

	return nil
}
