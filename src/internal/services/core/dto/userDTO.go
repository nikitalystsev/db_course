package dto

type UserSignInDTO struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserSignUpDTO struct {
	Fio         string `json:"fio"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserTokensDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredAt    int64  `json:"expired_at"`
}
