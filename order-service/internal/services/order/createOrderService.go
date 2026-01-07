package order

import (
	"context"
	"database/sql"
	"fmt"

	productServiceDto "pkg/services/product-service/dto"

	"github.com/CROWNIX/go-utils/apperror"
	"github.com/CROWNIX/go-utils/databases/sqlx"
	"github.com/CROWNIX/go-utils/utils/generic"
)

func (s *orderService) CreateOrder(ctx context.Context, input CreateOrderServiceInput) (output *CreateOrderServiceOutput, err error) {
	productIds := generic.TransformSlice(input.Items, func(item Item) uint64 {
		return item.ProductId
	})

	err = s.tx.DoTxContext(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted}, func(ctx context.Context, tx sqlx.RDBMS) error {
		products, err := s.productService.GetProductByIds(ctx, productIds)
		if err != nil {
			return err
		}

		if len(input.Items) != len(products) {
			return apperror.NotFound("One of product not found")
		}

		subTotal, err := calculateTotalPriceAndValidateOrder(input, products)
		if err != nil {
			return err
		}

		if input.GrandTotal != subTotal {
			return apperror.BadRequest("Grand total does not match the calculated total")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return
}

func calculateTotalPriceAndValidateOrder(input CreateOrderServiceInput, products []productServiceDto.GetDetailProductOutput) (subTotal uint64, err error) {
	productOrders := map[uint64]Item{}

	for _, item := range input.Items {
		productOrders[item.ProductId] = item
	}

	for _, product := range products {
		productOrder := productOrders[product.ID]

		if productOrder.Quantity > product.Stock {
			return 0, apperror.BadRequest(fmt.Sprintf("Unable to create order for product '%s': The product is out of stock", product.Name))
		}

		if product.MaximumPurchase != nil {
			if productOrder.Quantity < uint32(product.MinimumPurchase) || productOrder.Quantity > uint32(*product.MaximumPurchase) {
				return 0, apperror.BadRequest(fmt.Sprintf("Unable to create order for product '%s': The quantity must be between the minimum purchase and the maximum purchase", product.Name))
			}
		} else {
			if productOrder.Quantity < uint32(product.MinimumPurchase) {
				return 0, apperror.BadRequest(fmt.Sprintf("Unable to create order for product '%s': Quantity cannot be less than minimum purchase", product.Name))
			}
		}

		subTotal += uint64(productOrder.Quantity) * product.Price

	}
	return subTotal, nil
}
