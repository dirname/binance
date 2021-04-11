package spotclient

import (
	"encoding/json"
	binance "github.com/dirname/Binance"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"github.com/shopspring/decimal"
	"net/http"
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
	if _, ok := parser["fills"]; ok {
		result := NewOrderResponseFull{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	if _, ok := parser["status"]; ok {
		result := NewOrderResponseResult{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := NewOrderResponseACK{}
	err = json.Unmarshal(res, &result)
	return parser, err
}
