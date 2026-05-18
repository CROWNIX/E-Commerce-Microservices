package users

import (
	"auth-service/internal/models"
	"context"
)

func (r *userRepository) CreateUser(ctx context.Context, input CreateUserInput) (err error) {
    query := r.DB.Sq.
        Insert(models.UserTableName).
        Columns(
            models.UserField.Username,
            models.UserField.Email,
            models.UserField.Password,
        ).
        Values(
            input.Username,
            input.Email,
            input.Password,
        )

    _, err = r.DB.RDBMS.ExecSq(ctx, query, true)
    if err != nil {
        return err
    }

    return nil
}
