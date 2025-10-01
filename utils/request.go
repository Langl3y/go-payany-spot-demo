package utils

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"spot_demo/constants"
)

func NewHttpRequest(url string, method constants.HttpMethod,
	headers map[string]string, payload []byte, params map[string]any) (*http.Request, error) {

	var req *http.Request
	var err error

	if method == constants.GET {
		req, err = newGetRequest(url, params)
	} else if method == constants.POST {
		req, err = newPostRequest(url, payload)
	} else {
		return nil, fmt.Errorf("invalid method")
	}

	if err != nil {
		return nil, err
	}

	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	if method == constants.POST {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func newPostRequest(url string, payload []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	return req, nil
}

func newGetRequest(baseURL string, params map[string]any) (*http.Request, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, fmt.Sprintf("%v", v))
	}
	u.RawQuery = q.Encode()

	return http.NewRequest("GET", u.String(), nil)
}
