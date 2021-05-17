package account

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/dirname/binance"
	"github.com/dirname/binance/futures/usd/client"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"net/http"
)

// ListenKeyResponse listen key response
type ListenKeyResponse struct {
	ListenKey string `json:"listenKey"`
}

// listenKeyBuilder responsible to get ws account assets information
type listenKeyBuilder struct {
	Builder *binance.PrivateUrlBuilder
}

// NewListenKeyBuilder Initializer
func NewListenKeyBuilder(host, appKey, appSecret string) *listenKeyBuilder {
	return &listenKeyBuilder{
		Builder: binance.NewPrivateUrlBuilder(host, appKey, appSecret),
	}
}

// CreateListenKey Start a new user data stream. The stream will close after 60 minutes. If the account has an active listenKey, that listenKey will be returned and its validity will be extended for 60 minutes.
func (b *listenKeyBuilder) CreateListenKey() (string, error) {
	var err error
	req, err := b.Builder.Build(http.MethodPost, "/fapi/v1/listenKey", "", false, false, 0)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		if err == nil {
			err = errors.New(result.Message)
		}
		return result.Message, err
	}
	result := ListenKeyResponse{}
	err = json.Unmarshal(res, &result)
	return result.ListenKey, err
}

// PingListenKey Keepalive a user data stream to prevent a time out. User data streams will close after 60 minutes. It's recommended to send a ping about every 30 minutes.
func (b *listenKeyBuilder) PingListenKey(key string) (interface{}, error) {
	var err error
	if key == "" {
		err = errors.New(futuresclient.MissParameters)
		return nil, err
	}
	params := "listenKey=" + key
	req, err := b.Builder.Build(http.MethodPut, "/fapi/v1/listenKey", params, false, false, 0)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result.Message, err
	}
	if bytes.Equal(res, []byte{123, 125}) {
		return true, err
	}
	return nil, err
}

// DeleteListenKey Close out a user data stream.
func (b *listenKeyBuilder) DeleteListenKey(key string) (interface{}, error) {
	var err error
	if key == "" {
		err = errors.New(futuresclient.MissParameters)
		return nil, err
	}
	params := "listenKey=" + key
	req, err := b.Builder.Build(http.MethodDelete, "/fapi/v1/listenKey", params, false, false, 0)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result.Message, err
	}
	if bytes.Equal(res, []byte{123, 125}) {
		return true, err
	}
	return nil, err
}