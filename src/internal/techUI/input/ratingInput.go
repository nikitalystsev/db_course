package input

import (
	"SmartShopper-services/core/dto"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Review() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите отзыв: ")

	review, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	review = strings.TrimSpace(review)

	return review, nil
}

func Rating() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите оценку: ")

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

func RatingParams() (dto.RatingDTO, error) {
	var (
		rating dto.RatingDTO
		err    error
	)

	if rating.Review, err = Review(); err != nil {
		return dto.RatingDTO{}, err
	}
	if rating.Rating, err = Rating(); err != nil {
		return dto.RatingDTO{}, err
	}

	return rating, err
}
