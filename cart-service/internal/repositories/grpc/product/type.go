package product

import (
	"github.com/CROWNIX/go-utils/databases"
)

type GetProduct struct {
	ID         uint64                   `db:"id"`
	Name       string                   `db:"name"`
	Images     databases.JSON[[]string] `db:"images"`
	Price      uint64                   `db:"price"`
	FinalPrice uint64                   `db:"final_price"`
}
