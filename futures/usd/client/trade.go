package futuresclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dirname/binance"
	"github.com/dirname/binance/futures/usd/client/define/order"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"github.com/shopspring/decimal"
	"net/http"
	"strings"
	"time"
)

// TradeClient responsible to trading information
type TradeClient struct {
	Builder *binance.PrivateUrlBuilder
}

// NewTradeClient Initializer
func NewTradeClient(host, appKey, appSecret string) *TradeClient {
	return &TradeClient{
		Builder: binance.NewPrivateUrlBuilder(host, appKey, appSecret),
	}
}

//PositionSideDual change user's position mode
func (t *TradeClient) PositionSideDual(dualSidePosition bool, recv time.Duration) (interface{}, error) {
	var err error
	val := ""
	if dualSidePosition == false {
		val = "false"
	} else {
		val = "true"
	}
	params := strings.Join([]string{"dualSidePosition=", val}, "")
	req, err := t.Builder.Build(http.MethodPost, "/fapi/v1/positionSide/dual", params, true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	result := model.APIErrorResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//GetPositionSideDual Get user's position mode (Hedge Mode or One-way Mode ) on EVERY symbol
func (t *TradeClient) GetPositionSideDual(recv time.Duration) (interface{}, error) {
	var err error
	req, err := t.Builder.Build(http.MethodGet, "/fapi/v1/positionSide/dual", "", true, true, recv)
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
	result := DualSidePositionResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//ChangeMultiAssetsMode Change user's Multi-Assets mode (Multi-Assets Mode or Single-Asset Mode) on Every symbol
func (t *TradeClient) ChangeMultiAssetsMode(margin bool, recv time.Duration) (interface{}, error) {
	var err error
	val := ""
	if margin == false {
		val = "false"
	} else {
		val = "true"
	}
	params := strings.Join([]string{"multiAssetsMargin=", val}, "")
	req, err := t.Builder.Build(http.MethodPost, "/fapi/v1/multiAssetsMargin", params, true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	result := model.APIErrorResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//GetMultiAssetsMargin Get user's Multi-Assets mode (Multi-Assets Mode or Single-Asset Mode) on Every symbol
func (t *TradeClient) GetMultiAssetsMargin(recv time.Duration) (interface{}, error) {
	var err error
	req, err := t.Builder.Build(http.MethodGet, "/fapi/v1/multiAssetsMargin", "", true, true, recv)
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
	result := MultiAssetsMarginResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//NewOrder Send in a new order.
func (t *TradeClient) NewOrder(symbol, side, positionSide, ordersType, reduceOnly, newClientOrderID, closePosition, timeInForce, workingType, priceProtect, newOrderRespType string, quantity, price, stopPrice, activationPrice, callbackRate decimal.Decimal, recv time.Duration) (interface{}, error) {
	var err error
	params, err := buildOrder(symbol, side, positionSide, ordersType, reduceOnly, newClientOrderID, closePosition, timeInForce, workingType, priceProtect, newOrderRespType, quantity, price, stopPrice, activationPrice, callbackRate)
	if err != nil {
		return nil, err
	}
	req, err := t.Builder.Build(http.MethodPost, "/fapi/v1/order", params, true, true, recv)
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
	switch newOrderRespType {
	case order.Result:
		result := NewOrderResponseResult{}
		err = json.Unmarshal(res, &result)
		return result, err
	default:
		result := NewOrderResponseACK{}
		err = json.Unmarshal(res, &result)
		return parser, err
	}
	return parser, err
}

//BatchNewOrder batch new order
func (t *TradeClient) BatchNewOrder(batchOrders OrderBatch, recv time.Duration) (interface{}, error) {
	var err error
	if batchOrders == nil {
		err = errors.New(MissParameters)
		return nil, err
	}
	param, err := json.Marshal(batchOrders)
	if err != nil {
		return nil, err
	}
	params := fmt.Sprintf("batchOrders=%s", param)
	req, err := t.Builder.Build(http.MethodGet, "/fapi/v1/batchOrders", params, true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	result := BatchNewOrderResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//TestNewOrder new test order
func (t *TradeClient) TestNewOrder(symbol, side, positionSide, ordersType, reduceOnly, newClientOrderID, closePosition, timeInForce, workingType, priceProtect, newOrderRespType string, quantity, price, stopPrice, activationPrice, callbackRate decimal.Decimal, recv time.Duration) (interface{}, error) {
	var err error
	params, err := buildOrder(symbol, side, positionSide, ordersType, reduceOnly, newClientOrderID, closePosition, timeInForce, workingType, priceProtect, newOrderRespType, quantity, price, stopPrice, activationPrice, callbackRate)
	if err != nil {
		return nil, err
	}
	req, err := t.Builder.Build(http.MethodPost, "/fapi/v1/order/test", params, true, true, recv)
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
	return parser, err
}

//GetOrder Check an order's status.
func (t *TradeClient) GetOrder(symbol, origClientOrderID string, orderID int64, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" {
		err = errors.New(SymbolEmpty)
		return nil, err
	}
	if orderID <= 0 && origClientOrderID == "" {
		err = errors.New(MissParameters)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s", symbol)
	if origClientOrderID != "" {
		params += fmt.Sprintf("&origClientOrderId=%s", origClientOrderID)
	}
	if orderID > 0 {
		params += fmt.Sprintf("&orderId=%d", orderID)
	}
	req, err := t.Builder.Build(http.MethodGet, "/fapi/v1/order", params, true, true, recv)
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
	result := OrderBookResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//DeleteOrder delete order
func (t *TradeClient) DeleteOrder(symbol, origClientOrderID string, orderID int64, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" {
		err = errors.New(SymbolEmpty)
		return nil, err
	}
	if orderID <= 0 && origClientOrderID == "" {
		err = errors.New(MissParameters)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s", symbol)
	if origClientOrderID != "" {
		params += fmt.Sprintf("&origClientOrderId=%s", origClientOrderID)
	}
	if orderID > 0 {
		params += fmt.Sprintf("&orderId=%d", orderID)
	}
	req, err := t.Builder.Build(http.MethodDelete, "/fapi/v1/order", params, true, true, recv)
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
	result := DeleteOrderResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//DeleteAllOrder Cancel an active order
func (t *TradeClient) DeleteAllOrder(symbol string, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" {
		err = errors.New(SymbolEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s", symbol)
	req, err := t.Builder.Build(http.MethodDelete, "/fapi/v1/allOpenOrders", params, true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	result := model.APIErrorResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//BatchDeleteOrder batch delete order
func (t *TradeClient) BatchDeleteOrder(symbol string, orderIDList []int64, origClientOrderID []string, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" {
		err = errors.New(SymbolEmpty)
		return nil, err
	}
	if len(orderIDList) == 0 || len(origClientOrderID) == 0 {
		err = errors.New(MissParameters)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s", symbol)
	if len(orderIDList) > 0 {
		param, err := json.Marshal(orderIDList)
		if err != nil {
			return nil, err
		}
		params += fmt.Sprintf("&orderIdList=%s", param)
	}
	if len(origClientOrderID) > 0 {
		param, err := json.Marshal(orderIDList)
		if err != nil {
			return nil, err
		}
		params += fmt.Sprintf("&origClientOrderIdList=%s", param)
	}
	req, err := t.Builder.Build(http.MethodDelete, "/fapi/v1/batchOrders", params, true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	result := BatchDeleteOrderResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//CountDownDeleteAllOrder Cancel all open orders of the specified symbol at the end of the specified countdown.
func (t *TradeClient) CountDownDeleteAllOrder(symbol string, countDownTime int64, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" {
		err = errors.New(SymbolEmpty)
		return nil, err
	}
	if countDownTime < 0 {
		err = errors.New(CountDownTimeInvalid)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s&countdownTime=%d", symbol, countDownTime)
	req, err := t.Builder.Build(http.MethodPost, "/fapi/v1/countdownCancelAll", params, true, true, recv)
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
	result := CountDownDeleteOrderResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//GetOpenOrder get open order
func (t *TradeClient) GetOpenOrder(symbol, origClientOrderID string, orderID int64, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" {
		err = errors.New(SymbolEmpty)
		return nil, err
	}
	if orderID <= 0 && origClientOrderID == "" {
		err = errors.New(MissParameters)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s", symbol)
	if origClientOrderID != "" {
		params += fmt.Sprintf("&origClientOrderId=%s", origClientOrderID)
	}
	if orderID > 0 {
		params += fmt.Sprintf("&orderId=%d", orderID)
	}
	req, err := t.Builder.Build(http.MethodDelete, "/fapi/v1/openOrder", params, true, true, recv)
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
	result := OpenOrderResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//GetAllOpenOrder Get all open orders on a symbol. Careful when accessing this with no symbol.
func (t *TradeClient) GetAllOpenOrder(symbol string, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" {
		err = errors.New(SymbolEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s", symbol)
	req, err := t.Builder.Build(http.MethodGet, "/fapi/v1/openOrders", params, true, true, recv)
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
	result := AllOpenOrderResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//GetAllOrder Get all account orders; active, canceled, or filled.
func (t *TradeClient) GetAllOrder(symbol string, orderID, startTime, endTime int64, limit int32, recv time.Duration) (interface{}, error) {
	var err error
	params := ""
	if symbol != "" {
		params = fmt.Sprintf("symbol=%s", symbol)
	}
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	if limit > 1000 {
		limit = 1000
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	if orderID > 0 {
		params += fmt.Sprintf("&orderId=%d", orderID)
	}
	req, err := t.Builder.Build(http.MethodGet, "/fapi/v1/allOrders", params, true, true, recv)
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
	result := AllOrderResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//GetBalance get user balance
func (t *TradeClient) GetBalance(recv time.Duration) (interface{}, error) {
	var err error
	req, err := t.Builder.Build(http.MethodGet, "/fapi/v2/balance", "", true, true, recv)
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
	result := UserBalanceResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//GetUserInfo Get current account information.
func (t *TradeClient) GetUserInfo(recv time.Duration) (interface{}, error) {
	var err error
	req, err := t.Builder.Build(http.MethodGet, "/fapi/v2/account", "", true, true, recv)
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
	result := UserInfoResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//Leverage Change user's initial leverage of specific symbol market.
func (t *TradeClient) Leverage(symbol string, leverage int32, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" {
		err = errors.New(SymbolEmpty)
		return nil, err
	}
	if leverage <= 0 {
		err = errors.New(LeverageInvalid)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s&leverage=%d", symbol, leverage)
	req, err := t.Builder.Build(http.MethodPost, "/fapi/v1/leverage", params, true, true, recv)
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
	result := LeverageResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//MarginType change margin type
func (t *TradeClient) MarginType(symbol, marginType string, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" {
		err = errors.New(SymbolEmpty)
		return nil, err
	}
	if marginType == "" {
		err = errors.New(marginType)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s&marginType=%s", symbol, marginType)
	req, err := t.Builder.Build(http.MethodPost, "/fapi/v1/marginType", params, true, true, recv)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	result := model.APIErrorResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//IsolatedPositionMargin Only for isolated symbol
func (t *TradeClient) IsolatedPositionMargin(symbol, positionSide string, amount decimal.Decimal, sideType int32, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" || amount.LessThanOrEqual(decimal.NewFromInt(0)) || (sideType <= 0 && sideType > 2) {
		err = errors.New(MissParameters)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s&amount=%s&type=%d", symbol, amount, sideType)
	if positionSide != "" {
		params += fmt.Sprintf("&positionSide=%s", positionSide)
	}
	req, err := t.Builder.Build(http.MethodPost, "/fapi/v1/positionMargin", params, true, true, recv)
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
	result := IsolatedPositionMarginResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//IsolatedPositionHistory get isolated margin position history
func (t *TradeClient) IsolatedPositionHistory(symbol string, sideType, limit int32, startTime, endTime int64, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" {
		err = errors.New(SymbolEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s", symbol)
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	if sideType > 0 && sideType < 2 {
		params += fmt.Sprintf("&type=%d", sideType)
	}
	req, err := t.Builder.Build(http.MethodGet, "/fapi/v1/positionMargin/history", params, true, true, recv)
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
	result := IsolatedPositionHistoryResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//PositionRisk Get current position information.
func (t *TradeClient) PositionRisk(symbol string, recv time.Duration) (interface{}, error) {
	var err error
	params := ""
	if symbol != "" {
		params = fmt.Sprintf("symbol=%s", symbol)
	}
	req, err := t.Builder.Build(http.MethodGet, "/fapi/v2/positionRisk", params, true, true, recv)
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
	result := PositionRiskResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//TradeLists Get trades for a specific account and symbol.
func (t *TradeClient) TradeLists(symbol string, startTime, endTime, formID int64, limit int32, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" {
		err = errors.New(SymbolEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s", symbol)
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	if limit > 1000 {
		limit = 1000
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	if formID > 0 {
		params += fmt.Sprintf("&fromId=%d", formID)
	}
	req, err := t.Builder.Build(http.MethodGet, "/fapi/v1/userTrades", params, true, true, recv)
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
	result := TradeListResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//Incomes Get user profit and loss
func (t *TradeClient) Incomes(symbol, incomeType string, startTime, endTime int64, limit int32, recv time.Duration) (interface{}, error) {
	var err error
	params := ""
	if symbol != "" {
		params = fmt.Sprintf("symbol=%s", symbol)
	}
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	if limit > 1000 {
		limit = 1000
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	if incomeType != "" {
		params += fmt.Sprintf("&incomeType=%s", incomeType)
	}
	req, err := t.Builder.Build(http.MethodGet, "/fapi/v1/income", params, true, true, recv)
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
	result := IncomeResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//LeverageBracket Notional and Leverage Brackets
func (t *TradeClient) LeverageBracket(symbol string, recv time.Duration) (interface{}, error) {
	var err error
	params := ""
	if symbol != "" {
		params = fmt.Sprintf("symbol=%s", symbol)
	}
	req, err := t.Builder.Build(http.MethodGet, "/fapi/v1/leverageBracket", params, true, true, recv)
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
	if symbol != "" {
		result := LeverageBracketResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	var result []LeverageBracketResponse
	err = json.Unmarshal(res, &result)
	return result, err
}

//ADLQuantile adl quantile
func (t *TradeClient) ADLQuantile(symbol string, recv time.Duration) (interface{}, error) {
	var err error
	params := ""
	if symbol != "" {
		params = fmt.Sprintf("symbol=%s", symbol)
	}
	req, err := t.Builder.Build(http.MethodGet, "/fapi/v1/leverageBracket", params, true, true, recv)
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
	result := ADLQuantileResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//ForceOrders User's Force Orders
func (t *TradeClient) ForceOrders(symbol, autoCloseType string, startTime, endTime int64, limit int32, recv time.Duration) (interface{}, error) {
	var err error
	params := ""
	if symbol != "" {
		params = fmt.Sprintf("symbol=%s", symbol)
	}
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	if limit > 100 {
		limit = 100
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	if autoCloseType != "" {
		params += fmt.Sprintf("&autoCloseType=%s", autoCloseType)
	}
	req, err := t.Builder.Build(http.MethodGet, "/fapi/v1/forceOrders", params, true, true, recv)
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
	result := ForceOrdersResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//QuantitativeTradingRules quantitative trading rules
func (t *TradeClient) QuantitativeTradingRules(symbol string, recv time.Duration) (interface{}, error) {
	var err error
	params := ""
	if symbol != "" {
		params = fmt.Sprintf("symbol=%s", symbol)
	}
	req, err := t.Builder.Build(http.MethodGet, "/fapi/v1/apiTradingStatus", params, true, true, recv)
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
	result := QuantitativeTradingRulesResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//CommissionRate user commission rate
func (t *TradeClient) CommissionRate(symbol string, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" {
		err = errors.New(SymbolEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s", symbol)
	req, err := t.Builder.Build(http.MethodGet, "/fapi/v1/commissionRate", params, true, true, recv)
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
	result := CommissionRateResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}
