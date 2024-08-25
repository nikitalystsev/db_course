package impl

import (
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"SmartShopper-services/intfRepo"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ShopRepo struct {
	db *sqlx.DB
}

func NewShopRepo(db *sqlx.DB) intfRepo.IShopRepo {
	return &ShopRepo{db: db}
}

func (sr *ShopRepo) Create(ctx context.Context, shop *models.ShopModel) error {
	fmt.Println("call shopRepo.Create")

	query := `insert into ss.shop values ($1, $2, $3, $4, $5, $6)`

	result, err := sr.db.ExecContext(ctx, query, shop.ID, shop.RetailerID, shop.Title,
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

	fmt.Println("Все заебись. создали Магазинчик")
	return nil
}

func (sr *ShopRepo) DeleteByID(ctx context.Context, ID uuid.UUID) error {
	query := `delete from ss.shop where id = $1`

	result, err := sr.db.ExecContext(ctx, query, ID)
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

	err := sr.db.GetContext(ctx, &shop, query, ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrShopDoesNotExists
	}

	return &shop, nil
}

func (sr *ShopRepo) GetByAddress(ctx context.Context, shopAddress string) (*models.ShopModel, error) {
	fmt.Println("call shopRepo.GetByAddress")

	query := `select id, retailer_id, title, address, phone_number, fio_director from ss.shop where address = $1`

	var shop models.ShopModel

	err := sr.db.GetContext(ctx, &shop, query, shopAddress)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		fmt.Println("ебаная ошибка получения по адресу")
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("блять нет магазина. Нам же лучше")
		return nil, errs.ErrShopDoesNotExists
	}

	fmt.Println("все тип топ. получили по адресу")

	return &shop, nil
}
