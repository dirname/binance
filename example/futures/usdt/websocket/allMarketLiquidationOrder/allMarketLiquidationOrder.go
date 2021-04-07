package main

import (
	"fmt"
	"github.com/dirname/Binance/futures/usdt/websocket"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
)

func main() {
	client := futuresusdt.NewUSDTFuturesAllMarketLiquidationOrderWebsocketClient("!forceOrder@arr")
	client.SetHandler(func() {
		//client.Subscribe(123, "!forceOrder@arr")
		//client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresusdt.AllMarketLiquidationOrderResponse:
			logger.Info("AllMarketLiquidationOrder Response: %v", response.(futuresusdt.AllMarketLiquidationOrderResponse))
		case futuresusdt.AllMarketLiquidationOrderCombinedResponse:
			logger.Info("AllMarketLiquidationOrderCombinedResponse: %v", response.(futuresusdt.AllMarketLiquidationOrderCombinedResponse))
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

	//client.Unsubscribe(123, "!forceOrder@arr")
	client.Close()
	logger.Info("Client closed")
}
