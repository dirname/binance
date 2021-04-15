package main

import (
	"fmt"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"github.com/dirname/Binance/spot/websocket/market"
	"time"
)

func main() {
	client := spotclient.NewSpotDiffDepthDepthWebsocketClient("btcusdt@depth@100ms")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@depth@100ms", "ltcusdt@depth@100ms")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case spotclient.DiffBookDepthResponse:
			logger.Info("DiffDepth Response: %v", response.(spotclient.DiffBookDepthResponse))
		case spotclient.DiffBookDepthCombinedResponse:
			logger.Info("DiffDepthCombinedResponse: %v", response.(spotclient.DiffBookDepthCombinedResponse))
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

	client.Unsubscribe(123, "btcusdt@depth@100ms", "ltcusdt@depth@100ms")
	client.Close()
	logger.Info("Client closed")

}
