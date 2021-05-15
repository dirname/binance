package spotclient

import (
	"encoding/json"
	"errors"
	"fmt"
	binance "github.com/dirname/binance"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"github.com/dirname/binance/spot/client/define/orderRespType"
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

// TestNewOrder Test new order creation and signature/recvWindow long. Creates and validates a new order but does not send it into the matching engine.
func (t *TradeClient) TestNewOrder(symbol, side, orderType, timeInForce, newClientOderID, newOrderRespType string, quantity, quoteOrderQTY, price, stopPrice, icebergQTY decimal.Decimal, recv time.Duration) (interface{}, error) {
	var err error
	params, err := buildOrder(symbol, side, orderType, timeInForce, newClientOderID, newOrderRespType, quantity, quoteOrderQTY, price, stopPrice, icebergQTY)
	if err != nil {
		return nil, err
	}
	req, err := t.Builder.Build(http.MethodPost, "/api/v3/order/test", params, true, true, recv)
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

// NewOrder Send in a new order.
func (t *TradeClient) NewOrder(symbol, side, orderType, timeInForce, newClientOderID, newOrderRespType string, quantity, quoteOrderQTY, price, stopPrice, icebergQTY decimal.Decimal, recv time.Duration) (interface{}, error) {
	var err error
	params, err := buildOrder(symbol, side, orderType, timeInForce, newClientOderID, newOrderRespType, quantity, quoteOrderQTY, price, stopPrice, icebergQTY)
	if err != nil {
		return nil, err
	}
	req, err := t.Builder.Build(http.MethodPost, "/api/v3/order", params, true, true, recv)
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
	case orderRespType.Result:
		result := NewOrderResponseResult{}
		err = json.Unmarshal(res, &result)
		return result, err
	case orderRespType.Full:
		result := NewOrderResponseFull{}
		err = json.Unmarshal(res, &result)
		return result, err
	case orderRespType.ACK:
		result := NewOrderResponseACK{}
		err = json.Unmarshal(res, &result)
		return parser, err
	}
	return parser, err
}

// DeleteOrder delete a order
func (t *TradeClient) DeleteOrder(symbol, origClientOrderID, newClientOrderID string, orderID int64, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" || orderID <= 0 && origClientOrderID == "" {
		err = errors.New(MissParameters)
		return nil, err
	}
	params := strings.Join([]string{"symbol=", symbol}, "")
	if origClientOrderID != "" {
		params += strings.Join([]string{"&origClientOrderId=", origClientOrderID}, "")
	}
	if orderID > 0 {
		params += fmt.Sprintf("&orderId=%d", orderID)
	}
	params += buildNewClientOderID(newClientOrderID)
	req, err := t.Builder.Build(http.MethodDelete, "/api/v3/order", params, true, true, recv)
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

// DeleteOpenOrders delete all open orders
func (t *TradeClient) DeleteOpenOrders(symbol string, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" {
		err = errors.New(MissParameters)
		return nil, err
	}
	params := strings.Join([]string{"symbol=", symbol}, "")
	req, err := t.Builder.Build(http.MethodDelete, "/api/v3/openOrders", params, true, true, recv)
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
	result := DeleteOpenOrderResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// GetOrder get a order info
func (t *TradeClient) GetOrder(symbol, origClientOrderID string, orderID int64, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" || orderID <= 0 && origClientOrderID == "" {
		err = errors.New(MissParameters)
		return nil, err
	}
	params := strings.Join([]string{"symbol=", symbol}, "")
	if origClientOrderID != "" {
		params += strings.Join([]string{"&origClientOrderId=", origClientOrderID}, "")
	}
	if orderID > 0 {
		params += fmt.Sprintf("&orderId=%d", orderID)
	}
	req, err := t.Builder.Build(http.MethodGet, "/api/v3/order", params, true, true, recv)
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
	result := GetOrderResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// GetOpenOrder get all open orders
func (t *TradeClient) GetOpenOrder(symbol string, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" {
		err = errors.New(MissParameters)
		return nil, err
	}
	params := strings.Join([]string{"symbol=", symbol}, "")
	req, err := t.Builder.Build(http.MethodGet, "/api/v3/openOrders", params, true, true, recv)
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
	result := GetOpenOrdersResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// GetAllOrder get all orders
func (t *TradeClient) GetAllOrder(symbol string, orderID, startTime, endTime int64, limit int32, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" {
		err = errors.New(MissParameters)
		return nil, err
	}
	params := strings.Join([]string{"symbol=", symbol}, "")
	if orderID > 0 {
		params += fmt.Sprintf("&orderId=%d", orderID)
	}
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	if limit > 0 {
		params += fmt.Sprintf("&limi=%d", limit)
	}
	req, err := t.Builder.Build(http.MethodGet, "/api/v3/allOrders", params, true, true, recv)
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
	result := GetAllOrdersResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// NewOCO new oco order
func (t *TradeClient) NewOCO(symbol, listClientOrderID, side, limitClientOrderId, stopClientOrderId, stopLimitTimeInForce, newOrderRespType string, quantity, price, limitIcebergQty, stopPrice, stopLimitPrice, stopIcebergQty decimal.Decimal, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" || side == "" || quantity.LessThanOrEqual(decimal.NewFromInt(0)) || price.LessThanOrEqual(decimal.NewFromInt(0)) || stopPrice.LessThanOrEqual(decimal.NewFromInt(0)) {
		err = errors.New(MissParameters)
		return nil, err
	}
	if !stopLimitPrice.LessThanOrEqual(decimal.NewFromInt(0)) {
		if stopLimitTimeInForce == "" {
			err = errors.New(MissParameters)
			return nil, err
		}
	}
	params := strings.Join([]string{"symbol=", symbol, "&side=", side}, "")
	params += fmt.Sprintf("&quantity=%s&price=%s&stopPrice=%s", quantity, price, stopPrice)
	params += strings.Join([]string{buildStringParam("listClientOrderId", listClientOrderID), buildStringParam("limitClientOrderId", limitClientOrderId), buildStringParam("stopClientOrderId", stopClientOrderId), buildStringParam("stopLimitTimeInForce", stopLimitTimeInForce), buildStringParam("newOrderRespType", newOrderRespType)}, "")
	params += strings.Join([]string{buildDecimalParam("limitIcebergQty", limitIcebergQty), buildDecimalParam("stopLimitPrice", stopLimitPrice), buildDecimalParam("stopIcebergQty", stopIcebergQty)}, "")
	req, err := t.Builder.Build(http.MethodPost, "/api/v3/order/oco", params, true, true, recv)
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
	result := NewOCOOrderResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// DeleteOCO delete oco order
func (t *TradeClient) DeleteOCO(symbol, listClientOrderID, newClientOrderID string, orderListID int64, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" || (orderListID <= 0 && listClientOrderID == "") {
		err = errors.New(MissParameters)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s", symbol)
	params += strings.Join([]string{buildStringParam("listClientOrderId", listClientOrderID), buildStringParam("newClientOrderId", newClientOrderID), buildInt64Param("orderListId", orderListID)}, "")
	req, err := t.Builder.Build(http.MethodDelete, "/api/v3/orderList", params, true, true, recv)
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
	result := DeleteOCOResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// GetOCOOrder get a oco order info
func (t *TradeClient) GetOCOOrder(origClientOrderID string, orderListID int64, recv time.Duration) (interface{}, error) {
	var err error
	params := strings.Join([]string{buildStringParam("origClientOrderId", origClientOrderID), buildInt64Param("orderListId", orderListID)}, "")
	req, err := t.Builder.Build(http.MethodGet, "/api/v3/orderList", params, true, true, recv)
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
	result := GetOCOResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// GetAllOCOOrder get all oco orders info
func (t *TradeClient) GetAllOCOOrder(formID, startTime, endTime int64, limit int32, recv time.Duration) (interface{}, error) {
	var err error
	if formID > 0 && (startTime > 0 || endTime > 0) {
		err = errors.New(InvalidParameters)
		return nil, err
	}
	params := strings.Join([]string{buildInt64Param("formId", formID), buildInt64Param("startTime", startTime), buildInt64Param("endTime", endTime)}, "")
	if limit > 1000 {
		limit = 1000
	}
	if limit < 0 {
		limit = 500
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	req, err := t.Builder.Build(http.MethodGet, "/api/v3/allOrderList", params, true, true, recv)
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
	result := GetAllOCOResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// GetOpenOCOOrder get open oco order
func (t *TradeClient) GetOpenOCOOrder(recv time.Duration) (interface{}, error) {
	var err error
	req, err := t.Builder.Build(http.MethodGet, "/api/v3/openOrderList", "", true, true, recv)
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
	result := GetOpenOCOResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// GetAccountInfo get account information
func (t *TradeClient) GetAccountInfo(recv time.Duration) (interface{}, error) {
	var err error
	req, err := t.Builder.Build(http.MethodGet, "/api/v3/account", "", true, true, recv)
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
	result := AccountInfoResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// GetAccountTradeList get account trade list
func (t *TradeClient) GetAccountTradeList(symbol string, startTime, endTime, formID int64, limit int32, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" {
		err = errors.New(MissParameters)
		return nil, err
	}
	params := strings.Join([]string{"symbol=", symbol, buildInt64Param("startTime", startTime), buildInt64Param("endTime", endTime), buildInt64Param("formId", formID)}, "")
	if limit > 1000 {
		limit = 1000
	}
	if limit < 0 {
		limit = 500
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	req, err := t.Builder.Build(http.MethodGet, "/api/v3/myTrades", params, true, true, recv)
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
	result := AccountTradesListResponse{}
	err = json.Unmarshal(res, &result)
	return result, err

}
