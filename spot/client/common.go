package spotclient

import (
	"encoding/json"
	binance "github.com/dirname/Binance"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
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

// GetSAPISystemStatus get system status
func (c *CommonClient) GetSAPISystemStatus() (interface{}, error) {
	var err error
	req, err := c.Builder.Build("GET", "/sapi/v1/system/status", "")
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
	result := SystemStatusResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}

// Ping test server
func (c *CommonClient) Ping() (interface{}, error) {
	var err error
	req, err := c.Builder.Build(http.MethodGet, "/api/v3/ping", "")
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
	req, err := c.Builder.Build(http.MethodGet, "/api/v3/time", "")
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
