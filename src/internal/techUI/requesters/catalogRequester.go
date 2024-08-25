package requesters

import (
	"SmartShopper-services/core/dto"
	"SmartShopper-services/core/models"
	"SmartShopper/internal/techUI/input"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

const catalogMenu = `Меню каталога:
	1 -- Вывести товары
	2 -- Вывести информацию о товаре
	3 -- Сравнить цену на товар
	0 -- Вернуться в главное меню
`
const (
	pageLimit = 10

	productsKey      = "products"
	productParamsKey = "productParams"
)

func (r *Requester) processCatalogActions() error {
	var (
		menuItem int
		err      error
	)
	r.cache.Set(productParamsKey, dto.ProductParamsDTO{Limit: pageLimit, Offset: 0})
	r.cache.Set(productsKey, make([]uuid.UUID, 0))

	for {
		fmt.Printf("\n\n%s", catalogMenu)

		if menuItem, err = input.MenuItem(); err != nil {
			fmt.Printf("\n\n%s\n", err.Error())
			continue
		}

		switch menuItem {
		case 1:
			if err = r.viewPage(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 2:
			if err = r.viewProduct(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 3:
			if err = r.comparePriceOnProduct(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 0:
			r.cache.Delete(productParamsKey)
			r.cache.Delete(productsKey)
			return nil
		default:
			fmt.Printf("\n\nНеверный пункт меню!\n")
		}
	}
}

func (r *Requester) viewPage() error {
	var productParams dto.ProductParamsDTO
	if err := r.cache.Get(productParamsKey, &productParams); err != nil {
		return err
	}

	var productPagesID []uuid.UUID
	if err := r.cache.Get(productsKey, &productPagesID); err != nil {
		return err
	}

	request := HTTPRequest{
		Method: http.MethodGet,
		URL:    r.baseURL + "/products",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		QueryParams: map[string]string{
			"limit":  fmt.Sprintf("%d", productParams.Limit),
			"offset": fmt.Sprintf("%d", productParams.Offset),
		},
		Timeout: 10 * time.Second,
	}

	response, err := SendRequest(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		var info string
		if err = json.Unmarshal(response.Body, &info); err != nil {
			return err
		}
		return errors.New(info)
	}

	var products []*models.ProductModel
	if err = json.Unmarshal(response.Body, &products); err != nil {
		return err
	}

	printProducts(products, productParams.Offset)
	copyProductIDsToArray(&productPagesID, products)
	r.cache.Set(productsKey, productPagesID)
	r.cache.Set(
		productParamsKey,
		dto.ProductParamsDTO{
			Limit:  pageLimit,
			Offset: productParams.Offset + pageLimit,
		},
	)

	return nil
}

func (r *Requester) viewProduct() error {
	var products []uuid.UUID
	if err := r.cache.Get(productsKey, &products); err != nil {
		return err
	}

	num, err := input.ProductPagesNumber()
	if err != nil {
		return err
	}

	if num > len(products)-1 || num < 0 { // num -- это индекс
		return errors.New("номер товара выходит из диапазона выведенных значений")
	}

	productID := products[num]

	request := HTTPRequest{
		Method: http.MethodGet,
		URL:    r.baseURL + fmt.Sprintf("/products/%s", productID.String()),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Timeout: 10 * time.Second,
	}

	response, err := SendRequest(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		var info string
		if err = json.Unmarshal(response.Body, &info); err != nil {
			return err
		}
		return errors.New(info)
	}

	var product *dto.ProductDTO
	if err = json.Unmarshal(response.Body, &product); err != nil {
		return err
	}

	printProduct(product, num)

	return nil
}

func (r *Requester) comparePriceOnProduct() error {
	var products []uuid.UUID
	if err := r.cache.Get(productsKey, &products); err != nil {
		return err
	}

	num, err := input.ProductPagesNumber()
	if err != nil {
		return err
	}

	if num > len(products)-1 || num < 0 { // num -- это индекс
		return errors.New("номер товара выходит из диапазона выведенных значений")
	}

	productID := products[num]

	request := HTTPRequest{
		Method: http.MethodGet,
		URL:    r.baseURL + "/sales",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		QueryParams: map[string]string{
			"product_id": productID.String(),
		},
		Timeout: 10 * time.Second,
	}

	response, err := SendRequest(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		var info string
		if err = json.Unmarshal(response.Body, &info); err != nil {
			return err
		}
		return errors.New(info)
	}

	var sales []*models.SaleProductModel
	if err = json.Unmarshal(response.Body, &sales); err != nil {
		return err
	}

	printSales(sales)

	return nil
}

func printProducts(products []*models.ProductModel, offset int) {
	titleWidth := 60
	authorWidth := 60

	fmt.Printf("\n\n%-5s %-60s %-60s\n", "No.", "Название товара", "Категория")
	fmt.Println(strings.Repeat("-", 5+1+titleWidth+1+authorWidth))

	for i, product := range products {
		fmt.Printf("%-5d %-60s %-60s\n", offset+i, truncate(product.Name, titleWidth), truncate(product.Categories, authorWidth))
	}
}

func truncate(s string, maxLength int) string {
	if len(s) > maxLength {
		return s[:maxLength-3] + "..."
	}
	return s
}

func copyProductIDsToArray(productIDs *[]uuid.UUID, products []*models.ProductModel) {
	for _, product := range products {
		*productIDs = append(*productIDs, product.ID)
	}
}

// Функция для вывода информации о продукте
func printProduct(product *dto.ProductDTO, num int) {
	fmt.Printf("\n\nProduct №%d:\n", num)
	fmt.Println(strings.Repeat("-", 150))
	fmt.Printf("Retailer:       %s\n", product.Retailer)
	fmt.Printf("Distributor:    %s\n", product.Distributor)
	fmt.Printf("Manufacturer:   %s\n", product.Manufacturer)
	fmt.Printf("Name:           %s\n", product.Name)
	fmt.Printf("Categories:     %s\n", product.Categories)
	fmt.Printf("Brand:          %s\n", product.Brand)
	fmt.Printf("Compound:       %s\n", product.Compound)
	fmt.Printf("Gross Mass:     %.2f\n", product.GrossMass)
	fmt.Printf("Net Mass:       %.2f\n", product.NetMass)
	fmt.Printf("Package Type:   %s\n", product.PackageType)
}

func printSales(sales []*models.SaleProductModel) {
	titleWidth := 60
	authorWidth := 60

	fmt.Printf("\n\n%-5s %-60s %-60s\n", "No.", "Цена товара", "Валюта")
	fmt.Println(strings.Repeat("-", 5+1+titleWidth+1+authorWidth))

	for i, saleProduct := range sales {
		fmt.Printf("%-5d %-60s %-60s\n", i, truncate(fmt.Sprintf("%f", saleProduct.Price), titleWidth), truncate(saleProduct.Currency, authorWidth))
	}
}
