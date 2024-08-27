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
	"time"
)
import "github.com/go-redis/redis/v8"

type UserRepo struct {
	db     *sqlx.DB
	client *redis.Client
}

func NewUserRepo(db *sqlx.DB, client *redis.Client) intfRepo.IUserRepo {
	return &UserRepo{db: db, client: client}
}

func (ur *UserRepo) Create(ctx context.Context, user *models.UserModel) error {
	query := `insert into ss.user values ($1, $2, $3, $4, $5)`

	result, err := ur.db.ExecContext(ctx, query, user.ID, user.Fio, user.PhoneNumber, user.Password, user.RegistrationDate)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("userRepo.Create expected 1 row affected")
	}

	return nil
}

func (ur *UserRepo) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*models.UserModel, error) {
	query := `select id, fio, phone_number, password, registration_date from ss.user where phone_number = $1`

	var user models.UserModel
	err := ur.db.GetContext(ctx, &user, query, phoneNumber)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrUserDoesNotExists
	}

	return &user, nil
}

func (ur *UserRepo) SaveRefreshToken(ctx context.Context, id uuid.UUID, token string, ttl time.Duration) error {
	if err := ur.client.Set(ctx, token, id.String(), ttl).Err(); err != nil {
		return err
	}

	return nil
}

func (ur *UserRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (*models.UserModel, error) {
	IDStr, err := ur.client.Get(ctx, refreshToken).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}
	if errors.Is(err, redis.Nil) {
		return nil, errs.ErrUserDoesNotExists
	}

	var ID uuid.UUID
	if ID, err = uuid.Parse(IDStr); err != nil {
		return nil, err
	}

	query := `select id, fio, phone_number, password, registration_date from ss.user where id = $1`

	var user models.UserModel
	err = ur.db.GetContext(ctx, &user, query, ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrUserDoesNotExists
	}

	return &user, nil
}
