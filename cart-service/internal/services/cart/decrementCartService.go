package cart

import (
	"context"
	"errors"

	"github.com/CROWNIX/go-utils/apperror"
	"github.com/CROWNIX/go-utils/databases"
)

func (s *cartService) DecrementCart(ctx context.Context, userID uint64, productID uint64) (err error) {
	total, err := s.cartRepositoryReader.CountCartByUserAndProductId(ctx, userID, productID)
	if err != nil {
		return err
	}

	if total <= 0 {
		return apperror.NotFound("Cart not found")
	}

	product, err := s.productService.GetDetailProduct(ctx, productID)
	if err != nil {
		if errors.Is(err, databases.ErrNoRowFound) {
			return apperror.NotFound("Product not found")
		}

		return err
	}

	quantity, err := s.cartRepositoryReader.GetQuantityCartByUserAndProductId(ctx, userID, productID)
	if err != nil {
		return err
	}

	if product.MinimumPurchase == quantity {
		return apperror.BadRequest("Cannot decrement cart. Reached minimum purchase limit")
	}

	err = s.cartRepositoryWriter.DecrementCart(ctx, userID, productID)
	if err != nil {
		return err
	}

	return nil
}
