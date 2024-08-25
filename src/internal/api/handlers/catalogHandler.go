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

	sales, err := h.saleProductService.GetByProductID(c.Request.Context(), productID)
	if err != nil && errors.Is(err, errs.ErrSaleProductDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	salesDTO, err := h.copySalesToDTO(sales)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
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

	certificates, err := h.certificateService.GetByProductID(c.Request.Context(), productID)
	if err != nil && errors.Is(err, errs.ErrCertificateDoesNotExists) {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, certificates)
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
