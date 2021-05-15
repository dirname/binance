package main

import (
	"github.com/dirname/binance/config"
	"github.com/dirname/binance/futures/usd/client"
	logger "github.com/dirname/binance/logging"
)

var commonClient *futuresclient.CommonClient

func init() {
	logger.Enable(false)
	commonClient = futuresclient.NewCommonClient(config.USDFuturesRestHost)
}

func main() {
	ping()
	getServerTime()
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
