package carts

import "context"

type CartRepositoryReaderInterfaces interface {
	CountCartByUserAndProductId(context.Context, uint64, uint64) (uint8, error)
	GetQuantityCartByUserAndProductId(context.Context, uint64, uint64) (uint8, error)
}

type CartRepositoryWriterInterfaces interface {
	CreateCart(context.Context, CreateCartInput) error
	DeleteCart(context.Context, uint64, uint64) error
	IncrementCart(context.Context, uint64, uint64) error
}