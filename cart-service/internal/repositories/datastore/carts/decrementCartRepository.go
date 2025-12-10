package carts

import (
	"cart-service/internal/models"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (r *cartRepository) DecrementCart(ctx context.Context, userID uint64, productID uint64) (err error) {
	query := r.DB.Sq.
		Update(models.CartTableName).
		Set(models.CartField.Quantity, squirrel.Expr(fmt.Sprintf("%s - 1", models.CartField.Quantity))).
		Where(squirrel.Eq{
			models.CartField.UserID:    userID,
			models.CartField.ProductID: productID,
		})

	_, err = r.DB.RDBMS.ExecSq(ctx, query, true)

	if err != nil {
		return err
	}

	return nil
}