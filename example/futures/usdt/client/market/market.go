package main

import (
	"github.com/dirname/binance/config"
	"github.com/dirname/binance/futures/usd/client"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
)

var marketClient *futuresclient.MarketClient

func init() {
	logger.Enable(false)
	marketClient = futuresclient.NewMarketClient(config.USDFuturesRestHost, config.AppKey)
}

func main() {
	getExchangeInfo()
	getOrderBook()
	getRecentOrder()
}

func getExchangeInfo() {
	response, err := marketClient.GetExchangeInfo()
	if err != nil {
		logger.Error("getExchangeInfo err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getExchangeInfo API error: %v", response.(model.APIErrorResponse))
	case futuresclient.ExchangeInfoResponse:
		logger.Info("getExchangeInfo: %v", response.(futuresclient.ExchangeInfoResponse))
	default:
		logger.Info("getExchangeInfo Unknown response: %v", response)
	}
}

func getOrderBook() {
	response, err := marketClient.GetOrderBook("BTCUSDT", 5)
	if err != nil {
		logger.Error("getOrderBook err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getOrderBook API error: %v", response.(model.APIErrorResponse))
	case futuresclient.OrderBookResponse:
		logger.Info("getOrderBook: %v", response.(futuresclient.OrderBookResponse))
	default:
		logger.Info("getOrderBook Unknown response: %v", response)
	}
}

func getRecentOrder() {
	response, err := marketClient.GetRecentTrades("BTCUSDT", 5)
	if err != nil {
		logger.Error("getRecentOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getRecentOrder API error: %v", response.(model.APIErrorResponse))
	case futuresclient.RecentTradesListResponse:
		logger.Info("getRecentOrder: %v", response.(futuresclient.RecentTradesListResponse))
	default:
		logger.Info("getRecentOrder Unknown response: %v", response)
	}
}
