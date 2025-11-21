package categories

import (
	"category-service/internal/infra"

	"github.com/google/wire"
)

type categoryRepository struct {
	DB *infra.DB
}

func NewCategoryRepository(db *infra.DB) *categoryRepository {
	return &categoryRepository{
		DB: db,
	}
}

var SetWire = wire.NewSet(
	NewCategoryRepository,
	wire.Bind(new(CategoryRepositoryReaderInterfaces), new(*categoryRepository)),
	wire.Bind(new(CategoryRepositoryWriterInterfaces), new(*categoryRepository)),
)
