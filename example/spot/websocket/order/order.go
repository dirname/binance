package main

import (
	"fmt"
	"github.com/dirname/binance/config"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"github.com/dirname/binance/spot/websocket/account"
	"time"
)

func main() {
	listenKeyBuilder := spotclient.NewListenKeyBuilder(config.SpotRestHost, config.AppKey, config.AppSecret)
	key, err := listenKeyBuilder.CreateSpotListenKey()
	if err != nil {
		logger.Error("Failed to create spot listen key: %s", err.Error())
		return
	}
	client := spotclient.NewOrderWebsocketClient(key)
	client.SetReadTimerInterval(30 * time.Minute)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetKeepAliveInterval(1 * time.Second)
	client.SetHandler(func() {
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case spotclient.ListStatus:
			logger.Info("ListUpdate Response: %v", response.(spotclient.ListStatus))
		case spotclient.ListCombinedStatus:
			logger.Info("ListCombinedUpdate Response: %v", response.(spotclient.ListCombinedStatus))
		case spotclient.ExecutionReport:
			logger.Info("ExecutionUpdate Response: %v", response.(spotclient.ExecutionReport))
		case spotclient.ExecutionCombinedReport:
			logger.Info("ExecutionUpdateCombinedResponse: %v", response.(spotclient.ExecutionCombinedReport))
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
