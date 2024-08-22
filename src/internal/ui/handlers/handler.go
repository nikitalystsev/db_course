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
	shopService        intf.IShopService
	supplierService    intf.ISupplierService
	userService        intf.IUserService
	tokenManager       auth.ITokenManager
	accessTokenTTL     time.Duration
	refreshTokenTTL    time.Duration
}

func NewHandler(
	//productService intf.IProductService,
	//saleProductService intf.ISaleProductService,
	//shopService intf.IShopService,
	//supplierService intf.ISupplierService,
	userService intf.IUserService,
	tokenManager auth.ITokenManager,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
) *Handler {
	return &Handler{
		//productService:     productService,
		//saleProductService: saleProductService,
		//shopService:        shopService,
		//supplierService:    supplierService,
		userService:     userService,
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
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
