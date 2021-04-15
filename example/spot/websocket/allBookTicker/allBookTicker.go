package main

import (
	"fmt"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	spotclient "github.com/dirname/Binance/spot/websocket/market"
	"time"
)

func main() {
	client := spotclient.NewSpotAllBookTickerWebsocketClient("!bookTicker")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		//client.Subscribe(123, "!bookTicker")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case spotclient.AllBookTickerResponse:
			logger.Info("AllBookTicker Response: %v", response.(spotclient.AllBookTickerResponse))
		case spotclient.AllBookTickerCombinedResponse:
			logger.Info("AllBookTickerCombinedResponse: %v", response.(spotclient.AllBookTickerCombinedResponse))
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

	//client.Unsubscribe(123, "!bookTicker")
	client.Close()
	logger.Info("Client closed")
}
