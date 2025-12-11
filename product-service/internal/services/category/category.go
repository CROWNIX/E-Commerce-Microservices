package category

import (
	"product-service/internal/repositories/datastore/categories"

	"github.com/CROWNIX/go-utils/databases/sqlx"
	"github.com/google/wire"
)

type categoryService struct {
	categoryRepositoryReader categories.CategoryRepositoryReaderInterfaces
	categoryRepositoryWriter categories.CategoryRepositoryWriterInterfaces
	tx                       sqlx.Tx
}

type OptionParams struct {
	CategoryRepositoryReader categories.CategoryRepositoryReaderInterfaces
	CategoryRepositoryWriter categories.CategoryRepositoryWriterInterfaces
	Tx                       sqlx.Tx
}

func New(opts OptionParams) *categoryService {
	return &categoryService{
		categoryRepositoryReader: opts.CategoryRepositoryReader,
		categoryRepositoryWriter: opts.CategoryRepositoryWriter,
		tx:                       opts.Tx,
	}
}

var SetWire = wire.NewSet(
	wire.Struct(new(OptionParams), "*"),
	New,
	wire.Bind(new(CategoryServiceInterfaces), new(*categoryService)),
)
