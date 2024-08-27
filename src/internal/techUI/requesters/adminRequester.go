package requesters

import (
	"SmartShopper/internal/techUI/input"
	"fmt"
)

const adminMainMenu = `Главное меню:
	1 -- Перейти в каталог товаров
	2 -- Добавить новый магазин
	3 -- Перейти к обработке магазинов
	0 -- выйти
`

func (r *Requester) processAdminActions() error {
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
		fmt.Printf("\n\n%s", adminMainMenu)

		if menuItem, err = input.MenuItem(); err != nil {
			fmt.Printf("\n\n%s\n", err.Error())
			continue
		}

		switch menuItem {
		case 1:
			if err = r.processAdminCatalogActions(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 2:
			if err = r.addNewShop(); err != nil {
				fmt.Printf("\n\n%s\n", err.Error())
			}
		case 3:
			if err = r.processAdminShopCatalogActions(); err != nil {
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
