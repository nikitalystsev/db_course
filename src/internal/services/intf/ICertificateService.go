package intf

import (
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
)

type ICertificateService interface {
	GetByProductID(ctx context.Context, productID uuid.UUID) ([]*models.CertificateModel, error)
}
