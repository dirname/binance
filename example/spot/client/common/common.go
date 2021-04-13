package main

import (
	"github.com/dirname/Binance/config"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"github.com/dirname/Binance/spot/client"
)

var commonClient *spotclient.CommonClient

func init() {
	logger.Enable(false)
	commonClient = spotclient.NewCommonClient(config.SpotRestHost)
}

func main() {
	getSAPISystemStatus()
	ping()
	getServerTime()
}

func getSAPISystemStatus() {
	response, err := commonClient.GetSAPISystemStatus()
	if err != nil {
		logger.Error("getSAPISystemStatus err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getSAPISystemStatus SAPI error: %v", response.(model.APIErrorResponse))
	case spotclient.SystemStatusResponse:
		logger.Info("getSAPISystemStatus: %v", response.(spotclient.SystemStatusResponse))
	default:
		logger.Info("getSAPISystemStatus Unknown response: %v", response)
	}
}

func ping() {
	response, err := commonClient.Ping()
	if err != nil {
		logger.Error("ping response err: %s", err.Error())
	}
	if response != nil {
		logger.Info("ping server successfully")
	}
}

func getServerTime() {
	response, err := commonClient.GetServerTime()
	if err != nil {
		logger.Error("getServerTime error: %s", err.Error())
	}
	logger.Info("serverTime: %d", response)
}
