package main

import (
	"fmt"
	"github.com/dirname/binance/futures/usd/websocket/market"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"time"
)

func main() {
	client := futuresclient.NewUSDFuturesDiffBookDepthWebsocketClient("btcusdt@depth")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@depth", "ltcusdt@depth")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresclient.DiffBookDepthResponse:
			logger.Info("DiffBookDepth Response: %v", response.(futuresclient.DiffBookDepthResponse))
		case futuresclient.DiffBookDepthCombinedResponse:
			logger.Info("DiffBookDepthCombinedResponse: %v", response.(futuresclient.DiffBookDepthCombinedResponse))
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

	client.Unsubscribe(123, "btcusdt@depth", "ltcusdt@depth")
	client.Close()
	logger.Info("Client closed")
}
