package dto

type UserSignInDTO struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}
