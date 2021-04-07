package main

import (
	"fmt"
	"github.com/dirname/Binance/futures/usdt"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
)

func main() {
	client := futuresusdt.NewUSDTFuturesAllMarketPriceWebsocketClient("!markPrice@arr")
	client.SetHandler(func() {
		client.Subscribe(123, "!markPrice@arr@1s", "!markPrice@3s")
		//client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresusdt.AggTradeResponse:
			logger.Info("AllMarket Response: %v", response.(futuresusdt.AggTradeResponse))
		case futuresusdt.AggTradeCombinedResponse:
			logger.Info("AllMarketCombinedResponse: %v", response.(futuresusdt.AggTradeCombinedResponse))
		case model.WebsocketCommonResponse:
			logger.Info("Websocket Response: %v", response.(model.WebsocketCommonResponse))
		case model.WebsocketErrorResponse:
			logger.Info("Websocket Error Response: %v", response.(model.WebsocketErrorResponse))
		default:
			//logger.Info("Unknown Response: %v", response)
		}
	})
	client.Connect(true)
	fmt.Scanln()

	client.Unsubscribe(123, "!markPrice@arr@1s", "!markPrice@3s")
	client.Close()
	logger.Info("Client closed")
}
