package client

import (
	"encoding/json"
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
	req, err := m.Builder.Build(http.MethodGet, "/fapi/v3/exchangeInfo", "")
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
