package products

import (
	"context"

	productServiceDto "pkg/services/product-service/dto"
)

type ProductServiceInterfaces interface {
	GetDetailProduct(context.Context, uint64) (productServiceDto.GetDetailProductOutput, error)
	CountProductByIds(ctx context.Context, productIds []uint64) (uint32, error)
	GetProductByIds(ctx context.Context, productIds []uint64) ([]productServiceDto.GetDetailProductOutput, error)
}
