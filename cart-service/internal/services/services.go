package services

import (
	"cart-service/internal/services/cart"
)

type Service struct {
	CartService     cart.CartServiceInterfaces
}

func NewService(
	cartService cart.CartServiceInterfaces,
) *Service {
	return &Service{
		CartService:     cartService,
	}
}
