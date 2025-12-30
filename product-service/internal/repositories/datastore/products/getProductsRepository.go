package products

import (
	"context"
	"database/sql"
	"fmt"
	"product-service/internal/models"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
)

func (r *productRepository) GetProducts(ctx context.Context, input GetProductsInput) (output GetProductsOutput, err error) {
	selectQuery := r.DB.Sq.
		Select(
			models.ProductField.ID,
			models.ProductField.Name,
			models.ProductField.Images,
			models.ProductField.Price,
			models.ProductField.FinalPrice,
		).
		From(models.ProductTableName).
		Where(squirrel.Expr(fmt.Sprintf("%s IS NULL", models.ProductField.DeletedAt)))

	countQuery := r.DB.Sq.Select("COUNT(" + models.ProductField.ID + ") AS total").
		From(models.ProductTableName).
		Where(squirrel.Expr(fmt.Sprintf("%s IS NULL", models.ProductField.DeletedAt)))

	if input.CategoryID != nil {
		var parentID *int64
		checkQuery := r.DB.Sq.
			Select(models.CategoryField.ParentID).
			From(models.CategoryTableName).
			Where(squirrel.Eq{models.CategoryField.ID: *input.CategoryID}).
			Where(squirrel.Expr(fmt.Sprintf("%s IS NULL", models.CategoryField.DeletedAt)))

		err = r.DB.RDBMS.QuerySq(ctx, checkQuery, false, func(rows *sql.Rows) error {
			if rows.Next() {
				return rows.Scan(&parentID)
			}
			return nil
		})

		if err != nil {
			return output, err
		}

		var categoryIDs []uint64
		if parentID == nil {
			categoryIDs = append(categoryIDs, *input.CategoryID)

			childQuery := r.DB.Sq.
				Select(models.CategoryField.ID).
				From(models.CategoryTableName).
				Where(squirrel.Eq{models.CategoryField.ParentID: *input.CategoryID}).
				Where(squirrel.Expr(fmt.Sprintf("%s IS NULL", models.CategoryField.DeletedAt)))

			err = r.DB.RDBMS.QuerySq(ctx, childQuery, true, func(rows *sql.Rows) error {
				for rows.Next() {
					var childID uint64
					if err := rows.Scan(&childID); err != nil {
						return err
					}
					categoryIDs = append(categoryIDs, childID)
				}
				return nil
			})
			if err != nil {
				return output, err
			}

		} else {
			categoryIDs = []uint64{*input.CategoryID}
		}

		selectQuery = selectQuery.Where(squirrel.Eq{models.ProductField.CategoryID: categoryIDs})
		countQuery = countQuery.Where(squirrel.Eq{models.ProductField.CategoryID: categoryIDs})
	}

	selectQuery = input.Sorting.BuildSquirrel(selectQuery)

	output.PaginationOutput, err = r.DB.RDBMS.QuerySqPagination(ctx, countQuery, selectQuery, true, input.Pagination, func(rows *sql.Rows) error {
		for rows.Next() {
			var product GetProduct
			if err := dbscan.ScanRow(&product, rows); err != nil {
				return err
			}
			output.Items = append(output.Items, product)
		}
		return nil
	})
	if err != nil {
		return output, err
	}

	return
}