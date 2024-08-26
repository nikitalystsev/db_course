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
	var shopDTO dto.ShopDTO
	shopDTO.RetailerID = uuid.Nil
	shopDTO.Title = c.Query("title")
	shopDTO.Address = c.Query("address")
	shopDTO.PhoneNumber = c.Query("phone_number")
	shopDTO.FioDirector = c.Query("fio_director")
	limit := c.Query("limit")
	offset := c.Query("offset")

	if limit == "" || offset == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid parameter")
		return
	}
	var err error
	shopDTO.Limit, err = strconv.Atoi(limit)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid limit")
		return
	}
	shopDTO.Offset, err = strconv.Atoi(offset)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid offset")
		return
	}

	var shops []*models.ShopModel
	shops, err = h.shopService.GetByParams(c.Request.Context(), &shopDTO)
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
