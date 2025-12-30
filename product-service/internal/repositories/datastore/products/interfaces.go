package products

import "context"

type ProductRepositoryReaderInterfaces interface {
	GetProducts(context.Context, GetProductsInput) (GetProductsOutput, error)
	GetDetailProduct(context.Context, uint64) (GetDetailProductOutput, error)
	GetStockForUpdate(context.Context, uint64) (uint32, error)
}

type ProductRepositoryWriterInterfaces interface {
}