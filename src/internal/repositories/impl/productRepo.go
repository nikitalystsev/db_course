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

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) intfRepo.IProductRepo {
	return &ProductRepo{db: db}
}

func (pr *ProductRepo) Create(ctx context.Context, product *models.ProductModel) error {
	query := `insert into ss.product values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	result, err := pr.db.ExecContext(ctx, query, product.ID, product.RetailerID, product.DistributorID,
		product.ManufacturerID, product.Name, product.Categories, product.Brand, product.Compound,
		product.GrossMass, product.NetMass, product.PackageType,
	)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("productRepo.Create expected 1 row affected")
	}

	return nil
}

func (pr *ProductRepo) GetByID(ctx context.Context, ID uuid.UUID) (*models.ProductModel, error) {
	query := `select id, retailer_id, distributor_id, manufacturer_id, name, categories, brand, compound, gross_mass, net_mass, package_type from ss.product where id = $1`

	var product models.ProductModel
	err := pr.db.GetContext(ctx, &product, query, ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrProductDoesNotExists
	}

	return &product, nil
}

func (pr *ProductRepo) DeleteByID(ctx context.Context, ID uuid.UUID) error {
	query := `delete from ss.product where id = $1`

	result, err := pr.db.ExecContext(ctx, query, ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("productRepo.DeleteByID expected 1 row affected")
	}

	return nil
}

func (pr *ProductRepo) GetPage(ctx context.Context, limit, offset int) ([]*models.ProductModel, error) {
	query := `select id, retailer_id, distributor_id, manufacturer_id, name, categories, brand, compound, gross_mass, net_mass, package_type from ss.product limit $1 offset $2`

	var products []*models.ProductModel
	err := pr.db.SelectContext(ctx, &products, query, limit, offset)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrProductDoesNotExists
	}

	return products, nil
}
