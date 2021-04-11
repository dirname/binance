package main

import (
	"github.com/dirname/Binance/config"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"github.com/dirname/Binance/spot/client"
)

var tradeClient *spotclient.TradeClient

func init() {
	logger.Enable(false)
	tradeClient = spotclient.NewTradeClient(config.SpotRestHost, config.AppKey, config.AppSecret)
}

func main() {
	testNewOrder()
}

func testNewOrder() {
	response, err := tradeClient.TestNewOrder()
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
