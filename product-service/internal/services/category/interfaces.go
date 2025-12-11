package category

import (
	"context"
)

type CategoryServiceInterfaces interface {
	GetCategories(context.Context, GetCategoriesInput) (GetCategoriesOutput, error)
	GetParentCategory(context.Context, uint64) (GetParentCategoryOutput, error)
}
