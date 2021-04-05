package main

import (
	"fmt"
	"github.com/dirname/Binance/futures/usdt"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
)

func main() {
	client := usdt.NewUSDTFuturesMarketPriceWebsocketClient("btcusdt@markPrice")
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@markPrice", "ltcusdt@markPrice")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case usdt.MarketPriceResponse:
			logger.Info("MarketPrice: %v", response.(usdt.MarketPriceResponse))
		case usdt.MarketPriceCombinedResponse:
			logger.Info("MarketCombinedResponse: %v", response.(usdt.MarketPriceCombinedResponse))
		case model.WebsocketCommonResponse:
			logger.Info("Websocket Response: %v", response.(model.WebsocketCommonResponse))
		case model.WebsocketErrorResponse:
			logger.Info("Websocket Error Response: %v", response.(model.WebsocketErrorResponse))
		default:
			logger.Info("Unknown Response: %v", response)
		}
	})
	client.Connect(true)
	fmt.Scanln()

	client.Unsubscribe(123, "btcusdt@markPrice", "ltcusdt@markPrice")
	client.Close()
	logger.Info("Client closed")
}
