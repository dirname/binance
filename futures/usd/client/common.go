package futuresclient

import (
	"encoding/json"
	"github.com/dirname/binance"
	logger "github.com/dirname/binance/logging"
	"net/http"
)

// CommonClient responsible to get common information
type CommonClient struct {
	Builder *binance.PublicUrlBuilder
}

// NewCommonClient Initializer
func NewCommonClient(host string) *CommonClient {
	return &CommonClient{
		Builder: binance.NewPublicUrlBuilder(host),
	}
}

// Ping test server
func (c *CommonClient) Ping() (interface{}, error) {
	var err error
	req, err := c.Builder.Build(http.MethodGet, "/fapi/v1/ping", "")
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser interface{}
	err = json.Unmarshal(res, &parser)
	return parser, err
}

// GetServerTime get server's time
func (c *CommonClient) GetServerTime() (int64, error) {
	var err error
	req, err := c.Builder.Build(http.MethodGet, "/fapi/v1/time", "")
	if err != nil {
		logger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["serverTime"]; ok {
		result := &struct {
			ServerTime int64 `json:"serverTime"`
		}{}
		err = json.Unmarshal(res, result)
		return result.ServerTime, err
	}
	return 0, err
}
