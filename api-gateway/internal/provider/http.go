package provider

import (
	"github.com/CROWNIX/E-Commerce-Microservices/api-gateway/internal/config"
	"github.com/CROWNIX/E-Commerce-Microservices/api-gateway/internal/controller"
	"github.com/gin-gonic/gin"
)

func BootstrapHttp(cfg *config.Config, router *gin.Engine) {
	appController := controller.NewAppController()
	gatewayController := controller.NewGatewayController(authMiddleware, cfg)

	appController.Route(router)
	gatewayController.Route(router)
}
