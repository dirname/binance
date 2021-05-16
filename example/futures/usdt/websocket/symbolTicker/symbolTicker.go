package main

import (
	"fmt"
	"github.com/dirname/binance/futures/usd/websocket/market"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"time"
)

func main() {
	client := futuresclient.NewUSDFuturesSymbolTickerWebsocketClient("btcusdt@ticker")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@ticker", "ltcusdt@ticker")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresclient.SymbolTickerResponse:
			logger.Info("SymbolTicker: %v", response.(futuresclient.SymbolTickerResponse))
		case futuresclient.SymbolTickerCombinedResponse:
			logger.Info("SymbolTickerCombinedResponse: %v", response.(futuresclient.SymbolTickerCombinedResponse))
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
