package services

import (
	"product-service/internal/services/product"
)

type Service struct {
	ProductService  product.ProductServiceInterfaces
}

func NewService(
	productService product.ProductServiceInterfaces,
) *Service {
	return &Service{
		ProductService:  productService,
	}
}
