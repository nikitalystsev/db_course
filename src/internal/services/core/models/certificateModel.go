package models

import (
	"github.com/google/uuid"
	"time"
)

type CertificateModel struct {
	ID                uuid.UUID `json:"id" db:"id"`
	ProductID         uuid.UUID `json:"product_id" db:"product_id"`
	Type              string    `json:"type" db:"type"`
	Number            string    `json:"number" db:"number"`
	NormativeDocument string    `json:"normative_document" db:"normative_document"`
	StatusCompliance  bool      `json:"status_compliance" db:"status_compliance"`
	RegistrationDate  time.Time `json:"registration_date" db:"registration_date"`
	ExpirationDate    time.Time `json:"expiration_date" db:"expiration_date"`
}
