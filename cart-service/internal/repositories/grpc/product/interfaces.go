package product

import (
	"context"

	productServiceDto "pkg/services/product-service/dto"
)

type ProductServiceInterfaces interface {
	GetDetailProduct(context.Context, uint64) (productServiceDto.GetDetailProductOutput, error)
}
