package services

import (
	"product-service/internal/services/category"
	"product-service/internal/services/product"
)

type Service struct {
	ProductService  product.ProductServiceInterfaces
	CategoryService  category.CategoryServiceInterfaces
}

func NewService(
	productService product.ProductServiceInterfaces,
	categoryService category.CategoryServiceInterfaces,
) *Service {
	return &Service{
		ProductService:  productService,
		CategoryService:  categoryService,
	}
}
