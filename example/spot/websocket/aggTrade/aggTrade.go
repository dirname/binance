package main

import (
	"fmt"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"github.com/dirname/Binance/spot/websocket/market"
	"time"
)

func main() {
	client := spotclient.NewSpotAggTradeWebsocketClient("btcusdt@aggTrade")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@aggTrade", "ltcusdt@aggTrade")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case spotclient.AggTradeResponse:
			logger.Info("AggTrade Response: %v", response.(spotclient.AggTradeResponse))
		case spotclient.AggTradeCombinedResponse:
			logger.Info("AggTradeCombinedResponse: %v", response.(spotclient.AggTradeCombinedResponse))
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

	client.Unsubscribe(123, "btcusdt@aggTrade", "ltcusdt@aggTrade")
	client.Close()
	logger.Info("Client closed")

}
