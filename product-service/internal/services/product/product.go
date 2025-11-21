package product

import (
	"product-service/internal/repositories/datastore/products"

	"github.com/CROWNIX/go-utils/databases/sqlx"
	"github.com/google/wire"
)

type productService struct {
	productRepositoryReader products.ProductRepositoryReaderInterfaces
	productRepositoryWriter products.ProductRepositoryWriterInterfaces
	tx                      sqlx.Tx
}

type OptionParams struct {
	ProductRepositoryReader products.ProductRepositoryReaderInterfaces
	ProductRepositoryWriter products.ProductRepositoryWriterInterfaces
	Tx                      sqlx.Tx
}

func New(opts OptionParams) *productService {
	return &productService{
		productRepositoryReader: opts.ProductRepositoryReader,
		productRepositoryWriter: opts.ProductRepositoryWriter,
		tx:                      opts.Tx,
	}
}

var SetWire = wire.NewSet(
	wire.Struct(new(OptionParams), "*"),
	New,
	wire.Bind(new(ProductServiceInterfaces), new(*productService)),
)
