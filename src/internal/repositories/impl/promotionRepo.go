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

type PromotionRepo struct {
	db *sqlx.DB
}

func NewPromotionRepo(db *sqlx.DB) intfRepo.IPromotionRepo {
	return &PromotionRepo{db: db}
}

func (pr *PromotionRepo) GetByID(ctx context.Context, ID uuid.UUID) (*models.PromotionModel, error) {
	query := `select id, type, description, discount_size, start_date, end_date from ss.promotion where id = $1`

	var promotion models.PromotionModel
	err := pr.db.GetContext(ctx, &promotion, query, ID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrPromotionDoesNotExists
	}

	return &promotion, nil
}
