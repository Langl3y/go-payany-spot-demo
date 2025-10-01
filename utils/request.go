package utils

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"spot_demo/common"
)

var httpClient = &http.Client{}

func HttpRequest(req *http.Request) (*http.Response, error) {
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NewHttpRequest(url string, method common.HttpMethod,
	headers map[string]string, payload *bytes.Buffer, params map[string]any) (*http.Request, error) {

	var req *http.Request
	var err error

	if method == common.GET {
		req, err = newGetRequest(url, params)
	} else if method == common.POST {
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

	if method == common.POST {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func newPostRequest(url string, payload *bytes.Buffer) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, payload)
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
