package main

import (
	"github.com/dirname/Binance/config"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
	"github.com/dirname/Binance/spot/client"
)

var walletClient *spotclient.WalletClient

func init() {
	logger.Enable(false)
	walletClient = spotclient.NewWalletClient(config.SpotRestHost, config.AppKey, config.AppSecret)
}

func main() {
	//getAllCoinsInfo()
	//getSpotSnapshot()
	//getMarginSnapshot()
	//getFuturesSnapshot()
	disableFastWithdraw()
	enableFastWithdraw()
}

func getAllCoinsInfo() {
	response, err := walletClient.GetAllCoinsInfo(0)
	if err != nil {
		logger.Error("getWAPISystemStatus err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getAllCoinsInfo API error: %v", response.(model.APIErrorResponse))
	case spotclient.AllCoinInfoResponse:
		logger.Info("getAllCoinsInfo: %v", response.(spotclient.AllCoinInfoResponse))
	default:
		logger.Info("getAllCoinsInfo Unknown response: %v", response)
	}
}

func getSpotSnapshot() {
	response, err := walletClient.GetSpotSnapshot(0, 0, 30, 0)
	if err != nil {
		logger.Error("getSpotSnapshot err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getSpotSnapshot API error: %v", response.(model.APIErrorResponse))
	case spotclient.SpotSnapshotResponse:
		logger.Info("getSpotSnapshot: %v", response.(spotclient.SpotSnapshotResponse))
	default:
		logger.Info("getSpotSnapshot Unknown response: %v", response)
	}
}

func getMarginSnapshot() {
	response, err := walletClient.GetMarginSnapshot(0, 0, 30, 0)
	if err != nil {
		logger.Error("getMarginSnapshot err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getMarginSnapshot API error: %v", response.(model.APIErrorResponse))
	case spotclient.MarginSnapshotResponse:
		logger.Info("getMarginSnapshot: %v", response.(spotclient.MarginSnapshotResponse))
	default:
		logger.Info("getMarginSnapshot Unknown response: %v", response)
	}
}

func getFuturesSnapshot() {
	response, err := walletClient.GetFuturesSnapshot(0, 0, 30, 0)
	if err != nil {
		logger.Error("getFuturesSnapshot err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getFuturesSnapshot API error: %v", response.(model.APIErrorResponse))
	case spotclient.FuturesSnapshotResponse:
		logger.Info("getFuturesSnapshot: %v", response.(spotclient.FuturesSnapshotResponse))
	default:
		logger.Info("getFuturesSnapshot Unknown response: %v", response)
	}
}

func disableFastWithdraw() {
	response, err := walletClient.FastWithdrawSwitch(false, 0)
	if err != nil {
		logger.Error("disableFastWithdrawSwitch err: %s", err.Error())
	}
	logger.Info("%v", response)
}

func enableFastWithdraw() {
	response, err := walletClient.FastWithdrawSwitch(true, 0)
	if err != nil {
		logger.Error("enableFastWithdraw err: %s", err.Error())
	}
	logger.Info("%v", response)
}
