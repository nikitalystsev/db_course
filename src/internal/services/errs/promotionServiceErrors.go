package errs

import "errors"

var (
	ErrPromotionAlreadyExist  = errors.New("ошибка promotionService! Акция уже существует")
	ErrPromotionDoesNotExists = errors.New("ошибка promotionService! Акции не существует")
)
