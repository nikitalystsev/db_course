package intf

import (
	"SmartShopper-services/core/dto"
	"SmartShopper-services/core/models"
	"context"
	"github.com/google/uuid"
)

type ICertificateService interface {
	Create(ctx context.Context, certificate *models.CertificateModel) error
	GetByProductID(ctx context.Context, productID uuid.UUID) ([]*models.CertificateModel, error)
	GetByID(ctx context.Context, ID uuid.UUID) (*models.CertificateModel, error)
	Update(ctx context.Context, certificate *models.CertificateModel) error
	DeleteByID(ctx context.Context, ID uuid.UUID) error
	GetCertificateStatisticsByProductID(ctx context.Context, productID uuid.UUID) (*dto.CertificateStatisticsDTO, error)
}
