//go:build wireinject
// +build wireinject

package main

import (
	"auth-service/internal/config"
	"auth-service/internal/infra"
	"auth-service/internal/repositories/datastore/users"
	"auth-service/internal/services"
	"auth-service/internal/services/auth"
	"auth-service/internal/utils"

	"github.com/google/wire"
)

func LoadServices() (*services.Service, func()) {
	wire.Build(
		// INFRASTRUCTURE
		infra.NewMysql,
		infra.ProvideTx,

		// REPOSITORY DATASTORE LAYER
		users.SetWire,

		config.NewRedisClient,
		utils.NewTokenUtil,

		// SERVICE LAYER
		auth.SetWire,

		services.NewService,
	)
	return nil, nil
}
