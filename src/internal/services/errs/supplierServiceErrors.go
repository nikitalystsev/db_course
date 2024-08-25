package errs

import "errors"

var (
	ErrRetailerAlreadyExist      = errors.New("ошибка supplierService! Ритейлер уже существует")
	ErrRetailerDoesNotExists     = errors.New("ошибка supplierService! Ритейлера не существует")
	ErrDistributorAlreadyExist   = errors.New("ошибка supplierService! Дистрибьютор уже существует")
	ErrDistributorDoesNotExists  = errors.New("ошибка supplierService! Дистрибьютора не существует")
	ErrManufacturerAlreadyExist  = errors.New("ошибка supplierService! Производитель уже существует")
	ErrManufacturerDoesNotExists = errors.New("ошибка supplierService! Производителя не существует")
)
