package dto

import (
	"github.com/google/uuid"
	"time"
)

type CertificateDTO struct {
	ProductID         uuid.UUID `json:"product_id"`
	Type              string    `json:"type"`
	Number            string    `json:"number"`
	NormativeDocument string    `json:"normative_document"`
	StatusCompliance  bool      `json:"status_compliance"`
	RegistrationDate  time.Time `json:"registration_date"`
	ExpirationDate    time.Time `json:"expiration_date"`
}

type CertificateStatisticsDTO struct {
	ProductID              uuid.UUID `json:"product_id" db:"product_id"`
	CountValidCertificates int       `json:"count_valid_certificates" db:"count_valid_certificates"`
	TotalCountCertificates int       `json:"total_count_certificates" db:"total_count_certificates"`
}

type CertificateStatusDTO struct {
	StatusCompliance bool `json:"status_compliance"`
}
