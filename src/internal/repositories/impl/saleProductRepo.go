package impl

import (
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"SmartShopper-services/intfRepo"
	"context"
	"fmt"
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

	fmt.Println("productID: ", productID)

	var sales []*models.SaleProductModel
	err := spr.db.SelectContext(ctx, &sales, query, productID)
	if err != nil {
		fmt.Println("ошибка выполения запроса")
		return nil, err
	}
	if len(sales) == 0 {
		return nil, errs.ErrSaleProductDoesNotExists
	}

	return sales, nil
}
