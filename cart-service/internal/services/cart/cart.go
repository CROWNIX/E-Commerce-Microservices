package cart

import (
	"cart-service/internal/repositories/datastore/carts"
	"cart-service/internal/repositories/grpc/products"

	"github.com/CROWNIX/go-utils/databases/sqlx"
	"github.com/google/wire"
)

type cartService struct {
	cartRepositoryReader carts.CartRepositoryReaderInterfaces
	cartRepositoryWriter carts.CartRepositoryWriterInterfaces
	productService       products.ProductServiceInterfaces
	tx                   sqlx.Tx
}

type OptionParams struct {
	CartRepositoryReader carts.CartRepositoryReaderInterfaces
	CartRepositoryWriter carts.CartRepositoryWriterInterfaces
	ProductService       products.ProductServiceInterfaces
	Tx                   sqlx.Tx
}

func New(opts OptionParams) *cartService {
	return &cartService{
		cartRepositoryReader: opts.CartRepositoryReader,
		cartRepositoryWriter: opts.CartRepositoryWriter,
		productService:       opts.ProductService,
		tx:                   opts.Tx,
	}
}

var SetWire = wire.NewSet(
	wire.Struct(new(OptionParams), "*"),
	New,
	wire.Bind(new(CartServiceInterfaces), new(*cartService)),
)
