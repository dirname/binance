package main

import (
	"fmt"
	"github.com/dirname/binance/futures/usd/websocket/market"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"time"
)

func main() {
	client := futuresclient.NewUSDFuturesAllMarketLiquidationOrderWebsocketClient("!forceOrder@arr")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		//client.Subscribe(123, "!forceOrder@arr")
		//client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresclient.AllMarketLiquidationOrderResponse:
			logger.Info("AllMarketLiquidationOrder Response: %v", response.(futuresclient.AllMarketLiquidationOrderResponse))
		case futuresclient.AllMarketLiquidationOrderCombinedResponse:
			logger.Info("AllMarketLiquidationOrderCombinedResponse: %v", response.(futuresclient.AllMarketLiquidationOrderCombinedResponse))
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
