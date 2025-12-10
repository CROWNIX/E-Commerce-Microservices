package carts

import (
	"cart-service/internal/models"
	"context"

	"github.com/Masterminds/squirrel"
)

func (r *cartRepository) CreateCart(ctx context.Context, input CreateCartInput) (err error) {
	query := r.DB.Sq.
		Insert(models.CartTableName).
		Columns(
			models.CartField.UserID,
			models.CartField.ProductID,
			models.CartField.Quantity,
		).
		Values(
			input.UserID,
			input.ProductID,
			input.Quantity,
		)

	_, err = r.DB.RDBMS.ExecSq(ctx, query, true)
	if err != nil {
		return err
	}

	return nil
}

func (r *cartRepository) DeleteCart(ctx context.Context, cartID uint64) (err error) {
	query := r.DB.Sq.
		Delete(models.CartTableName).
		Where(squirrel.Eq{
			models.CartField.ID: cartID,
		})

	_, err = r.DB.RDBMS.ExecSq(ctx, query, true)
	if err != nil {
		return err
	}

	return nil
}
