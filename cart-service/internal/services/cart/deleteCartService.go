package cart

import (
	"context"

	"github.com/CROWNIX/go-utils/apperror"
)

func (c *cartService) DeleteCart(ctx context.Context, userID uint64, productID uint64) (err error) {
	total, err := c.cartRepositoryReader.CountCartByUserAndProductId(ctx, userID, productID)
	if err != nil {
		return err
	}

	if total <= 0 {
		return apperror.NotFound("Cart not found")
	}

	err = c.cartRepositoryWriter.DeleteCart(ctx, userID, productID)
	if err != nil {
		return err
	}

	return nil
}
