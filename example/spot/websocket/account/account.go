package main

import (
	"fmt"
	"github.com/dirname/Binance/config"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"github.com/dirname/Binance/spot/websocket/account"
	"time"
)

func main() {
	listenKeyBuilder := spotclient.NewListenKeyBuilder(config.SpotRestHost, config.AppKey, config.AppSecret)
	key, err := listenKeyBuilder.CreateSpotListenKey()
	if err != nil {
		logger.Error("Failed to create spot listen key: %s", err.Error())
		return
	}
	client := spotclient.NewAccountWebsocketClient(key)
	client.SetReadTimerInterval(30 * time.Minute)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetKeepAliveInterval(1 * time.Second)
	client.SetHandler(func() {
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case spotclient.AccountPosition:
			logger.Info("AccountUpdate Response: %v", response.(spotclient.AccountPosition))
		case spotclient.AccountCombinedPosition:
			logger.Info("AccountUpdateCombinedResponse: %v", response.(spotclient.AccountCombinedPosition))
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