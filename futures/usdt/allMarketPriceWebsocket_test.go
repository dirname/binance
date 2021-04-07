package futuresusdt

import (
	binance "github.com/dirname/Binance"
	"github.com/dirname/Binance/model"
	"reflect"
	"testing"
)

func TestFuturesAllMarketPriceWebsocketClient_GetCombined(t *testing.T) {
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
		{"TestFuturesAllMarketPriceWebsocketClient", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			b:  false,
			id: 0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &FuturesAllMarketPriceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
		})
	}
}

func TestFuturesAllMarketPriceWebsocketClient_GetSubscribe(t *testing.T) {
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
		{"TestFuturesAllMarketPriceWebsocketClient", fields{WebsocketClient: binance.WebsocketClient{}}, args{id: 0}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &FuturesAllMarketPriceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
		})
	}
}

func TestFuturesAllMarketPriceWebsocketClient_SetCombined(t *testing.T) {
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
		{"TestFuturesAllMarketPriceWebsocketClient", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			b:  false,
			id: 0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &FuturesAllMarketPriceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
		})
	}
}

func TestFuturesAllMarketPriceWebsocketClient_SetHandler(t *testing.T) {
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
		{"TestFuturesAllMarketPriceWebsocketClient", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			connectHandler:  nil,
			responseHandler: nil,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &FuturesAllMarketPriceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
		})
	}
}

func TestFuturesAllMarketPriceWebsocketClient_Subscribe(t *testing.T) {
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
		{"TestFuturesAllMarketPriceWebsocketClient", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			id:     0,
			params: nil,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &FuturesAllMarketPriceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
		})
	}
}

func TestFuturesAllMarketPriceWebsocketClient_Unsubscribe(t *testing.T) {
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
		{"TestFuturesAllMarketPriceWebsocketClient", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			id:     0,
			params: nil,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &FuturesAllMarketPriceWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
		})
	}
}

func TestFuturesAllMarketPriceWebsocketClient_handleMessage(t *testing.T) {
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
		{"TestFuturesAllMarketPriceWebsocketClient", fields{WebsocketClient: binance.WebsocketClient{}}, args{msg: []byte("{\"code\":0,\"msg\":\"Unknown property\",\"id\":0}")}, model.WebsocketErrorResponse{
			Code:    0,
			Message: "Unknown property",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &FuturesAllMarketPriceWebsocketClient{
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

func TestNewUSDTFuturesAllMarketPriceWebsocketClient(t *testing.T) {
	type args struct {
		streams []string
	}
	tests := []struct {
		name string
		args args
		want *FuturesAllMarketPriceWebsocketClient
	}{
		{"TestNewUSDTFuturesAllMarketPriceWebsocketClient", args{streams: []string{"BTCUSDT@markPrice"}}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUSDTFuturesAllMarketPriceWebsocketClient(tt.args.streams...); reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUSDTFuturesAllMarketPriceWebsocketClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
