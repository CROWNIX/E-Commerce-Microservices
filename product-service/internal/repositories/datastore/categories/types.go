package categories

import (
	"product-service/internal/models"

	"github.com/CROWNIX/go-utils/utils/primitive"
)

type GetCategoriesInput struct {
	Sorting primitive.Sorting
}

type GetCategoriesOutput struct {
	Items []models.Category
}

type GetCategoryChildren struct {
	ID   uint64 `db:"id"`
	Name string `db:"name"`
}

type GetParentCategoryOutput struct {
	ID       uint64                `db:"id"`
	Name     string                `db:"name"`
	ParentID *uint64               `db:"parent_id"`
	Children []GetCategoryChildren `db:"children"`
}
