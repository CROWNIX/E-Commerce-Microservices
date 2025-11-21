package users

import (
	"auth-service/internal/infra"

	"github.com/google/wire"
)

type userRepository struct {
	DB *infra.DB
}

func NewUserRepository(db *infra.DB) *userRepository {
	return &userRepository{
		DB: db,
	}
}

var SetWire = wire.NewSet(
	NewUserRepository,
	wire.Bind(new(UserRepositoryReaderInterfaces), new(*userRepository)),
	wire.Bind(new(UserRepositoryWriterInterfaces), new(*userRepository)),
)
