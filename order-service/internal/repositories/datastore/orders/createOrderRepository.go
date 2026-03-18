package orders

import (
	"context"

	utilSql "github.com/CROWNIX/go-utils/databases/sqlx"
	"order-service/internal/models"
)

func (r *orderRepository) CreateOrder(ctx context.Context, tx utilSql.RDBMS, input CreateOrderInput) (id uint64, err error){
	query := r.DB.Sq. 
		Insert(models.OrderTableName). 
		Columns(
			models.OrderField.UserID,
			models.OrderField.AddressID,
			models.OrderField.PaymentMethodID,
			models.OrderField.Status,
			models.OrderField.PaymentStatus,
		). 
		Values(
			input.UserID,
			input.AddressID,
			input.PaymentMethodID,
			input.Status,
			input.PaymentStatus,
		)

	result, err := tx.ExecSq(ctx, query, true)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}