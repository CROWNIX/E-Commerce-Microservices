package category

import (
	"category-service/internal/models"
	"category-service/internal/repositories/datastore/categories"
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

func (s *categoryService) GetParentCategory(ctx context.Context, categoryID uint64) (output GetParentCategoryOutput, err error) {
	parentCategoryOutput, err := s.categoryRepositoryReader.GetParentCategory(ctx, categoryID)

	if err != nil {
		return output, err
	}

	output = GetParentCategoryOutput{
		ID:   parentCategoryOutput.ID,
		Name: parentCategoryOutput.Name,
		Children: generic.TransformSlice(parentCategoryOutput.Children, func(children categories.GetCategoryChildren) GetCategoryChildren {
			return GetCategoryChildren{
				ID:   children.ID,
				Name: children.Name,
			}
		}),
	}

	return
}
