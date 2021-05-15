package main

import (
	"github.com/dirname/binance/config"
	"github.com/dirname/binance/futures/usd/client"
	"github.com/dirname/binance/futures/usd/client/define/contractType"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
)

var marketClient *futuresclient.MarketClient

func init() {
	logger.Enable(false)
	marketClient = futuresclient.NewMarketClient(config.USDFuturesRestHost, config.AppKey)
}

func main() {
	getExchangeInfo()
	getOrderBook()
	getRecentOrder()
	getOldTradeLookUp()
	getAggregateTrades()
	getCandlestick()
	getContractCandlestick()
	getIndexCandlestick()
	getMarketPriceCandlestick()
	getMarketPrice()
	getFundingRateHistory()
	get24hTickerPriceChange()
	getSymbolTicker()
	getSymbolOrderBookTicker()
	getOpenInterest()
	getOpenInterestStatistics()
	getTopTradeAccountsRatio()
	getTopTradePositionsRatio()
	getLongShortRatio()
	getTakerBuySellVolume()
	getHistoricalBLVTNavCandlestick()
	getCompositeIndexSymbol()
}

func getExchangeInfo() {
	response, err := marketClient.GetExchangeInfo()
	if err != nil {
		logger.Error("getExchangeInfo err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getExchangeInfo API error: %v", response.(model.APIErrorResponse))
	case futuresclient.ExchangeInfoResponse:
		logger.Info("getExchangeInfo: %v", response.(futuresclient.ExchangeInfoResponse))
	default:
		logger.Info("getExchangeInfo Unknown response: %v", response)
	}
}

func getOrderBook() {
	response, err := marketClient.GetOrderBook("BTCUSDT", 5)
	if err != nil {
		logger.Error("getOrderBook err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getOrderBook API error: %v", response.(model.APIErrorResponse))
	case futuresclient.OrderBookResponse:
		logger.Info("getOrderBook: %v", response.(futuresclient.OrderBookResponse))
	default:
		logger.Info("getOrderBook Unknown response: %v", response)
	}
}

func getRecentOrder() {
	response, err := marketClient.GetRecentTrades("BTCUSDT", 5)
	if err != nil {
		logger.Error("getRecentOrder err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getRecentOrder API error: %v", response.(model.APIErrorResponse))
	case futuresclient.RecentTradesListResponse:
		logger.Info("getRecentOrder: %v", response.(futuresclient.RecentTradesListResponse))
	default:
		logger.Info("getRecentOrder Unknown response: %v", response)
	}
}

func getOldTradeLookUp() {
	response, err := marketClient.GetOldTradeLookUp("BTCUSDT", 5, 0, 0, 0)
	if err != nil {
		logger.Error("getOldTradeLookUp err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getOldTradeLookUp API error: %v", response.(model.APIErrorResponse))
	case futuresclient.OlderTradeLookUpResponse:
		logger.Info("getOldTradeLookUp: %v", response.(futuresclient.OlderTradeLookUpResponse))
	default:
		logger.Info("getOldTradeLookUp Unknown response: %v", response)
	}
}

func getAggregateTrades() {
	response, err := marketClient.GetAggregateTrades("BTCUSDT", 5, 0, 0, 0)
	if err != nil {
		logger.Error("getAggregateTrades err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getAggregateTrades API error: %v", response.(model.APIErrorResponse))
	case futuresclient.AggregateTradeResponse:
		logger.Info("getAggregateTrades: %v", response.(futuresclient.AggregateTradeResponse))
	default:
		logger.Info("getAggregateTrades Unknown response: %v", response)
	}
}

func getCandlestick() {
	response, err := marketClient.GetCandlestick("BTCUSDT", "1m", 0, 0, 0)
	if err != nil {
		logger.Error("getCandlestick err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getCandlestick API error: %v", response.(model.APIErrorResponse))
	case futuresclient.CandlestickResponse:
		logger.Info("getCandlestick: %v", response.(futuresclient.CandlestickResponse))
	default:
		logger.Info("getCandlestick Unknown response: %v", response)
	}
}

func getContractCandlestick() {
	response, err := marketClient.GetContractCandlestick("BTCUSDT", contractType.Perpetual, "1m", 0, 0, 0)
	if err != nil {
		logger.Error("getContractCandlestick err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getContractCandlestick API error: %v", response.(model.APIErrorResponse))
	case futuresclient.ContractCandlestickResponse:
		logger.Info("getContractCandlestick: %v", response.(futuresclient.ContractCandlestickResponse))
	default:
		logger.Info("getContractCandlestick Unknown response: %v", response)
	}
}

func getIndexCandlestick() {
	response, err := marketClient.GetIndexCandlestick("BTCUSDT", "1m", 0, 0, 0)
	if err != nil {
		logger.Error("getIndexCandlestick err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getIndexCandlestick API error: %v", response.(model.APIErrorResponse))
	case futuresclient.IndexCandlestickResponse:
		logger.Info("getIndexCandlestick: %v", response.(futuresclient.IndexCandlestickResponse))
	default:
		logger.Info("getIndexCandlestick Unknown response: %v", response)
	}
}

func getMarketPriceCandlestick() {
	response, err := marketClient.GetMarketPriceCandlestick("BTCUSDT", "1m", 0, 0, 0)
	if err != nil {
		logger.Error("getMarketPriceCandlestick err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getMarketPriceCandlestick API error: %v", response.(model.APIErrorResponse))
	case futuresclient.MarketPriceCandlestickResponse:
		logger.Info("getMarketPriceCandlestick: %v", response.(futuresclient.MarketPriceCandlestickResponse))
	default:
		logger.Info("getMarketPriceCandlestick Unknown response: %v", response)
	}
}

func getMarketPrice() {
	response, err := marketClient.GetMarketPrice("BTCUSDT")
	if err != nil {
		logger.Error("getMarketPrice err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getMarketPrice API error: %v", response.(model.APIErrorResponse))
	case futuresclient.MarketPriceResponse:
		logger.Info("getMarketPrice: %v", response.(futuresclient.MarketPriceResponse))
	default:
		logger.Info("getMarketPrice Unknown response: %v", response)
	}
}

func getFundingRateHistory() {
	response, err := marketClient.GetFundingRateHistory("BTCUSDT", 0, 0, 0)
	if err != nil {
		logger.Error("getFundingRateHistory err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getFundingRateHistory API error: %v", response.(model.APIErrorResponse))
	case futuresclient.FundingRateResponse:
		logger.Info("getFundingRateHistory: %v", response.(futuresclient.FundingRateResponse))
	default:
		logger.Info("getFundingRateHistory Unknown response: %v", response)
	}
}

func get24hTickerPriceChange() {
	response, err := marketClient.Get24hTickerPriceChange("BTCUSDT")
	if err != nil {
		logger.Error("get24hTickerPriceChange err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("get24hTickerPriceChange API error: %v", response.(model.APIErrorResponse))
	case futuresclient.TickerPriceChangeStatisticsResponse:
		logger.Info("get24hTickerPriceChange: %v", response.(futuresclient.TickerPriceChangeStatisticsResponse))
	default:
		logger.Info("get24hTickerPriceChange Unknown response: %v", response)
	}
}

func getSymbolTicker() {
	response, err := marketClient.GetSymbolTicker("BTCUSDT")
	if err != nil {
		logger.Error("getSymbolTicker err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getSymbolTicker API error: %v", response.(model.APIErrorResponse))
	case futuresclient.SymbolPriceTickerResponse:
		logger.Info("getSymbolTicker: %v", response.(futuresclient.SymbolPriceTickerResponse))
	default:
		logger.Info("getSymbolTicker Unknown response: %v", response)
	}
}

func getSymbolOrderBookTicker() {
	response, err := marketClient.GetSymbolOrderBookTicker("BTCUSDT")
	if err != nil {
		logger.Error("getSymbolOrderBookTicker err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getSymbolOrderBookTicker API error: %v", response.(model.APIErrorResponse))
	case futuresclient.SymbolOrderBookTickerResponse:
		logger.Info("getSymbolOrderBookTicker: %v", response.(futuresclient.SymbolOrderBookTickerResponse))
	default:
		logger.Info("getSymbolOrderBookTicker Unknown response: %v", response)
	}
}

func getOpenInterest() {
	response, err := marketClient.GetOpenInterest("BTCUSDT")
	if err != nil {
		logger.Error("getOpenInterest err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getOpenInterest API error: %v", response.(model.APIErrorResponse))
	case futuresclient.OpenInterestResponse:
		logger.Info("getOpenInterest: %v", response.(futuresclient.OpenInterestResponse))
	default:
		logger.Info("getOpenInterest Unknown response: %v", response)
	}
}

func getOpenInterestStatistics() {
	response, err := marketClient.GetOpenInterestStatistics("BTCUSDT", "5m", 0, 0, 0)
	if err != nil {
		logger.Error("getOpenInterestStatistics err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getOpenInterestStatistics API error: %v", response.(model.APIErrorResponse))
	case futuresclient.OpenInterestStatisticsResponse:
		logger.Info("getOpenInterestStatistics: %v", response.(futuresclient.OpenInterestStatisticsResponse))
	default:
		logger.Info("getOpenInterestStatistics Unknown response: %v", response)
	}
}

func getTopTradeAccountsRatio() {
	response, err := marketClient.GetTopTradeAccountsRatio("BTCUSDT", "5m", 0, 0, 0)
	if err != nil {
		logger.Error("getTopTradeAccountsRatio err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getTopTradeAccountsRatio API error: %v", response.(model.APIErrorResponse))
	case futuresclient.TopTraderAccountsRatioResponse:
		logger.Info("getTopTradeAccountsRatio: %v", response.(futuresclient.TopTraderAccountsRatioResponse))
	default:
		logger.Info("getTopTradeAccountsRatio Unknown response: %v", response)
	}
}

func getTopTradePositionsRatio() {
	response, err := marketClient.GetTopTradePositionsRatio("BTCUSDT", "5m", 0, 0, 0)
	if err != nil {
		logger.Error("getTopTradePositionsRatio err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getTopTradePositionsRatio API error: %v", response.(model.APIErrorResponse))
	case futuresclient.TopTraderPositionsRatioResponse:
		logger.Info("getTopTradePositionsRatio: %v", response.(futuresclient.TopTraderPositionsRatioResponse))
	default:
		logger.Info("getTopTradePositionsRatio Unknown response: %v", response)
	}
}

func getLongShortRatio() {
	response, err := marketClient.GetLongShortRatio("BTCUSDT", "5m", 0, 0, 0)
	if err != nil {
		logger.Error("getLongShortRatio err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getLongShortRatio API error: %v", response.(model.APIErrorResponse))
	case futuresclient.LongShortRatioResponse:
		logger.Info("getLongShortRatio: %v", response.(futuresclient.LongShortRatioResponse))
	default:
		logger.Info("getLongShortRatio Unknown response: %v", response)
	}
}

func getTakerBuySellVolume() {
	response, err := marketClient.GetTakerBuySellVolume("BTCUSDT", "5m", 0, 0, 0)
	if err != nil {
		logger.Error("getTakerBuySellVolume err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getTakerBuySellVolume API error: %v", response.(model.APIErrorResponse))
	case futuresclient.TakerBuySellVolumeResponse:
		logger.Info("getTakerBuySellVolume: %v", response.(futuresclient.TakerBuySellVolumeResponse))
	default:
		logger.Info("getTakerBuySellVolume Unknown response: %v", response)
	}
}

func getHistoricalBLVTNavCandlestick() {
	response, err := marketClient.GetHistoricalBLVTNavCandlestick("BTCDOWN", "1m", 0, 0, 0)
	if err != nil {
		logger.Error("getHistoricalBLVTNavCandlestick err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getHistoricalBLVTNavCandlestick API error: %v", response.(model.APIErrorResponse))
	case futuresclient.HistoricalBLVTNavCandlestickResponse:
		logger.Info("getHistoricalBLVTNavCandlestick: %v", response.(futuresclient.HistoricalBLVTNavCandlestickResponse))
	default:
		logger.Info("getHistoricalBLVTNavCandlestick Unknown response: %v", response)
	}
}

func getCompositeIndexSymbol() {
	response, err := marketClient.GetCompositeIndexSymbol("DEFIUSDT")
	if err != nil {
		logger.Error("getCompositeIndexSymbol err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("getCompositeIndexSymbol API error: %v", response.(model.APIErrorResponse))
	case futuresclient.CompositeIndexSymbolResponse:
		logger.Info("getCompositeIndexSymbol: %v", response.(futuresclient.CompositeIndexSymbolResponse))
	default:
		logger.Info("getCompositeIndexSymbol Unknown response: %v", response)
	}
}
