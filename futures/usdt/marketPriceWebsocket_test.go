package FuturesUSDT

import (
	binance "github.com/dirname/Binance"
	"github.com/dirname/Binance/model"
	"reflect"
	"testing"
)

func TestFuturesMarketPriceWebsocketClient_GetCombined(t *testing.T) {
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
		{"TestFuturesMarketPriceWebsocketClient", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			b:  false,
			id: 0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &FuturesMarketPriceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
		})
	}
}

func TestFuturesMarketPriceWebsocketClient_GetSubscribe(t *testing.T) {
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
		{"TestFuturesMarketPriceWebsocketClient", fields{WebsocketClient: binance.WebsocketClient{}}, args{id: 0}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &FuturesMarketPriceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
		})
	}
}

func TestFuturesMarketPriceWebsocketClient_SetCombined(t *testing.T) {
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
		{"TestFuturesMarketPriceWebsocketClient", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			b:  false,
			id: 0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &FuturesMarketPriceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
		})
	}
}

func TestFuturesMarketPriceWebsocketClient_SetHandler(t *testing.T) {
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
		{"TestFuturesMarketPriceWebsocketClient", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			connectHandler:  nil,
			responseHandler: nil,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &FuturesMarketPriceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
		})
	}
}

func TestFuturesMarketPriceWebsocketClient_Subscribe(t *testing.T) {
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
		{"TestFuturesMarketPriceWebsocketClient", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			id:     0,
			params: nil,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &FuturesMarketPriceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
		})
	}
}

func TestFuturesMarketPriceWebsocketClient_Unsubscribe(t *testing.T) {
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
		{"TestFuturesMarketPriceWebsocketClient", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			id:     0,
			params: nil,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &FuturesMarketPriceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
		})
	}
}

func TestFuturesMarketPriceWebsocketClient_handleMessage(t *testing.T) {
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
		{"TestFuturesMarketPriceWebsocketClient", fields{WebsocketClient: binance.WebsocketClient{}}, args{msg: []byte("{\"code\":0,\"msg\":\"Unknown property\",\"id\":0}")}, model.WebsocketErrorResponse{
			Code:    0,
			Message: "Unknown property",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &FuturesMarketPriceWebsocketClient{
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

func TestNewUSDTFuturesMarketPriceWebsocketClient(t *testing.T) {
	type args struct {
		streams []string
	}
	tests := []struct {
		name string
		args args
		want *FuturesMarketPriceWebsocketClient
	}{
		{"TestNewUSDTFuturesMarketPriceWebsocketClient", args{streams: []string{"BTCUSDT@markPrice"}}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUSDTFuturesMarketPriceWebsocketClient(tt.args.streams...); reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUSDTFuturesMarketPriceWebsocketClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
