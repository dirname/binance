package main

import (
	"fmt"
	"github.com/dirname/Binance/futures/usdt/websocket/market"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"time"
)

func main() {
	client := futuresusdt.NewUSDTFuturesBLVTInfoClient("TRXDOWN@tokenNav")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "TRXDOWN@tokenNav", "TRXDOWN@tokenNav")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresusdt.BLVTInfoResponse:
			logger.Info("BLVTInfo Response: %v", response.(futuresusdt.BLVTInfoResponse))
		case futuresusdt.BLVTInfoCombinedResponse:
			logger.Info("BLVTInfoCombinedResponse: %v", response.(futuresusdt.BLVTInfoCombinedResponse))
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
	client.Unsubscribe(123, "TRXDOWN@tokenNav", "TRXDOWN@tokenNav")
	client.Close()
	logger.Info("Client closed")
}
