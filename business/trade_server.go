package business

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"spot_demo/caches"
	"spot_demo/common"
	"spot_demo/models/request/payload"
	"spot_demo/models/request/response"
	"spot_demo/utils"
	"strconv"
	"time"
)

func PutSpotLimit(account string, password string) (string, error) {
	var requestPayload payload.PutLimit
	var requestResponse response.TradeServer[response.PutLimitError, response.PutLimit]

	val, err := caches.GetUserID(account)
	if err != nil {
		return "", err
	}

	userID, _ := strconv.Atoi(val)
	if userID < 1 {
		result, _ := LogIn(account, password)
		userID, _ = strconv.Atoi(result)
	}

	rand.Seed(time.Now().UnixNano())
	side := rand.Intn(common.Bid-common.Ask+1) + common.Ask
	amount := rand.Intn(common.MaxAmount-common.MinAmount+1) + common.MinAmount
	price := rand.Intn(common.MaxPrice-common.MinPrice+1) + common.MinPrice

	requestPayload.ID = 1
	requestPayload.Method = common.Method
	requestPayload.Params = &[]interface{}{
		userID,
		common.WalletId,
		common.AssetPair,
		side,
		strconv.Itoa(amount),
		strconv.Itoa(price),
		"0.001",
		"0.001",
		"demo",
		common.FeeAsset,
		"0.001",
		0,
	}
	jsonPayload, _ := json.Marshal(requestPayload)

	req, err := utils.NewHttpRequest(common.PutStopLimitUrl, common.POST, nil,
		bytes.NewBuffer(jsonPayload), nil)
	if err != nil {
		return "", fmt.Errorf("failed to build HTTP request: %w", err)
	}

	resp, err := utils.HttpRequest(req)
	if err != nil {
		return "", fmt.Errorf("failed to build HTTP request: %w", err)
	}

	defer func(Body io.ReadCloser) {
		_err := Body.Close()
		if _err != nil {
			fmt.Print("error reading response")
		}
	}(resp.Body)

	if respErr := json.NewDecoder(resp.Body).Decode(&requestResponse); respErr != nil {
		fmt.Print("service unavailable")
	}

	if requestResponse.Error != nil {
		return fmt.Sprintf("User %d - message: %s", userID, requestResponse.Error.Message), nil
	} else {
		log := "User %d placed %s order: %d BTC at %d USDT"
		if side == 1 {
			return fmt.Sprintf(log, userID, "Sell", amount, price), nil
		} else {
			return fmt.Sprintf(log, userID, "Buy", amount, price), nil
		}
	}
}
