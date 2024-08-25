package intfRepo

import (
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
)

type ICertificateRepo interface {
	GetByProductID(ctx context.Context, productID uuid.UUID) ([]*models.CertificateModel, error)
}
