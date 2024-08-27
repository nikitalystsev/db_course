package input

import (
	"SmartShopper-services/core/dto"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func CertificateType() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите тип сертификата: ")

	title, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	title = strings.TrimSpace(title)

	return title, nil
}

func CertificateNumber() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите номер сертификата: ")

	title, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	title = strings.TrimSpace(title)

	return title, nil
}

func NormativeDocument() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите нормативный документ: ")

	title, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	title = strings.TrimSpace(title)

	return title, nil
}

func IsCompliance() (bool, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Действующий ли сертификат?(Y/N): ")

	isWithParams, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	isWithParams = strings.TrimSpace(isWithParams)
	if isWithParams != "n" && isWithParams != "N" {
		return true, nil
	}

	return false, nil
}

func RegistrationDate() (time.Time, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите дату регистрации сертификата (в формате YYYY-MM-DD): ")

	dateStr, err := reader.ReadString('\n')
	if err != nil {
		return time.Time{}, err
	}

	dateStr = strings.TrimSpace(dateStr)

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

func ExpirationDate() (time.Time, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите дату регистрации окончания действия сертификата (в формате YYYY-MM-DD): ")

	dateStr, err := reader.ReadString('\n')
	if err != nil {
		return time.Time{}, err
	}

	dateStr = strings.TrimSpace(dateStr)

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

func CertificateParams() (dto.CertificateDTO, error) {
	var (
		certificateDTO dto.CertificateDTO
		err            error
	)

	if certificateDTO.Type, err = CertificateType(); err != nil {
		return dto.CertificateDTO{}, err
	}
	if certificateDTO.Number, err = CertificateNumber(); err != nil {
		return dto.CertificateDTO{}, err
	}
	if certificateDTO.NormativeDocument, err = NormativeDocument(); err != nil {
		return dto.CertificateDTO{}, err
	}
	if certificateDTO.StatusCompliance, err = IsCompliance(); err != nil {
		return dto.CertificateDTO{}, err
	}
	if certificateDTO.RegistrationDate, err = RegistrationDate(); err != nil {
		return dto.CertificateDTO{}, err
	}
	if certificateDTO.ExpirationDate, err = ExpirationDate(); err != nil {
		return dto.CertificateDTO{}, err
	}

	return certificateDTO, nil
}

func CertificateCatalogNumber() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите номер сертификата из каталога: ")

	numStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	numStr = strings.TrimSpace(numStr)

	numInt, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, err
	}

	return numInt, nil
}
