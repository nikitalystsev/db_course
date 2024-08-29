package impl

import (
	"SmartShopper-services/core/dto"
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"SmartShopper-services/intf"
	"SmartShopper-services/intfRepo"
	"context"
	"errors"
	"github.com/google/uuid"
)

type CertificateService struct {
	certificateRepo intfRepo.ICertificateRepo
}

func NewCertificateService(certificateRepo intfRepo.ICertificateRepo) intf.ICertificateService {
	return &CertificateService{certificateRepo: certificateRepo}
}

func (cs *CertificateService) Create(ctx context.Context, certificate *models.CertificateModel) error {
	err := cs.certificateRepo.Create(ctx, certificate)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CertificateService) GetByProductID(ctx context.Context, productID uuid.UUID) ([]*models.CertificateModel, error) {
	certificates, err := cs.certificateRepo.GetByProductID(ctx, productID)
	if err != nil && errors.Is(err, errs.ErrCertificateDoesNotExists) {
		return nil, err
	}

	if certificates == nil {
		return nil, errs.ErrCertificateDoesNotExists
	}

	return certificates, nil
}

func (cs *CertificateService) DeleteByID(ctx context.Context, ID uuid.UUID) error {
	err := cs.certificateRepo.DeleteByID(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CertificateService) GetCertificateStatisticsByProductID(ctx context.Context, productID uuid.UUID) (*dto.CertificateStatisticsDTO, error) {
	certificates, err := cs.certificateRepo.GetByProductID(ctx, productID)
	if err != nil && !errors.Is(err, errs.ErrCertificateDoesNotExists) {
		return nil, err
	}
	if certificates == nil {
		return nil, errs.ErrCertificateDoesNotExists
	}

	certificateStatisticsDTO := dto.CertificateStatisticsDTO{
		ProductID:              productID,
		TotalCountCertificates: len(certificates),
	}

	for _, certificate := range certificates {
		if certificate.StatusCompliance {
			certificateStatisticsDTO.CountValidCertificates += 1
		}
	}

	return &certificateStatisticsDTO, nil
}
