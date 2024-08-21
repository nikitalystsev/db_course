package errs

import "errors"

var (
	ErrRetailerAlreadyExist      = errors.New("[!] supplierService error! Retailer already exists")
	ErrRetailerDoesNotExists     = errors.New("[!] supplierService error! Retailer does not exist")
	ErrDistributorAlreadyExist   = errors.New("[!] supplierService error! Distributor already exists")
	ErrDistributorDoesNotExists  = errors.New("[!] supplierService error! Distributor does not exist")
	ErrManufacturerAlreadyExist  = errors.New("[!] supplierService error! Manufacturer already exists")
	ErrManufacturerDoesNotExists = errors.New("[!] supplierService error! Manufacturer does not exist")
)
