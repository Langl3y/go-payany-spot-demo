package business

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"spot_demo/caches"
	"spot_demo/common"
	"spot_demo/models/request/payload"
	"spot_demo/models/request/response"
	"spot_demo/utils"
	"strconv"
)

func LogIn(account string, password string) (string, error) {
	var requestPayload payload.LogIn
	var requestResponse response.User[response.LogIn]

	requestPayload.Account = account
	requestPayload.Password = password
	jsonPayload, _ := json.Marshal(requestPayload)

	req, err := utils.NewHttpRequest(common.LoginUrl, common.POST, nil,
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

	userID := requestResponse.Data.UserID
	token := requestResponse.Data.Token
	expireTime := int64(requestResponse.Data.ExpireTime)

	caches.SetUserData(account, userID, token, expireTime)
	return strconv.Itoa(userID), nil
}
