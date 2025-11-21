package services

import (
	"category-service/internal/services/category"
)

type Service struct {
	CategoryService category.CategoryServiceInterfaces
}

func NewService(
	categoryService category.CategoryServiceInterfaces,
) *Service {
	return &Service{
		CategoryService: categoryService,
	}
}
