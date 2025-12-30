package product

import "context"

type ProductServiceInterfaces interface {
	GetDetailProduct(context.Context, uint64) (GetDetailProductOutput, error)
}