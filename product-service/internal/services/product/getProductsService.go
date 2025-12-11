package product

import (
	"context"
	"product-service/internal/repositories/datastore/products"

	"github.com/CROWNIX/go-utils/utils/generic"
)

func (s *productService) GetProducts(ctx context.Context, input GetProductsInput) (output GetProductsOutput, err error) {
	productsOutput, err := s.productRepositoryReader.GetProducts(ctx, products.GetProductsInput{
		Pagination: input.Pagination,
		Sorting:    input.Sorting,
		CategoryID: input.CategoryID,
	})

	if err != nil {
		return output, err
	}

	output = GetProductsOutput{
		PaginationOutput: productsOutput.PaginationOutput,
		Items: generic.TransformSlice(productsOutput.Items, func(product products.GetProduct) GetProduct {
			return GetProduct(product)
		}),
	}

	return
}