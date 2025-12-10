package carts

import (
	"cart-service/internal/models"
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
)

func (r *cartRepository) GetQuantityCartByUserAndProductId(ctx context.Context, userID uint64, productID uint64) (quantity uint8, err error) {
	query := r.DB.Sq.
		Select(
			models.CartField.Quantity,
		).
		From(models.CartTableName).
		Where(squirrel.Eq{
			models.CartField.UserID:    userID,
			models.CartField.ProductID: productID,
		})

	err = r.DB.RDBMS.QuerySq(ctx, query, true, func(rows *sql.Rows) error {
		if rows.Next() {
			return dbscan.ScanRow(&quantity, rows)
		}

		return nil
	})

	if err != nil {
		return quantity, err
	}

	return quantity, nil
}