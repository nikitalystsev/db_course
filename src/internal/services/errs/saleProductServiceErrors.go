package errs

import "errors"

var (
	ErrSaleProductAlreadyExist  = errors.New("ошибка saleProductService! Продажа уже существует")
	ErrSaleProductDoesNotExists = errors.New("ошибка SaleProductService! Продажи не существует")
)
