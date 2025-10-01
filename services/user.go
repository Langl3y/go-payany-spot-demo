package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"spot_demo/constants"
	"spot_demo/models/response"
	"spot_demo/utils"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func logIn(account string, password string) (string, error) {
	contentType := "application/json"
	bodyContent := fmt.Sprintf(`{"account": "%s", "password": "%s"}`, account, password)

	body := []byte(bodyContent)
	resp, _ := http.Post(constants.LoginUrl, contentType, bytes.NewBuffer(body))

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	var result response.LogInResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result.Code == 401 {
		return "", fmt.Errorf("incorrect account or password")
	}

	userId := result.Data.UserId
	token := result.Data.Token
	expireTime := int64(result.Data.ExpireTime)

	_, _ = setUserData(account, userId, token, expireTime)
	return strconv.Itoa(userId), nil
}

func setUserData(account string, userId int, token string, expireTime int64) (string, error) {
	redisClient := utils.NewRedisClient()

	data := map[string]interface{}{
		"user_id": userId,
		"token":   token,
	}

	expiresIn := expireTime - utils.CurrentTimestamp()
	redisClient.HSet(context.Background(), account, data)
	redisClient.Expire(context.Background(), account, time.Duration(expiresIn)*time.Second)

	val, _ := redisClient.HGet(context.Background(), account, "user_id").Result()
	return val, nil
}

func GetUserId(account string, password string) (string, error) {
	redisClient := utils.NewRedisClient()

	val, err := redisClient.HGet(context.Background(), account, "user_id").Result()
	fmt.Println(val)
	if errors.Is(redis.Nil, err) {
		return logIn(account, password)
	} else {
		return val, nil
	}
}
