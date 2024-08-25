package errs

import "errors"

var (
	ErrRatingAlreadyExist  = errors.New("ошибка ratingService! Оценка уже существует")
	ErrRatingDoesNotExists = errors.New("ошибка ratingService! Оценки не существует")
)
