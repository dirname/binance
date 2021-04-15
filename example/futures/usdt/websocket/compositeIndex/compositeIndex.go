package main

import (
	"fmt"
	"github.com/dirname/Binance/futures/usdt/websocket/market"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"time"
)

func main() {
	client := futuresusdt.NewUSDTFuturesCompositeIndexWebsocketClient("defiusdt@compositeIndex")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "defiusdt@compositeIndex", "defiusdt@compositeIndex")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresusdt.CompositeIndexResponse:
			logger.Info("CompositeIndex: %v", response.(futuresusdt.CompositeIndexResponse))
		case futuresusdt.CompositeIndexCombinedResponse:
			logger.Info("CompositeIndexCombinedResponse: %v", response.(futuresusdt.CompositeIndexCombinedResponse))
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
