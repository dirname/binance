package main

import (
	"fmt"
	"github.com/dirname/binance/futures/usd/websocket/market"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"time"
)

func main() {
	client := futuresclient.NewUSDFuturesBLVTCandlestickWebsocketClient("TRXDOWN@nav_kline_1m")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "TRXDOWN@nav_kline_5m", "TRXDOWN@nav_kline_1m")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresclient.BLVTCandlestickResponse:
			logger.Info("BLVTCandlestick Response: %v", response.(futuresclient.BLVTCandlestickResponse))
		case futuresclient.BLVTCandlestickCombinedResponse:
			logger.Info("BLVTCandlestickCombinedResponse: %v", response.(futuresclient.BLVTCandlestickCombinedResponse))
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

	client.Unsubscribe(123, "TRXDOWN@nav_kline_5m", "TRXDOWN@nav_kline_1m")
	client.Close()
	logger.Info("Client closed")
}
