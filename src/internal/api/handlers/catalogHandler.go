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

	retailer, err := h.getRetailerByID(product.RetailerID)
	if err != nil && errors.Is(err, errs.ErrRetailerDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	distributor, err := h.getDistributorByID(product.DistributorID)
	if err != nil && errors.Is(err, errs.ErrDistributorDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	manufacturer, err := h.getManufacturerByID(product.ManufacturerID)
	if err != nil && errors.Is(err, errs.ErrDistributorDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	productDTO := &dto.ProductDTO{
		Retailer:     retailer.Title,
		Distributor:  distributor.Title,
		Manufacturer: manufacturer.Title,
		Name:         product.Name,
		Categories:   product.Categories,
		Brand:        product.Brand,
		Compound:     product.Compound,
		GrossMass:    product.GrossMass,
		NetMass:      product.NetMass,
		PackageType:  product.PackageType,
	}

	c.JSON(http.StatusOK, productDTO)
}

func (h *Handler) getProducts(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	if limit == "" || offset == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid parameter")
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid limit")
		return
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid offset")
		return
	}

	var products []*models.ProductModel
	products, err = h.productService.GetPage(c.Request.Context(), limitInt, offsetInt)
	if err != nil && errors.Is(err, errs.ErrProductDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}
