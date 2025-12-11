package orders

import "context"

type OrderRepositoryReaderInterfaces interface {
}

type OrderRepositoryWriterInterfaces interface {
	CreateOrder(context.Context, CreateOrderInput) (error)
}