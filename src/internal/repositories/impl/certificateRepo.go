package impl

import (
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"SmartShopper-services/intfRepo"
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CertificateRepo struct {
	db *sqlx.DB
}

func NewCertificateRepo(db *sqlx.DB) intfRepo.ICertificateRepo {
	return &CertificateRepo{db: db}
}

func (cr *CertificateRepo) GetByProductID(ctx context.Context, productID uuid.UUID) ([]*models.CertificateModel, error) {
	query := `select id, product_id, type, number, normative_document, status_compliance, registration_date, expiration_date from ss.certificate_compliance where product_id = $1`

	var certificates []*models.CertificateModel
	err := cr.db.SelectContext(ctx, &certificates, query, productID)
	if err != nil {
		return nil, err
	}
	if len(certificates) == 0 {
		return nil, errs.ErrCertificateDoesNotExists
	}

	return certificates, nil
}
