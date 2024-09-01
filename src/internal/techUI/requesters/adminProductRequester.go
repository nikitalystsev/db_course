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

const adminProductMenu = `Меню обработки товара:
	1 -- Сравнить цену на товар
	2 -- Посмотреть сертификаты соответствия на товар
	3 -- Добавить сертификат соответствия на товар
	4 -- Удалить сертификат соответствия на товар
	5 -- Обновить статус соответствия на товар
	0 -- Вернуться в главное меню
`

func (r *Requester) processAdminProductActions(productID uuid.UUID, num int) error {
	var (
		menuItem int
		err      error
	)
	r.cache.Set(certificateKey, make([]uuid.UUID, 0))

	if err = r.viewProduct(productID, num); err != nil {
		return err
	}

	for {
		fmt.Printf("\n\n%s", adminProductMenu)

		if menuItem, err = input.MenuItem(); err != nil {
			fmt.Printf("\n\n%s\n", err.Error())
			continue
		}

		switch menuItem {
		case 1:
			if err = r.comparePriceOnProduct(productID, num); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 2:
			if err = r.viewCertificatesCompliance(productID, num); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 3:
			if err = r.addCertificate(productID, num); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 4:
			if err = r.deleteCertificate(productID, num); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 5:
			if err = r.updateCertificate(productID, num); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 0:
			return nil
		default:
			fmt.Printf("\n\nНеверный пункт меню!\n")
		}
	}
}

func (r *Requester) addCertificate(productID uuid.UUID, num int) error {
	var tokens dto.UserTokensDTO
	if err := r.cache.Get(tokensKey, &tokens); err != nil {
		return err
	}

	certificateDTO, err := input.CertificateParams()
	if err != nil {
		return err
	}
	certificateDTO.ProductID = productID

	request := HTTPRequest{
		Method: http.MethodPost,
		URL:    r.baseURL + "/api/certificates",
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", tokens.AccessToken),
		},
		Body:    certificateDTO,
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

	fmt.Printf("\n\nСертификат был успешно добавлен!\n")

	return nil
}

func (r *Requester) deleteCertificate(productID uuid.UUID, num int) error {
	var certificateIDs []uuid.UUID
	if err := r.cache.Get(certificateKey, &certificateIDs); err != nil {
		return err
	}

	var tokens dto.UserTokensDTO
	if err := r.cache.Get(tokensKey, &tokens); err != nil {
		return err
	}

	num, err := input.CertificateCatalogNumber()
	if err != nil {
		return err
	}

	if num > len(certificateIDs)-1 || num < 0 { // num -- это индекс
		return errors.New("номер сертификата выходит из диапазона выведенных значений")
	}

	certificateID := certificateIDs[num]

	request := HTTPRequest{
		Method: http.MethodDelete,
		URL:    r.baseURL + fmt.Sprintf("/api/certificates/%s", certificateID.String()),
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

	fmt.Printf("\n\nСертификат был успешно удален!\n")

	return nil
}

func (r *Requester) updateCertificate(productID uuid.UUID, num int) error {
	var certificateIDs []uuid.UUID
	if err := r.cache.Get(certificateKey, &certificateIDs); err != nil {
		return err
	}

	var tokens dto.UserTokensDTO
	if err := r.cache.Get(tokensKey, &tokens); err != nil {
		return err
	}

	num, err := input.CertificateCatalogNumber()
	if err != nil {
		return err
	}

	if num > len(certificateIDs)-1 || num < 0 { // num -- это индекс
		return errors.New("номер сертификата выходит из диапазона выведенных значений")
	}

	certificateID := certificateIDs[num]

	certificateStatusDTO, err := input.CertificateStatusParam()
	if err != nil {
		return err
	}

	request := HTTPRequest{
		Method: http.MethodPut,
		URL:    r.baseURL + fmt.Sprintf("/api/certificates/%s", certificateID.String()),
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", tokens.AccessToken),
		},
		Body:    certificateStatusDTO,
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

	fmt.Printf("\n\nСтатус сертификата №%d был успешно обновлен!\n", num)

	return nil
}
