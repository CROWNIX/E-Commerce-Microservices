package orders

import (
	"context"
	"order-service/internal/models"
)

func (r *orderRepository) CreateOrderDetail(ctx context.Context, input CreateOrderDetailInput) (err error){
	query := r.DB.Sq. 
		Insert(models.OrderTableName). 
		Columns(
			models.OrderDetailField.OrderID,
			models.OrderDetailField.ProductID,
			models.OrderDetailField.Quantity,
		)

	for _, item := range input.Items {
		query = query.Values(
			input.OrderID,
			item.ProductId,
			item.Quantity,
		)
	}

	_, err = r.DB.RDBMS.ExecSq(ctx, query, true)
	if err != nil {
		return err
	}

	return nil
}