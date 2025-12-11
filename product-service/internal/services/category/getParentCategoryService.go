package category

import (
	"product-service/internal/repositories/datastore/categories"
	"context"

	"github.com/CROWNIX/go-utils/utils/generic"
)

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
