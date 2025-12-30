package products

import (
	"context"
	"database/sql"
	"fmt"
	"product-service/internal/models"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
)

func (r *productRepository) GetStockForUpdate(ctx context.Context, productID uint64) (stock uint32, err error) {
	query := r.DB.Sq.
		Select(
			models.ProductField.Stock,
		).
		From(models.ProductTableName).
		Where(squirrel.Eq{models.ProductField.ID: productID}).
		Where(squirrel.Expr(fmt.Sprintf("%s IS NULL", models.ProductField.DeletedAt))). 
		Suffix("FOR UPDATE")

	err = r.DB.RDBMS.QuerySq(ctx, query, true, func(rows *sql.Rows) error {
		if rows.Next() {
			return dbscan.ScanRow(&stock, rows)
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return stock, nil
}
