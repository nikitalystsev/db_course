package requesters

import (
	"SmartShopper-services/core/dto"
	"SmartShopper/internal/techUI/input"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const userMainMenu = `Главное меню:
	1 -- Перейти в каталог товаров
	2 -- Добавить новый магазин
	3 -- Перейти к обработке магазинов
	0 -- выйти
`
const tokensKey = "tokens"

func (r *Requester) processUserActions() error {
	var (
		menuItem int
		err      error
	)
	stopRefresh := make(chan struct{})
	if err = r.signIn(stopRefresh); err != nil {
		fmt.Printf("\n\n%s\n", err.Error())
		return err
	}

	for {
		fmt.Printf("\n\n%s", userMainMenu)

		if menuItem, err = input.MenuItem(); err != nil {
			fmt.Printf("\n\n%s\n", err.Error())
			continue
		}

		switch menuItem {
		case 1:
			if err = r.processCatalogActions(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 2:
			if err = r.addNewShop(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 3:
			if err = r.processShopCatalogActions(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 0:
			close(stopRefresh)
			r.cache.Delete("tokens")
			fmt.Printf("\n\nВы успешно вышли из системы!\n")
			return nil
		default:
			fmt.Printf("\n\nНеверный пункт меню!\n")
		}

	}
}

func (r *Requester) addNewShop() error {
	var tokens dto.UserTokensDTO
	if err := r.cache.Get(tokensKey, &tokens); err != nil {
		return err
	}

	var shopDTO dto.ShopDTO
	fmt.Printf("\n\nДля добавления нового магазина необходимо ввести " +
		"данные Ритейлера, с которым он сотрудничает\n")

	shopParams, err := input.RetailerParams()
	if err != nil {
		return err
	}
	shopDTO.Retailer = *shopParams

	shopDTO.ShopParams, err = input.ShopParams()
	if err != nil {
		return err
	}

	request := HTTPRequest{
		Method: http.MethodPost,
		URL:    r.baseURL + "/api/shops",
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", tokens.AccessToken),
		},
		Body:    shopDTO,
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

	fmt.Printf("\n\nМагазин был успешно добавлен!\n")

	return nil
}

func (r *Requester) signIn(stopRefresh <-chan struct{}) error {
	readerSignInDTO, err := input.SignInParams()
	if err != nil {
		return err
	}

	request := HTTPRequest{
		Method: http.MethodPost,
		URL:    r.baseURL + "/auth/sign-in",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:    readerSignInDTO,
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

	var tokens dto.UserTokensDTO
	if err = json.Unmarshal(response.Body, &tokens); err != nil {
		return err
	}

	r.cache.Set("tokens", tokens)

	fmt.Printf("\n\nВы успешно вошли в систему!\n")

	go r.Refreshing(r.accessTokenTTL, stopRefresh)

	return nil
}

func (r *Requester) Refresh() error {
	var tokens dto.UserTokensDTO
	if err := r.cache.Get("tokens", &tokens); err != nil {
		return err
	}

	request := HTTPRequest{
		Method: http.MethodPost,
		URL:    r.baseURL + "/auth/refresh",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:    tokens.RefreshToken,
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

	if err = json.Unmarshal(response.Body, &tokens); err != nil {
		return err
	}

	r.cache.Set("tokens", tokens)

	//fmt.Printf("\n\nSuccessful refresh tokens!\n")

	return nil
}

func (r *Requester) Refreshing(interval time.Duration, stopRefresh <-chan struct{}) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := r.Refresh(); err != nil {
				fmt.Printf("\n\nОшибка обновления токенов: %v\n", err)
			}
		case <-stopRefresh:
			return
		}
	}
}
