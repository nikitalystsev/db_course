package dto

import "github.com/google/uuid"

type ProductParamsDTO struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ProductDTO struct {
	Retailer              string  `json:"retailer"`
	Distributor           string  `json:"distributor"`
	Manufacturer          string  `json:"manufacturer"`
	Name                  string  `json:"name"`
	Categories            string  `json:"categories"`
	Brand                 string  `json:"brand"`
	Compound              string  `json:"compound"`
	GrossMass             float32 `json:"gross_mass"`
	NetMass               float32 `json:"net_mass"`
	PackageType           string  `json:"package_type"`
	CertificatesStatistic string  `json:"certificates_statistic"`
}

type ProductCertificateDTO struct {
	ID                    uuid.UUID `json:"id" db:"id"`
	RetailerID            uuid.UUID `json:"retailer_id" db:"retailer_id"`
	DistributorID         uuid.UUID `json:"distributor_id" db:"distributor_id"`
	ManufacturerID        uuid.UUID `json:"manufacturer_id" db:"manufacturer_id"`
	Name                  string    `json:"name" db:"name"`
	Categories            string    `json:"categories" db:"categories"`
	Brand                 string    `json:"brand" db:"brand"`
	Compound              string    `json:"compound" db:"compound"`
	GrossMass             float32   `json:"gross_mass" db:"gross_mass"`
	NetMass               float32   `json:"net_mass" db:"net_mass"`
	PackageType           string    `json:"package_type" db:"package_type"`
	CertificatesStatistic string    `json:"certificates_statistic"`
}
