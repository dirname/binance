package spotclient

import (
	"encoding/json"
	binance "github.com/dirname/binance"
	"github.com/dirname/binance/config"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
)

// SpotCandlestickWebsocketClient responsible to handle market price data from websocket
type SpotCandlestickWebsocketClient struct {
	binance.WebsocketClient
}

// NewSpotCandlestickWebsocketClient Factory function
func NewSpotCandlestickWebsocketClient(streams ...string) *SpotCandlestickWebsocketClient {
	c := new(SpotCandlestickWebsocketClient)
	c.WebsocketClient.Init(config.SpotWssHost, streams...)
	return c
}

// Subscribe subscribe market price data
func (u *SpotCandlestickWebsocketClient) Subscribe(id uint, params ...string) {

	req := model.WebsocketCommon{
		Method: "SUBSCRIBE",
		Params: params,
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket subscribed, params: %s id: %d", params, id)
}

// Unsubscribe unsubscribe market price data
func (u *SpotCandlestickWebsocketClient) Unsubscribe(id uint, params ...string) {
	req := model.WebsocketCommon{
		Method: "UNSUBSCRIBE",
		Params: params,
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket unsubscribed, params: %s id: %d", params, id)
}

// SetCombined set combined, it is true when stream's length is greater than one, and false if stream's length is equal to one
func (u *SpotCandlestickWebsocketClient) SetCombined(b bool, id uint) {
	req := model.WebsocketCommon{
		Method: "SET_PROPERTY",
		Params: []interface{}{"combined", b},
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket set combined, params: %v", req)
}

// GetCombined get combined
func (u *SpotCandlestickWebsocketClient) GetCombined(id uint) {
	req := model.WebsocketCommon{
		Method: "GET_PROPERTY",
		Params: []interface{}{"combined"},
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket set combined, params: %v", req)
}

// GetSubscribe get subscribed list
func (u *SpotCandlestickWebsocketClient) GetSubscribe(id uint) {
	req := model.WebsocketCommon{
		Method: "LIST_SUBSCRIPTIONS",
		Params: nil,
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket get subscribed, params: %v", req)
}

// SetHandler set callback handler
func (u *SpotCandlestickWebsocketClient) SetHandler(connectHandler binance.ConnectedHandler, responseHandler binance.ResponseHandler) {
	u.WebsocketClient.SetConnectedHandler(connectHandler)
	u.WebsocketClient.SetResponseHandler(responseHandler)
	u.WebsocketClient.SetMessageHandler(u.handleMessage)
}

func (u *SpotCandlestickWebsocketClient) handleMessage(msg []byte) (interface{}, error) {
	var parser map[string]interface{}
	var err error
	err = json.Unmarshal(msg, &parser)
	if _, ok := parser["stream"]; ok {
		result := CandlestickCombinedResponse{}
		err = json.Unmarshal(msg, &result)
		return result, err
	}
	if _, ok := parser["result"]; ok {
		result := model.WebsocketCommonResponse{}
		err = json.Unmarshal(msg, &result)
		return result, err
	}
	if _, ok := parser["e"]; ok {
		result := CandlestickResponse{}
		err = json.Unmarshal(msg, &result)
		return result, err
	}
	if _, ok := parser["code"]; ok {
		result := model.WebsocketErrorResponse{}
		err = json.Unmarshal(msg, &result)
		return result, err
	}
	return parser, err
}
