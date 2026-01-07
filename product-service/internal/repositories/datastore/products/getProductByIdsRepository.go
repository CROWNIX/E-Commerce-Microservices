package products

import (
	"context"
	"database/sql"
	"fmt"
	"product-service/internal/models"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
)

func (r *productRepository) GetProductByIds(ctx context.Context, productIds []uint64) (products []GetProductByIdsOutput, err error) {
	query := r.DB.Sq.	
		Select(
			models.ProductField.ID,
			models.ProductField.Name,
			models.ProductField.Images,
			models.ProductField.Description,
			models.ProductField.Price,
			models.ProductField.Stock,
			models.ProductField.FinalPrice,
			models.ProductField.DiscountPercent,
			models.ProductField.MinimumPurchase,
			models.ProductField.MaximumPurchase,
		).
		From(models.ProductTableName).
		Where(squirrel.Eq{models.ProductField.ID: productIds}).
		Where(squirrel.Expr(fmt.Sprintf("%s IS NULL", models.ProductField.DeletedAt))).
		Suffix("FOR UPDATE")

	err = r.DB.RDBMS.QuerySq(ctx, query, true, func(rows *sql.Rows) error {
		if rows.Next() {
			return dbscan.ScanRow(&products, rows)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return products, nil
}
