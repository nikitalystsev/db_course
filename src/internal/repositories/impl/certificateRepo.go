package impl

import (
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"SmartShopper-services/intfRepo"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CertificateRepo struct {
	db *sqlx.DB
}

func NewCertificateRepo(db *sqlx.DB) intfRepo.ICertificateRepo {
	return &CertificateRepo{db: db}
}

func (cr *CertificateRepo) Create(ctx context.Context, certificate *models.CertificateModel) error {
	query := `insert into ss.certificate_compliance values ($1, $2, $3, $4, $5, $6, $7, $8)`
	result, err := cr.db.ExecContext(ctx, query, certificate.ID, certificate.ProductID, certificate.Type,
		certificate.Number, certificate.NormativeDocument, certificate.StatusCompliance,
		certificate.RegistrationDate, certificate.ExpirationDate)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("certificateRepo.Create expected 1 row affected")
	}

	return nil
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

func (cr *CertificateRepo) DeleteByID(ctx context.Context, ID uuid.UUID) error {
	query := `delete from ss.certificate_compliance where id = $1`

	result, err := cr.db.ExecContext(ctx, query, ID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("CertificateRepo.DeleteByID expected 1 row affected")
	}

	return nil
}
