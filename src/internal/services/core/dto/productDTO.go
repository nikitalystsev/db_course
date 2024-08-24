package dto

type ProductParamsDTO struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ProductDTO struct {
	Retailer     string  `json:"retailer"`
	Distributor  string  `json:"distributor"`
	Manufacturer string  `json:"manufacturer"`
	Name         string  `json:"name"`
	Categories   string  `json:"categories"`
	Brand        string  `json:"brand"`
	Compound     string  `json:"compound"`
	GrossMass    float32 `json:"gross_mass"`
	NetMass      float32 `json:"net_mass"`
	PackageType  string  `json:"package_type"`
}
