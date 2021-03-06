package main

import (
	"fmt"
	"github.com/dirname/binance/futures/usd/websocket/market"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"time"
)

func main() {
	client := futuresclient.NewUSDFuturesCompositeIndexWebsocketClient("defiusdt@compositeIndex")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "defiusdt@compositeIndex", "defiusdt@compositeIndex")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresclient.CompositeIndexResponse:
			logger.Info("CompositeIndex: %v", response.(futuresclient.CompositeIndexResponse))
		case futuresclient.CompositeIndexCombinedResponse:
			logger.Info("CompositeIndexCombinedResponse: %v", response.(futuresclient.CompositeIndexCombinedResponse))
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

	client.Unsubscribe(123, "defiusdt@compositeIndex", "defiusdt@compositeIndex")
	client.Close()
	logger.Info("Client closed")
}
