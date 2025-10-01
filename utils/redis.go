package utils

import (
	"spot_demo/constants"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	once        sync.Once
)

func NewRedisClient() *redis.Client {
	once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     constants.RedisAddr,
			Password: constants.RedisPassword,
			DB:       constants.RedisDB,
		})
	})

	return redisClient
}
