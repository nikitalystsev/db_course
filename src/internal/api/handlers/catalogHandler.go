package handlers

import (
	"SmartShopper-services/core/dto"
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

	var productDTO *dto.ProductDTO

	if h.inUseCache {
		cacheKey := "product:" + ID.String()
		cachedProduct, err := h.cache.Get(c.Request.Context(), cacheKey).Result()
		if err == nil {
			if err = json.Unmarshal([]byte(cachedProduct), &productDTO); err == nil {
				c.JSON(http.StatusOK, productDTO)
				return
			}
		}
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

	stat, err := h.getCertificateStatisticsByProductID(product.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	productDTO = &dto.ProductDTO{
		Retailer:              retailer.Title,
		Distributor:           distributor.Title,
		Manufacturer:          manufacturer.Title,
		Name:                  product.Name,
		Categories:            product.Categories,
		Brand:                 product.Brand,
		Compound:              product.Compound,
		GrossMass:             product.GrossMass,
		NetMass:               product.NetMass,
		PackageType:           product.PackageType,
		CertificatesStatistic: stat,
	}

	if h.inUseCache {
		cacheKey := "product:" + ID.String()
		productJSON, _ := json.Marshal(productDTO)
		h.cache.Set(c.Request.Context(), cacheKey, productJSON, 0)
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

	if h.inUseCache {
		cacheKey := fmt.Sprintf("products:limit:%d:offset:%d", limitInt, offsetInt)
		cachedData, err := h.cache.Get(c.Request.Context(), cacheKey).Result()
		if err == nil {
			if err = json.Unmarshal([]byte(cachedData), &products); err == nil {
				c.JSON(http.StatusOK, products)
				return
			}
		}
	}

	products, err = h.productService.GetPage(c.Request.Context(), limitInt, offsetInt)
	if err != nil && errors.Is(err, errs.ErrProductDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	productsCertificates, err := h.addInfoAboutCertificates(products)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if h.inUseCache {
		cacheKey := fmt.Sprintf("products:limit:%d:offset:%d", limitInt, offsetInt)
		jsonData, _ := json.Marshal(products)
		h.cache.Set(c.Request.Context(), cacheKey, jsonData, 0)
	}

	c.JSON(http.StatusOK, productsCertificates)
}

func (h *Handler) addInfoAboutCertificates(products []*models.ProductModel) ([]*dto.ProductCertificateDTO, error) {
	productCertificatesDTO := make([]*dto.ProductCertificateDTO, len(products))

	for i, product := range products {
		stat, err := h.getCertificateStatisticsByProductID(product.ID)
		if err != nil {
			return nil, err
		}
		productCertificatesDTO[i] = &dto.ProductCertificateDTO{
			ID:                    product.ID,
			RetailerID:            product.RetailerID,
			DistributorID:         product.DistributorID,
			ManufacturerID:        product.ManufacturerID,
			Name:                  product.Name,
			Categories:            product.Categories,
			Brand:                 product.Brand,
			Compound:              product.Compound,
			GrossMass:             product.GrossMass,
			NetMass:               product.NetMass,
			PackageType:           product.PackageType,
			CertificatesStatistic: stat,
		}
	}

	return productCertificatesDTO, nil
}

func (h *Handler) getCertificateStatisticsByProductID(productID uuid.UUID) (string, error) {
	certificateStatisticsDTO, err := h.certificateService.GetCertificateStatisticsByProductID(
		context.Background(),
		productID,
	)
	if err != nil && !errors.Is(err, errs.ErrCertificateDoesNotExists) {
		return "", err
	}

	if certificateStatisticsDTO == nil {
		return "Не сертифицирован", nil
	}

	if certificateStatisticsDTO.TotalCountCertificates == 0 {
		return "Не сертифицирован", nil
	}

	if certificateStatisticsDTO.CountValidCertificates == 0 {
		return fmt.Sprintf(
			"Не соответствует ни одному из %d сертификатов",
			certificateStatisticsDTO.CountValidCertificates,
		), nil
	}

	return fmt.Sprintf(
		"Cоответствует %d/%d сертификатам",
		certificateStatisticsDTO.CountValidCertificates,
		certificateStatisticsDTO.TotalCountCertificates,
	), err
}
