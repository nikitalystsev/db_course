package errs

import "errors"

var (
	ErrRatingAlreadyExist  = errors.New("[!] RatingService error! Rating already exists")
	ErrRatingDoesNotExists = errors.New("[!] RatingService error! Rating does not exist")
)
