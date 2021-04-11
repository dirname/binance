package main

import (
	"github.com/dirname/Binance/config"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"github.com/dirname/Binance/spot/client"
	"github.com/dirname/Binance/spot/client/orderSide"
	"github.com/dirname/Binance/spot/client/orderType"
	"github.com/dirname/Binance/spot/client/timeInForce"
	"github.com/shopspring/decimal"
)

var tradeClient *spotclient.TradeClient

func init() {
	logger.Enable(false)
	tradeClient = spotclient.NewTradeClient(config.SpotRestHost, config.AppKey, config.AppSecret)
}

func main() {
	//testNewOrder()
	//newOrder()
}

func testNewOrder() {
	response, err := tradeClient.TestNewOrder("BTCUSDT", orderSide.Buy, orderType.Limit, timeInForce.FOK, "", "", decimal.NewFromFloat(0.001), decimal.Decimal{}, decimal.NewFromFloat(59700.12), decimal.Decimal{}, decimal.Decimal{}, 0)
	if err != nil {
		logger.Error("testNewOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("testNewOrder API error: %v", response.(model.APIErrorResponse))
	default:
		logger.Info("testNewOrder response: %v", response)
	}
}

func newOrder() {
	response, err := tradeClient.NewOrder("CFXUSDT", orderSide.Sell, orderType.Limit, timeInForce.GTC, "", "", decimal.NewFromFloat(10.0), decimal.Decimal{}, decimal.NewFromFloat(1.205), decimal.Decimal{}, decimal.Decimal{}, 0)
	if err != nil {
		logger.Error("testNewOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("newOrder API error: %v", response.(model.APIErrorResponse))
	case spotclient.NewOrderResponseACK:
		logger.Info("newOrder ACK: %v", response.(spotclient.NewOrderResponseACK))
	case spotclient.NewOrderResponseFull:
		logger.Info("newOrder Full: %v", response.(spotclient.NewOrderResponseFull))
	case spotclient.NewOrderResponseResult:
		logger.Info("newOrder RESULT: %v", response.(spotclient.NewOrderResponseResult))
	default:
		logger.Info("newOrder response: %v", response)
	}
}
