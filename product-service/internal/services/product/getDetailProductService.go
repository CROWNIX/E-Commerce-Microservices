package product

import (
	"context"
	"errors"

	"github.com/CROWNIX/go-utils/apperror"
	"github.com/CROWNIX/go-utils/databases"
)

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
