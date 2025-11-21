package cart

import (
	"context"
)

type CartServiceInterfaces interface {
	CreateCart(context.Context, CreateCartInput) error
	DeleteCart(context.Context, uint64, uint64) error
}
