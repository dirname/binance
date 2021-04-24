package main

import (
	"fmt"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"github.com/dirname/binance/spot/websocket/market"
	"time"
)

func main() {
	client := spotclient.NewSpotSymbolMiniTickerWebsocketClient("btcusdt@miniTicker")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@miniTicker", "ltcusdt@miniTicker")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case spotclient.SymbolMiniTickerResponse:
			logger.Info("SymbolMiniTicker Response: %v", response.(spotclient.SymbolMiniTickerResponse))
		case spotclient.SymbolMiniTickerCombinedResponse:
			logger.Info("SymbolMiniTickerCombinedResponse: %v", response.(spotclient.SymbolMiniTickerCombinedResponse))
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

	client.Unsubscribe(123, "btcusdt@miniTicker", "ltcusdt@miniTicker")
	client.Close()
	logger.Info("Client closed")

}
