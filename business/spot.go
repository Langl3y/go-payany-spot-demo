package business

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hft_backend/constants"
	"hft_backend/models/payload"
	"hft_backend/models/response"
	"hft_backend/utils"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func LogIn(account string, password string) (string, error) {
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

	_, _ = setData(account, userId, token, expireTime)
	return strconv.Itoa(userId), nil
}

func setData(account string, userId int, token string, expireTime int64) (string, error) {
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
		return LogIn(account, password)
	} else {
		return val, nil
	}
}

func PutSpotLimit(account string, password string) {
	var requestPayload payload.StopLimit
	var responseData response.PutLimitResponse

	redisClient := utils.NewRedisClient()
	val, _ := redisClient.HGet(context.Background(), account, "user_id").Result()
	if val == "" {
		val, _ = GetUserId(account, password)
	}

	userId, _ := strconv.Atoi(val)
	if userId < 1 {
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	side := rand.Intn(constants.Bid-constants.Ask+1) + constants.Ask
	amount := rand.Intn(constants.MaxAmount-constants.MinAmount+1) + constants.MinAmount
	price := rand.Intn(constants.MaxPrice-constants.MinPrice+1) + constants.MinPrice

	requestPayload.Id = 1
	requestPayload.Method = constants.Method

	requestPayload.Params = []interface{}{
		userId,
		constants.WalletId,
		constants.AssetPair,
		side,
		strconv.Itoa(amount),
		strconv.Itoa(price),
		"0.001",
		"0.001",
		"demo",
		constants.FeeAsset,
		"0.001",
		0,
	}

	jsonPayload, _ := json.Marshal(requestPayload)
	resp, _ := http.Post(constants.PutStopLimitUrl, "application/json", bytes.NewBuffer(jsonPayload))

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil || resp.Body == nil {
			fmt.Print("Service Unavailable")
			os.Exit(2)
		}
	}(resp.Body)

	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		fmt.Printf("Failed to place order: %s", err)
		os.Exit(3)
	}

	if responseData.Error != nil {
		exceptionObj := &responseData.Error

		if exceptionObj != nil {
			message := responseData.Error.Message
			fmt.Printf(`User %d - message: %s`, userId, message)
		} else {
			fmt.Print("Service Unavailable")
		}
	} else {
		log := "User %d placed %s order: %d BTC at %d USDT"
		if side == 1 {
			fmt.Printf(log, userId, "Sell", amount, price)
		} else {
			fmt.Printf(log, userId, "Buy", amount, price)
		}
	}

	fmt.Println()
	return
}
