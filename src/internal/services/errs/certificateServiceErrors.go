package errs

import "errors"

var (
	ErrCertificateAlreadyExist  = errors.New("ошибка certificateService! Сертификат уже существует")
	ErrCertificateDoesNotExists = errors.New("ошибка certificateService! Сертификата не существует")
)
