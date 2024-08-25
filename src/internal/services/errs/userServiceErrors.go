package errs

import "errors"

var (
	ErrUserAlreadyExist             = errors.New("ошибка userService! Пользователь уже существует")
	ErrEmptyUserFio                 = errors.New("ошибка userService! Пустое ФИО")
	ErrEmptyUserPassword            = errors.New("ошибка userService! Пустой пароль")
	ErrInvalidUserPasswordLen       = errors.New("ошибка userService! Некорректная длина пароля")
	ErrEmptyUserPhoneNumber         = errors.New("ошибка userService! Пустой номер телефона")
	ErrInvalidUserPhoneNumberLen    = errors.New("ошибка userService! Некорректная длина номера телефона")
	ErrInvalidUserPhoneNumberFormat = errors.New("ошибка userService! Некорректный формат номера телефона")
	ErrUserDoesNotExists            = errors.New("ошибка userService! Пользователя не существует")
	ErrUserObjectIsNil              = errors.New("ошибка userService! Объект пользователя nil")
)
