package impl

import (
	"SmartShopper-services/core/dto"
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"SmartShopper-services/intf"
	"SmartShopper-services/intfRepo"
	"SmartShopper-services/pkg/auth"
	"SmartShopper-services/pkg/hash"
	"context"
	"errors"
	"github.com/google/uuid"
	"strconv"
	"time"
)

const (
	UserPhoneNumberLen = 11
	UserPasswordLen    = 10
)

type UserService struct {
	userRepo        intfRepo.IUserRepo
	tokenManager    auth.ITokenManager
	hasher          hash.IPasswordHasher
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewUserService(
	userRepo intfRepo.IUserRepo,
	tokenManager auth.ITokenManager,
	hasher hash.IPasswordHasher,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
) intf.IUserService {
	return &UserService{
		userRepo:        userRepo,
		tokenManager:    tokenManager,
		hasher:          hasher,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

// SignUp Зарегистрироваться
func (us *UserService) SignUp(ctx context.Context, user *models.UserModel) error {
	if user == nil {
		return errs.ErrUserObjectIsNil
	}
	err := us.baseValidation(ctx, user)
	if err != nil {
		return err
	}

	hashedPassword, err := us.hasher.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	err = us.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) SignIn(ctx context.Context, user *dto.UserSignInDTO) (models.Tokens, error) {
	if user == nil {
		return models.Tokens{}, errs.ErrUserObjectIsNil
	}

	exitingUser, err := us.userRepo.GetByPhoneNumber(ctx, user.PhoneNumber)
	if err != nil && !errors.Is(err, errs.ErrUserDoesNotExists) {
		return models.Tokens{}, err
	}

	if exitingUser == nil {
		return models.Tokens{}, errs.ErrUserDoesNotExists
	}

	err = us.hasher.Compare(exitingUser.Password, user.Password)
	if err != nil {
		return models.Tokens{}, err
	}

	return us.createTokens(ctx, exitingUser.ID, "") // пока что пустая роль
}

func (us *UserService) RefreshTokens(ctx context.Context, refreshToken string) (models.Tokens, error) {
	existingReader, err := us.userRepo.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		return models.Tokens{}, err
	}

	return us.createTokens(ctx, existingReader.ID, "") // пока что пустая роль
}

func (us *UserService) baseValidation(ctx context.Context, user *models.UserModel) error {
	existingUser, err := us.userRepo.GetByPhoneNumber(ctx, user.PhoneNumber)

	if err != nil && !errors.Is(err, errs.ErrUserDoesNotExists) {
		return err
	}

	if existingUser != nil {
		return errs.ErrUserAlreadyExist
	}

	if user.Fio == "" {
		return errs.ErrEmptyUserFio
	}

	if user.PhoneNumber == "" {
		return errs.ErrEmptyUserPhoneNumber
	}

	if user.Password == "" {
		return errs.ErrEmptyUserPassword
	}

	if len(user.Password) != UserPasswordLen {
		return errs.ErrInvalidUserPasswordLen
	}

	if len(user.PhoneNumber) != UserPhoneNumberLen {
		return errs.ErrInvalidUserPhoneNumberLen
	}

	_, err = strconv.Atoi(user.PhoneNumber)
	if err != nil {
		return errs.ErrInvalidUserPhoneNumberFormat
	}

	return nil
}

func (us *UserService) createTokens(ctx context.Context, userID uuid.UUID, userRole string) (models.Tokens, error) {
	var (
		res models.Tokens
		err error
	)
	res.AccessToken, err = us.tokenManager.NewJWT(userID, userRole, us.accessTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = us.tokenManager.NewRefreshToken()
	if err != nil {
		return res, err
	}

	err = us.userRepo.SaveRefreshToken(ctx, userID, res.RefreshToken, us.refreshTokenTTL)
	if err != nil {
		return res, err
	}

	return res, nil
}
