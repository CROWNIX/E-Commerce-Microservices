package carts

import (
	"cart-service/internal/models"
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
)

func (r *cartRepository) CountCartByUserAndProductId(ctx context.Context, userID uint64, productID uint64) (total uint8, err error) {
	query := r.DB.Sq.
		Select(
			"COUNT(" + models.CartField.ID + ") AS total",
		).
		From(models.CartTableName).
		Where(squirrel.Eq{models.CartField.UserID: userID}).
		Where(squirrel.Eq{models.CartField.ProductID: productID})

	err = r.DB.RDBMS.QuerySq(ctx, query, true, func(rows *sql.Rows) error {
		if rows.Next() {
			return dbscan.ScanRow(&total, rows)
		}

		return nil
	})

	if err != nil {
		return total, err
	}

	return total, nil
}

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

func (r *cartRepository) DeleteCart(ctx context.Context, userID uint64, productID uint64) (err error) {
	query := r.DB.Sq.
		Delete(models.CartTableName).
		Where(squirrel.Eq{
			models.CartField.UserID: userID,
			models.CartField.ProductID: productID,
		})

	_, err = r.DB.RDBMS.ExecSq(ctx, query, true)
	if err != nil {
		return err
	}

	return nil
}
