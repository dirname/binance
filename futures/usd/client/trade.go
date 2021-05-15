package futuresclient

import (
	"encoding/json"
	"errors"
	"github.com/dirname/binance"
	"github.com/dirname/binance/futures/usd/client/orderRespType"
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
func (t *TradeClient) PositionSideDual(symbol string, recv time.Duration) (interface{}, error) {
	var err error
	if symbol == "" {
		err = errors.New(MissParameters)
		return nil, err
	}
	params := strings.Join([]string{"symbol=", symbol}, "")
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
	req, err := t.Builder.Build(http.MethodPost, "/fapi/v1/positionSide/dual", params, true, true, recv)
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

func (t *TradeClient) NewOrder(symbol, side, positionSide, ordersType, reduceOnly, newClientOrderID, closePosition, timeInForce, workingType, priceProtect, newOrderRespType string, quantity, price, stopPrice, activationPrice, callbackRate decimal.Decimal, recv time.Duration) (interface{}, error) {
	var err error
	params, err := buildOrder(symbol, side, positionSide, ordersType, reduceOnly, newClientOrderID, closePosition, timeInForce, workingType, priceProtect, newOrderRespType, quantity, price, stopPrice, activationPrice, callbackRate)
	if err != nil {
		return nil, err
	}
	req, err := t.Builder.Build(http.MethodPost, "/fapi/v3/order", params, true, true, recv)
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
	case orderRespType.ACK:
		result := NewOrderResponseACK{}
		err = json.Unmarshal(res, &result)
		return parser, err
	}
	return parser, err
}
