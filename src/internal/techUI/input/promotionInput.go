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

func Type() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите тип акции: ")

	title, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	title = strings.TrimSpace(title)

	return title, nil
}

func Description() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите описание акции: ")

	author, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	author = strings.TrimSpace(author)

	return author, nil
}

func DiscountSize() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите размер скидки (в процента): ")

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

func StartDate() (time.Time, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите дату начала акции (в формате YYYY-MM-DD): ")

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

func EndDate() (time.Time, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите дату конца акции (в формате YYYY-MM-DD): ")

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

func PromotionParams() (dto.PromotionDTO, error) {
	var (
		newPromotion dto.PromotionDTO
		err          error
	)

	if newPromotion.Type, err = Type(); err != nil {
		return dto.PromotionDTO{}, err
	}
	if newPromotion.Description, err = Description(); err != nil {
		return dto.PromotionDTO{}, err
	}
	if newPromotion.DiscountSize, err = DiscountSize(); err != nil {
		return dto.PromotionDTO{}, err
	}
	if newPromotion.StartDate, err = StartDate(); err != nil {
		return dto.PromotionDTO{}, err
	}
	if newPromotion.EndDate, err = EndDate(); err != nil {
		return dto.PromotionDTO{}, err
	}

	return newPromotion, nil
}
