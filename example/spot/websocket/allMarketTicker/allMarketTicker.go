package main

import (
	"fmt"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	spotclient "github.com/dirname/Binance/spot/websocket/market"
)

func main() {
	client := spotclient.NewSpotAllMarketTickerWebsocketClient("!ticker@arr")
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
