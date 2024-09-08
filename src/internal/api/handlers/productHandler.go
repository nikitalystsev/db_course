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
)

func (h *Handler) getSalesByProductID(c *gin.Context) {
	productIDStr := c.Query("product_id")
	if productIDStr == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid parameter")
		return
	}

	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid product_id")
		return
	}

	var salesDTO []*dto.SaleProductDTO

	if h.inUseCache {
		cacheKey := "sales:" + productID.String()
		cachedSales, err := h.cache.Get(c.Request.Context(), cacheKey).Result()
		if err == nil {
			if err := json.Unmarshal([]byte(cachedSales), &salesDTO); err == nil {
				c.JSON(http.StatusOK, salesDTO)
				return
			}
		}
	}

	sales, err := h.saleProductService.GetByProductID(c.Request.Context(), productID)
	if err != nil && errors.Is(err, errs.ErrSaleProductDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	salesDTO, err = h.copySalesToDTO(sales)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if h.inUseCache {
		cacheKey := "sales:" + productID.String()
		salesJSON, _ := json.Marshal(salesDTO)
		h.cache.Set(c.Request.Context(), cacheKey, salesJSON, 0)
	}

	c.JSON(http.StatusOK, salesDTO)
}

func (h *Handler) getCertificatesByProductID(c *gin.Context) {
	productIDStr := c.Query("product_id")
	if productIDStr == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid parameter")
		return
	}

	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid product_id")
		return
	}

	var certificates []*models.CertificateModel

	if h.inUseCache {
		cacheKey := fmt.Sprintf("certificates:%s", productID)
		cashedData, err := h.cache.Get(c.Request.Context(), cacheKey).Result()
		if err == nil {
			if err = json.Unmarshal([]byte(cashedData), &certificates); err == nil {
				c.JSON(http.StatusOK, certificates)
				return
			}
		}
	}

	certificates, err = h.certificateService.GetByProductID(c.Request.Context(), productID)
	if err != nil && errors.Is(err, errs.ErrCertificateDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if h.inUseCache {
		cacheKey := fmt.Sprintf("certificates:%s", productID)
		jsonData, _ := json.Marshal(certificates)
		h.cache.Set(c.Request.Context(), cacheKey, jsonData, 0)
	}

	c.JSON(http.StatusOK, certificates)
}

func (h *Handler) addCertificate(c *gin.Context) {
	var certificateDTO dto.CertificateDTO
	if err := c.BindJSON(&certificateDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	certificate := &models.CertificateModel{
		ID:                uuid.New(),
		ProductID:         certificateDTO.ProductID,
		Type:              certificateDTO.Type,
		Number:            certificateDTO.Number,
		NormativeDocument: certificateDTO.NormativeDocument,
		StatusCompliance:  certificateDTO.StatusCompliance,
		RegistrationDate:  certificateDTO.RegistrationDate,
		ExpirationDate:    certificateDTO.ExpirationDate,
	}

	err := h.certificateService.Create(c.Request.Context(), certificate)
	if err != nil && errors.Is(err, errs.ErrCertificateAlreadyExist) {
		c.AbortWithStatusJSON(http.StatusConflict, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusCreated)
}

func (h *Handler) deleteCertificateByID(c *gin.Context) {
	certificateID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.certificateService.DeleteByID(c.Request.Context(), certificateID)
	if err != nil && errors.Is(err, errs.ErrCertificateDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) updateCertificateByID(c *gin.Context) {
	certificateID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var certificateStatusDTO dto.CertificateStatusDTO
	if err = c.BindJSON(&certificateStatusDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	certificate, err := h.certificateService.GetByID(c.Request.Context(), certificateID)
	if err != nil && errors.Is(err, errs.ErrCertificateDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	certificate.StatusCompliance = certificateStatusDTO.StatusCompliance
	err = h.certificateService.Update(c.Request.Context(), certificate)
	if err != nil && errors.Is(err, errs.ErrCertificateDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}
func (h *Handler) getRetailerByID(retailerID uuid.UUID) (*models.SupplierModel, error) {
	retailer, err := h.supplierService.GetRetailerByID(context.Background(), retailerID)
	if err != nil {
		return nil, err
	}

	return retailer, nil
}

func (h *Handler) getDistributorByID(distributorID uuid.UUID) (*models.SupplierModel, error) {
	retailer, err := h.supplierService.GetDistributorByID(context.Background(), distributorID)
	if err != nil {
		return nil, err
	}

	return retailer, nil
}

func (h *Handler) getManufacturerByID(manufacturerID uuid.UUID) (*models.SupplierModel, error) {
	retailer, err := h.supplierService.GetManufacturerByID(context.Background(), manufacturerID)
	if err != nil {
		return nil, err
	}

	return retailer, nil
}

func (h *Handler) getShopByID(shopID uuid.UUID) (*models.ShopModel, error) {
	shop, err := h.shopService.GetByID(context.Background(), shopID)
	if err != nil {
		return nil, err
	}

	return shop, nil
}

func (h *Handler) getPromotionByID(promotionID uuid.UUID) (*models.PromotionModel, error) {
	if promotionID == uuid.Nil {
		return nil, errs.ErrPromotionDoesNotExists
	}

	promotion, err := h.promotionService.GetByID(context.Background(), promotionID)
	if err != nil {
		return nil, err
	}

	return promotion, nil
}

func (h *Handler) copySalesToDTO(sales []*models.SaleProductModel) ([]*dto.SaleProductDTO, error) {
	var salesDTO []*dto.SaleProductDTO

	for _, saleProduct := range sales {
		saleDTO, err := h.copySaleToDTO(saleProduct)
		if err != nil {
			return nil, err
		}
		salesDTO = append(salesDTO, saleDTO)
	}

	return salesDTO, nil
}

func (h *Handler) copySaleToDTO(sale *models.SaleProductModel) (*dto.SaleProductDTO, error) {
	shop, err := h.getShopByID(sale.ShopID)
	if err != nil {
		return nil, err
	}
	var saleProduct dto.SaleProductDTO
	saleProduct.ShopTitle = shop.Title
	saleProduct.ShopAddress = shop.Address

	promotion, err := h.getPromotionByID(sale.PromotionID)
	if err != nil && !errors.Is(err, errs.ErrPromotionDoesNotExists) {
		return nil, err
	}

	if promotion == nil {
		saleProduct.PromotionType = "Нет акции"
		saleProduct.PromotionDescription = "Нет акции"
		saleProduct.PromotionDiscountSize = nil
	} else {
		saleProduct.PromotionType = promotion.Type
		saleProduct.PromotionDescription = promotion.Description
		saleProduct.PromotionDiscountSize = &promotion.DiscountSize
	}

	saleProduct.Price = sale.Price
	saleProduct.Currency = sale.Currency
	saleProduct.SettingDate = sale.SettingDate
	saleProduct.AvgRating = sale.AvgRating

	return &saleProduct, nil
}
