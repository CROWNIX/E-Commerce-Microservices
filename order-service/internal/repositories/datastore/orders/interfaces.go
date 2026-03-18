package orders

import (
	"context"
	utilSql "github.com/CROWNIX/go-utils/databases/sqlx"
)

type OrderRepositoryReaderInterfaces interface {
}

type OrderRepositoryWriterInterfaces interface {
	CreateOrder(context.Context, utilSql.RDBMS, CreateOrderInput) (uint64, error)
	CreateOrderDetail(context.Context, CreateOrderDetailInput) error
}