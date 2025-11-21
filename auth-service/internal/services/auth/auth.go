package auth

import (
	"auth-service/internal/repositories/datastore/users"
	"auth-service/internal/utils"

	"github.com/CROWNIX/go-utils/databases/sqlx"
	"github.com/google/wire"
)

type authService struct {
	userRepositoryReader users.UserRepositoryReaderInterfaces
	userRepositoryWriter users.UserRepositoryWriterInterfaces
	tx                   sqlx.Tx
	TokenUtil            *utils.TokenUtil
}

type OptionParams struct {
	UserRepositoryReader users.UserRepositoryReaderInterfaces
	UserRepositoryWriter users.UserRepositoryWriterInterfaces
	Tx                   sqlx.Tx
	TokenUtil            *utils.TokenUtil
}

func New(opts OptionParams) *authService {
	return &authService{
		userRepositoryReader: opts.UserRepositoryReader,
		userRepositoryWriter: opts.UserRepositoryWriter,
		tx:                   opts.Tx,
		TokenUtil:            opts.TokenUtil,
	}
}

var SetWire = wire.NewSet(
	wire.Struct(new(OptionParams), "*"),
	New,
	wire.Bind(new(AuthServiceInterfaces), new(*authService)),
)
