package input

import (
	"SmartShopper-services/core/dto"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//func IsWithParams() (bool, error) {
//	reader := bufio.NewReader(os.Stdin)
//
//	fmt.Printf("Would you like to enter search parameters?(Y/N): ")
//
//	isWithParams, err := reader.ReadString('\n')
//	if err != nil {
//		return false, err
//	}
//
//	isWithParams = strings.TrimSpace(isWithParams)
//	if isWithParams != "y" && isWithParams != "Y" {
//		return true, nil
//	}
//
//	return false, nil
//}

func RetailerTitle() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите название Ритейлера: ")

	retailerTitle, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	retailerTitle = strings.TrimSpace(retailerTitle)

	return retailerTitle, nil
}

func RetailerAddress() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите адрес Ритейлера: ")

	author, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	author = strings.TrimSpace(author)

	return author, nil
}

func RetailerPhoneNumber() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите номер телефона Ритейлера: ")

	publisher, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	publisher = strings.TrimSpace(publisher)

	return publisher, nil
}

func RetailerFioRepresentative() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите ФИО представителя: ")

	rarity, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	rarity = strings.TrimSpace(rarity)

	return rarity, nil
}

func RetailerParams() (*dto.SupplierDTO, error) {
	var (
		retailer dto.SupplierDTO
		err      error
	)

	if retailer.Title, err = RetailerTitle(); err != nil {
		return nil, err
	}
	if retailer.Address, err = RetailerAddress(); err != nil {
		return nil, err
	}
	if retailer.PhoneNumber, err = RetailerPhoneNumber(); err != nil {
		return nil, err
	}
	if retailer.FioRepresentative, err = RetailerFioRepresentative(); err != nil {
		return nil, err
	}
	return &retailer, nil
}

func Genre() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Input genre: ")

	genre, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	genre = strings.TrimSpace(genre)

	return genre, nil
}

func PublishingYear() (uint, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Input publishing year: ")

	yearStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	yearStr = strings.TrimSpace(yearStr)

	yearInt, err := strconv.Atoi(yearStr)
	if err != nil {
		return 0, err
	}

	year := uint(yearInt)

	return year, nil
}

func Language() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Input language: ")

	language, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	language = strings.TrimSpace(language)

	return language, nil
}

func AgeLimit() (uint, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Input age limit: ")

	ageLimitStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	ageLimitStr = strings.TrimSpace(ageLimitStr)

	ageLimitInt, err := strconv.Atoi(ageLimitStr)
	if err != nil {
		return 0, err
	}

	ageLimit := uint(ageLimitInt)

	return ageLimit, nil
}

func CopiesNumber() (uint, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Input book's copies number: ")

	copiesNumStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	copiesNumStr = strings.TrimSpace(copiesNumStr)

	copiesNumInt, err := strconv.Atoi(copiesNumStr)
	if err != nil {
		return 0, err
	}

	copiesNum := uint(copiesNumInt)

	return copiesNum, nil
}
