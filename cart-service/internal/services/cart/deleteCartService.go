package cart

import (
	"context"

	"github.com/CROWNIX/go-utils/apperror"
)

func (c *cartService) DeleteCart(ctx context.Context, userID uint64, cartID uint64) (err error) {
	total, err := c.cartRepositoryReader.CountCartByUserAndProductId(ctx, userID, cartID)
	if err != nil {
		return err
	}

	if total <= 0 {
		return apperror.NotFound("Cart not found")
	}

	err = c.cartRepositoryWriter.DeleteCart(ctx, cartID)
	if err != nil {
		return err
	}

	return nil
}
