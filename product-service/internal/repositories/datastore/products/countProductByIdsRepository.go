package products

import (
	"context"
	"database/sql"
	"fmt"
	"product-service/internal/models"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
)

func (r *productRepository) CountProductByIds(ctx context.Context, productIds[]uint64) (total uint32, err error) {
	query := r.DB.Sq.	
		Select(
			"COUNT(id)",
		).
		From(models.ProductTableName).
		
		Where(squirrel.Eq{models.ProductField.ID: productIds}).
		Where(squirrel.Expr(fmt.Sprintf("%s IS NULL", models.ProductField.DeletedAt)))

	err = r.DB.RDBMS.QuerySq(ctx, query, true, func(rows *sql.Rows) error {
		if rows.Next() {
			return dbscan.ScanRow(&total, rows)
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return total, nil
}
