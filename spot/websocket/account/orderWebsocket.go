package spotclient

import (
	"encoding/json"
	binance "github.com/dirname/binance"
	"github.com/dirname/binance/config"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
)

// OrderWebsocketClient responsible to handle market price data from websocket
type OrderWebsocketClient struct {
	binance.WebsocketClient
}

// NewOrderWebsocketClient Factory function
func NewOrderWebsocketClient(streams ...string) *OrderWebsocketClient {
	c := new(OrderWebsocketClient)
	c.WebsocketClient.Init(config.SpotWssHost, streams...)
	return c
}

// Subscribe subscribe market price data
func (u *OrderWebsocketClient) Subscribe(id uint, params ...string) {

	req := model.WebsocketCommon{
		Method: "SUBSCRIBE",
		Params: params,
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket subscribed, params: %s id: %d", params, id)
}

// Unsubscribe unsubscribe market price data
func (u *OrderWebsocketClient) Unsubscribe(id uint, params ...string) {
	req := model.WebsocketCommon{
		Method: "UNSUBSCRIBE",
		Params: params,
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket unsubscribed, params: %s id: %d", params, id)
}

// SetCombined set combined, it is true when stream's length is greater than one, and false if stream's length is equal to one
func (u *OrderWebsocketClient) SetCombined(b bool, id uint) {
	req := model.WebsocketCommon{
		Method: "SET_PROPERTY",
		Params: []interface{}{"combined", b},
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket set combined, params: %v", req)
}

// GetCombined get combined
func (u *OrderWebsocketClient) GetCombined(id uint) {
	req := model.WebsocketCommon{
		Method: "GET_PROPERTY",
		Params: []interface{}{"combined"},
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket set combined, params: %v", req)
}

// GetSubscribe get subscribed list
func (u *OrderWebsocketClient) GetSubscribe(id uint) {
	req := model.WebsocketCommon{
		Method: "LIST_SUBSCRIPTIONS",
		Params: nil,
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket get subscribed, params: %v", req)
}

// SetHandler set callback handler
func (u *OrderWebsocketClient) SetHandler(connectHandler binance.ConnectedHandler, responseHandler binance.ResponseHandler) {
	u.WebsocketClient.SetConnectedHandler(connectHandler)
	u.WebsocketClient.SetResponseHandler(responseHandler)
	u.WebsocketClient.SetMessageHandler(u.handleMessage)
}

func (u *OrderWebsocketClient) handleMessage(msg []byte) (interface{}, error) {
	var parser map[string]interface{}
	var err error
	err = json.Unmarshal(msg, &parser)
	if _, ok := parser["stream"]; ok {
		judge := OrderCombinedJudge{}
		err = json.Unmarshal(msg, &judge)
		if judge.Data.EventType == "executionReport" {
			result := ExecutionCombinedReport{}
			err = json.Unmarshal(msg, &result)
			return result, err
		}
		if judge.Data.EventType == "listStatus" {
			result := ListCombinedStatus{}
			err = json.Unmarshal(msg, &result)
			return result, err
		}
		return judge, err
	}
	if _, ok := parser["result"]; ok {
		result := model.WebsocketCommonResponse{}
		err = json.Unmarshal(msg, &result)
		return result, err
	}
	if _, ok := parser["s"]; ok {
		judge := OrderJudge{}
		err = json.Unmarshal(msg, &judge)
		if judge.EventType == "executionReport" {
			result := ExecutionReport{}
			err = json.Unmarshal(msg, &result)
			return result, err
		}
		if judge.EventType == "listStatus" {
			result := ListStatus{}
			err = json.Unmarshal(msg, &result)
			return result, err
		}
		return judge, err
	}
	if _, ok := parser["code"]; ok {
		result := model.WebsocketErrorResponse{}
		err = json.Unmarshal(msg, &result)
		return result, err
	}
	return parser, err
}
