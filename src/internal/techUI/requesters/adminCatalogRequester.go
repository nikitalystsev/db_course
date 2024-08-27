package requesters

import (
	"SmartShopper-services/core/dto"
	"SmartShopper/internal/techUI/input"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

const adminCatalogMenu = `Меню каталога:
	1 -- Вывести товары
	2 -- Перейти к товару
	0 -- Вернуться в главное меню
`

func (r *Requester) processAdminCatalogActions() error {
	var (
		menuItem int
		err      error
	)
	r.cache.Set(productParamsKey, dto.ProductParamsDTO{Limit: pageLimit, Offset: 0})
	r.cache.Set(productsKey, make([]uuid.UUID, 0))

	for {
		fmt.Printf("\n\n%s", adminCatalogMenu)

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
			if err = r.processAdminProduct(); err != nil {
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

func (r *Requester) processAdminProduct() error {
	var productPagesID []uuid.UUID
	if err := r.cache.Get(productsKey, &productPagesID); err != nil {
		return err
	}

	num, err := input.ProductPagesNumber()
	if err != nil {
		return err
	}

	if num > len(productPagesID)-1 || num < 0 { // num -- это индекс
		return errors.New("номер товара выходит из диапазона выведенных значений")
	}

	productID := productPagesID[num]

	err = r.processAdminProductActions(productID, num)
	if err != nil {
		return err
	}

	return nil
}
