package main

import (
	"fmt"
	"github.com/dirname/binance/futures/usd/websocket/market"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"time"
)

func main() {
	client := futuresusdt.NewUSDFuturesAggTradeWebsocketClient("btcusdt@aggTrade")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@aggTrade", "ltcusdt@aggTrade")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresusdt.AggTradeResponse:
			logger.Info("AggTrade Response: %v", response.(futuresusdt.AggTradeResponse))
		case futuresusdt.AggTradeCombinedResponse:
			logger.Info("AggTradeCombinedResponse: %v", response.(futuresusdt.AggTradeCombinedResponse))
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
