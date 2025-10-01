package business

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"spot_demo/constants"
	"spot_demo/models/payload"
	"spot_demo/models/response"
	"spot_demo/services"
	"spot_demo/utils"
	"strconv"
	"time"
)

func PutSpotLimit(account string, password string) {
	var requestPayload payload.StopLimit
	var responseData response.PutLimitResponse

	redisClient := utils.NewRedisClient()
	val, _ := redisClient.HGet(context.Background(), account, "user_id").Result()
	if val == "" {
		val, _ = services.GetUserId(account, password)
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
