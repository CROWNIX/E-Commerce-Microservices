package jwtutils

import (
	"context"
	"errors"
	"fmt"
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
	fmt.Println("token", jwtToken)
	fmt.Println("secret", j.config.Secret)
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.config.Secret), nil
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("Invalid Token")
	}

	fmt.Println("token parsed")

	claims := token.Claims.(jwt.MapClaims)
	fmt.Println("claims", claims)

	expire := claims["expire"].(float64)
	if int64(expire) < time.Now().UnixMilli() {
		return nil, errors.New("Invalid Token")
	}

	fmt.Println("start get redis")
	fmt.Printf("Redis client = %#v\n", j.Redis)
	fmt.Printf("ctx nil? %v\n", ctx == nil)
	fmt.Printf("jwtToken = %s\n", jwtToken)
	result, err := j.Redis.Exists(ctx, jwtToken).Result()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println("success get token from redis")

	if result == 0 {
		return nil, errors.New("Invalid Token")
	}
	fmt.Println("token exists")

	return claims, nil
}
