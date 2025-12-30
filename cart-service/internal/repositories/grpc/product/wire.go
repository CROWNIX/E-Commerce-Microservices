package product

import (
	"github.com/google/wire"
)

var SetWire = wire.NewSet(
	NewProductRepository,
)
