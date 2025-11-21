package auth

import "context"

type AuthServiceInterfaces interface {
	Register(context.Context, RegisterInput) (error)
	Login(context.Context, LoginInput) (string, error)
}
