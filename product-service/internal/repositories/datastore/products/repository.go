package products

import (
	"product-service/internal/infra"

	"github.com/google/wire"
)

type productRepository struct {
	DB *infra.DB
}

func NewProductRepository(db *infra.DB) *productRepository {
	return &productRepository{
		DB: db,
	}
}

var SetWire = wire.NewSet(
	NewProductRepository,
	wire.Bind(new(ProductRepositoryReaderInterfaces), new(*productRepository)),
	wire.Bind(new(ProductRepositoryWriterInterfaces), new(*productRepository)),
)
