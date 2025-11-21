package categories

import (
	"category-service/internal/models"
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

func (r *categoryRepository) GetParentCategory(ctx context.Context, categoryID uint64) (output GetParentCategoryOutput, err error) {
	query := r.DB.Sq.
		Select(
			models.CategoryField.ID,
			models.CategoryField.Name,
			models.CategoryField.ParentID,
		).
		From(models.CategoryTableName).
		Where(squirrel.Eq{models.CategoryField.ID: categoryID}).
		Where(squirrel.Expr(fmt.Sprintf("%s IS NULL", models.CategoryField.DeletedAt)))

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
		return output, nil
	}

	parentID := output.ID

	if output.ParentID != nil {
		parentID = *output.ParentID
		parentQuery := r.DB.Sq.
			Select(
				models.CategoryField.ID,
				models.CategoryField.Name,
			).
			From(models.CategoryTableName).
			Where(squirrel.Eq{models.CategoryField.ID: parentID}).
			Where(squirrel.Expr(fmt.Sprintf("%s IS NULL", models.CategoryField.DeletedAt)))

		err = r.DB.RDBMS.QuerySq(ctx, parentQuery, true, func(rows *sql.Rows) error {
			if rows.Next() {
				return dbscan.ScanRow(&output, rows)
			}
			return nil
		})
		if err != nil {
			return output, err
		}
	}

	childQuery := r.DB.Sq.
		Select(
			models.CategoryField.ID,
			models.CategoryField.Name,
		).
		From(models.CategoryTableName).
		Where(squirrel.Eq{models.CategoryField.ParentID: parentID}).
		Where(squirrel.Expr(fmt.Sprintf("%s IS NULL", models.CategoryField.DeletedAt)))

	err = r.DB.RDBMS.QuerySq(ctx, childQuery, true, func(rows *sql.Rows) error {
		for rows.Next() {
			var children GetCategoryChildren
			if err := dbscan.ScanRow(&children, rows); err != nil {
				return err
			}
			output.Children = append(output.Children, children)
		}
		return nil
	})
	if err != nil {
		return output, err
	}

	return output, nil
}
