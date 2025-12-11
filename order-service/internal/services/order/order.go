package order

import (
	"order-service/internal/repositories/datastore/orders"

	"github.com/CROWNIX/go-utils/databases/sqlx"
	"github.com/google/wire"
)

type orderService struct {
	orderRepositoryReader    orders.OrderRepositoryReaderInterfaces
	orderRepositoryWriter    orders.OrderRepositoryWriterInterfaces
	tx                      sqlx.Tx
}

type OptionParams struct {
	OrderRepositoryReader    orders.OrderRepositoryReaderInterfaces
	OrderRepositoryWriter    orders.OrderRepositoryWriterInterfaces
	Tx                      sqlx.Tx
}

func New(opts OptionParams) *orderService {
	return &orderService{
		orderRepositoryReader:    opts.OrderRepositoryReader,
		orderRepositoryWriter:    opts.OrderRepositoryWriter,
		tx:                      opts.Tx,
	}
}

var SetWire = wire.NewSet(
	wire.Struct(new(OptionParams), "*"),
	New,
	wire.Bind(new(OrderServiceInterfaces), new(*orderService)),
)
