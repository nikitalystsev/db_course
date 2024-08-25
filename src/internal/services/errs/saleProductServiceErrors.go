package errs

import "errors"

var (
	ErrSaleProductAlreadyExist  = errors.New("[!] SaleProductService error! SaleProduct already exists")
	ErrSaleProductDoesNotExists = errors.New("[!] SaleProductService error! SaleProduct does not exist")
)
