package input

import (
	"SmartShopper-services/core/dto"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsWithParams() (bool, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Хотите ли вы ввести параметры поиска?(Y/N): ")

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

func ShopTitle() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите название магазина: ")

	retailerTitle, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	retailerTitle = strings.TrimSpace(retailerTitle)

	return retailerTitle, nil
}

func ShopAddress() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите адрес магазина: ")

	author, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	author = strings.TrimSpace(author)

	return author, nil
}

func ShopPhoneNumber() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите номер телефона магазина: ")

	publisher, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	publisher = strings.TrimSpace(publisher)

	return publisher, nil
}

func ShopFioDirector() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите ФИО директора магазина: ")

	rarity, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	rarity = strings.TrimSpace(rarity)

	return rarity, nil
}

func ShopParams() (dto.ShopParamsDTO, error) {
	var (
		shop dto.ShopParamsDTO
		err  error
	)

	if shop.Title, err = ShopTitle(); err != nil {
		return dto.ShopParamsDTO{}, err
	}
	if shop.Address, err = ShopAddress(); err != nil {
		return dto.ShopParamsDTO{}, err
	}
	if shop.PhoneNumber, err = ShopPhoneNumber(); err != nil {
		return dto.ShopParamsDTO{}, err
	}
	if shop.FioDirector, err = ShopFioDirector(); err != nil {
		return dto.ShopParamsDTO{}, err
	}

	return shop, nil
}

func ShopPagesNumber() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите номер магазина из списка: ")

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
