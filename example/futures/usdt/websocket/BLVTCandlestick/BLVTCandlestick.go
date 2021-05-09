package main

import (
	"fmt"
	"github.com/dirname/binance/futures/usd/websocket/market"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"time"
)

func main() {
	client := futuresusdt.NewUSDFuturesBLVTCandlestickWebsocketClient("TRXDOWN@nav_kline_1m")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "TRXDOWN@nav_kline_5m", "TRXDOWN@nav_kline_1m")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresusdt.BLVTCandlestickResponse:
			logger.Info("BLVTCandlestick Response: %v", response.(futuresusdt.BLVTCandlestickResponse))
		case futuresusdt.BLVTCandlestickCombinedResponse:
			logger.Info("BLVTCandlestickCombinedResponse: %v", response.(futuresusdt.BLVTCandlestickCombinedResponse))
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
