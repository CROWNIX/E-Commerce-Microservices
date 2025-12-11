package config

import (

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:      GetConfig().Redis.Address, 
		DB:        GetConfig().Redis.Db,
	})

	return client
}
