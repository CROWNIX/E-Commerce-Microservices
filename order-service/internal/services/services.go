package services

import "order-service/internal/services/order"

type Service struct {
	OrderService order.OrderServiceInterfaces
}

func NewService(
	OrderService order.OrderServiceInterfaces,
) *Service {
	return &Service{
		OrderService: OrderService,
	}
}
