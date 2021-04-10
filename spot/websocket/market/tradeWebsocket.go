package spotclient

import (
	"encoding/json"
	binance "github.com/dirname/Binance"
	"github.com/dirname/Binance/config"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
)

// SpotTradeWebsocketClient responsible to handle market price data from websocket
type SpotTradeWebsocketClient struct {
	binance.WebsocketClient
}

// NewSpotTradeWebsocketClient Factory function
func NewSpotTradeWebsocketClient(streams ...string) *SpotTradeWebsocketClient {
	c := new(SpotTradeWebsocketClient)
	c.WebsocketClient.Init(config.SpotWssHost, streams...)
	return c
}

// Subscribe subscribe market price data
func (u *SpotTradeWebsocketClient) Subscribe(id uint, params ...string) {

	req := model.WebsocketCommon{
		Method: "SUBSCRIBE",
		Params: params,
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket subscribed, params: %s id: %d", params, id)
}

// Unsubscribe unsubscribe market price data
func (u *SpotTradeWebsocketClient) Unsubscribe(id uint, params ...string) {
	req := model.WebsocketCommon{
		Method: "UNSUBSCRIBE",
		Params: params,
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket unsubscribed, params: %s id: %d", params, id)
}

// SetCombined set combined, it is true when stream's length is greater than one, and false if stream's length is equal to one
func (u *SpotTradeWebsocketClient) SetCombined(b bool, id uint) {
	req := model.WebsocketCommon{
		Method: "SET_PROPERTY",
		Params: []interface{}{"combined", b},
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket set combined, params: %v", req)
}

// GetCombined get combined
func (u *SpotTradeWebsocketClient) GetCombined(id uint) {
	req := model.WebsocketCommon{
		Method: "GET_PROPERTY",
		Params: []interface{}{"combined"},
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket set combined, params: %v", req)
}

// GetSubscribe get subscribed list
func (u *SpotTradeWebsocketClient) GetSubscribe(id uint) {
	req := model.WebsocketCommon{
		Method: "LIST_SUBSCRIPTIONS",
		Params: nil,
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket get subscribed, params: %v", req)
}

// SetHandler set callback handler
func (u *SpotTradeWebsocketClient) SetHandler(connectHandler binance.ConnectedHandler, responseHandler binance.ResponseHandler) {
	u.WebsocketClient.SetConnectedHandler(connectHandler)
	u.WebsocketClient.SetResponseHandler(responseHandler)
	u.WebsocketClient.SetMessageHandler(u.handleMessage)
}

func (u *SpotTradeWebsocketClient) handleMessage(msg []byte) (interface{}, error) {
	var parser map[string]interface{}
	var err error
	err = json.Unmarshal(msg, &parser)
	if _, ok := parser["stream"]; ok {
		result := TradeCombinedResponse{}
		err = json.Unmarshal(msg, &result)
		return result, err
	}
	if _, ok := parser["result"]; ok {
		result := model.WebsocketCommonResponse{}
		err = json.Unmarshal(msg, &result)
		return result, err
	}
	if _, ok := parser["e"]; ok {
		result := TradeResponse{}
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
