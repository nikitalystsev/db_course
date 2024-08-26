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

func Currency() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите валюту цены: ")

	publisher, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	publisher = strings.TrimSpace(publisher)

	return publisher, nil
}

func Price() (float32, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите цену товара: ")

	numStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	numStr = strings.TrimSpace(numStr)

	numFloat, err := strconv.ParseFloat(numStr, 32)
	if err != nil {
		return 0, err
	}

	return float32(numFloat), nil
}

func SettingDate() (time.Time, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите дату установки цены (в формате YYYY-MM-DD): ")

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

func PriceParams() (dto.PriceDTO, error) {
	var (
		priceDTO dto.PriceDTO
		err      error
	)

	if priceDTO.Currency, err = Currency(); err != nil {
		return dto.PriceDTO{}, err
	}
	if priceDTO.Price, err = Price(); err != nil {
		return dto.PriceDTO{}, err
	}
	if priceDTO.SettingDate, err = SettingDate(); err != nil {
		return dto.PriceDTO{}, err
	}
	return priceDTO, nil
}

func SaleProductPrice() (float32, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите новую цену: ")

	numStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	numStr = strings.TrimSpace(numStr)

	numFloat, err := strconv.ParseFloat(numStr, 32)
	if err != nil {
		return 0, err
	}

	return float32(numFloat), nil
}
