package services

import (
	"auth-service/internal/services/auth"
)

type Service struct {
	AuthService auth.AuthServiceInterfaces
}

func NewService(
	authService auth.AuthServiceInterfaces,
) *Service {
	return &Service{
		AuthService: authService,
	}
}
