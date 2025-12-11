package category

import (
	"product-service/internal/models"
	"product-service/internal/repositories/datastore/categories"
	"context"

	"github.com/CROWNIX/go-utils/utils/generic"
)

func (s *categoryService) GetCategories(ctx context.Context, input GetCategoriesInput) (output GetCategoriesOutput, err error) {
	categoryOutput, err := s.categoryRepositoryReader.GetCategories(ctx, categories.GetCategoriesInput{
		Sorting: input.Sorting,
	})

	if err != nil {
		return output, err
	}

	output = GetCategoriesOutput{
		Items: generic.TransformSlice(categoryOutput.Items, func(category models.Category) models.Category {
			return models.Category(category)
		}),
	}

	return
}