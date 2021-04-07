package main

import (
	"fmt"
	"github.com/dirname/Binance/futures/usdt"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
)

func main() {
	client := FuturesUSDT.NewUSDTFuturesAggTradeWebsocketClient("btcusdt@aggTrade")
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@aggTrade", "ltcusdt@aggTrade")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case FuturesUSDT.AggTradeResponse:
			logger.Info("AggTrade Response: %v", response.(FuturesUSDT.AggTradeResponse))
		case FuturesUSDT.AggTradeCombinedResponse:
			logger.Info("AggTradeCombinedResponse: %v", response.(FuturesUSDT.AggTradeCombinedResponse))
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

	client.Unsubscribe(123, "btcusdt@aggTrade", "ltcusdt@aggTrade")
	client.Close()
	logger.Info("Client closed")

}
