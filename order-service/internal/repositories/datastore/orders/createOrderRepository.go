package orders

import (
	"context"
	"order-service/internal/models"
)

func (r *orderRepository) CreateOrder(ctx context.Context, input CreateOrderInput) (id uint64, err error){
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

	result, err := r.DB.RDBMS.ExecSq(ctx, query, true)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}