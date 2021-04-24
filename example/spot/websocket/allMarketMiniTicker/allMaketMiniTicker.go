package main

import (
	"fmt"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	spotclient "github.com/dirname/binance/spot/websocket/market"
	"time"
)

func main() {
	client := spotclient.NewSpotAllMarketMiniTickerWebsocketClient("!miniTicker@arr")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "!miniTicker@arr")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case spotclient.AllMarketMiniTickerResponse:
			logger.Info("AllMarketMiniTicker Response: %v", response.(spotclient.AllMarketMiniTickerResponse))
		case spotclient.AllMarketMiniTickerCombinedResponse:
			logger.Info("AllMarketMiniTickerCombinedResponse: %v", response.(spotclient.AllMarketMiniTickerCombinedResponse))
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

	client.Unsubscribe(123, "!miniTicker@arr")
	client.Close()
	logger.Info("Client closed")
}
