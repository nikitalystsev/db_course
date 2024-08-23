package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MenuItem() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите пункт меню: ")

	menuItemStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	menuItemStr = strings.TrimSpace(menuItemStr)

	menuItemInt, err := strconv.Atoi(menuItemStr)
	if err != nil {
		return 0, err
	}

	return menuItemInt, nil
}
