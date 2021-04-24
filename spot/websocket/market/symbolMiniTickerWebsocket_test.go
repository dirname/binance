package spotclient

import (
	binance "github.com/dirname/binance"
	"github.com/dirname/binance/model"
	"reflect"
	"testing"
)

func TestNewSpotSymbolMiniTickerWebsocketClient(t *testing.T) {
	type args struct {
		streams []string
	}
	tests := []struct {
		name string
		args args
		want *SpotSymbolMiniTickerWebsocketClient
	}{
		{"TestNewSpotSymbolMiniTickerWebsocketClient", args{streams: []string{"BTCUSDT@markPrice"}}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSpotSymbolMiniTickerWebsocketClient(tt.args.streams...); reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpotSymbolMiniTickerWebsocketClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpotSymbolMiniTickerWebsocketClient_GetCombined(t *testing.T) {
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
		{"TestSpotSymbolMiniTickerWebsocketClient_GetCombined", fields{WebsocketClient: binance.WebsocketClient{}}, args{0}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &SpotSymbolMiniTickerWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			u.GetCombined(tt.args.id)
		})
	}
}

func TestSpotSymbolMiniTickerWebsocketClient_GetSubscribe(t *testing.T) {
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
		{"TestSpotSymbolMiniTickerWebsocketClient_GetSubscribe", fields{WebsocketClient: binance.WebsocketClient{}}, args{0}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &SpotSymbolMiniTickerWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			u.GetSubscribe(tt.args.id)
		})
	}
}

func TestSpotSymbolMiniTickerWebsocketClient_SetCombined(t *testing.T) {
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
		{"TestSpotSymbolMiniTickerWebsocketClient_SetCombined", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			b:  false,
			id: 0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &SpotSymbolMiniTickerWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			u.SetCombined(tt.args.b, tt.args.id)
		})
	}
}

func TestSpotSymbolMiniTickerWebsocketClient_SetHandler(t *testing.T) {
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
		{"TestSpotSymbolMiniTickerWebsocketClient_SetHandler", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			connectHandler:  nil,
			responseHandler: nil,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &SpotSymbolMiniTickerWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			u.SetHandler(tt.args.connectHandler, tt.args.responseHandler)
		})
	}
}

func TestSpotSymbolMiniTickerWebsocketClient_Subscribe(t *testing.T) {
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
		{"TestSpotSymbolMiniTickerWebsocketClient_Subscribe", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			id:     0,
			params: nil,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &SpotSymbolMiniTickerWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			u.Subscribe(tt.args.id, tt.args.params...)
		})
	}
}

func TestSpotSymbolMiniTickerWebsocketClient_Unsubscribe(t *testing.T) {
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
		{"TestSpotSymbolMiniTickerWebsocketClient_Unsubscribe", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			id:     0,
			params: nil,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &SpotSymbolMiniTickerWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			u.Unsubscribe(tt.args.id, tt.args.params...)
		})
	}
}

func TestSpotSymbolMiniTickerWebsocketClient_handleMessage(t *testing.T) {
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
		{"TestSpotSymbolMiniTickerWebsocketClient_handleMessage", fields{WebsocketClient: binance.WebsocketClient{}}, args{msg: []byte("{\"code\":0,\"msg\":\"Unknown property\",\"id\":0}")}, model.WebsocketErrorResponse{
			Code:    0,
			Message: "Unknown property",
		}, false},
		{"TestSpotSymbolMiniTickerWebsocketClient_handleMessage", fields{WebsocketClient: binance.WebsocketClient{}}, args{msg: []byte("{\"stream\":\"test\"}")}, SymbolMiniTickerCombinedResponse{
			StreamName: "test",
		}, false},
		{"TestSpotSymbolMiniTickerWebsocketClient_handleMessage", fields{WebsocketClient: binance.WebsocketClient{}}, args{msg: []byte("{\"e\":\"test\"}")}, SymbolMiniTickerResponse{
			EventType: "test",
		}, false},
		{"TestSpotSymbolMiniTickerWebsocketClient_handleMessage", fields{WebsocketClient: binance.WebsocketClient{}}, args{msg: []byte("{\"result\":\"test\"}")}, model.WebsocketCommonResponse{
			Result: "test",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &SpotSymbolMiniTickerWebsocketClient{
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
