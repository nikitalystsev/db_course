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

type PromotionRepo struct {
	db     *sqlx.DB
	getter *trmsqlx.CtxGetter
}

func NewPromotionRepo(db *sqlx.DB) intfRepo.IPromotionRepo {
	return &PromotionRepo{db: db, getter: trmsqlx.DefaultCtxGetter}
}

func (pr *PromotionRepo) Create(ctx context.Context, promotion *models.PromotionModel) error {
	query := `insert into ss.promotion values ($1, $2, $3, $4, $5, $6)`

	result, err := pr.getter.DefaultTrOrDB(ctx, pr.db).ExecContext(ctx, query, promotion.ID, promotion.Type, promotion.Description,
		promotion.DiscountSize, promotion.StartDate, promotion.EndDate)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("promotionRepo.Create expected 1 row affected")
	}

	return nil
}

func (pr *PromotionRepo) GetByID(ctx context.Context, ID uuid.UUID) (*models.PromotionModel, error) {
	query := `select id, type, description, discount_size, start_date, end_date from ss.promotion where id = $1`

	var promotion models.PromotionModel
	err := pr.getter.DefaultTrOrDB(ctx, pr.db).GetContext(ctx, &promotion, query, ID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrPromotionDoesNotExists
	}

	return &promotion, nil
}
