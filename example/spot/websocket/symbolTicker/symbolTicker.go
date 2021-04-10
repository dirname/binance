package main

import (
	"fmt"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"github.com/dirname/Binance/spot/websocket/market"
)

func main() {
	client := spotclient.NewSpotSymbolTickerWebsocketClient("btcusdt@ticker")
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@ticker", "ltcusdt@ticker")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case spotclient.SymbolTickerResponse:
			logger.Info("SymbolTicker Response: %v", response.(spotclient.SymbolTickerResponse))
		case spotclient.SymbolTickerCombinedResponse:
			logger.Info("SymbolTickerCombinedResponse: %v", response.(spotclient.SymbolTickerCombinedResponse))
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

	client.Unsubscribe(123, "btcusdt@ticker", "ltcusdt@ticker")
	client.Close()
	logger.Info("Client closed")

}
