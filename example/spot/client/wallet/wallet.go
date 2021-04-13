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
	//getDepositHistoryNetwork()
	//getDepositHistory()
	//getWithdrawHistoryNetwork()
	//getWithdraw()
	//getDepositAddressNetwork()
	//getDepositAddress()
	//getSAPIAccountStatus()
	//getSAPIAccountAPIStatus()
	//getSAPIDustLog()
	//dustTransfer()
	//getDividendRecord()
	//getSAPIAssetDetail()
	//getSAPITradeFee()
	//universalTransfer()
	//getUniversalTransferRecord()
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
	logger.Info("getDepositHistory response: %v", response)
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
	logger.Info("getWithdraw response: %v", response)
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

func getSAPIAccountStatus() {
	response, err := walletClient.SAPIAccountStatus(0)
	if err != nil {
		logger.Error("getSAPIAccountStatus err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getSAPIAccountStatus API error: %v", response.(model.APIErrorResponse))
	case spotclient.SAPIAccountResponse:
		logger.Info("getSAPIAccountStatus: %v", response.(spotclient.SAPIAccountResponse))
	default:
		logger.Info("getSAPIAccountStatus Unknown response: %v", response)
	}
}

func getSAPIAccountAPIStatus() {
	response, err := walletClient.SAPIAccountAPIStatus(0)
	if err != nil {
		logger.Error("getSAPIAccountAPIStatus err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getSAPIAccountAPIStatus API error: %v", response.(model.APIErrorResponse))
	case spotclient.AccountSAPIStatusResponse:
		logger.Info("getSAPIAccountAPIStatus: %v", response.(spotclient.AccountSAPIStatusResponse))
	default:
		logger.Info("getSAPIAccountStatus Unknown response: %v", response)
	}
}

func getSAPIDustLog() {
	response, err := walletClient.SAPIDustLog(0)
	if err != nil {
		logger.Error("getSAPIDustLog err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getSAPIDustLog API error: %v", response.(model.APIErrorResponse))
	case spotclient.DustLogSAPIResponse:
		logger.Info("getSAPIDustLog: %v", response.(spotclient.DustLogSAPIResponse))
	default:
		logger.Info("getSAPIAccountStatus Unknown response: %v", response)
	}
}

func dustTransfer() {
	response, err := walletClient.DustTransfer([]string{"BTC"}, 0)
	if err != nil {
		logger.Error("dustTransfer err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("dustTransfer API error: %v", response.(model.APIErrorResponse))
	case spotclient.DustTransferResponse:
		logger.Info("dustTransfer: %v", response.(spotclient.DustTransferResponse))
	default:
		logger.Info("getSAPIAccountStatus Unknown response: %v", response)
	}
}

func getDividendRecord() {
	response, err := walletClient.AccountDividendRecord("", 0, 0, 30, 0)
	if err != nil {
		logger.Error("getDividendRecord err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getDividendRecord API error: %v", response.(model.APIErrorResponse))
	case spotclient.AccountDividendRecordResponse:
		logger.Info("getDividendRecord: %v", response.(spotclient.AccountDividendRecordResponse))
	default:
		logger.Info("getSAPIAccountStatus Unknown response: %v", response)
	}
}

func getSAPIAssetDetail() {
	response, err := walletClient.SAPIAssetDetail(0)
	if err != nil {
		logger.Error("getSAPIAssetDetail err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getSAPIAssetDetail API error: %v", response.(model.APIErrorResponse))
	case spotclient.SAPIAssetDetailResponse:
		logger.Info("getSAPIAssetDetail: %v", response.(spotclient.SAPIAssetDetailResponse))
	default:
		logger.Info("getSAPIAccountStatus Unknown response: %v", response)
	}
}

func getSAPITradeFee() {
	response, err := walletClient.SAPITradeFee("", 0)
	if err != nil {
		logger.Error("getSAPITradeFee err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getSAPITradeFee API error: %v", response.(model.APIErrorResponse))
	case spotclient.SAPITradeFeeResponse:
		logger.Info("getSAPITradeFee: %v", response.(spotclient.SAPITradeFeeResponse))
	default:
		logger.Info("getSAPITradeFee Unknown response: %v", response)
	}
}

func universalTransfer() {
	response, err := walletClient.UniversalTransfer("UMFUTURE_MAIN", "USDT", "10", 0)
	if err != nil {
		logger.Error("universalTransfer err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("universalTransfer API error: %v", response.(model.APIErrorResponse))
	case spotclient.UniversalTransferResponse:
		logger.Info("universalTransfer: %v", response.(spotclient.UniversalTransferResponse))
	default:
		logger.Info("universalTransfer Unknown response: %v", response)
	}
}

func getUniversalTransferRecord() {
	response, err := walletClient.UniversalTransferRecord("UMFUTURE_MAIN", 0, 0, 0, 0, 0)
	if err != nil {
		logger.Error("getUniversalTransferRecord err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getUniversalTransferRecord API error: %v", response.(model.APIErrorResponse))
	case spotclient.UniversalTransferRecordResponse:
		logger.Info("getUniversalTransferRecord: %v", response.(spotclient.UniversalTransferRecordResponse))
	default:
		logger.Info("getUniversalTransferRecord Unknown response: %v", response)
	}
}
