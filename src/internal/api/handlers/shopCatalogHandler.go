package handlers

import (
	"SmartShopper-services/core/dto"
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

func (h *Handler) getShops(c *gin.Context) {
	var shopParamsDTO dto.ShopParamsDTO
	shopParamsDTO.Title = c.Query("title")
	shopParamsDTO.Address = c.Query("address")
	shopParamsDTO.PhoneNumber = c.Query("phone_number")
	shopParamsDTO.FioDirector = c.Query("fio_director")
	limit := c.Query("limit")
	offset := c.Query("offset")

	if limit == "" || offset == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid parameter")
		return
	}
	var err error
	shopParamsDTO.Limit, err = strconv.Atoi(limit)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid limit")
		return
	}
	shopParamsDTO.Offset, err = strconv.Atoi(offset)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid offset")
		return
	}

	var shops []*models.ShopModel
	shops, err = h.shopService.GetByParams(c.Request.Context(), &shopParamsDTO)
	if err != nil && errors.Is(err, errs.ErrShopDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, shops)
}

func (h *Handler) deleteShopByID(c *gin.Context) {
	shopID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id")
		return
	}

	err = h.shopService.DeleteByID(c.Request.Context(), shopID)
	if err != nil && errors.Is(err, errs.ErrShopDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
