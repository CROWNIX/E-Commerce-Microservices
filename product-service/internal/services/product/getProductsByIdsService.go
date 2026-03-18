package product

import (
	"context"
	"fmt"
)

func (s *productService) GetProductsByIds(ctx context.Context, ids []uint64) ([]GetDetailProductOutput, error) {
	productsOutput, err := s.productRepositoryReader.GetProductByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	result := make([]GetDetailProductOutput, 0, len(productsOutput))
	for _, p := range productsOutput {
		result = append(result, GetDetailProductOutput{
			ID:              p.ID,
			Name:            p.Name,
			Images:          p.Images,
			Description:     p.Description,
			Price:           p.Price,
			Stock:           p.Stock,
			FinalPrice:      p.FinalPrice,
			DiscountPercent: p.DiscountPercent,
			MinimumPurchase: p.MinimumPurchase,
			MaximumPurchase: p.MaximumPurchase,
		})
	}

	return result, nil
}
