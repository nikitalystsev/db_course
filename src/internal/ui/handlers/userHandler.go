package handlers

import (
	"SmartShopper-services/core/dto"
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (h *Handler) signUp(c *gin.Context) {
	fmt.Println("call signUp")

	var inp dto.UserSignUpDTO
	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	user := models.UserModel{
		ID:               uuid.New(),
		Fio:              inp.Fio,
		PhoneNumber:      inp.PhoneNumber,
		Password:         inp.Password,
		RegistrationDate: time.Now(),
	}

	err := h.userService.SignUp(c.Request.Context(), &user)
	if err != nil && errors.Is(err, errs.ErrUserObjectIsNil) {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	if err != nil && errors.Is(err, errs.ErrUserAlreadyExist) {
		c.AbortWithStatusJSON(http.StatusConflict, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		println(err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) signIn(c *gin.Context) {
	println("call signIn")

	var inp dto.UserSignInDTO
	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.userService.SignIn(c.Request.Context(), &inp)
	if err != nil && errors.Is(err, errs.ErrUserDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil && errors.Is(err, errors.New("wrong password")) {
		c.AbortWithStatusJSON(http.StatusConflict, err.Error())
		return
	}
	if err != nil && errors.Is(err, errs.ErrUserObjectIsNil) {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.ReaderTokensDTO{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
		ExpiredAt:    time.Now().Add(h.accessTokenTTL).UnixMilli(),
	})
}

func (h *Handler) refresh(c *gin.Context) {
	var inp string
	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid input body")
		return
	}

	res, err := h.userService.RefreshTokens(c.Request.Context(), inp)
	if err != nil && errors.Is(err, errs.ErrUserDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.ReaderTokensDTO{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
		ExpiredAt:    time.Now().Add(h.accessTokenTTL).UnixMilli(),
	})
}
