package main

import (
	"github.com/dirname/Binance/config"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"github.com/dirname/Binance/spot/client"
)

var marketClient *spotclient.MarketClient

func init() {
	logger.Enable(false)
	marketClient = spotclient.NewMarketClient(config.SpotRestHost, config.AppKey)
}

func main() {
	//getExchangeInfo()
	//getOrderBook()
	//getRecentTrades()
	//getOlderTrades()
	//getAggTrades()
	//getCandlestick()
	//getAveragePrice()
	//get24hTickerPriceChange()
	getSymbolTickerPrice()
	getSymbolOrderBook()
}

func getExchangeInfo() {
	response, err := marketClient.GetExchangeInfo()
	if err != nil {
		logger.Error("getExchangeInfo err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getExchangeInfo API error: %v", response.(model.APIErrorResponse))
	case spotclient.ExchangeInfoResponse:
		logger.Info("getExchangeInfo: %v", response.(spotclient.ExchangeInfoResponse))
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
	case spotclient.OrderBookResponse:
		logger.Info("getOrderBook: %v", response.(spotclient.OrderBookResponse))
	default:
		logger.Info("getOrderBook Unknown response: %v", response)
	}
}

func getRecentTrades() {
	response, err := marketClient.GetRecentTrades("BTCUSDT", 5)
	if err != nil {
		logger.Error("getRecentTrades err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getRecentTrades API error: %v", response.(model.APIErrorResponse))
	case spotclient.RecentTradesListResponse:
		logger.Info("getRecentTrades: %v", response.(spotclient.RecentTradesListResponse))
	default:
		logger.Info("getRecentTrades Unknown response: %v", response)
	}
}

func getOlderTrades() {
	response, err := marketClient.GetOldTradeLookUp("BTCUSDT", 5, 0)
	if err != nil {
		logger.Error("getOlderTrades err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getOlderTrades API error: %v", response.(model.APIErrorResponse))
	case spotclient.OlderTradeLookUpResponse:
		logger.Info("getOlderTrades: %v", response.(spotclient.OlderTradeLookUpResponse))
	default:
		logger.Info("getOlderTrades Unknown response: %v", response)
	}
}

func getAggTrades() {
	response, err := marketClient.GetAggregateTrades("BTCUSDT", 5, 0, 0, 0)
	if err != nil {
		logger.Error("getAggTrades err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getAggTrades API error: %v", response.(model.APIErrorResponse))
	case spotclient.AggregateTradeResponse:
		logger.Info("getAggTrades: %v", response.(spotclient.AggregateTradeResponse))
	default:
		logger.Info("getAggTrades Unknown response: %v", response)
	}
}

func getCandlestick() {
	response, err := marketClient.GetCandlestick("BTCUSDT", "1m", 0, 0, 0)
	if err != nil {
		logger.Error("getCandlestick err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getCandlestick API error: %v", response.(model.APIErrorResponse))
	case spotclient.CandlestickResponse:
		logger.Info("getCandlestick: %v", response.(spotclient.CandlestickResponse))
	default:
		logger.Info("getCandlestick Unknown response: %v", response)
	}
}

func getAveragePrice() {
	response, err := marketClient.GetAveragePrice("BTCUSDT")
	if err != nil {
		logger.Error("getAveragePrice err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getAveragePrice API error: %v", response.(model.APIErrorResponse))
	case spotclient.CurrentAveragePriceResponse:
		logger.Info("getAveragePrice: %v", response.(spotclient.CurrentAveragePriceResponse))
	default:
		logger.Info("getAveragePrice Unknown response: %v", response)
	}
}

func get24hTickerPriceChange() {
	response, err := marketClient.Get24hTickerPriceChange("")
	if err != nil {
		logger.Error("get24hTickerPriceChange err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("get24hTickerPriceChange API error: %v", response.(model.APIErrorResponse))
	case spotclient.TickerPriceChangeStatisticsResponse:
		logger.Info("get24hTickerPriceChange: %v", response.(spotclient.TickerPriceChangeStatisticsResponse))
	case []spotclient.TickerPriceChangeStatisticsResponse:
		logger.Info("get24hTickerPriceChange Array: %v", response.([]spotclient.TickerPriceChangeStatisticsResponse))
	default:
		logger.Info("get24hTickerPriceChange Unknown response: %v", response)
	}
}

func getSymbolTickerPrice() {
	response, err := marketClient.GetSymbolTickerPrice("LTCBTC")
	if err != nil {
		logger.Error("getSymbolTickerPrice err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getSymbolTickerPrice API error: %v", response.(model.APIErrorResponse))
	case spotclient.SymbolPriceTickerResponse:
		logger.Info("getSymbolTickerPrice: %v", response.(spotclient.SymbolPriceTickerResponse))
	case []spotclient.SymbolPriceTickerResponse:
		logger.Info("getSymbolTickerPrice Array: %v", response.([]spotclient.SymbolPriceTickerResponse))
	default:
		logger.Info("getSymbolTickerPrice Unknown response: %v", response)
	}
}

func getSymbolOrderBook() {
	response, err := marketClient.GetSymbolOrderBookTicker("LTCBTC")
	if err != nil {
		logger.Error("getSymbolOrderBook err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getSymbolOrderBook API error: %v", response.(model.APIErrorResponse))
	case spotclient.SymbolOrderBookTickerResponse:
		logger.Info("getSymbolOrderBook: %v", response.(spotclient.SymbolOrderBookTickerResponse))
	case []spotclient.SymbolOrderBookTickerResponse:
		logger.Info("getSymbolOrderBook Array: %v", response.([]spotclient.SymbolOrderBookTickerResponse))
	default:
		logger.Info("getSymbolOrderBook Unknown response: %v", response)
	}
}
