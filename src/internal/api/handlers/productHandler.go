package handlers

import (
	"SmartShopper-services/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (h *Handler) getProductByID(c *gin.Context) {
	IDStr := c.Param("id")
	ID, err := uuid.Parse(IDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.productService.GetByID(c.Request.Context(), ID)
	if errors.Is(err, errs.ErrProductDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}
