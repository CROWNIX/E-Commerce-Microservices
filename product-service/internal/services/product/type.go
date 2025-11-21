package product

import (
	"github.com/CROWNIX/go-utils/databases"
	"github.com/CROWNIX/go-utils/utils/primitive"
)

type GetProductsInput struct {
	Pagination primitive.PaginationInput
	Sorting    primitive.Sorting
	CategoryID *uint64
}

type GetProduct struct {
	ID         uint64
	Name       string
	Images     databases.JSON[[]string]
	Price      uint64
	FinalPrice uint64
}

type GetProductsOutput struct {
	PaginationOutput primitive.PaginationOutput
	Items            []GetProduct
}

type GetDetailProductOutput struct {
	ID              uint64
	Name            string
	Images          databases.JSON[[]string]
	Description     string
	Price           uint64
	Stock           uint32
	FinalPrice      uint64
	DiscountPercent uint8
	MinimumPurchase uint8
	MaximumPurchase *uint8
}
