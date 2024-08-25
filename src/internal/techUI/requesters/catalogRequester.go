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
	2 -- Вывести информацию о товаре
	3 -- Сравнить цену на товар
	4 -- Посмотреть сертификаты соответствия на товар
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
		case 4:
			if err = r.viewCertificatesCompliance(); err != nil {
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

	var salesDTO []*dto.SaleProductDTO
	if err = json.Unmarshal(response.Body, &salesDTO); err != nil {
		return err
	}

	printSales(salesDTO, num)

	return nil
}

func (r *Requester) viewCertificatesCompliance() error {
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
		URL:    r.baseURL + "/certificates",
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

	var certificates []*models.CertificateModel
	if err = json.Unmarshal(response.Body, &certificates); err != nil {
		return err
	}

	printCertificates(certificates, num)

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

func printProduct(product *dto.ProductDTO, num int) {
	t := table.NewWriter()
	t.SetTitle(fmt.Sprintf("Товар №%d", num))
	t.SetStyle(table.StyleBold)
	t.Style().Format.Header = text.FormatTitle

	t.AppendRow(table.Row{"Ритейлер", product.Retailer})
	t.AppendRow(table.Row{"Дистрибьютор", product.Distributor})
	t.AppendRow(table.Row{"Производитель", product.Manufacturer})
	t.AppendRow(table.Row{"Название", product.Name})
	t.AppendRow(table.Row{"Категория", product.Categories})
	t.AppendRow(table.Row{"Бренд", product.Brand})
	t.AppendRow(table.Row{"Состав", product.Compound})
	t.AppendRow(table.Row{"Масса брутто", product.GrossMass})
	t.AppendRow(table.Row{"Масса нетто", product.NetMass})
	t.AppendRow(table.Row{"Тип упаковки", product.PackageType})

	fmt.Println(t.Render())
}

func printSales(salesDTO []*dto.SaleProductDTO, num int) {
	t := table.NewWriter()
	t.SetTitle(fmt.Sprintf("Места продажи товара №%d", num))
	t.SetStyle(table.StyleBold)
	t.Style().Format.Header = text.FormatTitle
	t.AppendHeader(
		table.Row{"No.",
			"Магазин",
			"Адрес",
			"Цена",
			"Валюта",
			"Тип акции",
			"Размер скидки",
			"Средний рейтинг",
		},
	)

	for i, saleProduct := range salesDTO {
		var discountSize string
		if saleProduct.PromotionDiscountSize == nil {
			discountSize = fmt.Sprintf("%d", 0)
		} else {
			discountSize = fmt.Sprintf("%d", *saleProduct.PromotionDiscountSize)
		}
		var avgRating string
		if saleProduct.AvgRating == nil {
			avgRating = "Не оценен"
		} else {
			avgRating = fmt.Sprintf("%f", *saleProduct.AvgRating)
		}
		t.AppendRow(
			table.Row{
				i,
				saleProduct.ShopTitle,
				saleProduct.ShopAddress,
				saleProduct.Price,
				saleProduct.Currency,
				saleProduct.PromotionType,
				discountSize,
				avgRating,
			},
		)
	}

	fmt.Println(t.Render())
}

func printCertificates(certificates []*models.CertificateModel, num int) {
	t := table.NewWriter()
	t.SetTitle(fmt.Sprintf("Сертификаты соответствия товара №%d", num))
	t.SetStyle(table.StyleBold)
	t.Style().Format.Header = text.FormatTitle
	t.AppendHeader(
		table.Row{"No.",
			"Тип сертификата",
			"Номер сертификата",
			"Статус соответствия",
			"Дата регистрации",
			"Дата окончания действия",
		},
	)

	for i, certificate := range certificates {
		var statusCompliance = "Действующий"
		if !certificate.StatusCompliance {
			statusCompliance = "Недействующий"
		}
		t.AppendRow(
			table.Row{
				i,
				certificate.Type,
				certificate.Number,
				statusCompliance,
				certificate.RegistrationDate,
				certificate.ExpirationDate,
			},
		)
	}

	fmt.Println(t.Render())
}
