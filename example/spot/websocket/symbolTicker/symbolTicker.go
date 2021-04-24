package main

import (
	"fmt"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"github.com/dirname/binance/spot/websocket/market"
	"time"
)

func main() {
	client := spotclient.NewSpotSymbolTickerWebsocketClient("btcusdt@ticker")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
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
