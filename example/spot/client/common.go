package main

import (
	"github.com/dirname/Binance/config"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"github.com/dirname/Binance/spot/client"
)

var client *spotclient.CommonClient

func init() {
	client = spotclient.NewCommonClient(config.SpotRestHost)
}

func main() {
	getWAPISystemStatus()
	getSAPISystemStatus()
	ping()
	getServerTime()
	getExchangeInfo()
}

func getWAPISystemStatus() {
	response, err := client.GetWAPISystemStatus()
	if err != nil {
		logger.Error("getWAPISystemStatus err: %s", err.Error())
	}
	switch response.(type) {
	case model.WAPIErrorResponse:
		logger.Info("getWAPISystemStatus WAPI error: %v", response.(model.WAPIErrorResponse))
	case spotclient.SystemStatusResponse:
		logger.Info("getWAPISystemStatus: %v", response.(spotclient.SystemStatusResponse))
	default:
		logger.Info("getWAPISystemStatus Unknown response: %v", response)
	}
}

func getSAPISystemStatus() {
	response, err := client.GetSAPISystemStatus()
	if err != nil {
		logger.Error("getSAPISystemStatus err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getSAPISystemStatus WAPI error: %v", response.(model.APIErrorResponse))
	case spotclient.SystemStatusResponse:
		logger.Info("getSAPISystemStatus: %v", response.(spotclient.SystemStatusResponse))
	default:
		logger.Info("getSAPISystemStatus Unknown response: %v", response)
	}
}

func ping() {
	response, err := client.Ping()
	if err != nil {
		logger.Error("ping response err: %s", err.Error())
	}
	if response != nil {
		logger.Info("ping server successfully")
	}
}

func getServerTime() {
	response, err := client.GetServerTime()
	if err != nil {
		logger.Error("getServerTime error: %s", err.Error())
	}
	logger.Info("serverTime: %d", response)
}

func getExchangeInfo() {
	response, err := client.GetExchangeInfo()
	if err != nil {
		logger.Error("getExchangeInfo err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getExchangeInfo API error: %v", response.(model.APIErrorResponse))
	case spotclient.ExchangeInfoResponse:
		logger.Info("getExchangeInfo: %v", response.(spotclient.ExchangeInfoResponse))
	default:
		logger.Info("getExchangeInfo Unknown response: %v", response)
	}
}
