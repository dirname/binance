package main

import (
	"fmt"
	"github.com/dirname/binance/futures/usdt/websocket/market"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"time"
)

func main() {
	client := futuresusdt.NewUSDTFuturesAllMarketTickerWebsocketClient("!ticker@arr")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresusdt.AllMarketTickerResponse:
			logger.Info("AllMarketTicker Response: %v", response.(futuresusdt.AllMarketTickerResponse))
		case futuresusdt.AllMarketTickerCombinedResponse:
			logger.Info("AllMarketTickerCombinedResponse: %v", response.(futuresusdt.AllMarketTickerCombinedResponse))
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

	client.Close()
	logger.Info("Client closed")
}
