package cart

import (
	"cart-service/internal/repositories/datastore/carts"
	"context"
	"fmt"
	"log/slog"

	"github.com/CROWNIX/go-utils/apperror"
)

func (s *cartService) CreateCart(ctx context.Context, input CreateCartInput) (err error) {
	total, err := s.cartRepositoryReader.CountCartByUserAndProductId(ctx, input.UserId, input.ProductId)
	
	if err != nil {
		slog.Debug(fmt.Sprintf("Failed to count cart by user and product id: %s", err.Error()))
		return apperror.BadRequest("Quantity cannot be less than minimum purchase")
	}
	
	if total > 0 {
		return apperror.Conflict("Product has been added to cart")
	}
	
	product, err := s.productService.GetDetailProduct(ctx, input.ProductId)
	if err != nil {
		slog.Debug(fmt.Sprintf("Failed to get detail product from product service: %s", err.Error()))
		return apperror.InternalServer("Internal Server Error")
	}

	if product.MaximumPurchase != nil {
		if input.Quantity < product.MinimumPurchase || input.Quantity > *product.MaximumPurchase {
			return apperror.BadRequest("The quantity must be between the minimum purchase and the maximum purchase")
		}
	} else {
		if input.Quantity < product.MinimumPurchase {
			return apperror.BadRequest("Quantity cannot be less than minimum purchase")
		}
	}

	err = s.cartRepositoryWriter.CreateCart(ctx, carts.CreateCartInput{
		UserID:    input.UserId,
		ProductID: input.ProductId,
		Quantity:  input.Quantity,
	})

	return
}