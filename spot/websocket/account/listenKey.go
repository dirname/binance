package spotclient

import (
	"bytes"
	"encoding/json"
	"errors"
	binance "github.com/dirname/binance"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"github.com/dirname/binance/spot/client"
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

// CreateSpotListenKey Start a new user data stream. The stream will close after 60 minutes. If the account has an active listenKey, that listenKey will be returned and its validity will be extended for 60 minutes.
func (b *listenKeyBuilder) CreateSpotListenKey() (string, error) {
	var err error
	req, err := b.Builder.Build(http.MethodPost, "/api/v3/userDataStream", "", false, false, 0)
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

// PingSpotListenKey Keepalive a user data stream to prevent a time out. User data streams will close after 60 minutes. It's recommended to send a ping about every 30 minutes.
func (b *listenKeyBuilder) PingSpotListenKey(key string) (interface{}, error) {
	var err error
	if key == "" {
		err = errors.New(spotclient.MissParameters)
		return nil, err
	}
	params := "listenKey=" + key
	req, err := b.Builder.Build(http.MethodPut, "/api/v3/userDataStream", params, false, false, 0)
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

// DeleteSpotListenKey Close out a user data stream.
func (b *listenKeyBuilder) DeleteSpotListenKey(key string) (interface{}, error) {
	var err error
	if key == "" {
		err = errors.New(spotclient.MissParameters)
		return nil, err
	}
	params := "listenKey=" + key
	req, err := b.Builder.Build(http.MethodDelete, "/api/v3/userDataStream", params, false, false, 0)
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

// CreateMarginListenKey Start a new user data stream. The stream will close after 60 minutes. If the account has an active listenKey, that listenKey will be returned and its validity will be extended for 60 minutes.
func (b *listenKeyBuilder) CreateMarginListenKey() (string, error) {
	var err error
	req, err := b.Builder.Build(http.MethodPost, "/sapi/v1/userDataStream", "", false, false, 0)
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

// PingMarginListenKey Keepalive a user data stream to prevent a time out. User data streams will close after 60 minutes. It's recommended to send a ping about every 30 minutes.
func (b *listenKeyBuilder) PingMarginListenKey(key string) (interface{}, error) {
	var err error
	if key == "" {
		err = errors.New(spotclient.MissParameters)
		return nil, err
	}
	params := "listenKey=" + key
	req, err := b.Builder.Build(http.MethodPut, "/sapi/v1/userDataStream", params, false, false, 0)
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

// DeleteMarginListenKey Close out a user data stream.
func (b *listenKeyBuilder) DeleteMarginListenKey(key string) (interface{}, error) {
	var err error
	if key == "" {
		err = errors.New(spotclient.MissParameters)
		return nil, err
	}
	params := "listenKey=" + key
	req, err := b.Builder.Build(http.MethodDelete, "/sapi/v1/userDataStream", params, false, false, 0)
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

// CreateIsolatedListenKey Start a new user data stream. The stream will close after 60 minutes. If the account has an active listenKey, that listenKey will be returned and its validity will be extended for 60 minutes.
func (b *listenKeyBuilder) CreateIsolatedListenKey(symbol string) (string, error) {
	var err error
	if symbol == "" {
		err = errors.New(spotclient.MissParameters)
		return "", err
	}
	params := "symbol=" + symbol
	req, err := b.Builder.Build(http.MethodPost, "/sapi/v1/userDataStream/isolated", params, false, false, 0)
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

// PingIsolatedListenKey Keepalive a user data stream to prevent a time out. User data streams will close after 60 minutes. It's recommended to send a ping about every 30 minutes.
func (b *listenKeyBuilder) PingIsolatedListenKey(symbol, key string) (interface{}, error) {
	var err error
	if key == "" || symbol == "" {
		err = errors.New(spotclient.MissParameters)
		return nil, err
	}
	params := "listenKey=" + key + "&symbol=" + symbol
	req, err := b.Builder.Build(http.MethodPut, "/sapi/v1/userDataStream/isolated", params, false, false, 0)
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

// DeleteIsolatedListenKey Close out a user data stream.
func (b *listenKeyBuilder) DeleteIsolatedListenKey(symbol, key string) (interface{}, error) {
	var err error
	if key == "" || symbol == "" {
		err = errors.New(spotclient.MissParameters)
		return nil, err
	}
	params := "listenKey=" + key + "&symbol=" + symbol
	req, err := b.Builder.Build(http.MethodDelete, "/sapi/v1/userDataStream/isolated", params, false, false, 0)
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
