package impl

import (
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"SmartShopper-services/intfRepo"
	"context"
	"database/sql"
	"errors"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SaleProductRepo struct {
	db     *sqlx.DB
	getter *trmsqlx.CtxGetter
}

func NewSaleProductRepo(db *sqlx.DB) intfRepo.ISaleProductRepo {
	return &SaleProductRepo{db: db, getter: trmsqlx.DefaultCtxGetter}
}

func (spr *SaleProductRepo) Create(ctx context.Context, saleProduct *models.SaleProductModel) error {
	query := `insert into ss.sale_product (id, shop_id, product_id, promotion_id, price, currency, setting_date) values ($1, $2, $3, $4, $5, $6, $7)`

	var promotionID interface{}
	if saleProduct.PromotionID == uuid.Nil {
		promotionID = nil
	} else {
		promotionID = saleProduct.PromotionID
	}

	result, err := spr.getter.DefaultTrOrDB(ctx, spr.db).ExecContext(ctx, query, saleProduct.ID, saleProduct.ShopID, saleProduct.ProductID,
		promotionID, saleProduct.Price, saleProduct.Currency, saleProduct.SettingDate,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("saleProductRepo.Create expected 1 row affected")
	}

	return nil
}

func (spr *SaleProductRepo) GetByProductID(ctx context.Context, productID uuid.UUID) ([]*models.SaleProductModel, error) {
	query := `select id, shop_id, product_id, promotion_id, price, currency, setting_date, avg_rating from ss.sale_product where product_id = $1`

	var sales []*models.SaleProductModel
	err := spr.getter.DefaultTrOrDB(ctx, spr.db).SelectContext(ctx, &sales, query, productID)
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
	err := spr.getter.DefaultTrOrDB(ctx, spr.db).SelectContext(ctx, &sales, query, shopID)
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
	err := spr.getter.DefaultTrOrDB(ctx, spr.db).GetContext(ctx, &sale, query, ID)
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

	result, err := spr.getter.DefaultTrOrDB(ctx, spr.db).ExecContext(ctx, query, saleProduct.Price, saleProduct.ID)
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
