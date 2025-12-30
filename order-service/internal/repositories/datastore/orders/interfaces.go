package orders

import "context"

type OrderRepositoryReaderInterfaces interface {
}

type OrderRepositoryWriterInterfaces interface {
	CreateOrder(context.Context, CreateOrderInput) (uint64, error)
	CreateOrderDetail(context.Context, CreateOrderDetailInput) error
}