package product

import (
	"context"
)

type ProductServiceInterfaces interface {
	GetProducts(context.Context, GetProductsInput) (GetProductsOutput, error)
	GetDetailProduct(context.Context, uint64) (GetDetailProductOutput, error)
}
