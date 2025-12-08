package controller

import (
	"github.com/CROWNIX/E-Commerce-Microservices/api-gateway/internal/config"
	"github.com/CROWNIX/E-Commerce-Microservices/api-gateway/internal/proxy"
	"github.com/CROWNIX/E-Commerce-Microservices/api-gateway/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type GatewayController struct {
	authMiddleware *middleware.AuthMiddleware
	cfg            *config.Config
}

func NewGatewayController(authMiddleware *middleware.AuthMiddleware, cfg *config.Config) *GatewayController {
	return &GatewayController{
		authMiddleware: authMiddleware,
		cfg:            cfg,
	}
}

func (c *GatewayController) Route(r *gin.Engine) {
	// PROXIES
	authProxy := proxy.NewReverseProxy(c.cfg.AuthURL, "/v1")
	productProxy := proxy.NewReverseProxy(c.cfg.ProductServicetURL, "/v1/products")
	categoryProxy := proxy.NewReverseProxy(c.cfg.CategoryServiceURL, "/v1")
	cartProxy := proxy.NewReverseProxy(c.cfg.CartServicetURL, "/v1")

	r.POST("/login", authProxy)
	r.POST("/register", authProxy)
	r.Any("/products/*path", productProxy)
	r.POST("/carts", c.authMiddleware.JwtMiddlewareUsingRedis, cartProxy)
	r.GET("/categories", categoryProxy)
}
