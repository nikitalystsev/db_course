package errs

import "errors"

var (
	ErrUserAlreadyExist             = errors.New("[!] userService error! User already exists")
	ErrEmptyUserFio                 = errors.New("[!] userService error! Empty User fio")
	ErrEmptyUserPassword            = errors.New("[!] userService error! Empty User password")
	ErrInvalidUserPasswordLen       = errors.New("[!] userService error! Invalid User password len")
	ErrEmptyUserPhoneNumber         = errors.New("[!] userService error! Empty User phoneNumber")
	ErrInvalidUserPhoneNumberLen    = errors.New("[!] userService error! Invalid User phoneNumber len")
	ErrInvalidUserPhoneNumberFormat = errors.New("[!] userService error! Invalid User phoneNumber format")
	ErrUserDoesNotExists            = errors.New("[!] userService error! User does not exist")
	ErrUserObjectIsNil              = errors.New("[!] userService error! User object is nil")
)
