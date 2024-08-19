package wbapi

import (
	"net/http"
)

const BaseUrl = "https://supplies-api.wildberries.ru/api/v1"

type WbClient struct {
	BaseUrl    string
	ApiToken   string
	HTTPClient *http.Client
}

func NewWbClient(apiToken string) *WbClient {
	return &WbClient{
		BaseUrl:    BaseUrl,
		ApiToken:   apiToken,
		HTTPClient: &http.Client{},
	}
}
