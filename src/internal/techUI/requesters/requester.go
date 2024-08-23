package requesters

import (
	"SmartShopper/internal/techUI/input"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-humble/locstor"
	"net/http"
	"time"
)

const mainMenu = `Главное меню:
	1 -- Зарегистрироваться
	2 -- Войти
	0 -- Остановить выполнение программы
`

type Requester struct {
	localStorage    *locstor.DataStore
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	baseURL         string
}

func NewRequester(accessTokenTTL, refreshTokenTTL time.Duration, port string) *Requester {
	return &Requester{
		localStorage:    locstor.NewDataStore(locstor.JSONEncoding),
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
