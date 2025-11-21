//go:build wireinject
// +build wireinject

package main

import (
	"category-service/internal/infra"
	"category-service/internal/repositories/datastore/categories"
	"category-service/internal/services"
	"category-service/internal/services/category"

	"github.com/google/wire"
)

func LoadServices() (*services.Service, func()) {
	wire.Build(
		// INFRASTRUCTURE
		infra.NewMysql,
		infra.ProvideTx,

		// REPOSITORY DATASTORE LAYER
		categories.SetWire,

		// SERVICE LAYER
		category.SetWire,

		services.NewService,
	)
	return nil, nil
}
