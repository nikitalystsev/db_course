package requesters

import (
	"SmartShopper/internal/techUI/input"
	myCache "SmartShopper/pkg/cache"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

const mainMenu = `Главное меню:
	1 -- Зарегистрироваться
	2 -- Войти
	3 -- Войти как администратор
	4 -- Перейти к каталогу товаров
	0 -- Остановить выполнение программы
`

type Requester struct {
	cache           myCache.ICache
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	baseURL         string
}

func NewRequester(accessTokenTTL, refreshTokenTTL time.Duration, port string) *Requester {
	return &Requester{
		cache:           myCache.NewCache(),
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
		baseURL:         "http://localhost:" + port,
	}
}

func (r *Requester) Run() {
	var (
		menuItem int
		err      error
	)
	for {
		fmt.Printf("\n\n%s", mainMenu)

		if menuItem, err = input.MenuItem(); err != nil {
			fmt.Printf("\n\n%s\n", err)
			continue
		}

		switch menuItem {
		case 1:
			if err = r.signUp(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 2:
			if err = r.processUserActions(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 4:
			if err = r.processCatalogActions(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 0:
			os.Exit(0)
		default:
			fmt.Printf("\n\nНеверный пункт меню!\n")
		}
	}
}

func (r *Requester) signUp() error {
	userSignUpDTO, err := input.SignUpParams()
	if err != nil {
		return err
	}

	request := HTTPRequest{
		Method: http.MethodPost,
		URL:    r.baseURL + "/auth/sign-up",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:    userSignUpDTO,
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

	fmt.Printf("\n\nВы были успешно зарегистрированы!\n")

	return nil
}
