package categories

import "context"

type CategoryRepositoryReaderInterfaces interface {
	GetCategories(context.Context, GetCategoriesInput) (GetCategoriesOutput, error)
	GetParentCategory(context.Context, uint64) (GetParentCategoryOutput, error)
}

type CategoryRepositoryWriterInterfaces interface {
}