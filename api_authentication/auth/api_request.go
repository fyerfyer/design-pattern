package auth

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

type APIRequest struct {
	baseURL   string
	token     string
	appID     string
	timestamp int64
}

func NewAPIRequest(baseURL, token, appID string, timestamp int64) *APIRequest {
	return &APIRequest{
		baseURL:   baseURL,
		token:     token,
		appID:     appID,
		timestamp: timestamp,
	}
}

func NewAPIRequestFromFullURL(fullURL string) (*APIRequest, error) {
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		return nil, err
	}

	qureyParam := parsedURL.Query()
	appID := qureyParam.Get("appid")
	token := qureyParam.Get("token")
	timestampStr := qureyParam.Get("timestamp")

	if appID == "" || token == "" || timestampStr == "" {
		return nil, errors.New("missing required parameters in URL")
	}

	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return nil, errors.New("invalid timestamp")
	}

	return &APIRequest{
		baseURL:   strings.Split(fullURL, "?")[0], // Base URL without query parameters
		appID:     appID,
		token:     token,
		timestamp: timestamp,
	}, nil
}

func (r *APIRequest) GetBaseUrl() string {
	return r.baseURL
}

func (r *APIRequest) GetToken() string {
	return r.token
}

func (r *APIRequest) GetAppID() string {
	return r.appID
}

func (r *APIRequest) GetTimeStamp() int64 {
	return r.timestamp
}
