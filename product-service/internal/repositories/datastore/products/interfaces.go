package products

import (
	"context"
)

type ProductRepositoryReaderInterfaces interface {
	GetProducts(context.Context, GetProductsInput) (GetProductsOutput, error)
	GetDetailProduct(context.Context, uint64) (GetDetailProductOutput, error)
	GetStockForUpdate(context.Context, uint64) (uint32, error)
	CountProductByIds(context.Context, []uint64) (uint32, error)
	GetProductByIds(context.Context, []uint64) ([]GetProductByIdsOutput, error)
}

type ProductRepositoryWriterInterfaces interface {
}