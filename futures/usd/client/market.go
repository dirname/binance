package futuresclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dirname/binance"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"net/http"
)

// MarketClient responsible to get market information
type MarketClient struct {
	Builder *binance.PublicUrlBuilder
	AppKey  string
}

// NewMarketClient Initializer
func NewMarketClient(host, appKey string) *MarketClient {
	return &MarketClient{
		Builder: binance.NewPublicUrlBuilder(host),
		AppKey:  appKey,
	}
}

// GetExchangeInfo Current exchange trading rules and symbol information
func (m *MarketClient) GetExchangeInfo() (interface{}, error) {
	var err error
	req, err := m.Builder.Build(http.MethodGet, "/fapi/v1/exchangeInfo", "")
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := ExchangeInfoResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//GetOrderBook get depth info  Valid limits:[5, 10, 20, 50, 100, 500, 1000]
func (m *MarketClient) GetOrderBook(symbol string, limit int32) (interface{}, error) {
	var err error
	if symbol == "" {
		err := errors.New(SymbolEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s", symbol)
	if limit > 1000 {
		limit = 1000
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	req, err := m.Builder.Build(http.MethodGet, "/fapi/v1/depth", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := OrderBookResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// GetRecentTrades Get recent trades.
func (m *MarketClient) GetRecentTrades(symbol string, limit int32) (interface{}, error) {
	var err error
	if symbol == "" {
		err := errors.New(SymbolEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s", symbol)
	if limit > 1000 {
		limit = 1000
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	req, err := m.Builder.Build(http.MethodGet, "/fapi/v1/trades", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := RecentTradesListResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// GetOldTradeLookUp Get older market trades.
func (m *MarketClient) GetOldTradeLookUp(symbol string, limit int32, formID, startTime, endTime int64) (interface{}, error) {
	var err error
	if symbol == "" {
		err := errors.New(SymbolEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s", symbol)
	if limit > 1000 {
		limit = 1000
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	if formID > 0 {
		params += fmt.Sprintf("&formId=%d", formID)
	}
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	req, err := m.Builder.Build(http.MethodGet, "/fapi/v1/historicalTrades", params)
	req.Header.Set("X-MBX-APIKEY", m.AppKey)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := OlderTradeLookUpResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// GetAggregateTrades Get compressed, aggregate trades. Trades that fill at the time, from the same order, with the same price will have the quantity aggregated.
func (m *MarketClient) GetAggregateTrades(symbol string, limit int32, formID, startTime, endTime int64) (interface{}, error) {
	var err error
	if symbol == "" {
		err := errors.New(SymbolEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s", symbol)
	if limit > 1000 {
		limit = 1000
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	if formID > 0 {
		params += fmt.Sprintf("&formId=%d", formID)
	}
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	req, err := m.Builder.Build(http.MethodGet, "/fapi/v1/aggTrades", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := AggregateTradeResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// GetCandlestick Kline/candlestick bars for a symbol and are uniquely identified by their open time. It is not recommended to use this method, this method may have type errors, please use websocket subscription to obtain data.
func (m *MarketClient) GetCandlestick(symbol, interval string, limit int32, startTime, endTime int64) (interface{}, error) {
	var err error
	if symbol == "" {
		err := errors.New(SymbolEmpty)
		return nil, err
	}
	if interval == "" {
		err := errors.New(IntervalEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s&interval=%s", symbol, interval)
	if limit > 1500 {
		limit = 1500
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	req, err := m.Builder.Build(http.MethodGet, "/fapi/v1/klines", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := CandlestickResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// GetContractCandlestick Kline/candlestick bars for a specific contract type. Kline are uniquely identified by their open time.
func (m *MarketClient) GetContractCandlestick(pair, contractType, interval string, limit int32, startTime, endTime int64) (interface{}, error) {
	var err error
	if pair == "" {
		err := errors.New(PairEmpty)
		return nil, err
	}
	if interval == "" {
		err := errors.New(IntervalEmpty)
		return nil, err
	}
	if contractType == "" {
		err := errors.New(ContractEmpty)
		return nil, err
	}
	params := fmt.Sprintf("pair=%s&interval=%s&contractType=%s", pair, interval, contractType)
	if limit > 1500 {
		limit = 1500
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	req, err := m.Builder.Build(http.MethodGet, "/fapi/v1/continuousKlines", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := ContractCandlestickResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// GetIndexCandlestick Kline/candlestick bars for the index price of a pair. Kline are uniquely identified by their open time.
func (m *MarketClient) GetIndexCandlestick(pair, interval string, limit int32, startTime, endTime int64) (interface{}, error) {
	var err error
	if pair == "" {
		err := errors.New(PairEmpty)
		return nil, err
	}
	if interval == "" {
		err := errors.New(IntervalEmpty)
		return nil, err
	}
	params := fmt.Sprintf("pair=%s&interval=%s", pair, interval)
	if limit > 1500 {
		limit = 1500
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	req, err := m.Builder.Build(http.MethodGet, "/fapi/v1/indexPriceKlines", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := IndexCandlestickResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// GetMarketPriceCandlestick Kline/candlestick bars for the mark price of a symbol. Kline are uniquely identified by their open time.
func (m *MarketClient) GetMarketPriceCandlestick(symbol, interval string, limit int32, startTime, endTime int64) (interface{}, error) {
	var err error
	if symbol == "" {
		err := errors.New(SymbolEmpty)
		return nil, err
	}
	if interval == "" {
		err := errors.New(IntervalEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s&interval=%s", symbol, interval)
	if limit > 1500 {
		limit = 1500
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	req, err := m.Builder.Build(http.MethodGet, "/fapi/v1/markPriceKlines", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := MarketPriceCandlestickResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// GetMarketPrice Mark Price and Funding Rate
func (m *MarketClient) GetMarketPrice(symbol string) (interface{}, error) {
	var err error
	params := ""
	if symbol != "" {
		params = fmt.Sprintf("symbol=%s", symbol)
	}
	req, err := m.Builder.Build(http.MethodGet, "/fapi/v1/premiumIndex", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	if params == "" {
		var result []MarketPriceCandlestickResponse
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := MarketPriceResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//GetFundingRateHistory get funding rate history
func (m *MarketClient) GetFundingRateHistory(symbol string, startTime, endTime int64, limit int32) (interface{}, error) {
	var err error
	params := ""
	if symbol != "" {
		params = fmt.Sprintf("symbol=%s", symbol)
	}
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	if limit > 1000 {
		limit = 1000
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	req, err := m.Builder.Build(http.MethodGet, "/fapi/v1/fundingRate", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := FundingRateResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// Get24hTickerPriceChange 24 hour rolling window price change statistics. Careful when accessing this with no symbol.
func (m *MarketClient) Get24hTickerPriceChange(symbol string) (interface{}, error) {
	var err error
	params := ""
	if symbol != "" {
		params = fmt.Sprintf("symbol=%s", symbol)
	}
	req, err := m.Builder.Build(http.MethodGet, "/fapi/v1/ticker/24hr", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	if params == "" {
		var result []TickerPriceChangeStatisticsResponse
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := TickerPriceChangeStatisticsResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//GetSymbolTicker get symbol ticker
func (m *MarketClient) GetSymbolTicker(symbol string) (interface{}, error) {
	var err error
	params := ""
	if symbol != "" {
		params = fmt.Sprintf("symbol=%s", symbol)
	}
	req, err := m.Builder.Build(http.MethodGet, "/fapi/v1/ticker/price", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	if params == "" {
		var result []SymbolPriceTickerResponse
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := SymbolPriceTickerResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// GetSymbolOrderBookTicker Best price/qty on the order book for a symbol or symbols.
func (m *MarketClient) GetSymbolOrderBookTicker(symbol string) (interface{}, error) {
	var err error
	params := ""
	if symbol != "" {
		params = fmt.Sprintf("symbol=%s", symbol)
	}
	req, err := m.Builder.Build(http.MethodGet, "/fapi/v1/ticker/bookTicker", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	if params == "" {
		var result []SymbolOrderBookTickerResponse
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := SymbolOrderBookTickerResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//GetOpenInterest Get present open interest of a specific symbol.
func (m *MarketClient) GetOpenInterest(symbol string) (interface{}, error) {
	var err error
	if symbol == "" {
		err := errors.New(SymbolEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s", symbol)
	req, err := m.Builder.Build(http.MethodGet, "/fapi/v1/openInterest", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := OpenInterestResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//GetOpenInterestStatistics Get interest statistics
func (m *MarketClient) GetOpenInterestStatistics(symbol, period string, startTime, endTime int64, limit int32) (interface{}, error) {
	var err error
	if symbol == "" {
		err := errors.New(SymbolEmpty)
		return nil, err
	}
	if period == "" {
		err := errors.New(PeriodEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s&period=%s", symbol, period)
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	if limit > 500 {
		limit = 500
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	req, err := m.Builder.Build(http.MethodGet, "/futures/data/openInterestHist", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := OpenInterestStatisticsResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//GetTopTradeAccountsRatio get top trader long/short ratio
func (m *MarketClient) GetTopTradeAccountsRatio(symbol, period string, startTime, endTime int64, limit int32) (interface{}, error) {
	var err error
	if symbol == "" {
		err := errors.New(SymbolEmpty)
		return nil, err
	}
	if period == "" {
		err := errors.New(PeriodEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s&period=%s", symbol, period)
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	if limit > 500 {
		limit = 500
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	req, err := m.Builder.Build(http.MethodGet, "/futures/data/topLongShortAccountRatio", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := TopTraderAccountsRatioResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//GetTopTradePositionsRatio get top trader long/short ratio
func (m *MarketClient) GetTopTradePositionsRatio(symbol, period string, startTime, endTime int64, limit int32) (interface{}, error) {
	var err error
	if symbol == "" {
		err := errors.New(SymbolEmpty)
		return nil, err
	}
	if period == "" {
		err := errors.New(PeriodEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s&period=%s", symbol, period)
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	if limit > 500 {
		limit = 500
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	req, err := m.Builder.Build(http.MethodGet, "/futures/data/topLongShortPositionRatio", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := TopTraderPositionsRatioResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//GetLongShortRatio get long short ratio response
func (m *MarketClient) GetLongShortRatio(symbol, period string, startTime, endTime int64, limit int32) (interface{}, error) {
	var err error
	if symbol == "" {
		err := errors.New(SymbolEmpty)
		return nil, err
	}
	if period == "" {
		err := errors.New(PeriodEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s&period=%s", symbol, period)
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	if limit > 500 {
		limit = 500
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	req, err := m.Builder.Build(http.MethodGet, "/futures/data/globalLongShortAccountRatio", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := LongShortRatioResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//GetTakerBuySellVolume get taker buy/sell volume
func (m *MarketClient) GetTakerBuySellVolume(symbol, period string, startTime, endTime int64, limit int32) (interface{}, error) {
	var err error
	if symbol == "" {
		err := errors.New(SymbolEmpty)
		return nil, err
	}
	if period == "" {
		err := errors.New(PeriodEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s&period=%s", symbol, period)
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	if limit > 500 {
		limit = 500
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	req, err := m.Builder.Build(http.MethodGet, "/futures/data/takerlongshortRatio", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := TakerBuySellVolumeResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//GetHistoricalBLVTNavCandlestick The BLVT NAV system is based on Binance Futures, so the endpoint is based on fapi
func (m *MarketClient) GetHistoricalBLVTNavCandlestick(symbol, interval string, startTime, endTime int64, limit int32) (interface{}, error) {
	var err error
	if symbol == "" {
		err := errors.New(SymbolEmpty)
		return nil, err
	}
	if interval == "" {
		err := errors.New(IntervalEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s&interval=%s", symbol, interval)
	if startTime > 0 {
		params += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		params += fmt.Sprintf("&endTime=%d", endTime)
	}
	if limit > 1000 {
		limit = 1000
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	req, err := m.Builder.Build(http.MethodGet, "/fapi/v1/lvtKlines", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := HistoricalBLVTNavCandlestickResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

//GetCompositeIndexSymbol get composite index symbol
func (m *MarketClient) GetCompositeIndexSymbol(symbol string) (interface{}, error) {
	var err error
	params := ""
	if symbol != "" {
		params = fmt.Sprintf("symbol=%s", symbol)
	}
	req, err := m.Builder.Build(http.MethodGet, "/fapi/v1/indexInfo", params)
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}
	if params == "" {
		var result []CompositeIndexSymbolResponse
		err = json.Unmarshal(res, &result)
		return result, err
	}
	result := CompositeIndexSymbolResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}
