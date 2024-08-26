package input

import (
	"SmartShopper-services/core/dto"
	"bufio"
	"fmt"
	"os"
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

func DistributorTitle() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите название Дистрибьютора: ")

	retailerTitle, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	retailerTitle = strings.TrimSpace(retailerTitle)

	return retailerTitle, nil
}

func DistributorAddress() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите адрес Дистрибьютора: ")

	author, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	author = strings.TrimSpace(author)

	return author, nil
}

func DistributorPhoneNumber() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите номер телефона Дистрибьютора: ")

	publisher, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	publisher = strings.TrimSpace(publisher)

	return publisher, nil
}

func DistributorFioRepresentative() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите ФИО представителя: ")

	rarity, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	rarity = strings.TrimSpace(rarity)

	return rarity, nil
}

func DistributorParams() (*dto.SupplierDTO, error) {
	var (
		retailer dto.SupplierDTO
		err      error
	)

	if retailer.Title, err = DistributorTitle(); err != nil {
		return nil, err
	}
	if retailer.Address, err = DistributorAddress(); err != nil {
		return nil, err
	}
	if retailer.PhoneNumber, err = DistributorPhoneNumber(); err != nil {
		return nil, err
	}
	if retailer.FioRepresentative, err = DistributorFioRepresentative(); err != nil {
		return nil, err
	}
	return &retailer, nil
}

func ManufacturerTitle() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите название Производителя: ")

	retailerTitle, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	retailerTitle = strings.TrimSpace(retailerTitle)

	return retailerTitle, nil
}

func ManufacturerAddress() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите адрес Производителя: ")

	author, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	author = strings.TrimSpace(author)

	return author, nil
}

func ManufacturerPhoneNumber() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите номер телефона Производителя: ")

	publisher, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	publisher = strings.TrimSpace(publisher)

	return publisher, nil
}

func ManufacturerFioRepresentative() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите ФИО представителя: ")

	rarity, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	rarity = strings.TrimSpace(rarity)

	return rarity, nil
}

func ManufacturerParams() (*dto.SupplierDTO, error) {
	var (
		retailer dto.SupplierDTO
		err      error
	)

	if retailer.Title, err = ManufacturerTitle(); err != nil {
		return nil, err
	}
	if retailer.Address, err = ManufacturerAddress(); err != nil {
		return nil, err
	}
	if retailer.PhoneNumber, err = ManufacturerPhoneNumber(); err != nil {
		return nil, err
	}
	if retailer.FioRepresentative, err = ManufacturerFioRepresentative(); err != nil {
		return nil, err
	}
	return &retailer, nil
}
