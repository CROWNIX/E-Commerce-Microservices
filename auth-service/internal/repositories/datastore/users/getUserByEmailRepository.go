package users

import (
	"auth-service/internal/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/CROWNIX/go-utils/databases"
	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
)

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (output GetUserOutput, err error) {
    query := r.DB.Sq.
        Select(
            models.UserField.ID,
            models.UserField.Username,
            models.UserField.Email,
            models.UserField.Password,
        ).
        From(models.UserTableName).
        Where(squirrel.Eq{models.UserField.Email: email}).
        Where(squirrel.Expr(fmt.Sprintf("%s IS NULL", models.UserField.DeletedAt)))

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
