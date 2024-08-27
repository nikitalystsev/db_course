package requesters

import (
	"SmartShopper-services/core/dto"
	"SmartShopper/internal/techUI/input"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

const adminShopCatalogMenu = `Меню обработки магазинов:
	1 -- Вывести магазины
	2 -- Следующая страница
	3 -- Удалить магазин
	4 -- Перейти к магазину
	0 -- Вернуться в главное меню
`

func (r *Requester) processAdminShopCatalogActions() error {
	var (
		menuItem int
		err      error
	)
	r.cache.Set(shopParamsKey, dto.ShopParamsDTO{Limit: pageLimit, Offset: 0})
	r.cache.Set(shopsKey, make([]uuid.UUID, 0))

	for {
		fmt.Printf("\n\n%s", adminShopCatalogMenu)

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
			if err = r.deleteShop(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 4:
			if err = r.processShop(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 0:
			return nil
		default:
			fmt.Printf("\n\nНеверный пункт меню!\n")
		}

	}
}

func (r *Requester) deleteShop() error {
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

	shopID := shopPagesID[num]

	request := HTTPRequest{
		Method: http.MethodDelete,
		URL:    r.baseURL + fmt.Sprintf("/api/shops/%s", shopID.String()),
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", tokens.AccessToken),
		},
		Timeout: 10 * time.Second,
	}
	response, err := SendRequest(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusNoContent {
		var info string
		if err = json.Unmarshal(response.Body, &info); err != nil {
			return err
		}
		return errors.New(info)
	}

	fmt.Printf("\n\nМагазин был успешно удален!\n")

	return nil
}
