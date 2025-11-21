package utils

import (
	"category-service/internal/config"
	"category-service/internal/models"
	"context"
	"time"

	"github.com/CROWNIX/go-utils/apperror"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type TokenUtil struct {
	Redis *redis.Client
}

func NewTokenUtil(redisClient *redis.Client) *TokenUtil {
	return &TokenUtil{
		Redis: redisClient,
	}
}

func (t *TokenUtil) CreateToken(ctx context.Context, auth models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     auth.ID,
		"expire": time.Now().Add(time.Hour * 24 * 30).UnixMilli(),
	})

	jwtToken, err := token.SignedString([]byte(config.GetConfig().JWT.Secret))
	if err != nil {
		return "", err
	}

	_, err = t.Redis.SetEx(ctx, jwtToken, auth.ID, time.Hour*25*30).Result()
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (t *TokenUtil) ParseToken(ctx context.Context, jwtToken string) (*models.User, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JWT.Secret), nil
	})

	if err != nil {
		return nil, apperror.Unauthorized("Failed parse token")
	}

	claims := token.Claims.(jwt.MapClaims)

	expire := claims["expire"].(float64)
	if int64(expire) < time.Now().UnixMilli() {
		return nil, apperror.Unauthorized("Access token expired")
	}

	result, err := t.Redis.Exists(ctx, jwtToken).Result()
	if err != nil {
		return nil, err
	}

	if result == 0 {
		return nil, apperror.Unauthorized("Invalid access token")
	}

	idFloat, ok := claims["id"].(float64)
	if !ok {
		return nil, apperror.Unauthorized("Invalid token payload")
	}

	auth := &models.User{
		ID: uint64(idFloat),
	}

	return auth, nil
}
