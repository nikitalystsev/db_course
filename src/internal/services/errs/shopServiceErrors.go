package errs

import "errors"

var (
	ErrShopAlreadyExist  = errors.New("ошибка shopService! Магазин уже существует")
	ErrShopDoesNotExists = errors.New("ошибка shopService! Магазина не существует")
)
