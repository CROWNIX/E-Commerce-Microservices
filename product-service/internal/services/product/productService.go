package product

import (
	"context"
	"errors"
	"product-service/internal/repositories/datastore/products"

	"github.com/CROWNIX/go-utils/apperror"
	"github.com/CROWNIX/go-utils/databases"
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

func (s *productService) GetDetailProduct(ctx context.Context, productID uint64) (output GetDetailProductOutput, err error) {
	detailProductOutput, err := s.productRepositoryReader.GetDetailProduct(ctx, productID)

	if err != nil {
		if errors.Is(err, databases.ErrNoRowFound) {
			return output, apperror.NotFound("Product not found")
		}
		return output, err
	}

	output = GetDetailProductOutput{
		ID:              detailProductOutput.ID,
		Name:            detailProductOutput.Name,
		Images:          detailProductOutput.Images,
		Description:     detailProductOutput.Description,
		Price:           detailProductOutput.Price,
		Stock:           detailProductOutput.Stock,
		FinalPrice:      detailProductOutput.FinalPrice,
		DiscountPercent: detailProductOutput.DiscountPercent,
		MinimumPurchase: detailProductOutput.MinimumPurchase,
		MaximumPurchase: detailProductOutput.MaximumPurchase,
	}

	return
}
