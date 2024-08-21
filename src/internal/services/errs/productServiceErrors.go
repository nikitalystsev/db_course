package errs

import "errors"

var (
	ErrProductAlreadyExist  = errors.New("[!] productService error! Product already exists")
	ErrProductDoesNotExists = errors.New("[!] productService error! Product does not exist")
)
