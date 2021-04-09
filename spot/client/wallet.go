package spotclient

import (
	"encoding/json"
	"errors"
	"fmt"
	binance "github.com/dirname/Binance"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"github.com/shopspring/decimal"
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
// Todo wait verified
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

// SAPIWithdraw withdraw
func (w *WalletClient) SAPIWithdraw(coin, clientID, network, address, addressTag, name string, amount decimal.Decimal, transactionFeeFlag bool, recv time.Duration) (interface{}, error) {
	var err error
	var params string
	if coin == "" {
		err = errors.New("coin can not be empty")
		return nil, err
	}
	params += fmt.Sprintf("coin=%s", coin)
	if clientID != "" {
		params += fmt.Sprintf("&withdrawOrderId=%s", clientID)
	}
	if network != "" {
		params += fmt.Sprintf("&network=%s", network)
	}
	if address == "" {
		err = errors.New("address can not be empty")
		return nil, err
	}
	params += fmt.Sprintf("&address=%s", address)
	if addressTag != "" {
		params += fmt.Sprintf("&addressTag=%s", addressTag)
	}
	if amount.LessThanOrEqual(decimal.NewFromInt(0)) {
		err = errors.New("amount can not less than or equal zero")
		return nil, err
	}
	params += fmt.Sprintf("&transactionFeeFlag=%t", transactionFeeFlag)
	if name != "" {
		params += fmt.Sprintf("&name=%s", name)
	}
	req, err := w.Builder.Build(http.MethodPost, "/sapi/v1/capital/withdraw/apply", "", true, true, recv)
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
	result := SAPIWithdrawResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// WAPIWithdraw withdraw
func (w *WalletClient) WAPIWithdraw(coin, clientID, network, address, addressTag, name string, amount decimal.Decimal, transactionFeeFlag bool, recv time.Duration) (interface{}, error) {
	var err error
	var params string
	if coin == "" {
		err = errors.New("coin can not be empty")
		return nil, err
	}
	params += fmt.Sprintf("coin=%s", coin)
	if clientID != "" {
		params += fmt.Sprintf("&withdrawOrderId=%s", clientID)
	}
	if network != "" {
		params += fmt.Sprintf("&network=%s", network)
	}
	if address == "" {
		err = errors.New("address can not be empty")
		return nil, err
	}
	params += fmt.Sprintf("&address=%s", address)
	if addressTag != "" {
		params += fmt.Sprintf("&addressTag=%s", addressTag)
	}
	if amount.LessThanOrEqual(decimal.NewFromInt(0)) {
		err = errors.New("amount can not less than or equal zero")
		return nil, err
	}
	params += fmt.Sprintf("&transactionFeeFlag=%t", transactionFeeFlag)
	if name != "" {
		params += fmt.Sprintf("&name=%s", name)
	}
	req, err := w.Builder.Build(http.MethodPost, "/wapi/v3/withdraw.html", "", true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	result := SAPIWithdrawResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// DepositHistoryNetwork deposit history supporting network set status exclude [0,6,1] to use default.
func (w *WalletClient) DepositHistoryNetwork(coin string, status, offset, limit int32, startTime, endTime int64, recv time.Duration) (interface{}, error) {
	var err error
	var params string
	if coin != "" {
		params += fmt.Sprintf("coin=%s", coin)
	}
	if status == 0 || status == 6 || status == 1 {
		params += fmt.Sprintf("&status=%d", status)
	}
	if offset > 0 {
		params += fmt.Sprintf("&offset=%d", offset)
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	req, err := w.Builder.Build(http.MethodGet, "/sapi/v1/capital/deposit/hisrec", params, true, true, recv)
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
	result := DepositHistoryNetworkResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// DepositHistory deposit history
func (w *WalletClient) DepositHistory(asset string, status int32, startTime, endTime int64, recv time.Duration) (interface{}, error) {
	var err error
	var params string
	if asset != "" {
		params += fmt.Sprintf("asset=%s", asset)
	}
	if status == 0 || status == 6 || status == 1 {
		params += fmt.Sprintf("&status=%d", status)
	}
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	req, err := w.Builder.Build(http.MethodGet, "/wapi/v3/depositHistory.html", params, true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.WAPIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := DepositHistoryResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// WithdrawHistoryNetwork withdraw history supporting network set status exclude [0,6,1] to use default.
func (w *WalletClient) WithdrawHistoryNetwork(coin string, status, offset, limit int32, startTime, endTime int64, recv time.Duration) (interface{}, error) {
	var err error
	var params string
	if coin != "" {
		params += fmt.Sprintf("coin=%s", coin)
	}
	if status == 0 || status == 6 || status == 1 {
		params += fmt.Sprintf("&status=%d", status)
	}
	if offset > 0 {
		params += fmt.Sprintf("&offset=%d", offset)
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	req, err := w.Builder.Build(http.MethodGet, "/sapi/v1/capital/withdraw/history", params, true, true, recv)
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
	result := WithdrawHistoryNetworkResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// WithdrawHistory withdraw history
func (w *WalletClient) WithdrawHistory(asset string, status int32, startTime, endTime int64, recv time.Duration) (interface{}, error) {
	var err error
	var params string
	if asset != "" {
		params += fmt.Sprintf("asset=%s", asset)
	}
	if status == 0 || status == 6 || status == 1 {
		params += fmt.Sprintf("&status=%d", status)
	}
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	req, err := w.Builder.Build(http.MethodGet, "/wapi/v3/withdrawHistory.html", params, true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.WAPIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := WithdrawHistoryResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// DepositAddressNetwork deposit address supporting network
func (w *WalletClient) DepositAddressNetwork(coin, network string, recv time.Duration) (interface{}, error) {
	var err error
	var params string
	if coin == "" {
		err = errors.New("coin can not be empty")
		return nil, err
	}
	params = fmt.Sprintf("coin=%s", coin)
	if network != "" {
		params += fmt.Sprintf("&network=%s", network)
	}
	req, err := w.Builder.Build(http.MethodGet, "/sapi/v1/capital/deposit/address", params, true, true, recv)
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
	result := DepositAddressNetworkResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// DepositAddress deposit address
func (w *WalletClient) DepositAddress(asset string, status bool, recv time.Duration) (interface{}, error) {
	var err error
	var params string
	if asset == "" {
		err = errors.New("asset can not be empty")
		return nil, err
	}
	params = fmt.Sprintf("asset=%s", asset)
	params += fmt.Sprintf("&status=%t", status)
	req, err := w.Builder.Build(http.MethodGet, "/wapi/v3/depositAddress.html", params, true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser DepositAddressResponse
	err = json.Unmarshal(res, &parser)
	return parser, err
}
