package requesters

import (
	"SmartShopper-services/core/dto"
	"SmartShopper-services/core/models"
	"SmartShopper/internal/techUI/input"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"net/http"
	"time"
)

const catalogMenu = `Меню каталога:
	1 -- Вывести товары
	2 -- Перейти к товару
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
			if err = r.processProduct(); err != nil {
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

func (r *Requester) processProduct() error {
	var productPagesID []uuid.UUID
	if err := r.cache.Get(productsKey, &productPagesID); err != nil {
		return err
	}

	num, err := input.ProductPagesNumber()
	if err != nil {
		return err
	}

	if num > len(productPagesID)-1 || num < 0 { // num -- это индекс
		return errors.New("номер товара выходит из диапазона выведенных значений")
	}

	productID := productPagesID[num]

	err = r.processProductActions(productID, num)
	if err != nil {
		return err
	}

	return nil
}

func printProducts(products []*models.ProductModel, offset int) {
	t := table.NewWriter()
	t.SetTitle(fmt.Sprintf("Страница товаров №%d", offset/pageLimit+1))
	t.SetStyle(table.StyleBold)
	t.Style().Format.Header = text.FormatTitle
	t.AppendHeader(table.Row{"No.", "Название товара", "Категория"})

	for i, product := range products {
		t.AppendRow(table.Row{offset + i, product.Name, product.Categories})
	}
	fmt.Println(t.Render())
}

func copyProductIDsToArray(productIDs *[]uuid.UUID, products []*models.ProductModel) {
	for _, product := range products {
		*productIDs = append(*productIDs, product.ID)
	}
}
