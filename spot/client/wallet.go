package spotclient

import (
	"encoding/json"
	"fmt"
	binance "github.com/dirname/Binance"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"net/http"
	"time"
)

// WalletClient responsible to get wallet information
type WalletClient struct {
	Builder *binance.PrivateUrlBuilder
}

// NewWalletClient Initializer
func NewWalletClient(host, appKey, appSecret string) *WalletClient {
	return &WalletClient{
		Builder: binance.NewPrivateUrlBuilder(host, appKey, appSecret),
	}
}

// GetAllCoinsInfo all coins information
func (w *WalletClient) GetAllCoinsInfo(recv time.Duration) (interface{}, error) {
	var err error
	req, err := w.Builder.Build(http.MethodGet, "/sapi/v1/capital/config/getall", "", true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := AllCoinInfoResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// GetSpotSnapshot get daily account snapshot
func (w *WalletClient) GetSpotSnapshot(startTime, endTime int64, limit int32, recv time.Duration) (interface{}, error) {
	var err error
	params := "type=SPOT"
	if startTime != 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime != 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	if limit < 5 {
		limit = 5
	}
	if limit > 30 {
		limit = 30
	}
	params += fmt.Sprintf("&limit=%d", limit)
	req, err := w.Builder.Build(http.MethodGet, "/sapi/v1/accountSnapshot", params, true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser SpotSnapshotResponse
	err = json.Unmarshal(res, &parser)
	if parser.Code != 200 {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	return parser, err
}

// GetFuturesSnapshot get daily account snapshot
func (w *WalletClient) GetFuturesSnapshot(startTime, endTime int64, limit int32, recv time.Duration) (interface{}, error) {
	var err error
	params := "type=FUTURES"
	if startTime != 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime != 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	if limit < 5 {
		limit = 5
	}
	if limit > 30 {
		limit = 30
	}
	params += fmt.Sprintf("&limit=%d", limit)
	req, err := w.Builder.Build(http.MethodGet, "/sapi/v1/accountSnapshot", params, true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser FuturesSnapshotResponse
	err = json.Unmarshal(res, &parser)
	if parser.Code != 200 {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	return parser, err
}

// GetMarginSnapshot get daily account snapshot
func (w *WalletClient) GetMarginSnapshot(startTime, endTime int64, limit int32, recv time.Duration) (interface{}, error) {
	var err error
	params := "type=MARGIN"
	if startTime != 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime != 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	if limit < 5 {
		limit = 5
	}
	if limit > 30 {
		limit = 30
	}
	params += fmt.Sprintf("&limit=%d", limit)
	req, err := w.Builder.Build(http.MethodGet, "/sapi/v1/accountSnapshot", params, true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser MarginSnapshotResponse
	err = json.Unmarshal(res, &parser)
	if parser.Code != 200 {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	return parser, err
}

// FastWithdrawSwitch enabled/disabled fast withdraw
func (w *WalletClient) FastWithdrawSwitch(status bool, recv time.Duration) (interface{}, error) {
	path := ""
	if status {
		path = "/sapi/v1/account/disableFastWithdrawSwitch"
	} else {
		path = "/sapi/v1/account/enableFastWithdrawSwitch"
	}
	var err error
	req, err := w.Builder.Build(http.MethodPost, path, "", true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser interface{}
	err = json.Unmarshal(res, &parser)
	logger.Debug("msg: %s", string(res))
	return parser, err
}
