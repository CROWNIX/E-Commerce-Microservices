package order

import "context"

type OrderServiceInterfaces interface {
	CreateOrder(context.Context, CreateOrderServiceInput) (CreateOrderServiceOutput, error)
}
