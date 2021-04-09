package main

import (
	"fmt"
	"github.com/dirname/Binance/futures/usdt/websocket/market"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
)

func main() {
	client := futuresusdt.NewUSDTFuturesAllMarketMiniTickerWebsocketClient("!miniTicker@arr")
	client.SetHandler(func() {
		client.Subscribe(123, "!miniTicker@arr")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresusdt.AllMarketMiniTickerResponse:
			logger.Info("AllMarketMiniTicker Response: %v", response.(futuresusdt.AllMarketMiniTickerResponse))
		case futuresusdt.AllMarketMiniTickerCombinedResponse:
			logger.Info("AllMarketMiniTickerCombinedResponse: %v", response.(futuresusdt.AllMarketMiniTickerCombinedResponse))
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
