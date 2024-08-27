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

const shopMenu = `Меню обработки магазинов:
	1 -- Вывести магазины
	2 -- Следующая страница
	3 -- Добавить товар в магазин
	4 -- Поставить оценку товару в магазине
	5 -- Изменить цену на товар в магазине
	0 -- Вернуться в главное меню
`

const (
	shopsKey            = "shops"
	shopParamsKey       = "shopParams"
	shopProductSalesKey = "shopProductSales"
)

func (r *Requester) processShopActions() error {
	var (
		menuItem int
		err      error
	)
	r.cache.Set(shopParamsKey, dto.ShopParamsDTO{Limit: pageLimit, Offset: 0})
	r.cache.Set(shopsKey, make([]uuid.UUID, 0))

	for {
		fmt.Printf("\n\n%s", shopMenu)

		if menuItem, err = input.MenuItem(); err != nil {
			fmt.Printf("\n\n%s\n", err.Error())
			continue
		}

		switch menuItem {
		case 1:
			if err = r.viewFirstPageShops(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 2:
			if err = r.viewNextPage(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 3:
			if err = r.adNewProductInShop(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 4:
			if err = r.addRatingProductByShop(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 5:
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

func (r *Requester) viewFirstPageShops() error {
	var shopParams dto.ShopParamsDTO
	var shopPagesID []uuid.UUID

	var tokens dto.UserTokensDTO
	if err := r.cache.Get(tokensKey, &tokens); err != nil {
		return err
	}

	isWithParams, err := input.IsWithParams()
	if err != nil {
		return err
	}

	if isWithParams {
		if shopParams, err = input.ShopParams(); err != nil {
			return err
		}
	}
	shopParams.Limit = pageLimit
	shopParams.Offset = 0

	request := HTTPRequest{
		Method: http.MethodGet,
		URL:    r.baseURL + "/api/shops",
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", tokens.AccessToken),
		},
		QueryParams: map[string]string{
			"title":        shopParams.Title,
			"address":      shopParams.Address,
			"phone_number": shopParams.PhoneNumber,
			"fio_director": shopParams.FioDirector,
			"limit":        fmt.Sprintf("%d", shopParams.Limit),
			"offset":       fmt.Sprintf("%d", shopParams.Offset),
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

	var shops []*models.ShopModel
	if err = json.Unmarshal(response.Body, &shops); err != nil {
		return err
	}

	printShops(shops, shopParams.Offset)
	copyShopIDsToArray(&shopPagesID, shops)
	r.cache.Set(shopsKey, shopPagesID)
	r.cache.Set(
		shopParamsKey,
		dto.ShopParamsDTO{
			Limit:  pageLimit,
			Offset: shopParams.Offset + pageLimit,
		},
	)

	return nil
}

func (r *Requester) viewNextPage() error {
	var shopParams dto.ShopParamsDTO
	if err := r.cache.Get(shopParamsKey, &shopParams); err != nil {
		return err
	}

	var shopPagesID []uuid.UUID
	if err := r.cache.Get(shopsKey, &shopPagesID); err != nil {
		return err
	}

	var tokens dto.UserTokensDTO
	if err := r.cache.Get(tokensKey, &tokens); err != nil {
		return err
	}

	request := HTTPRequest{
		Method: http.MethodGet,
		URL:    r.baseURL + "/api/shops",
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", tokens.AccessToken),
		},
		QueryParams: map[string]string{
			"title":        shopParams.Title,
			"address":      shopParams.Address,
			"phone_number": shopParams.PhoneNumber,
			"fio_director": shopParams.FioDirector,
			"limit":        fmt.Sprintf("%d", shopParams.Limit),
			"offset":       fmt.Sprintf("%d", shopParams.Offset),
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

	var shops []*models.ShopModel
	if err = json.Unmarshal(response.Body, &shops); err != nil {
		return err
	}

	printShops(shops, shopParams.Offset)
	copyShopIDsToArray(&shopPagesID, shops)
	r.cache.Set(shopsKey, shopPagesID)
	r.cache.Set(
		shopParamsKey,
		dto.ShopParamsDTO{
			Limit:  pageLimit,
			Offset: shopParams.Offset + pageLimit,
		},
	)

	return nil
}

func (r *Requester) addRatingProductByShop() error {
	var tokens dto.UserTokensDTO
	if err := r.cache.Get(tokensKey, &tokens); err != nil {
		return err
	}

	if err := r.getSalesByShopID(tokens); err != nil {
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

func (r *Requester) getSalesByShopID(tokens dto.UserTokensDTO) error {
	var shopPagesID []uuid.UUID
	if err := r.cache.Get(shopsKey, &shopPagesID); err != nil {
		return err
	}

	num, err := input.ShopPagesNumber()
	if err != nil {
		return err
	}

	if num > len(shopPagesID)-1 || num < 0 { // num -- это индекс
		return errors.New("номер магазина выходит из диапазона выведенных значений")
	}

	shopID := shopPagesID[num]

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

func (r *Requester) changePriceOnsaleProduct() error {
	var tokens dto.UserTokensDTO
	if err := r.cache.Get(tokensKey, &tokens); err != nil {
		return err
	}

	if err := r.getSalesByShopID(tokens); err != nil {
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

func (r *Requester) adNewProductInShop() error {
	var tokens dto.UserTokensDTO
	if err := r.cache.Get(tokensKey, &tokens); err != nil {
		return err
	}

	var shopPagesID []uuid.UUID
	if err := r.cache.Get(shopsKey, &shopPagesID); err != nil {
		return err
	}

	num, err := input.ShopPagesNumber()
	if err != nil {
		return err
	}

	if num > len(shopPagesID)-1 || num < 0 { // num -- это индекс
		return errors.New("номер магазина выходит из диапазона выведенных значений")
	}

	var newSaleProductDTO dto.NewSaleProductDTO
	newSaleProductDTO.ShopID = shopPagesID[num]

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

func printShops(shops []*models.ShopModel, offset int) {
	t := table.NewWriter()
	t.SetTitle(fmt.Sprintf("Страница магазинов №%d", offset/pageLimit+1))
	t.SetStyle(table.StyleBold)
	t.Style().Format.Header = text.FormatTitle
	t.AppendHeader(table.Row{"No.", "Название", "Адрес", "Номер телефона", "ФИО директора"})

	for i, shop := range shops {
		t.AppendRow(table.Row{offset + i, shop.Title, shop.Address, shop.PhoneNumber, shop.FioDirector})
	}
	fmt.Println(t.Render())
}

func copyShopIDsToArray(shopIDs *[]uuid.UUID, shops []*models.ShopModel) {
	for _, shop := range shops {
		*shopIDs = append(*shopIDs, shop.ID)
	}
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
