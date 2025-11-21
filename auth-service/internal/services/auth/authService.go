package auth

import (
	"auth-service/internal/models"
	"auth-service/internal/repositories/datastore/users"
	"context"
	"errors"
	"strings"

	"github.com/CROWNIX/go-utils/apperror"
	"github.com/CROWNIX/go-utils/databases"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) Register(ctx context.Context, input RegisterInput) (err error) {
	total, err := s.userRepositoryReader.CountUserByEmail(ctx, input.Email)

	if err != nil {
		return err
	}

	if total > 0 {
		return apperror.Conflict("Email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	username := strings.Split(input.Email, "@")[0]

	err = s.userRepositoryWriter.CreateUser(ctx, users.CreateUserInput{
		Username: username,
		Email:    input.Email,
		Password: string(hashedPassword),
	})

	if err != nil {
		return err
	}

	return
}

func (s *authService) Login(ctx context.Context, input LoginInput) (token string, err error) {
	userOutput, err := s.userRepositoryReader.GetUserByEmail(ctx, input.Email)

	if err != nil {
		if errors.Is(err, databases.ErrNoRowFound) {
			return "", apperror.Conflict("Email atau password salah")
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userOutput.Password), []byte(input.Password)); err != nil {
		return "", apperror.Conflict("Email atau password salah")
	}

	token, err = s.TokenUtil.CreateToken(ctx, models.User{ID: userOutput.ID})
	if err != nil {
		return "", err
	}

	return token, nil
}
