package spotclient

import (
	binance "github.com/dirname/binance"
	"github.com/dirname/binance/model"
	"reflect"
	"testing"
)

func TestBalanceWebsocketClient_GetCombined(t *testing.T) {
	type fields struct {
		WebsocketClient binance.WebsocketClient
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestBalanceWebsocketClient_GetCombined", fields{binance.WebsocketClient{}}, args{0}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &BalanceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			u.GetCombined(tt.args.id)
		})
	}
}

func TestBalanceWebsocketClient_GetSubscribe(t *testing.T) {
	type fields struct {
		WebsocketClient binance.WebsocketClient
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestBalanceWebsocketClient_GetSubscribe", fields{binance.WebsocketClient{}}, args{0}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &BalanceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			u.GetSubscribe(tt.args.id)
		})
	}
}

func TestBalanceWebsocketClient_SetCombined(t *testing.T) {
	type fields struct {
		WebsocketClient binance.WebsocketClient
	}
	type args struct {
		b  bool
		id uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestBalanceWebsocketClient_SetCombined", fields{binance.WebsocketClient{}}, args{
			b:  false,
			id: 0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &BalanceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			u.SetCombined(tt.args.b, tt.args.id)
		})
	}
}

func TestBalanceWebsocketClient_SetHandler(t *testing.T) {
	type fields struct {
		WebsocketClient binance.WebsocketClient
	}
	type args struct {
		connectHandler  binance.ConnectedHandler
		responseHandler binance.ResponseHandler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestBalanceWebsocketClient_SetHandler", fields{binance.WebsocketClient{}}, args{
			connectHandler:  nil,
			responseHandler: nil,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &BalanceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			u.SetHandler(tt.args.connectHandler, tt.args.responseHandler)
		})
	}
}

func TestBalanceWebsocketClient_Subscribe(t *testing.T) {
	type fields struct {
		WebsocketClient binance.WebsocketClient
	}
	type args struct {
		id     uint
		params []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestBalanceWebsocketClient_Subscribe", fields{binance.WebsocketClient{}}, args{
			id:     0,
			params: nil,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &BalanceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			u.Subscribe(tt.args.id, tt.args.params...)
		})
	}
}

func TestBalanceWebsocketClient_Unsubscribe(t *testing.T) {
	type fields struct {
		WebsocketClient binance.WebsocketClient
	}
	type args struct {
		id     uint
		params []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestBalanceWebsocketClient_Unsubscribe", fields{binance.WebsocketClient{}}, args{
			id:     0,
			params: nil,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &BalanceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			u.Unsubscribe(tt.args.id, tt.args.params...)
		})
	}
}

func TestBalanceWebsocketClient_handleMessage(t *testing.T) {
	type fields struct {
		WebsocketClient binance.WebsocketClient
	}
	type args struct {
		msg []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestBalanceWebsocketClient_handleMessage", fields{WebsocketClient: binance.WebsocketClient{}}, args{msg: []byte("{\"code\":0,\"msg\":\"Unknown property\",\"id\":0}")}, model.WebsocketErrorResponse{
			Code:    0,
			Message: "Unknown property",
		}, false},
		{"TestAccountWebsocketClient_handleMessage", fields{WebsocketClient: binance.WebsocketClient{}}, args{msg: []byte("{\"stream\": \"test\", \"data\": {\"e\": \"balanceUpdate\"}}")}, BalanceCombinedUpdate{
			StreamName: "test",
			Data:       BalanceUpdate{EventType: "balanceUpdate"},
		}, false},
		{"TestBalanceWebsocketClient_handleMessage", fields{WebsocketClient: binance.WebsocketClient{}}, args{msg: []byte("{\"e\":\"balanceUpdate\",\"E\":1573200697110,\"a\":\"ABC\",\"d\":\"100.00000000\",\"T\":1573200697068}")}, BalanceUpdate{
			EventType:    "balanceUpdate",
			EventTime:    1573200697110,
			Asset:        "ABC",
			BalanceDelta: "100.00000000",
			ClearTime:    1573200697068,
		}, false},
		{"TestBalanceWebsocketClient_handleMessage", fields{WebsocketClient: binance.WebsocketClient{}}, args{msg: []byte("{\"result\":\"test\"}")}, model.WebsocketCommonResponse{
			Result: "test",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &BalanceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			got, err := u.handleMessage(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("handleMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBalanceWebsocketClient(t *testing.T) {
	type args struct {
		streams []string
	}
	tests := []struct {
		name string
		args args
		want *BalanceWebsocketClient
	}{
		{"TestNewBalanceWebsocketClient", args{streams: []string{"BTCUSDT@markPrice"}}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBalanceWebsocketClient(tt.args.streams...); reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBalanceWebsocketClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
