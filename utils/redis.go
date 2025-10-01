package utils

import (
	"spot_demo/common"
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
			Addr:     common.RedisAddr,
			Password: common.RedisPassword,
			DB:       common.RedisDB,
		})
	})

	return redisClient
}
