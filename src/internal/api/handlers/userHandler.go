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
	fmt.Println("call signIn")

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

	c.JSON(http.StatusOK, dto.UserTokensDTO{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
		ExpiredAt:    time.Now().Add(h.accessTokenTTL).UnixMilli(),
	})
}

func (h *Handler) addRetailerIfNotExist(c *gin.Context) {
	var retailerDTO dto.SupplierDTO
	if err := c.BindJSON(&retailerDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	retailer := &models.SupplierModel{
		ID:                uuid.New(),
		Title:             retailerDTO.Title,
		Address:           retailerDTO.Address,
		PhoneNumber:       retailerDTO.PhoneNumber,
		FioRepresentative: retailerDTO.FioRepresentative,
	}

	err := h.supplierService.CreateRetailer(c.Request.Context(), retailer)
	if err != nil && errors.Is(err, errs.ErrRetailerAlreadyExist) {
		c.AbortWithStatusJSON(http.StatusConflict, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, retailer.ID)
}

func (h *Handler) getRetailerByAddress(c *gin.Context) {
	address := c.Query("address")

	retailer, err := h.supplierService.GetRetailerByAddress(c.Request.Context(), address)
	if err != nil && errors.Is(err, errs.ErrRetailerDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, retailer.ID)
}

func (h *Handler) addShop(c *gin.Context) {
	var shopDTO dto.ShopDTO
	if err := c.BindJSON(&shopDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err := h.shopService.Create(c.Request.Context(), &shopDTO)
	if err != nil && errors.Is(err, errs.ErrShopAlreadyExist) {
		c.AbortWithStatusJSON(http.StatusConflict, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
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

	c.JSON(http.StatusOK, dto.UserTokensDTO{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
		ExpiredAt:    time.Now().Add(h.accessTokenTTL).UnixMilli(),
	})
}
