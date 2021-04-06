package usdt

import (
	"reflect"
	"testing"
)

func TestNewUSDTFuturesMarketPriceWebsocketClient(t *testing.T) {
	type args struct {
		streams []string
	}
	tests := []struct {
		name   string
		args   args
		refuse interface{}
	}{
		{"NewUSDTFuturesMarketPriceWebsocketClient", args{streams: []string{"btcusdt@arrTrade"}}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUSDTFuturesMarketPriceWebsocketClient(tt.args.streams...); reflect.DeepEqual(got, tt.refuse) {
				t.Errorf("NewUSDTFuturesMarketPriceWebsocketClient() = %v, refuse %v", got, tt.refuse)
			}
		})
	}
}