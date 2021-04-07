package main

import (
	"fmt"
	"github.com/dirname/Binance/futures/usdt/websocket"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
)

func main() {
	client := futuresusdt.NewUSDTFuturesContinuesCandlestickWebsocketClient("btcusdt_perpetual@continuousKline_1m")
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt_perpetual@continuousKline_1m", "btcusdt_perpetual@continuousKline_5m")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresusdt.ContinuesCandlestickResponse:
			logger.Info("ContinuesCandlestick Response: %v", response.(futuresusdt.ContinuesCandlestickResponse))
		case futuresusdt.ContinuesCandlestickCombinedResponse:
			logger.Info("ContinuesCandlestickCombinedResponse: %v", response.(futuresusdt.ContinuesCandlestickCombinedResponse))
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

	client.Unsubscribe(123, "btcusdt_perpetual@continuousKline_1m", "btcusdt_perpetual@continuousKline_5m")
	client.Close()
	logger.Info("Client closed")
}
