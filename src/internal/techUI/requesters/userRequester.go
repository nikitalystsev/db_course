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
	0 -- выйти
`

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
		case 0:
			close(stopRefresh)
			fmt.Println("\n\nВы успешно вышли из системы!")
			return nil
		default:
			fmt.Printf("\n\nНеверный пункт меню!\n")
		}

	}

}

func (r *Requester) signIn(stopRefresh <-chan struct{}) error {
	readerSignInDTO, err := input.SignInParams()
	if err != nil {
		return err
	}

	request := HTTPRequest{
		Method: "POST",
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

	var tokens dto.ReaderTokensDTO
	if err = json.Unmarshal(response.Body, &tokens); err != nil {
		return err
	}
	if err = r.localStorage.Save("tokens", tokens); err != nil {
		return err
	}

	fmt.Printf("\n\nAuthentication successful!\n")

	go r.Refreshing(r.accessTokenTTL, stopRefresh)

	return nil
}

func (r *Requester) Refresh() error {
	var tokens dto.ReaderTokensDTO
	if err := r.localStorage.Find("tokens", &tokens); err != nil {
		return err
	}
	request := HTTPRequest{
		Method: "POST",
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
	if err = r.localStorage.Save("tokens", tokens); err != nil {
		return err
	}

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
				fmt.Printf("\n\nerror refreshing tokens: %v\n", err)
			}
		case <-stopRefresh:
			return
		}
	}
}
