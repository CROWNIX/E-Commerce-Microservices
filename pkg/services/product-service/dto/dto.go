package dto 

import (
	"github.com/CROWNIX/go-utils/databases"
)

type GetDetailProductOutput struct {
	ID              uint64                   `db:"id"`
	Name            string                   `db:"name"`
	Images          databases.JSON[[]string] `db:"images"`
	Description     string                   `db:"description"`
	Price           uint64                   `db:"price"`
	Stock           uint32                   `db:"stock"`
	FinalPrice      uint64                   `db:"final_price"`
	DiscountPercent uint8                    `db:"discount_percent"`
	MinimumPurchase uint8                    `db:"minimum_purchase"`
	MaximumPurchase *uint8                   `db:"maximum_purchase"`
}