package main

import (
	"fmt"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"github.com/dirname/Binance/spot/websocket/market"
	"time"
)

func main() {
	client := spotclient.NewSpotTradeWebsocketClient("btcusdt@trade")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@trade", "ltcusdt@trade")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case spotclient.TradeResponse:
			logger.Info("Trade Response: %v", response.(spotclient.TradeResponse))
		case spotclient.TradeCombinedResponse:
			logger.Info("TradeCombinedResponse: %v", response.(spotclient.TradeCombinedResponse))
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

	client.Unsubscribe(123, "btcusdt@trade", "ltcusdt@trade")
	client.Close()
	logger.Info("Client closed")

}
