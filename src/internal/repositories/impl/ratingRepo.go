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

type RatingRepo struct {
	db *sqlx.DB
}

func NewRatingRepo(db *sqlx.DB) intfRepo.IRatingRepo {
	return &RatingRepo{db: db}
}

func (rr *RatingRepo) Create(ctx context.Context, rating *models.RatingModel) error {
	query := `insert into ss.rating values ($1, $2, $3, $4, $5)`

	result, err := rr.db.ExecContext(ctx, query, rating.ID, rating.UserID,
		rating.SaleProductID, rating.Review, rating.Rating)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("ratingRepo.Create expected 1 row affected")
	}

	return nil
}

func (rr *RatingRepo) GetByUserAndSale(ctx context.Context, userID, saleID uuid.UUID) (*models.RatingModel, error) {
	query := `select id, user_id, sale_product_id, review, rating from ss.rating where user_id = $1 and sale_product_id = $2`

	var rating models.RatingModel
	err := rr.db.GetContext(ctx, &rating, query, userID, saleID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrRatingDoesNotExists
	}

	return &rating, nil
}

func (rr *RatingRepo) GetByID(ctx context.Context, ID uuid.UUID) (*models.RatingModel, error) {
	query := `select id, user_id, sale_product_id, review, rating from ss.rating where id = $1`

	var rating models.RatingModel
	err := rr.db.GetContext(ctx, &rating, query, ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrRatingDoesNotExists
	}

	return &rating, nil
}

func (rr *RatingRepo) DeleteByID(ctx context.Context, ID uuid.UUID) error {
	query := `delete from ss.rating where id = $1`

	result, err := rr.db.ExecContext(ctx, query, ID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("ratingRepo.DeleteByID expected 1 row affected")
	}

	return nil
}
