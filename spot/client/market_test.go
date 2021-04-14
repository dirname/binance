package spotclient

import (
	binance "github.com/dirname/Binance"
	"reflect"
	"testing"
)

func TestMarketClient_Get24hTickerPriceChange(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	type args struct {
		symbol string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestMarketClient_Get24hTickerPriceChange", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{symbol: "123"}, nil, true},
		{"TestMarketClient_Get24hTickerPriceChange", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{symbol: ""}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.Get24hTickerPriceChange(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get24hTickerPriceChange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get24hTickerPriceChange() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetAggregateTrades(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	type args struct {
		symbol    string
		limit     int32
		formID    int64
		startTime int64
		endTime   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestMarketClient_GetAggregateTrades", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "",
			limit:     0,
			formID:    0,
			startTime: 0,
			endTime:   0,
		}, nil, true},
		{"TestMarketClient_GetAggregateTrades", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "test",
			limit:     2000,
			formID:    2000,
			startTime: 2000,
			endTime:   2000,
		}, AggregateTradeResponse{}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetAggregateTrades(tt.args.symbol, tt.args.limit, tt.args.formID, tt.args.startTime, tt.args.endTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAggregateTrades() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAggregateTrades() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetAveragePrice(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	type args struct {
		symbol string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestMarketClient_GetAveragePrice", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{""}, nil, true},
		{"TestMarketClient_GetAveragePrice", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{"test"}, CurrentAveragePriceResponse{
			Min:   0,
			Price: "",
		}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetAveragePrice(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAveragePrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAveragePrice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetCandlestick(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	type args struct {
		symbol    string
		interval  string
		limit     int32
		startTime int64
		endTime   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestMarketClient_GetCandlestick", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "test",
			interval:  "1min",
			limit:     11110,
			startTime: 11110,
			endTime:   11110,
		}, CandlestickResponse{}, true},
		{"TestMarketClient_GetCandlestick", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "test",
			interval:  "",
			limit:     0,
			startTime: 0,
			endTime:   0,
		}, nil, true},
		{"TestMarketClient_GetCandlestick", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "test",
			interval:  "",
			limit:     0,
			startTime: 0,
			endTime:   0,
		}, nil, true},
		{"TestMarketClient_GetCandlestick", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "",
			interval:  "",
			limit:     0,
			startTime: 0,
			endTime:   0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetCandlestick(tt.args.symbol, tt.args.interval, tt.args.limit, tt.args.startTime, tt.args.endTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCandlestick() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCandlestick() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetExchangeInfo(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{"TestMarketClient_GetExchangeInfo", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, ExchangeInfoResponse{}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetExchangeInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetExchangeInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExchangeInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetOldTradeLookUp(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	type args struct {
		symbol string
		limit  int32
		formID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestMarketClient_GetOldTradeLookUp", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol: "",
			limit:  0,
			formID: 0,
		}, nil, true},
		{"TestMarketClient_GetOldTradeLookUp", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol: "test",
			limit:  11110,
			formID: 11110,
		}, OlderTradeLookUpResponse{}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetOldTradeLookUp(tt.args.symbol, tt.args.limit, tt.args.formID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOldTradeLookUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOldTradeLookUp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetOrderBook(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	type args struct {
		symbol string
		limit  int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestMarketClient_GetOrderBook", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol: "",
			limit:  0,
		}, nil, true},
		{"TestMarketClient_GetOrderBook", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol: "test",
			limit:  1111110,
		}, OrderBookResponse{}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetOrderBook(tt.args.symbol, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrderBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrderBook() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetRecentTrades(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	type args struct {
		symbol string
		limit  int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestMarketClient_GetRecentTrades", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol: "",
			limit:  0,
		}, nil, true},
		{"TestMarketClient_GetRecentTrades", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol: "test",
			limit:  111111110,
		}, RecentTradesListResponse{}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetRecentTrades(tt.args.symbol, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRecentTrades() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRecentTrades() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetSymbolOrderBookTicker(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	type args struct {
		symbol string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestMarketClient_GetSymbolOrderBookTicker", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{"test"}, SymbolOrderBookTickerResponse{}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetSymbolOrderBookTicker(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSymbolOrderBookTicker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSymbolOrderBookTicker() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetSymbolTickerPrice(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	type args struct {
		symbol string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestMarketClient_GetSymbolTickerPrice", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{"test"}, SymbolPriceTickerResponse{}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetSymbolTickerPrice(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSymbolTickerPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSymbolTickerPrice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMarketClient(t *testing.T) {
	type args struct {
		host   string
		appKey string
	}
	tests := []struct {
		name string
		args args
		want *MarketClient
	}{
		{"TestNewMarketClient", args{
			host:   "",
			appKey: "",
		}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMarketClient(tt.args.host, tt.args.appKey); reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMarketClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
