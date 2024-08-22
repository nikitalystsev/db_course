package intf

import (
	"SmartShopper-services/core/dto"
	"SmartShopper-services/core/models"
	"context"
)

type IUserService interface {
	SignUp(ctx context.Context, user *models.UserModel) error
	SignIn(ctx context.Context, user *dto.UserSignInDTO) (models.Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (models.Tokens, error)
}
