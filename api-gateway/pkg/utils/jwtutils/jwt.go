package jwtutils

import (
	"context"
	"errors"
	"time"

	"github.com/CROWNIX/E-Commerce-Microservices/api-gateway/internal/config"
	"github.com/redis/go-redis/v9"

	"github.com/golang-jwt/jwt/v5"
)

type JwtUtil interface {
	ParseAndVerifyWithRedis(context.Context, string) (jwt.MapClaims, error)
}

type JWTClaims struct {
	jwt.RegisteredClaims
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
}

type JWTPayload struct {
	UserID int64
	Email  string
}

type jwtUtil struct {
	config *config.JWTConfig
	Redis  *redis.Client
}

func NewJwtUtil() JwtUtil {
	return &jwtUtil{
		config: &config.JWTConfig{
			Secret: config.GetConfig().JwtSecretKey,
		},
		Redis: redis.NewClient(&redis.Options{
			Addr:     config.GetConfig().RedisAddress,     
			DB:       config.GetConfig().RedisDb,
		}),
	}
}

func (j *jwtUtil) ParseAndVerifyWithRedis(ctx context.Context, jwtToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.config.Secret), nil
	})

	if err != nil {
		return nil, errors.New("Invalid Token")
	}


	claims := token.Claims.(jwt.MapClaims)

	expire := claims["expire"].(float64)
	if int64(expire) < time.Now().UnixMilli() {
		return nil, errors.New("Invalid Token")
	}

	result, err := j.Redis.Exists(ctx, jwtToken).Result()
	if err != nil {
		return nil, err
	}

	if result == 0 {
		return nil, errors.New("Invalid Token")
	}

	return claims, nil
}
