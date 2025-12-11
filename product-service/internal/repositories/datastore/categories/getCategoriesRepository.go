package categories

import (
	"product-service/internal/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
)

func (r *categoryRepository) GetCategories(ctx context.Context, input GetCategoriesInput) (output GetCategoriesOutput, err error) {
	query := r.DB.Sq.
		Select(
			models.CategoryField.ID,
			models.CategoryField.Name,
			models.CategoryField.Image,
			models.CategoryField.ParentID).
		From(models.CategoryTableName).
		Where(squirrel.Expr(fmt.Sprintf("%s IS NULL", models.CategoryField.ParentID))).
		Where(squirrel.Expr(fmt.Sprintf("%s IS NULL", models.CategoryField.DeletedAt)))

	query = input.Sorting.BuildSquirrel(query)

	err = r.DB.RDBMS.QuerySq(ctx, query, true, func(rows *sql.Rows) error {
		for rows.Next() {
			var category models.Category
			if err := dbscan.ScanRow(&category, rows); err != nil {
				return err
			}
			output.Items = append(output.Items, category)
		}

		return nil
	})

	if err != nil {
		return output, err
	}

	return
}