package main

import (
	"fmt"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"github.com/dirname/Binance/spot/websocket/market"
	"time"
)

func main() {
	client := spotclient.NewSpotCandlestickWebsocketClient("btcusdt@kline_1m")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@kline_1m", "ltcusdt@kline_1m")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case spotclient.CandlestickResponse:
			logger.Info("Candlestick Response: %v", response.(spotclient.CandlestickResponse))
		case spotclient.CandlestickCombinedResponse:
			logger.Info("CandlestickCombinedResponse: %v", response.(spotclient.CandlestickCombinedResponse))
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

	client.Unsubscribe(123, "btcusdt@kline_1m", "ltcusdt@kline_1m")
	client.Close()
	logger.Info("Client closed")

}
