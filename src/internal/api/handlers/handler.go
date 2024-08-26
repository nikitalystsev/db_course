package handlers

import (
	"SmartShopper-services/intf"
	"SmartShopper-services/pkg/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

type Handler struct {
	productService     intf.IProductService
	saleProductService intf.ISaleProductService
	promotionService   intf.IPromotionService
	shopService        intf.IShopService
	supplierService    intf.ISupplierService
	userService        intf.IUserService
	ratingService      intf.IRatingService
	certificateService intf.ICertificateService
	tokenManager       auth.ITokenManager
	accessTokenTTL     time.Duration
	refreshTokenTTL    time.Duration
}

func NewHandler(
	productService intf.IProductService,
	saleProductService intf.ISaleProductService,
	promotionService intf.IPromotionService,
	shopService intf.IShopService,
	supplierService intf.ISupplierService,
	userService intf.IUserService,
	ratingService intf.IRatingService,
	certificateService intf.ICertificateService,
	tokenManager auth.ITokenManager,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
) *Handler {
	return &Handler{
		productService:     productService,
		saleProductService: saleProductService,
		promotionService:   promotionService,
		shopService:        shopService,
		supplierService:    supplierService,
		userService:        userService,
		ratingService:      ratingService,
		certificateService: certificateService,
		tokenManager:       tokenManager,
		accessTokenTTL:     accessTokenTTL,
		refreshTokenTTL:    refreshTokenTTL,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	router := gin.Default()

	router.Use(h.corsSettings())

	authenticate := router.Group("/auth")
	{
		authenticate.POST("/sign-up", h.signUp)
		authenticate.POST("/sign-in", h.signIn)
		authenticate.POST("/refresh", h.refresh)
	}

	guest := router.Group("/")
	{
		guest.GET("/products/:id", h.getProductByID)
		guest.GET("/products", h.getProducts)
		guest.GET("/sales", h.getSalesByProductID)
		guest.GET("/certificates", h.getCertificatesByProductID)
	}

	registered := router.Group("/api", h.userIdentity)
	{
		registered.POST("/ratings", h.addSaleProductRating)
		registered.POST("/retailers", h.addRetailerIfNotExist)
		registered.GET("/retailers", h.getRetailerByAddress)
		registered.POST("/shops", h.addShop)
		registered.GET("/shops", h.getShops)
		registered.GET("/sales", h.getSalesByShopID)
		registered.POST("/sales", h.addSaleProductInShop)
		registered.PUT("/sales/:id", h.updateSaleProductPriceByID)
	}

	return router
}

func (h *Handler) corsSettings() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
		},
		AllowOrigins: []string{
			"*",
		},
		AllowCredentials: true,
		AllowHeaders: []string{
			"Authorization",
			"Content-Type",
		},
		ExposeHeaders: []string{
			"Content-Type",
		},
	})
}
