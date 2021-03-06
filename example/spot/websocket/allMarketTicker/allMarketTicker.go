package main

import (
	"fmt"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	spotclient "github.com/dirname/binance/spot/websocket/market"
	"time"
)

func main() {
	client := spotclient.NewSpotAllMarketTickerWebsocketClient("!ticker@arr")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "!ticker@arr")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case spotclient.AllMarketTickerResponse:
			logger.Info("AllMarketTicker Response: %v", response.(spotclient.AllMarketTickerResponse))
		case spotclient.AllMarketTickerCombinedResponse:
			logger.Info("AllMarketTickerCombinedResponse: %v", response.(spotclient.AllMarketTickerCombinedResponse))
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

	client.Unsubscribe(123, "!ticker@arr")
	client.Close()
	logger.Info("Client closed")
}
