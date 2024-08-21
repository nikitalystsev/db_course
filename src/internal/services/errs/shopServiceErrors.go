package errs

import "errors"

var (
	ErrShopAlreadyExist  = errors.New("[!] shopService error! Shop already exists")
	ErrShopDoesNotExists = errors.New("[!] shopService error! Shop does not exist")
)
