package main

import (
	"fmt"
	"github.com/dirname/Binance/futures/usdt/websocket/market"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"time"
)

func main() {
	client := futuresusdt.NewUSDTFuturesCandlestickWebsocketClient("btcusdt@kline_1m")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@kline_1m", "btcusdt@kline_5m")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresusdt.CandlestickResponse:
			logger.Info("Candlestick Response: %v", response.(futuresusdt.CandlestickResponse))
		case futuresusdt.CandlestickCombinedResponse:
			logger.Info("CandlestickCombinedResponse: %v", response.(futuresusdt.CandlestickCombinedResponse))
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

	client.Unsubscribe(123, "btcusdt@kline_1m", "btcusdt@kline_5m")
	client.Close()
	logger.Info("Client closed")
}
