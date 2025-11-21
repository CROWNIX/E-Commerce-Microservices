//go:build wireinject
// +build wireinject

package main

import (
	"product-service/internal/infra"

	"product-service/internal/repositories/datastore/products"
	"product-service/internal/services"
	"product-service/internal/services/product"

	"github.com/google/wire"
)

func LoadServices() (*services.Service, func()) {
	wire.Build(
		// INFRASTRUCTURE
		infra.NewMysql,
		infra.ProvideTx,

		// REPOSITORY DATASTORE LAYER
		products.SetWire,

		// SERVICE LAYER
		product.SetWire,

		services.NewService,
	)
	return nil, nil
}
