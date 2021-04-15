package main

import (
	"fmt"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"github.com/dirname/Binance/spot/websocket/market"
	"time"
)

func main() {
	client := spotclient.NewSpotPartialBookDepthWebsocketClient("btcusdt@depth10")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@depth10", "ltcusdt@depth10")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case spotclient.PartialBookDepthResponse:
			logger.Info("PartialBookTicker Response: %v", response.(spotclient.PartialBookDepthResponse))
		case spotclient.PartialBookDepthCombinedResponse:
			logger.Info("PartialBookTickerCombinedResponse: %v", response.(spotclient.PartialBookDepthCombinedResponse))
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

	client.Unsubscribe(123, "btcusdt@depth10", "ltcusdt@depth10")
	client.Close()
	logger.Info("Client closed")

}
