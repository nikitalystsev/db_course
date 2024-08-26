package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IsWithPromotion() (bool, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Применена ли к товару какая-либо акция? (Y/N): ")

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
