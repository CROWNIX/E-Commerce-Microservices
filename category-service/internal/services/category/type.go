package category

import (
	"category-service/internal/models"

	"github.com/CROWNIX/go-utils/utils/primitive"
)

type GetCategoriesInput struct {
	Sorting primitive.Sorting
}

type GetCategoriesOutput struct {
	Items []models.Category
}

type GetCategoryChildren struct {
	ID   uint64
	Name string
}

type GetParentCategoryOutput struct {
	ID       uint64
	Name     string
	Children []GetCategoryChildren
}
