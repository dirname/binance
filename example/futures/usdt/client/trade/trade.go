package main

import (
	"github.com/dirname/binance/config"
	"github.com/dirname/binance/futures/usd/client"
	"github.com/dirname/binance/futures/usd/client/define/autoCloseType"
	"github.com/dirname/binance/futures/usd/client/define/marginType"
	"github.com/dirname/binance/futures/usd/client/define/orderBool"
	"github.com/dirname/binance/futures/usd/client/define/orderRespType"
	"github.com/dirname/binance/futures/usd/client/define/orderSide"
	"github.com/dirname/binance/futures/usd/client/define/orderType"
	"github.com/dirname/binance/futures/usd/client/define/workingType"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"github.com/shopspring/decimal"
)

var tradeClient *futuresclient.TradeClient

func init() {
	logger.Enable(false)
	tradeClient = futuresclient.NewTradeClient(config.USDFuturesRestHost, config.AppKey, config.AppSecret)
}

func main() {
	//positionSideDual()
	//getPositionSideDual()
	//changeMultiAssetsMode()
	//getMultiAssetsMargin()
	//testNewOrder()
	//newOrder()
	//getOrder()
	//deleteOrder()
	//deleteAllOrder()
	//countDownOrder()
	//getOpenOrder()
	//getAllOpenOrder()
	//getAllOrder()
	//getBalance()
	//getUserInfo()
	//leverage()
	getPositionMargin()
	getPositionRisk()
	getTradeList()
	getIncome()
	getLeverageBracket()
	getADLQuantile()
	getForceOrders()
	getQuantitativeRules()
	getCommissionRate()
}

func getPositionMargin() {
	response, err := tradeClient.IsolatedPositionHistory("BTCUSDT", marginType.AddIsolatedPositionMargin, 0, 0, 0, 0)
	if err != nil {
		logger.Error("getPositionMargin err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getPositionMargin API error: %v", response.(model.APIErrorResponse))
	case futuresclient.IsolatedPositionHistoryResponse:
		logger.Info("getPositionMargin: %v", response.(futuresclient.IsolatedPositionHistoryResponse))
	default:
		logger.Info("getPositionMargin Unknown response: %v", response)
	}
}

func getPositionRisk() {
	response, err := tradeClient.PositionRisk("BTCUSDT", 0)
	if err != nil {
		logger.Error("getPositionRisk err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getPositionRisk API error: %v", response.(model.APIErrorResponse))
	case futuresclient.PositionRiskResponse:
		logger.Info("getPositionRisk: %v", response.(futuresclient.PositionRiskResponse))
	default:
		logger.Info("getPositionRisk Unknown response: %v", response)
	}
}

func getTradeList() {
	response, err := tradeClient.TradeLists("BTCUSDT", 0, 0, 0, 0, 0)
	if err != nil {
		logger.Error("getTradeList err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getTradeList API error: %v", response.(model.APIErrorResponse))
	case futuresclient.TradeListResponse:
		logger.Info("getTradeList: %v", response.(futuresclient.TradeListResponse))
	default:
		logger.Info("getTradeList Unknown response: %v", response)
	}
}

func getIncome() {
	response, err := tradeClient.Incomes("BTCUSDT", "", 0, 0, 0, 0)
	if err != nil {
		logger.Error("getIncome err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getIncome API error: %v", response.(model.APIErrorResponse))
	case futuresclient.IncomeResponse:
		logger.Info("getIncome: %v", response.(futuresclient.IncomeResponse))
	default:
		logger.Info("getIncome Unknown response: %v", response)
	}
}

func getLeverageBracket() {
	response, err := tradeClient.LeverageBracket("BTCUSDT", 0)
	if err != nil {
		logger.Error("getLeverageBracket err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getLeverageBracket API error: %v", response.(model.APIErrorResponse))
	case futuresclient.LeverageBracketResponse:
		logger.Info("getLeverageBracket: %v", response.(futuresclient.LeverageBracketResponse))
	default:
		logger.Info("getLeverageBracket Unknown response: %v", response)
	}
}

func getADLQuantile() {
	response, err := tradeClient.ADLQuantile("BTCUSDT", 0)
	if err != nil {
		logger.Error("getADLQuantile err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getADLQuantile API error: %v", response.(model.APIErrorResponse))
	case futuresclient.ADLQuantileResponse:
		logger.Info("getADLQuantile: %v", response.(futuresclient.ADLQuantileResponse))
	default:
		logger.Info("getADLQuantile Unknown response: %v", response)
	}
}

func getForceOrders() {
	response, err := tradeClient.ForceOrders("BTCUSDT", autoCloseType.ADL, 0, 0, 0, 0)
	if err != nil {
		logger.Error("getADLQuantile err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getForceOrders API error: %v", response.(model.APIErrorResponse))
	case futuresclient.ForceOrdersResponse:
		logger.Info("getForceOrders: %v", response.(futuresclient.ForceOrdersResponse))
	default:
		logger.Info("getForceOrders Unknown response: %v", response)
	}
}

func getQuantitativeRules() {
	response, err := tradeClient.QuantitativeTradingRules("BTCUSDT", 0)
	if err != nil {
		logger.Error("getQuantitativeRules err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getQuantitativeRules API error: %v", response.(model.APIErrorResponse))
	case futuresclient.QuantitativeTradingRulesResponse:
		logger.Info("getQuantitativeRules: %v", response.(futuresclient.QuantitativeTradingRulesResponse))
	default:
		logger.Info("getQuantitativeRules Unknown response: %v", response)
	}
}

func getCommissionRate() {
	response, err := tradeClient.CommissionRate("BTCUSDT", 0)
	if err != nil {
		logger.Error("getCommissionRate err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getCommissionRate API error: %v", response.(model.APIErrorResponse))
	case futuresclient.CommissionRateResponse:
		logger.Info("getCommissionRate: %v", response.(futuresclient.CommissionRateResponse))
	default:
		logger.Info("getCommissionRate Unknown response: %v", response)
	}
}

func positionSideDual() {
	response, err := tradeClient.PositionSideDual(true, 0)
	if err != nil {
		logger.Error("positionSideDual err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("positionSideDual API error: %v", response.(model.APIErrorResponse))
	case futuresclient.DualSidePositionResponse:
		logger.Info("positionSideDual: %v", response.(futuresclient.DualSidePositionResponse))
	default:
		logger.Info("positionSideDual Unknown response: %v", response)
	}
}

func getPositionSideDual() {
	response, err := tradeClient.GetPositionSideDual(0)
	if err != nil {
		logger.Error("getPositionSideDual err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getPositionSideDual API error: %v", response.(model.APIErrorResponse))
	case futuresclient.DualSidePositionResponse:
		logger.Info("getPositionSideDual: %v", response.(futuresclient.DualSidePositionResponse))
	default:
		logger.Info("getPositionSideDual Unknown response: %v", response)
	}
}

func changeMultiAssetsMode() {
	response, err := tradeClient.ChangeMultiAssetsMode(false, 0)
	if err != nil {
		logger.Error("changeMultiAssetsMode err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("changeMultiAssetsMode API error: %v", response.(model.APIErrorResponse))
	case futuresclient.DualSidePositionResponse:
		logger.Info("changeMultiAssetsMode: %v", response.(futuresclient.DualSidePositionResponse))
	default:
		logger.Info("changeMultiAssetsMode Unknown response: %v", response)
	}
}

func getMultiAssetsMargin() {
	response, err := tradeClient.GetMultiAssetsMargin(0)
	if err != nil {
		logger.Error("getMultiAssetsMargin err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getMultiAssetsMargin API error: %v", response.(model.APIErrorResponse))
	case futuresclient.MultiAssetsMarginResponse:
		logger.Info("getMultiAssetsMargin: %v", response.(futuresclient.MultiAssetsMarginResponse))
	default:
		logger.Info("getMultiAssetsMargin Unknown response: %v", response)
	}
}

func newOrder() {
	response, err := tradeClient.NewOrder("BTCUSDT", orderSide.Buy, orderSide.Long, orderType.Market, "", "", "", "", workingType.MarkPrice, orderBool.FALSE, orderRespType.ACK, decimal.NewFromInt(1), decimal.Decimal{}, decimal.Decimal{}, decimal.Decimal{}, decimal.Decimal{}, 0)
	if err != nil {
		logger.Error("newOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("newOrder API error: %v", response.(model.APIErrorResponse))
	default:
		logger.Info("newOrder response: %v", response)
	}
}

func testNewOrder() {
	response, err := tradeClient.TestNewOrder("BTCUSDT", orderSide.Buy, orderSide.Both, orderType.Market, orderBool.False, "", "", "", workingType.MarkPrice, orderBool.FALSE, orderRespType.ACK, decimal.NewFromInt(1), decimal.Decimal{}, decimal.Decimal{}, decimal.Decimal{}, decimal.Decimal{}, 0)
	if err != nil {
		logger.Error("testNewOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("testNewOrder API error: %v", response.(model.APIErrorResponse))
	default:
		logger.Info("testNewOrder response: %v", response)
	}
}

func getOrder() {
	response, err := tradeClient.GetOrder("BTCUSDT", "", 0, 0)
	if err != nil {
		logger.Error("getOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getOrder API error: %v", response.(model.APIErrorResponse))
	case futuresclient.OrderResponse:
		logger.Info("getOrder: %v", response.(futuresclient.OrderResponse))
	default:
		logger.Info("getOrder Unknown response: %v", response)
	}
}

func deleteOrder() {
	response, err := tradeClient.DeleteOrder("BTCUSDT", "", 0, 0)
	if err != nil {
		logger.Error("deleteOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("deleteOrder API error: %v", response.(model.APIErrorResponse))
	case futuresclient.DeleteOrderResponse:
		logger.Info("deleteOrder: %v", response.(futuresclient.DeleteOrderResponse))
	default:
		logger.Info("deleteOrder Unknown response: %v", response)
	}
}

func deleteAllOrder() {
	response, err := tradeClient.DeleteAllOrder("BTCUSDT", 0)
	if err != nil {
		logger.Error("deleteAllOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("deleteAllOrder API: %v", response.(model.APIErrorResponse))
	default:
		logger.Info("deleteAllOrder Unknown response: %v", response)
	}
}

func countDownOrder() {
	response, err := tradeClient.CountDownDeleteAllOrder("BTCUSDT", 1000, 0)
	if err != nil {
		logger.Error("countDownOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("countDownOrder API error: %v", response.(model.APIErrorResponse))
	case futuresclient.CountDownDeleteOrderResponse:
		logger.Info("countDownOrder: %v", response.(futuresclient.CountDownDeleteOrderResponse))
	default:
		logger.Info("countDownOrder Unknown response: %v", response)
	}
}

func getOpenOrder() {
	response, err := tradeClient.GetOpenOrder("BTCUSDT", "", 0, 0)
	if err != nil {
		logger.Error("getOpenOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getOpenOrder API error: %v", response.(model.APIErrorResponse))
	case futuresclient.OpenOrderResponse:
		logger.Info("getOpenOrder: %v", response.(futuresclient.OpenOrderResponse))
	default:
		logger.Info("getOpenOrder Unknown response: %v", response)
	}
}

func getAllOpenOrder() {
	response, err := tradeClient.GetAllOpenOrder("BTCUSDT", 0)
	if err != nil {
		logger.Error("getAllOpenOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getAllOpenOrder API error: %v", response.(model.APIErrorResponse))
	case futuresclient.AllOpenOrderResponse:
		logger.Info("getAllOpenOrder: %v", response.(futuresclient.AllOpenOrderResponse))
	default:
		logger.Info("getAllOpenOrder Unknown response: %v", response)
	}
}

func getAllOrder() {
	response, err := tradeClient.GetAllOrder("BTCUSDT", 0, 0, 0, 0, 0)
	if err != nil {
		logger.Error("getAllOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getAllOrder API error: %v", response.(model.APIErrorResponse))
	case futuresclient.AllOrderResponse:
		logger.Info("getAllOrder: %v", response.(futuresclient.AllOrderResponse))
	default:
		logger.Info("getAllOrder Unknown response: %v", response)
	}
}

func getBalance() {
	response, err := tradeClient.GetBalance(0)
	if err != nil {
		logger.Error("getBalance err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getBalance API error: %v", response.(model.APIErrorResponse))
	case futuresclient.UserBalanceResponse:
		logger.Info("getBalance: %v", response.(futuresclient.UserBalanceResponse))
	default:
		logger.Info("getBalance Unknown response: %v", response)
	}
}

func getUserInfo() {
	response, err := tradeClient.GetUserInfo(0)
	if err != nil {
		logger.Error("getUserInfo err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getUserInfo API error: %v", response.(model.APIErrorResponse))
	case futuresclient.UserInfoResponse:
		logger.Info("getUserInfo: %v", response.(futuresclient.UserInfoResponse))
	default:
		logger.Info("getUserInfo Unknown response: %v", response)
	}
}

func leverage() {
	response, err := tradeClient.Leverage("BTCUSDT", 125, 0)
	if err != nil {
		logger.Error("leverage err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("leverage API error: %v", response.(model.APIErrorResponse))
	case futuresclient.LeverageResponse:
		logger.Info("leverage: %v", response.(futuresclient.LeverageResponse))
	default:
		logger.Info("leverage Unknown response: %v", response)
	}
}
