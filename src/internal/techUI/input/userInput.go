package input

import (
	"SmartShopper-services/core/dto"
	"bufio"
	"fmt"
	"github.com/howeyc/gopass"
	"os"
	"strings"
)

func Fio() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите ваше ФИО: ")

	fio, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	fio = strings.TrimSpace(fio)

	return fio, nil
}

func PhoneNumber() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите ваш номер телефона: ")

	phoneNumber, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	phoneNumber = strings.TrimSpace(phoneNumber)

	return phoneNumber, nil
}

func Password() (string, error) {
	fmt.Print("Введите ваш пароль: ")

	silentPassword, err := gopass.GetPasswdMasked()
	if err != nil {
		return "", err
	}

	password := string(silentPassword)
	password = strings.TrimSpace(password)

	return password, nil
}

func SignUpParams() (*dto.UserSignUpDTO, error) {
	var (
		res dto.UserSignUpDTO
		err error
	)

	if res.Fio, err = Fio(); err != nil {
		return nil, err
	}
	if res.PhoneNumber, err = PhoneNumber(); err != nil {
		return nil, err
	}
	if res.Password, err = Password(); err != nil {
		return nil, err
	}

	return &res, nil
}

func SignInParams() (*dto.UserSignInDTO, error) {
	var (
		res dto.UserSignInDTO
		err error
	)

	if res.PhoneNumber, err = PhoneNumber(); err != nil {
		return nil, err
	}
	if res.Password, err = Password(); err != nil {
		return nil, err
	}

	return &res, nil
}
