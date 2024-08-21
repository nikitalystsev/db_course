package intf

import (
	"SmartShopper-services/core/dto"
	"SmartShopper-services/core/models"
	"context"
)

type IUserService interface {
	SignUp(ctx context.Context, reader *models.UserModel) error
	SignIn(ctx context.Context, reader *dto.UserSignInDTO) (models.Tokens, error)
}
