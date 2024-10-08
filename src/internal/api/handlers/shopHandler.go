package handlers

import (
	"SmartShopper-services/core/dto"
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (h *Handler) getSalesByShopID(c *gin.Context) {
	shopIDStr := c.Query("shop_id")
	shopID, err := uuid.Parse(shopIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	sales, err := h.saleProductService.GetByShopID(c.Request.Context(), shopID)
	if err != nil && errors.Is(err, errs.ErrSaleProductDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	salesDTO, err := h.copySalesToShopDTO(sales)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, salesDTO)
}

func (h *Handler) updateSaleProductPriceByID(c *gin.Context) {
	saleProductIDStr := c.Param("id")
	saleProductID, err := uuid.Parse(saleProductIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var newPrice float32
	if err = c.BindJSON(&newPrice); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	sale, err := h.saleProductService.GetByID(c.Request.Context(), saleProductID)
	if err != nil && errors.Is(err, errs.ErrSaleProductDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	sale.Price = newPrice

	err = h.saleProductService.Update(c.Request.Context(), sale)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) addSaleProductInShop(c *gin.Context) {
	var newSaleProductDTO dto.NewSaleProductDTO
	if err := c.BindJSON(&newSaleProductDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err := h.saleProductService.Create(c.Request.Context(), &newSaleProductDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) getProductForShopByID(productID uuid.UUID) (*models.ProductModel, error) {
	product, err := h.productService.GetByID(context.Background(), productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (h *Handler) copySalesToShopDTO(sales []*models.SaleProductModel) ([]*dto.SaleProductShopDTO, error) {
	var salesDTO []*dto.SaleProductShopDTO

	for _, saleProduct := range sales {
		saleDTO, err := h.copySaleToShopDTO(saleProduct)
		if err != nil {
			return nil, err
		}
		salesDTO = append(salesDTO, saleDTO)
	}

	return salesDTO, nil
}

func (h *Handler) copySaleToShopDTO(sale *models.SaleProductModel) (*dto.SaleProductShopDTO, error) {
	product, err := h.getProductForShopByID(sale.ProductID)
	if err != nil {
		return nil, err
	}
	var saleProduct dto.SaleProductShopDTO
	saleProduct.ID = sale.ID
	saleProduct.ProductName = product.Name
	saleProduct.ProductCategories = product.Categories

	promotion, err := h.getPromotionByID(sale.PromotionID)
	if err != nil && !errors.Is(err, errs.ErrPromotionDoesNotExists) {
		return nil, err
	}

	if promotion == nil {
		saleProduct.PromotionType = "Нет акции"
		saleProduct.PromotionDiscountSize = nil
	} else {
		saleProduct.PromotionType = promotion.Type
		saleProduct.PromotionDiscountSize = &promotion.DiscountSize
	}

	saleProduct.Price = sale.Price
	saleProduct.Currency = sale.Currency
	saleProduct.SettingDate = sale.SettingDate
	saleProduct.AvgRating = sale.AvgRating
	saleProduct.CertificatesStatistic, err = h.getCertificateStatisticsByProductID(product.ID)
	if err != nil {
		return nil, err
	}
	
	return &saleProduct, nil
}
