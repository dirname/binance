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
	params := ""
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
	if name != "" {
		params += fmt.Sprintf("&name=%s", name)
	}
	params += fmt.Sprintf("&transactionFeeFlag=%t", transactionFeeFlag)
	if amount.LessThanOrEqual(decimal.NewFromInt(0)) {
		err = errors.New("amount can not less than or equal zero")
		return nil, err
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
	params := ""
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
	if name != "" {
		params += fmt.Sprintf("&name=%s", name)
	}
	if amount.LessThanOrEqual(decimal.NewFromInt(0)) {
		err = errors.New("amount can not less than or equal zero")
		return nil, err
	}
	params += fmt.Sprintf("&transactionFeeFlag=%t", transactionFeeFlag)
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

// WAPIAccountStatus account status
func (w *WalletClient) WAPIAccountStatus(recv time.Duration) (interface{}, error) {
	var err error
	req, err := w.Builder.Build(http.MethodGet, "/wapi/v3/accountStatus.html", "", true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	result := WAPIAccountResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// SAPIAccountStatus account status
func (w *WalletClient) SAPIAccountStatus(recv time.Duration) (interface{}, error) {
	var err error
	req, err := w.Builder.Build(http.MethodGet, "/sapi/v1/account/status", "", true, true, recv)
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
	result := SAPIAccountResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// WAPIAccountAPIStatus account API status
func (w *WalletClient) WAPIAccountAPIStatus(recv time.Duration) (interface{}, error) {
	var err error
	req, err := w.Builder.Build(http.MethodGet, "/wapi/v3/apiTradingStatus.html", "", true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	result := AccountWAPIAPIStatusResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// SAPIAccountAPIStatus account API status
func (w *WalletClient) SAPIAccountAPIStatus(recv time.Duration) (interface{}, error) {
	var err error
	req, err := w.Builder.Build(http.MethodGet, "/sapi/v1/account/apiTradingStatus", "", true, true, recv)
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
	result := AccountSAPIStatusResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// WAPIDustLog Fetch small amounts of assets exchanged BNB records
func (w *WalletClient) WAPIDustLog(recv time.Duration) (interface{}, error) {
	var err error
	req, err := w.Builder.Build(http.MethodGet, "/wapi/v3/userAssetDribbletLog.html", "", true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	result := DustLogWAPIResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// SAPIDustLog Fetch small amounts of assets exchanged BNB records
func (w *WalletClient) SAPIDustLog(recv time.Duration) (interface{}, error) {
	var err error
	req, err := w.Builder.Build(http.MethodGet, "/sapi/v1/asset/dribblet", "", true, true, recv)
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
	result := DustLogSAPIResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// DustTransfer Convert dust assets to BNB.
func (w *WalletClient) DustTransfer(asset []string, recv time.Duration) (interface{}, error) {
	var err error
	var params string
	for _, v := range asset {
		params += fmt.Sprintf("&asset=%s", v)
	}
	if len(params) > 1 {
		params = params[1:]
	}
	req, err := w.Builder.Build(http.MethodPost, "/sapi/v1/asset/dust", params, true, true, recv)
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
	result := DustTransferResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// AccountDividendRecord Query asset dividend record.
func (w *WalletClient) AccountDividendRecord(asset string, startTime, endTime int64, limit int, recv time.Duration) (interface{}, error) {
	var err error
	var params string
	if asset != "" {
		params += fmt.Sprintf("asset=%s", asset)
	}
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	if limit > 500 {
		limit = 500
	}
	if limit < 0 {
		limit = 20
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	if len(params) > 1 {
		params = params[1:]
	}
	req, err := w.Builder.Build(http.MethodGet, "/sapi/v1/asset/assetDividend", params, true, true, recv)
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
	result := AccountDividendRecordResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// WAPIAssetDetail Fetch details of assets supported on Binance.
// Warning The result types returned by minWithdrawAmount and withdrawFee are sometimes inconsistent, leading to running bugs. Currently, it has been parsed as interface{}. It is recommended to use this method of SAPI instead of this method.
func (w *WalletClient) WAPIAssetDetail(recv time.Duration) (interface{}, error) {
	var err error
	req, err := w.Builder.Build(http.MethodGet, "/wapi/v3/assetDetail.html", "", true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	result := WAPIAssetDetailResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// SAPIAssetDetail Fetch details of assets supported on Binance.
func (w *WalletClient) SAPIAssetDetail(recv time.Duration) (interface{}, error) {
	var err error
	req, err := w.Builder.Build(http.MethodGet, "/sapi/v1/asset/assetDetail", "", true, true, recv)
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
	result := SAPIAssetDetailResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// WAPITradeFee Fetch trade fee, values in percentage.
func (w *WalletClient) WAPITradeFee(symbol string, recv time.Duration) (interface{}, error) {
	var err error
	params := ""
	if symbol != "" {
		params += fmt.Sprintf("symbol=%s", symbol)
	}
	req, err := w.Builder.Build(http.MethodGet, "/wapi/v3/tradeFee.html", params, true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	result := WAPITradeFeeResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// SAPITradeFee Fetch trade fee, values in percentage.
func (w *WalletClient) SAPITradeFee(symbol string, recv time.Duration) (interface{}, error) {
	var err error
	params := ""
	if symbol != "" {
		params += fmt.Sprintf("symbol=%s", symbol)
	}
	req, err := w.Builder.Build(http.MethodGet, "/sapi/v1/asset/tradeFee", params, true, true, recv)
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
	result := SAPITradeFeeResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// UniversalTransfer You need to enable Permits Universal Transfer option for the api key which requests this endpoint.
/*
typeName:
	MAIN_C2C Spot account transfer to C2C account
	MAIN_UMFUTURE Spot account transfer to USDⓈ-M Futures account
	MAIN_CMFUTURE Spot account transfer to COIN-M Futures account
	MAIN_MARGIN Spot account transfer to Margin（cross）account
	MAIN_MINING Spot account transfer to Mining account
	C2C_MAIN C2C account transfer to Spot account
	C2C_UMFUTURE C2C account transfer to USDⓈ-M Futures account
	C2C_MINING C2C account transfer to Mining account
	C2C_MARGIN C2C account transfer to Margin(cross) account
	UMFUTURE_MAIN USDⓈ-M Futures account transfer to Spot account
	UMFUTURE_C2C USDⓈ-M Futures account transfer to C2C account
	UMFUTURE_MARGIN USDⓈ-M Futures account transfer to Margin（cross）account
	CMFUTURE_MAIN COIN-M Futures account transfer to Spot account
	CMFUTURE_MARGIN COIN-M Futures account transfer to Margin(cross) account
	MARGIN_MAIN Margin（cross）account transfer to Spot account
	MARGIN_UMFUTURE Margin（cross）account transfer to USDⓈ-M Futures
	MARGIN_CMFUTURE Margin（cross）account transfer to COIN-M Futures
	MARGIN_MINING Margin（cross）account transfer to Mining account
	MARGIN_C2C Margin（cross）account transfer to C2C account
	MINING_MAIN Mining account transfer to Spot account
	MINING_UMFUTURE Mining account transfer to USDⓈ-M Futures account
	MINING_C2C Mining account transfer to C2C account
	MINING_MARGIN Mining account transfer to Margin(cross) account
*/
func (w *WalletClient) UniversalTransfer(typeName, asset, amount string, recv time.Duration) (interface{}, error) {
	var err error
	if typeName == "" || asset == "" || amount == "" {
		err = errors.New("miss params")
		return nil, err
	}
	params := ""
	params += fmt.Sprintf("type=%s&asset=%s&amount=%s", typeName, asset, amount)
	req, err := w.Builder.Build(http.MethodPost, "/sapi/v1/asset/transfer", params, true, true, recv)
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
	result := UniversalTransferResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// UniversalTransferRecord universal transfer record
/*
typeName:
	MAIN_C2C Spot account transfer to C2C account
	MAIN_UMFUTURE Spot account transfer to USDⓈ-M Futures account
	MAIN_CMFUTURE Spot account transfer to COIN-M Futures account
	MAIN_MARGIN Spot account transfer to Margin（cross）account
	MAIN_MINING Spot account transfer to Mining account
	C2C_MAIN C2C account transfer to Spot account
	C2C_UMFUTURE C2C account transfer to USDⓈ-M Futures account
	C2C_MINING C2C account transfer to Mining account
	C2C_MARGIN C2C account transfer to Margin(cross) account
	UMFUTURE_MAIN USDⓈ-M Futures account transfer to Spot account
	UMFUTURE_C2C USDⓈ-M Futures account transfer to C2C account
	UMFUTURE_MARGIN USDⓈ-M Futures account transfer to Margin（cross）account
	CMFUTURE_MAIN COIN-M Futures account transfer to Spot account
	CMFUTURE_MARGIN COIN-M Futures account transfer to Margin(cross) account
	MARGIN_MAIN Margin（cross）account transfer to Spot account
	MARGIN_UMFUTURE Margin（cross）account transfer to USDⓈ-M Futures
	MARGIN_CMFUTURE Margin（cross）account transfer to COIN-M Futures
	MARGIN_MINING Margin（cross）account transfer to Mining account
	MARGIN_C2C Margin（cross）account transfer to C2C account
	MINING_MAIN Mining account transfer to Spot account
	MINING_UMFUTURE Mining account transfer to USDⓈ-M Futures account
	MINING_C2C Mining account transfer to C2C account
	MINING_MARGIN Mining account transfer to Margin(cross) account
*/
func (w *WalletClient) UniversalTransferRecord(typeName string, startTime, endTime int64, current, size int32, recv time.Duration) (interface{}, error) {
	var err error
	if typeName == "" {
		err = errors.New("miss params")
		return nil, err
	}
	params := ""
	params += fmt.Sprintf("type=%s", typeName)
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	if current < 0 {
		current = 1
	}
	if size < 0 {
		size = 10
	}
	if size > 100 {
		size = 100
	}
	if current > 0 {
		params += fmt.Sprintf("&current=%d", current)
	}
	if size > 0 {
		params += fmt.Sprintf("&size=%d", size)
	}
	req, err := w.Builder.Build(http.MethodGet, "/sapi/v1/asset/transfer", params, true, true, recv)
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
	result := UniversalTransferRecordResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}
