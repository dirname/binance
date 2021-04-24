package spotclient

import (
	"encoding/json"
	"errors"
	"fmt"
	binance "github.com/dirname/binance"
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
	req, err := m.Builder.Build(http.MethodGet, "/api/v3/exchangeInfo", "")
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

// GetOrderBook get depth info  Valid limits:[5, 10, 20, 50, 100, 500, 1000, 5000]
func (m *MarketClient) GetOrderBook(symbol string, limit int32) (interface{}, error) {
	var err error
	if symbol == "" {
		err := errors.New(SymbolEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s", symbol)
	if limit > 5000 {
		limit = 5000
	}
	if limit > 0 {
		params += fmt.Sprintf("&limit=%d", limit)
	}
	req, err := m.Builder.Build(http.MethodGet, "/api/v3/depth", params)
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
	req, err := m.Builder.Build(http.MethodGet, "/api/v3/trades", params)
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
func (m *MarketClient) GetOldTradeLookUp(symbol string, limit int32, formID int64) (interface{}, error) {
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
	req, err := m.Builder.Build(http.MethodGet, "/api/v3/historicalTrades", params)
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
	req, err := m.Builder.Build(http.MethodGet, "/api/v3/aggTrades", params)
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
	if limit > 1000 {
		limit = 1000
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
	req, err := m.Builder.Build(http.MethodGet, "/api/v3/klines", params)
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

// GetAveragePrice Current average price for a symbol.
func (m *MarketClient) GetAveragePrice(symbol string) (interface{}, error) {
	var err error
	if symbol == "" {
		err := errors.New(SymbolEmpty)
		return nil, err
	}
	params := fmt.Sprintf("symbol=%s", symbol)
	req, err := m.Builder.Build(http.MethodGet, "/api/v3/avgPrice", params)
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
	result := CurrentAveragePriceResponse{}
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
	req, err := m.Builder.Build(http.MethodGet, "/api/v3/ticker/24hr", params)
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

// GetSymbolTickerPrice Latest price for a symbol or symbols.
func (m *MarketClient) GetSymbolTickerPrice(symbol string) (interface{}, error) {
	var err error
	params := ""
	if symbol != "" {
		params = fmt.Sprintf("symbol=%s", symbol)
	}
	req, err := m.Builder.Build(http.MethodGet, "/api/v3/ticker/price", params)
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
	req, err := m.Builder.Build(http.MethodGet, "/api/v3/ticker/bookTicker", params)
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
