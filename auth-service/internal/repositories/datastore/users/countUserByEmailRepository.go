package users

import (
	"auth-service/internal/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
)

func (r *userRepository) CountUserByEmail(ctx context.Context, email string) (total uint64, err error) {
    query := r.DB.Sq.
        Select(
            "COUNT(" + models.UserField.Email + ") AS total",
        ).
        From(models.UserTableName).
        Where(squirrel.Eq{models.UserField.Email: email}).
        Where(squirrel.Expr(fmt.Sprintf("%s IS NULL", models.UserField.DeletedAt)))

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
