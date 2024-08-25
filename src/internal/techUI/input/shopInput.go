package input

import (
	"SmartShopper-services/core/dto"
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strings"
)

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

func ShopParams() (*dto.ShopDTO, error) {
	var (
		shop dto.ShopDTO
		err  error
	)

	if shop.Title, err = ShopTitle(); err != nil {
		return nil, err
	}
	if shop.Address, err = ShopAddress(); err != nil {
		return nil, err
	}
	if shop.PhoneNumber, err = ShopPhoneNumber(); err != nil {
		return nil, err
	}
	if shop.FioDirector, err = ShopFioDirector(); err != nil {
		return nil, err
	}
	shop.RetailerID = uuid.Nil

	return &shop, nil
}
