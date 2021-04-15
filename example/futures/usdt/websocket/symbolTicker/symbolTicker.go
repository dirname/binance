package main

import (
	"fmt"
	"github.com/dirname/Binance/futures/usdt/websocket/market"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"time"
)

func main() {
	client := futuresusdt.NewUSDTFuturesSymbolTickerWebsocketClient("btcusdt@ticker")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@ticker", "ltcusdt@ticker")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresusdt.SymbolTickerResponse:
			logger.Info("SymbolTicker: %v", response.(futuresusdt.SymbolTickerResponse))
		case futuresusdt.SymbolTickerCombinedResponse:
			logger.Info("SymbolTickerCombinedResponse: %v", response.(futuresusdt.SymbolTickerCombinedResponse))
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
