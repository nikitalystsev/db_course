package requesters

//
//import (
//	"SmartShopper-services/core/dto"
//	"SmartShopper/internal/techUI/input"
//	"encoding/json"
//	"errors"
//	"fmt"
//	"github.com/google/uuid"
//	"net/http"
//	"time"
//)
//
////const adminMainMenu = `Главное меню:
////	1 -- Перейти в каталог товаров
////	2 -- Добавить новый магазин
////	3 -- Перейти к обработке магазинов
////	0 -- Выйти
////`
//
//const adminShopMenu = `Меню обработки магазинов:
//	1 -- Вывести магазины
//	2 -- Следующая страница
//	3 -- Добавить товар в магазин
//	4 -- Поставить оценку товару в магазине
//	5 -- Изменить цену на товар в магазине
//	0 -- Вернуться в главное меню
//`
//
//func (r *Requester) processAdminShopActions() error {
//	var (
//		menuItem int
//		err      error
//	)
//	r.cache.Set(shopParamsKey, dto.ShopParamsDTO{Limit: pageLimit, Offset: 0})
//	r.cache.Set(shopsKey, make([]uuid.UUID, 0))
//
//	for {
//		fmt.Printf("\n\n%s", adminShopMenu)
//
//		if menuItem, err = input.MenuItem(); err != nil {
//			fmt.Printf("\n\n%s\n", err.Error())
//			continue
//		}
//
//		switch menuItem {
//		case 1:
//			if err = r.viewFirstPageShops(); err != nil {
//				fmt.Printf("\n\n%s\n", err.Error())
//			}
//		case 2:
//			if err = r.viewNextPage(); err != nil {
//				fmt.Printf("\n\n%s\n", err.Error())
//			}
//		case 3:
//			if err = r.adNewProductInShop(); err != nil {
//				fmt.Printf("\n\n%s\n", err.Error())
//			}
//		case 4:
//			if err = r.addRatingProductByShop(); err != nil {
//				fmt.Printf("\n\n%s\n", err.Error())
//			}
//		case 5:
//			if err = r.changePriceOnsaleProduct(); err != nil {
//				fmt.Printf("\n\n%s\n", err.Error())
//			}
//		case 0:
//			return nil
//		default:
//			fmt.Printf("\n\nНеверный пункт меню!\n")
//		}
//
//	}
//}
//
//func (r *Requester) signInAsAdmin(stopRefresh <-chan struct{}) error {
//	readerSignInDTO, err := input.SignInParams()
//	if err != nil {
//		return err
//	}
//
//	request := HTTPRequest{
//		Method: http.MethodPost,
//		URL:    r.baseURL + "/auth/sign-in",
//		Headers: map[string]string{
//			"Content-Type": "application/json",
//		},
//		Body:    readerSignInDTO,
//		Timeout: 10 * time.Second,
//	}
//
//	response, err := SendRequest(request)
//	if err != nil {
//		return err
//	}
//
//	if response.StatusCode != http.StatusOK {
//		var info string
//		if err = json.Unmarshal(response.Body, &info); err != nil {
//			return err
//		}
//		return errors.New(info)
//	}
//
//	var tokens dto.UserTokensDTO
//	if err = json.Unmarshal(response.Body, &tokens); err != nil {
//		return err
//	}
//
//	r.cache.Set("tokens", tokens)
//
//	fmt.Printf("\n\nВы успешно вошли в систему как Администратор!\n")
//
//	go r.Refreshing(r.accessTokenTTL, stopRefresh)
//
//	return nil
//}
//
//func (r *Requester) deleteShop() error {
//	var tokens dto.UserTokensDTO
//	if err := r.cache.Get(tokensKey, &tokens); err != nil {
//		return err
//	}
//
//	var shopDTO dto.ShopDTO
//	fmt.Printf("\n\nДля добавления нового магазина необходимо ввести " +
//		"данные Ритейлера, с которым он сотрудничает\n")
//
//	shopParams, err := input.RetailerParams()
//	if err != nil {
//		return err
//	}
//	shopDTO.Retailer = *shopParams
//
//	shopDTO.ShopParams, err = input.ShopParams()
//	if err != nil {
//		return err
//	}
//
//	request := HTTPRequest{
//		Method: http.MethodPost,
//		URL:    r.baseURL + "/api/shops",
//		Headers: map[string]string{
//			"Content-Type":  "application/json",
//			"Authorization": fmt.Sprintf("Bearer %s", tokens.AccessToken),
//		},
//		Body:    shopDTO,
//		Timeout: 10 * time.Second,
//	}
//
//	response, err := SendRequest(request)
//	if err != nil {
//		return err
//	}
//
//	if response.StatusCode != http.StatusOK {
//		var info string
//		if err = json.Unmarshal(response.Body, &info); err != nil {
//			return err
//		}
//		return errors.New(info)
//	}
//
//	fmt.Printf("\n\nМагазин был успешно добавлен!\n")
//
//	return nil
//}
