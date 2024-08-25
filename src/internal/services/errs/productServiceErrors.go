package errs

import "errors"

var (
	ErrProductAlreadyExist  = errors.New("ошибка productService! Товар уже существует")
	ErrProductDoesNotExists = errors.New("ошибка productService! Товара не существует")
)
