package users

import "context"

type UserRepositoryReaderInterfaces interface {
	CountUserByEmail(context.Context, string) (uint64, error)
	GetUserByEmail(context.Context, string) (GetUserOutput, error)
}

type UserRepositoryWriterInterfaces interface {
	CreateUser(context.Context, CreateUserInput) error
}
