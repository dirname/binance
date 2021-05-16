package main

import (
	"fmt"
	"github.com/dirname/binance/futures/usd/websocket/market"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"time"
)

func main() {
	client := futuresclient.NewUSDFuturesSymbolMiniTickerWebsocketClient("btcusdt@miniTicker")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@miniTicker", "ltcusdt@miniTicker")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresclient.SymbolMiniTickerResponse:
			logger.Info("SymbolMiniTicker: %v", response.(futuresclient.SymbolMiniTickerResponse))
		case futuresclient.SymbolMiniTickerCombinedResponse:
			logger.Info("SymbolMiniTickerCombinedResponse: %v", response.(futuresclient.SymbolMiniTickerCombinedResponse))
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
