package main

import (
	"github.com/dirname/binance/config"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"github.com/dirname/binance/spot/client"
	"github.com/dirname/binance/spot/client/define/orderSide"
	"github.com/dirname/binance/spot/client/define/orderType"
	"github.com/dirname/binance/spot/client/define/timeInForce"
	"github.com/shopspring/decimal"
)

var tradeClient *spotclient.TradeClient

func init() {
	logger.Enable(false)
	tradeClient = spotclient.NewTradeClient(config.SpotRestHost, config.AppKey, config.AppSecret)
}

func main() {
	testNewOrder()
	newOrder()
	deleteOrder()
	deleteOpenOrder()
	getOrder()
	getOpenOrder()
	getAllOrder()
	newOCOOrder()
	deleteOCOOrder()
	getOCOOrder()
	getAllOCOOrder()
	getOpenOCOOrder()
	getAccountInfo()
	getAccountTradesList()
}

func testNewOrder() {
	response, err := tradeClient.TestNewOrder("BTCUSDT", orderSide.Buy, orderType.Limit, timeInForce.FOK, "", "", decimal.NewFromFloat(0.001), decimal.Decimal{}, decimal.NewFromFloat(59700.12), decimal.Decimal{}, decimal.Decimal{}, 0)
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

func newOrder() {
	response, err := tradeClient.NewOrder("CFXUSDT", orderSide.Sell, orderType.Limit, timeInForce.GTC, "", "", decimal.NewFromFloat(10.0), decimal.Decimal{}, decimal.NewFromFloat(1.21), decimal.Decimal{}, decimal.Decimal{}, 0)
	if err != nil {
		logger.Error("testNewOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("newOrder API error: %v", response.(model.APIErrorResponse))
	case spotclient.NewOrderResponseACK:
		logger.Info("newOrder ACK: %v", response.(spotclient.NewOrderResponseACK))
	case spotclient.NewOrderResponseFull:
		logger.Info("newOrder Full: %v", response.(spotclient.NewOrderResponseFull))
	case spotclient.NewOrderResponseResult:
		logger.Info("newOrder RESULT: %v", response.(spotclient.NewOrderResponseResult))
	default:
		logger.Info("newOrder response: %v", response)
	}
}

func deleteOrder() {
	response, err := tradeClient.DeleteOrder("CFXUSDT", "Tc5GhznrCVhmIfRticUpQ3", "", 0, 0)
	if err != nil {
		logger.Error("deleteOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("deleteOrder API error: %v", response.(model.APIErrorResponse))
	case spotclient.DeleteOrderResponse:
		logger.Info("deleteOrder: %v", response.(spotclient.DeleteOrderResponse))
	default:
		logger.Info("deleteOrder response: %v", response)
	}
}

func deleteOpenOrder() {
	response, err := tradeClient.DeleteOpenOrders("CFXUSDT", 0)
	if err != nil {
		logger.Error("deleteOpenOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("deleteOpenOrder API error: %v", response.(model.APIErrorResponse))
	case spotclient.DeleteOpenOrderResponse:
		logger.Info("deleteOpenOrder: %v", response.(spotclient.DeleteOpenOrderResponse))
	default:
		logger.Info("deleteOpenOrder response: %v", response)
	}
}

func getOrder() {
	response, err := tradeClient.GetOrder("CFXUSDT", "IfaBDncaEOmRhILXdzSPc7", 0, 0)
	if err != nil {
		logger.Error("getOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getOrder API error: %v", response.(model.APIErrorResponse))
	case spotclient.GetOrderResponse:
		logger.Info("getOrder: %v", response.(spotclient.GetOrderResponse))
	default:
		logger.Info("getOrder response: %v", response)
	}
}

func getOpenOrder() {
	response, err := tradeClient.GetOpenOrder("CFXUSDT", 0)
	if err != nil {
		logger.Error("getOpenOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getOpenOrder API error: %v", response.(model.APIErrorResponse))
	case spotclient.GetOpenOrdersResponse:
		logger.Info("getOpenOrder: %v", response.(spotclient.GetOpenOrdersResponse))
	default:
		logger.Info("getOpenOrder response: %v", response)
	}
}

func getAllOrder() {
	response, err := tradeClient.GetAllOrder("CFXUSDT", 0, 0, 0, 0, 0)
	if err != nil {
		logger.Error("getAllOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getAllOrder API error: %v", response.(model.APIErrorResponse))
	case spotclient.GetAllOrdersResponse:
		logger.Info("getAllOrder: %v", response.(spotclient.GetAllOrdersResponse))
	default:
		logger.Info("getAllOrder response: %v", response)
	}
}

func newOCOOrder() {
	response, err := tradeClient.NewOCO("CFXUSDT", "", orderSide.Buy, "", "", timeInForce.GTC, "", decimal.NewFromFloat(10), decimal.NewFromFloat(1.18), decimal.Decimal{}, decimal.NewFromFloat(1.24), decimal.NewFromFloat(1.23), decimal.Decimal{}, 0)
	if err != nil {
		logger.Error("newOCOOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("newOCOOrder API error: %v", response.(model.APIErrorResponse))
	case spotclient.NewOCOOrderResponse:
		logger.Info("newOCOOrder: %v", response.(spotclient.NewOCOOrderResponse))
	default:
		logger.Info("newOCOOrder response: %v", response)
	}
}

func deleteOCOOrder() {
	response, err := tradeClient.DeleteOCO("CFXUSDT", "urCFHREAQZY3JcK2MyNlAc", "", 0, 0)
	if err != nil {
		logger.Error("deleteOCOOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("deleteOCOOrder API error: %v", response.(model.APIErrorResponse))
	case spotclient.DeleteOCOResponse:
		logger.Info("deleteOCOOrder: %v", response.(spotclient.DeleteOCOResponse))
	default:
		logger.Info("deleteOCOOrder response: %v", response)
	}
}

func getOCOOrder() {
	response, err := tradeClient.GetOCOOrder("RE0K6NUJcmBtIueyGkJATj", 0, 0)
	if err != nil {
		logger.Error("getOCOOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getOCOOrder API error: %v", response.(model.APIErrorResponse))
	case spotclient.GetOCOResponse:
		logger.Info("getOCOOrder: %v", response.(spotclient.GetOCOResponse))
	default:
		logger.Info("getOCOOrder response: %v", response)
	}
}

func getAllOCOOrder() {
	response, err := tradeClient.GetAllOCOOrder(0, 0, 0, 0, 0)
	if err != nil {
		logger.Error("getAllOCOOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getAllOCOOrder API error: %v", response.(model.APIErrorResponse))
	case spotclient.GetAllOCOResponse:
		logger.Info("getAllOCOOrder: %v", response.(spotclient.GetAllOCOResponse))
	default:
		logger.Info("getAllOCOOrder response: %v", response)
	}
}

func getOpenOCOOrder() {
	response, err := tradeClient.GetOpenOCOOrder(0)
	if err != nil {
		logger.Error("getOpenOCOOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getOpenOCOOrder API error: %v", response.(model.APIErrorResponse))
	case spotclient.GetOpenOCOResponse:
		logger.Info("getOpenOCOOrder: %v", response.(spotclient.GetOpenOCOResponse))
	default:
		logger.Info("getOpenOCOOrder response: %v", response)
	}
}

func getAccountInfo() {
	response, err := tradeClient.GetAccountInfo(0)
	if err != nil {
		logger.Error("getAccountInfo err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getAccountInfo API error: %v", response.(model.APIErrorResponse))
	case spotclient.AccountInfoResponse:
		logger.Info("getAccountInfo: %v", response.(spotclient.AccountInfoResponse))
	default:
		logger.Info("getAccountInfo response: %v", response)
	}
}

func getAccountTradesList() {
	response, err := tradeClient.GetAccountTradeList("CFXUSDT", 0, 0, 0, 0, 0)
	if err != nil {
		logger.Error("getAccountTradesList err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getAccountTradesList API error: %v", response.(model.APIErrorResponse))
	case spotclient.AccountTradesListResponse:
		logger.Info("getAccountTradesList: %v", response.(spotclient.AccountTradesListResponse))
	default:
		logger.Info("getAccountTradesList response: %v", response)
	}
}
