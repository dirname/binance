package main

import (
	"fmt"
	"github.com/dirname/Binance/futures/usdt/websocket/market"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"time"
)

func main() {
	client := futuresusdt.NewUSDTFuturesSymbolBookTickerWebsocketClient("btcusdt@bookTicker")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@bookTicker", "ltcusdt@bookTicker")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresusdt.SymbolBookTickerResponse:
			logger.Info("SymbolBookTicker: %v", response.(futuresusdt.SymbolBookTickerResponse))
		case futuresusdt.SymbolBookTickerCombinedResponse:
			logger.Info("SymbolBookTickerCombinedResponse: %v", response.(futuresusdt.SymbolBookTickerCombinedResponse))
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

	client.Unsubscribe(123, "btcusdt@bookTicker", "ltcusdt@bookTicker")
	client.Close()
	logger.Info("Client closed")
}
