package requesters

import (
	"SmartShopper-services/core/dto"
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

const shopMenu = `Меню обработки магазина:
	1 -- Добавить товар в магазин
	2 -- Поставить оценку товару в магазине
	3 -- Изменить цену на товар в магазине
	0 -- Вернуться к магазинам
`

const shopProductSalesKey = "shopProductSales"

func (r *Requester) processShopActions(shopID uuid.UUID, num int) error {
	var (
		menuItem int
		err      error
	)

	for {
		if err = r.getSalesByShopID(shopID, num); err != nil {
			return err
		}

		fmt.Printf("\n\n%s", shopMenu)

		if menuItem, err = input.MenuItem(); err != nil {
			fmt.Printf("\n\n%s\n", err.Error())
			continue
		}

		switch menuItem {
		case 1:
			if err = r.adNewProductInShop(shopID); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 2:
			if err = r.addRatingProductByShop(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 3:
			if err = r.changePriceOnsaleProduct(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 0:
			return nil
		default:
			fmt.Printf("\n\nНеверный пункт меню!\n")
		}

	}
}

func (r *Requester) getSalesByShopID(shopID uuid.UUID, num int) error {
	var tokens dto.UserTokensDTO
	if err := r.cache.Get(tokensKey, &tokens); err != nil {
		return err
	}

	request := HTTPRequest{
		Method: http.MethodGet,
		URL:    r.baseURL + "/api/sales",
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", tokens.AccessToken),
		},
		QueryParams: map[string]string{
			"shop_id": shopID.String(),
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
	var sales []*dto.SaleProductShopDTO
	if err = json.Unmarshal(response.Body, &sales); err != nil {
		return err
	}

	printShopSales(sales, num)
	var shopProductSales []uuid.UUID
	copyShopProductIDsToArray(&shopProductSales, sales)
	r.cache.Set(shopProductSalesKey, shopProductSales)

	return nil
}

func (r *Requester) adNewProductInShop(shopID uuid.UUID) error {
	var tokens dto.UserTokensDTO
	var err error
	if err = r.cache.Get(tokensKey, &tokens); err != nil {
		return err
	}
	var newSaleProductDTO dto.NewSaleProductDTO
	newSaleProductDTO.ShopID = shopID

	fmt.Printf("\nДля добавления нового товара необходимо ввести " +
		"данные Дистрибьютора, который его распространяет\n")

	newSaleProductDTO.Suppliers[0], err = input.DistributorParams()
	if err != nil {
		return err
	}

	fmt.Printf("\nДля добавления нового товара необходимо ввести " +
		"данные Производителя, который его производит\n")

	newSaleProductDTO.Suppliers[1], err = input.ManufacturerParams()
	if err != nil {
		return err
	}

	newSaleProductDTO.Product, err = input.ProductParams()
	if err != nil {
		return err
	}

	newSaleProductDTO.Price, err = input.PriceParams()
	if err != nil {
		return err
	}

	isWithPromotion, err := input.IsWithPromotion()
	if err != nil {
		return err
	}

	if isWithPromotion {
		newSaleProductDTO.Promotion, err = input.PromotionParams()
		if err != nil {
			return err
		}
	}

	request := HTTPRequest{
		Method: http.MethodPost,
		URL:    r.baseURL + "/api/sales",
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", tokens.AccessToken),
		},
		Body:    newSaleProductDTO,
		Timeout: 10 * time.Second,
	}
	response, err := SendRequest(request)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusCreated {
		var info string
		if err = json.Unmarshal(response.Body, &info); err != nil {
			return err
		}
		return errors.New(info)
	}

	fmt.Printf("\n\nВы успешно добавили товар в магазин!\n")

	return nil
}

func (r *Requester) addRatingProductByShop() error {
	var tokens dto.UserTokensDTO
	if err := r.cache.Get(tokensKey, &tokens); err != nil {
		return err
	}

	var shopProductSales []uuid.UUID
	if err := r.cache.Get(shopProductSalesKey, &shopProductSales); err != nil {
		return err
	}

	num, err := input.ProductPagesNumber()
	if err != nil {
		return err
	}

	if num > len(shopProductSales)-1 || num < 0 { // num -- это индекс
		return errors.New("номер товара выходит из диапазона выведенных значений")
	}

	saleProductID := shopProductSales[num]

	ratingDTO, err := input.RatingParams()
	if err != nil {
		return err
	}
	ratingDTO.SaleProductID = saleProductID

	request := HTTPRequest{
		Method: http.MethodPost,
		URL:    r.baseURL + "/api/ratings",
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", tokens.AccessToken),
		},
		Body:    ratingDTO,
		Timeout: 10 * time.Second,
	}
	response, err := SendRequest(request)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusCreated {
		var info string
		if err = json.Unmarshal(response.Body, &info); err != nil {
			return err
		}
		return errors.New(info)
	}

	fmt.Printf("\n\nОтзыв был успешно добавлен!\n")

	return nil
}

func (r *Requester) changePriceOnsaleProduct() error {
	var tokens dto.UserTokensDTO
	if err := r.cache.Get(tokensKey, &tokens); err != nil {
		return err
	}

	var shopProductSales []uuid.UUID
	if err := r.cache.Get(shopProductSalesKey, &shopProductSales); err != nil {
		return err
	}

	num, err := input.ProductPagesNumber()
	if err != nil {
		return err
	}

	if num > len(shopProductSales)-1 || num < 0 { // num -- это индекс
		return errors.New("номер товара выходит из диапазона выведенных значений")
	}

	saleProductID := shopProductSales[num]

	newPrice, err := input.SaleProductPrice()
	if err != nil {
		return err
	}

	request := HTTPRequest{
		Method: http.MethodPut,
		URL:    r.baseURL + fmt.Sprintf("/api/sales/%s", saleProductID.String()),
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", tokens.AccessToken),
		},
		Body:    newPrice,
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

	fmt.Printf("\n\nЦена была успешно изменена!\n")

	return nil
}

func printShopSales(salesDTO []*dto.SaleProductShopDTO, num int) {
	t := table.NewWriter()
	t.SetTitle(fmt.Sprintf("Товары, продающиеся в магазине №%d", num))
	t.SetStyle(table.StyleBold)
	t.Style().Format.Header = text.FormatTitle
	t.AppendHeader(
		table.Row{"No.",
			"Название товара",
			"Категория",
			"Тип акции",
			"Размер скидки",
			"Цена",
			"Валюта",
			"Средний рейтинг",
			"Соответствие сертификатам",
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
				saleProduct.ProductName,
				saleProduct.ProductCategories,
				saleProduct.PromotionType,
				discountSize,
				saleProduct.Price,
				saleProduct.Currency,
				avgRating,
				saleProduct.CertificatesStatistic,
			},
		)
	}

	fmt.Println(t.Render())
}

func copyShopProductIDsToArray(salesIDs *[]uuid.UUID, sales []*dto.SaleProductShopDTO) {
	for _, saleProduct := range sales {
		*salesIDs = append(*salesIDs, saleProduct.ID)
	}
}
