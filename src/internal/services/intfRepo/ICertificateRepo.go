package intfRepo

import (
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
)

type ICertificateRepo interface {
	GetByProductID(ctx context.Context, productID uuid.UUID) ([]*models.CertificateModel, error)
	Create(ctx context.Context, certificate *models.CertificateModel) error
	DeleteByID(ctx context.Context, ID uuid.UUID) error
	GetByID(ctx context.Context, ID uuid.UUID) (*models.CertificateModel, error)
	Update(ctx context.Context, certificate *models.CertificateModel) error
}
