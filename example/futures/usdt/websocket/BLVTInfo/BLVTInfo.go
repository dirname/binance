package main

import (
	"fmt"
	"github.com/dirname/binance/futures/usd/websocket/market"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"time"
)

func main() {
	client := futuresclient.NewUSDFuturesBLVTInfoClient("TRXDOWN@tokenNav")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "TRXDOWN@tokenNav", "TRXDOWN@tokenNav")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresclient.BLVTInfoResponse:
			logger.Info("BLVTInfo Response: %v", response.(futuresclient.BLVTInfoResponse))
		case futuresclient.BLVTInfoCombinedResponse:
			logger.Info("BLVTInfoCombinedResponse: %v", response.(futuresclient.BLVTInfoCombinedResponse))
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
