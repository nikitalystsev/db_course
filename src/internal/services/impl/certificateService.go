package impl

import (
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
