package infRepo

import (
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
	"time"
)

type IUserRepo interface {
	Create(ctx context.Context, reader *models.UserModel) error
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*models.UserModel, error)
	SaveRefreshToken(ctx context.Context, id uuid.UUID, token string, ttl time.Duration) error
}
