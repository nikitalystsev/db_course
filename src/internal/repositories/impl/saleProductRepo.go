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

type SaleProductRepo struct {
	db *sqlx.DB
}

func NewSaleProductRepo(db *sqlx.DB) intfRepo.ISaleProductRepo {
	return &SaleProductRepo{db: db}
}

func (spr *SaleProductRepo) GetByProductID(ctx context.Context, productID uuid.UUID) ([]*models.SaleProductModel, error) {
	query := `select id, shop_id, product_id, promotion_id, price, currency, setting_date, avg_rating from ss.sale_product where product_id = $1`

	var sales []*models.SaleProductModel
	err := spr.db.SelectContext(ctx, &sales, query, productID)
	if err != nil {
		return nil, err
	}
	if len(sales) == 0 {
		return nil, errs.ErrSaleProductDoesNotExists
	}

	return sales, nil
}

func (spr *SaleProductRepo) GetByShopID(ctx context.Context, shopID uuid.UUID) ([]*models.SaleProductModel, error) {
	query := `select id, shop_id, product_id, promotion_id, price, currency, setting_date, avg_rating from ss.sale_product where shop_id = $1`

	var sales []*models.SaleProductModel
	err := spr.db.SelectContext(ctx, &sales, query, shopID)
	if err != nil {
		return nil, err
	}
	if len(sales) == 0 {
		return nil, errs.ErrSaleProductDoesNotExists
	}

	return sales, nil
}

func (spr *SaleProductRepo) GetByID(ctx context.Context, ID uuid.UUID) (*models.SaleProductModel, error) {
	query := `select id, shop_id, product_id, promotion_id, price, currency, setting_date, avg_rating from ss.sale_product where id = $1`

	var sale models.SaleProductModel
	err := spr.db.GetContext(ctx, &sale, query, ID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrSaleProductDoesNotExists
	}

	return &sale, nil
}

func (spr *SaleProductRepo) Update(ctx context.Context, saleProduct *models.SaleProductModel) error {
	query := `update ss.sale_product set price = $1 where id = $2`

	result, err := spr.db.ExecContext(ctx, query, saleProduct.Price, saleProduct.ID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("saleProductRepo.Update expected 1 row affected")
	}
	return nil
}
