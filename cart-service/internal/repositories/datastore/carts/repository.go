package carts

import (
	"cart-service/internal/infra"

	"github.com/google/wire"
)

type cartRepository struct {
	DB *infra.DB
}

func NewCartRepository(db *infra.DB) *cartRepository {
	return &cartRepository{
		DB: db,
	}
}

var SetWire = wire.NewSet(
	NewCartRepository,
	wire.Bind(new(CartRepositoryReaderInterfaces), new(*cartRepository)),
	wire.Bind(new(CartRepositoryWriterInterfaces), new(*cartRepository)),
)
