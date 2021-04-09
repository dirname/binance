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
	//disableFastWithdraw()
	//enableFastWithdraw()
	getDepositHistoryNetwork()
	getDepositHistory()
	getWithdrawHistoryNetwork()
	getWithdraw()
	getDepositAddressNetwork()
	getDepositAddress()
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

func getDepositHistoryNetwork() {
	response, err := walletClient.DepositHistoryNetwork("", -1, 0, 0, 0, 0, 0)
	if err != nil {
		logger.Error("getDepositHistoryNetwork err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getDepositHistoryNetwork API error: %v", response.(model.APIErrorResponse))
	case spotclient.DepositHistoryNetworkResponse:
		logger.Info("getDepositHistoryNetwork: %v", response.(spotclient.DepositHistoryNetworkResponse))
	default:
		logger.Info("getDepositHistoryNetwork Unknown response: %v", response)
	}
}

func getDepositHistory() {
	response, err := walletClient.DepositHistory("", -1, 0, 0, 0)
	if err != nil {
		logger.Error("getDepositHistory err: %s", err.Error())
	}
	switch response.(type) {
	case model.WAPIErrorResponse:
		logger.Info("getDepositHistory WAPI error: %v", response.(model.APIErrorResponse))
	case spotclient.DepositHistoryResponse:
		logger.Info("getDepositHistory: %v", response.(spotclient.DepositHistoryResponse))
	default:
		logger.Info("getDepositHistory Unknown response: %v", response)
	}
}

func getWithdrawHistoryNetwork() {
	response, err := walletClient.WithdrawHistoryNetwork("", -1, 0, 0, 0, 0, 0)
	if err != nil {
		logger.Error("getWithdrawHistoryNetwork err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getWithdrawHistoryNetwork API error: %v", response.(model.APIErrorResponse))
	case spotclient.WithdrawHistoryNetworkResponse:
		logger.Info("getWithdrawHistoryNetwork: %v", response.(spotclient.WithdrawHistoryNetworkResponse))
	default:
		logger.Info("getWithdrawHistoryNetwork Unknown response: %v", response)
	}
}

func getWithdraw() {
	response, err := walletClient.WithdrawHistory("", -1, 0, 0, 0)
	if err != nil {
		logger.Error("getWithdraw err: %s", err.Error())
	}
	switch response.(type) {
	case model.WAPIErrorResponse:
		logger.Info("getWithdraw WAPI error: %v", response.(model.APIErrorResponse))
	case spotclient.WithdrawHistoryResponse:
		logger.Info("getWithdraw: %v", response.(spotclient.WithdrawHistoryResponse))
	default:
		logger.Info("getWithdraw Unknown response: %v", response)
	}
}

func getDepositAddressNetwork() {
	response, err := walletClient.DepositAddressNetwork("BTC", "", 0)
	if err != nil {
		logger.Error("getDepositAddressNetwork err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getDepositAddressNetwork API error: %v", response.(model.APIErrorResponse))
	case spotclient.DepositAddressNetworkResponse:
		logger.Info("getDepositAddressNetwork: %v", response.(spotclient.DepositAddressNetworkResponse))
	default:
		logger.Info("getDepositAddressNetwork Unknown response: %v", response)
	}
}

func getDepositAddress() {
	response, err := walletClient.DepositAddress("BNB", false, 0)
	if err != nil {
		logger.Error("getDepositAddress err: %s", err.Error())
	}
	logger.Info("getDepositAddress response: %v", response)
}
