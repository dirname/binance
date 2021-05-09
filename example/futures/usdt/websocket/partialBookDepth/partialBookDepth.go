package main

import (
	"fmt"
	"github.com/dirname/binance/futures/usd/websocket/market"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"time"
)

func main() {
	client := futuresusdt.NewUSDFuturesPartialBookDepthWebsocketClient("btcusdt@depth10")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@depth10", "ltcusdt@depth10")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresusdt.PartialBookDepthResponse:
			logger.Info("PartialBookDepth: %v", response.(futuresusdt.PartialBookDepthResponse))
		case futuresusdt.PartialBookDepthCombinedResponse:
			logger.Info("PartialBookDepthCombinedResponse: %v", response.(futuresusdt.PartialBookDepthCombinedResponse))
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
