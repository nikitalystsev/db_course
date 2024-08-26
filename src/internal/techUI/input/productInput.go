package input

import (
	"SmartShopper-services/core/dto"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Name() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите название товара: ")

	title, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	title = strings.TrimSpace(title)

	return title, nil
}

func Categories() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите категории товара: ")

	author, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	author = strings.TrimSpace(author)

	return author, nil
}

func Brand() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите бренд товара: ")

	author, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	author = strings.TrimSpace(author)

	return author, nil
}

func Compound() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите состав: ")

	publisher, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	publisher = strings.TrimSpace(publisher)

	return publisher, nil
}

func GrossMass() (float32, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите массу брутто: ")

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

func NetMass() (float32, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите массу нетто: ")

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

func PackageType() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите тип упаковки: ")

	publisher, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	publisher = strings.TrimSpace(publisher)

	return publisher, nil
}

func ProductParams() (dto.NewProductDTO, error) {
	var (
		newProduct dto.NewProductDTO
		err        error
	)

	if newProduct.Name, err = Name(); err != nil {
		return dto.NewProductDTO{}, err
	}
	if newProduct.Categories, err = Categories(); err != nil {
		return dto.NewProductDTO{}, err
	}
	if newProduct.Brand, err = Brand(); err != nil {
		return dto.NewProductDTO{}, err
	}
	if newProduct.Compound, err = Compound(); err != nil {
		return dto.NewProductDTO{}, err
	}
	if newProduct.GrossMass, err = GrossMass(); err != nil {
		return dto.NewProductDTO{}, err
	}
	if newProduct.NetMass, err = NetMass(); err != nil {
		return dto.NewProductDTO{}, err
	}
	if newProduct.PackageType, err = PackageType(); err != nil {
		return dto.NewProductDTO{}, err
	}

	return newProduct, nil
}

func ProductPagesNumber() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите номер товара из каталога: ")

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
