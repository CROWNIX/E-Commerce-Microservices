package provider

import (
	"github.com/CROWNIX/E-Commerce-Microservices/api-gateway/pkg/middleware"
	"github.com/CROWNIX/E-Commerce-Microservices/api-gateway/pkg/utils/jwtutils"
)

var (
	jwtUtil        jwtutils.JwtUtil
	authMiddleware *middleware.AuthMiddleware
)

func BootstrapGlobal() {
	jwtUtil = jwtutils.NewJwtUtil()
	authMiddleware = middleware.NewAuthMiddleware(jwtUtil)
}
