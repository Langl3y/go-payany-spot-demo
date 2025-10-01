package caches

import (
	"context"
	"errors"
	"spot_demo/utils"
	"time"

	"github.com/redis/go-redis/v9"
)

func SetUserData(account string, userID int, token string, expireTime int64) {
	redisClient := utils.NewRedisClient()

	data := map[string]interface{}{
		"user_id": userID,
		"token":   token,
	}

	expiresIn := expireTime - utils.CurrentTimestamp()
	redisClient.HSet(context.Background(), account, data)
	redisClient.Expire(context.Background(), account, time.Duration(expiresIn)*time.Second)
}

func GetUserID(account string) (string, error) {
	redisClient := utils.NewRedisClient()
	val, err := redisClient.HGet(context.Background(), account, "user_id").Result()

	if errors.Is(redis.Nil, err) {
		return "", err
	} else {
		return val, nil
	}
}
