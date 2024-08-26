package impl

import (
	"SmartShopper-services/core/dto"
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

type ShopRepo struct {
	db     *sqlx.DB
	getter *trmsqlx.CtxGetter
}

func NewShopRepo(db *sqlx.DB) intfRepo.IShopRepo {
	return &ShopRepo{db: db, getter: trmsqlx.DefaultCtxGetter}
}

func (sr *ShopRepo) Create(ctx context.Context, shop *models.ShopModel) error {
	query := `insert into ss.shop values ($1, $2, $3, $4, $5, $6)`

	result, err := sr.getter.DefaultTrOrDB(ctx, sr.db).ExecContext(ctx, query, shop.ID, shop.RetailerID, shop.Title,
		shop.Address, shop.PhoneNumber, shop.FioDirector)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("shopRepo.Create expected 1 row affected")
	}
	return nil
}

func (sr *ShopRepo) DeleteByID(ctx context.Context, ID uuid.UUID) error {
	query := `delete from ss.shop where id = $1`

	result, err := sr.getter.DefaultTrOrDB(ctx, sr.db).ExecContext(ctx, query, ID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("shopRepo.DeleteByID expected 1 row affected")
	}

	return nil
}

func (sr *ShopRepo) GetByID(ctx context.Context, ID uuid.UUID) (*models.ShopModel, error) {
	query := `select id, retailer_id, title, address, phone_number, fio_director from ss.shop where id = $1`

	var shop models.ShopModel

	err := sr.getter.DefaultTrOrDB(ctx, sr.db).GetContext(ctx, &shop, query, ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrShopDoesNotExists
	}

	return &shop, nil
}

func (sr *ShopRepo) GetByAddress(ctx context.Context, shopAddress string) (*models.ShopModel, error) {
	query := `select id, retailer_id, title, address, phone_number, fio_director from ss.shop where address = $1`

	var shop models.ShopModel

	err := sr.getter.DefaultTrOrDB(ctx, sr.db).GetContext(ctx, &shop, query, shopAddress)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrShopDoesNotExists
	}

	return &shop, nil
}

func (sr *ShopRepo) GetByParams(ctx context.Context, params *dto.ShopDTO) ([]*models.ShopModel, error) {
	query := `select id, retailer_id, title, address, phone_number, fio_director 
			  from ss.shop 
	          where ($1 = '' or title ilike '%' || $1 || '%') and 
	                ($2 = '' or address ilike '%' || $2 || '%') and 
	                ($3 = '' or phone_number ilike '%' || $3 || '%') and 
	                ($4 = '' or fio_director ilike '%' || $4 || '%')
	          limit $5 offset $6`

	var shops []*models.ShopModel

	err := sr.getter.DefaultTrOrDB(ctx, sr.db).SelectContext(ctx, &shops, query,
		params.Title,
		params.Address,
		params.PhoneNumber,
		params.FioDirector,
		params.Limit,
		params.Offset,
	)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if len(shops) == 0 {
		return nil, errs.ErrShopDoesNotExists
	}

	return shops, nil
}
