package orders

import (
	"order-service/internal/infra"

	"github.com/google/wire"
)

type orderRepository struct {
	DB *infra.DB
}

func NewOrderRepository(db *infra.DB) *orderRepository {
	return &orderRepository{
		DB: db,
	}
}

var SetWire = wire.NewSet(
	NewOrderRepository,
	wire.Bind(new(OrderRepositoryReaderInterfaces), new(*orderRepository)),
	wire.Bind(new(OrderRepositoryWriterInterfaces), new(*orderRepository)),
)
