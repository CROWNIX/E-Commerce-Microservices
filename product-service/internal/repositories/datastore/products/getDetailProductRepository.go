package products

import (
	"context"
	"database/sql"
	"fmt"
	"product-service/internal/models"

	"github.com/CROWNIX/go-utils/databases"
	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
)

func (r *productRepository) GetDetailProduct(ctx context.Context, productID uint64) (output GetDetailProductOutput, err error) {
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
		Where(squirrel.Eq{models.ProductField.ID: productID}).
		Where(squirrel.Expr(fmt.Sprintf("%s IS NULL", models.ProductField.DeletedAt)))

	err = r.DB.RDBMS.QuerySq(ctx, query, true, func(rows *sql.Rows) error {
		if rows.Next() {
			return dbscan.ScanRow(&output, rows)
		}
		return nil
	})

	if err != nil {
		return output, err
	}

	if output.ID == 0 {
		return output, databases.ErrNoRowFound
	}

	return output, nil
}
