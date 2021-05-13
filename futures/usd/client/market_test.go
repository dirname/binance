package futuresclient

import (
	"github.com/dirname/binance"
	"github.com/dirname/binance/config"
	"github.com/dirname/binance/model"
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
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{symbol: "test"}, model.APIErrorResponse{}, false},
		{"TestMarketClient_Get24hTickerPriceChange", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{symbol: "test"}, nil, true},
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
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{
			symbol:    "test",
			limit:     2000,
			formID:    2000,
			startTime: 2000,
			endTime:   2000,
		}, model.APIErrorResponse{
			Code:    -1121,
			Message: "Invalid symbol.",
		}, false},
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
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
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

func TestMarketClient_GetCompositeIndexSymbol(t *testing.T) {
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
		{"TestMarketClient_GetCompositeIndexSymbol", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{symbol: ""}, nil, true},
		{"TestMarketClient_GetCompositeIndexSymbol", fields{
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{symbol: "test"}, nil, false},
		{"TestMarketClient_GetCompositeIndexSymbol", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{symbol: "test"}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetCompositeIndexSymbol(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCompositeIndexSymbol() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCompositeIndexSymbol() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetContractCandlestick(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	type args struct {
		symbol       string
		contractType string
		interval     string
		limit        int32
		startTime    int64
		endTime      int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestMarketClient_GetFundingRateHistory", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:       "",
			contractType: "",
			interval:     "",
			startTime:    0,
			endTime:      0,
			limit:        0,
		}, []FundingRateResponse{}, true},
		{"TestMarketClient_GetFundingRateHistory", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:       "",
			contractType: "test",
			interval:     "",
			startTime:    0,
			endTime:      0,
			limit:        0,
		}, []FundingRateResponse{}, true},
		{"TestMarketClient_GetFundingRateHistory", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:       "",
			contractType: "",
			interval:     "test",
			startTime:    0,
			endTime:      0,
			limit:        0,
		}, []FundingRateResponse{}, true},
		{"TestMarketClient_GetFundingRateHistory", fields{
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{
			symbol:       "test",
			contractType: "test",
			interval:     "test",
			startTime:    10000,
			endTime:      10000,
			limit:        10000,
		}, model.APIErrorResponse{}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetContractCandlestick(tt.args.symbol, tt.args.contractType, tt.args.interval, tt.args.limit, tt.args.startTime, tt.args.endTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContractCandlestick() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetContractCandlestick() got = %v, want %v", got, tt.want)
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

func TestMarketClient_GetFundingRateHistory(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	type args struct {
		symbol    string
		startTime int64
		endTime   int64
		limit     int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestMarketClient_GetFundingRateHistory", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "",
			startTime: 0,
			endTime:   0,
			limit:     0,
		}, nil, true},
		{"TestMarketClient_GetFundingRateHistory", fields{
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{
			symbol:    "test",
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, nil, false},
		{"TestMarketClient_GetFundingRateHistory", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "test",
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetFundingRateHistory(tt.args.symbol, tt.args.startTime, tt.args.endTime, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFundingRateHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFundingRateHistory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetHistoricalBLVTNavCandlestick(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	type args struct {
		symbol    string
		period    string
		startTime int64
		endTime   int64
		limit     int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestMarketClient_GetFundingRateHistory", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "",
			period:    "",
			startTime: 0,
			endTime:   0,
			limit:     0,
		}, FundingRateResponse{}, true},
		{"TestMarketClient_GetFundingRateHistory", fields{
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{
			symbol:    "test",
			period:    "test",
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, model.APIErrorResponse{}, false},
		{"TestMarketClient_GetFundingRateHistory", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "test",
			period:    "test",
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetHistoricalBLVTNavCandlestick(tt.args.symbol, tt.args.period, tt.args.startTime, tt.args.endTime, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHistoricalBLVTNavCandlestick() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHistoricalBLVTNavCandlestick() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetIndexCandlestick(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	type args struct {
		pair      string
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
		{"TestMarketClient_GetIndexCandlestick", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			pair:      "",
			interval:  "",
			startTime: 0,
			endTime:   0,
			limit:     0,
		}, IndexCandlestickResponse{}, true},
		{"TestMarketClient_GetIndexCandlestick", fields{
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{
			pair:      "test",
			interval:  "test",
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, model.APIErrorResponse{}, false},
		{"TestMarketClient_GetIndexCandlestick", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			pair:      "test",
			interval:  "test",
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetIndexCandlestick(tt.args.pair, tt.args.interval, tt.args.limit, tt.args.startTime, tt.args.endTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIndexCandlestick() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIndexCandlestick() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetLongShortRatio(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	type args struct {
		symbol    string
		period    string
		startTime int64
		endTime   int64
		limit     int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestMarketClient_GetLongShortRatio", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "",
			period:    "",
			startTime: 0,
			endTime:   0,
			limit:     0,
		}, LongShortRatioResponse{}, true},
		{"TestMarketClient_GetLongShortRatio", fields{
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{
			symbol:    "test",
			period:    "test",
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, nil, false},
		{"TestMarketClient_GetLongShortRatio", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "test",
			period:    "test",
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetLongShortRatio(tt.args.symbol, tt.args.period, tt.args.startTime, tt.args.endTime, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLongShortRatio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLongShortRatio() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetMarketPrice(t *testing.T) {
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
		{"TestMarketClient_GetMarketPrice", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol: "",
		}, nil, true},
		{"TestMarketClient_GetMarketPrice", fields{
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{
			symbol: "test",
		}, nil, false},
		{"TestMarketClient_GetMarketPrice", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol: "test",
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetMarketPrice(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMarketPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMarketPrice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetMarketPriceCandlestick(t *testing.T) {
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
		{"TestMarketClient_GetMarketPriceCandlestick", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "",
			interval:  "",
			startTime: 0,
			endTime:   0,
			limit:     0,
		}, MarketPriceCandlestickResponse{}, true},
		{"TestMarketClient_GetMarketPriceCandlestick", fields{
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{
			symbol:    "test",
			interval:  "test",
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, nil, false},
		{"TestMarketClient_GetMarketPriceCandlestick", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "test",
			interval:  "test",
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetMarketPriceCandlestick(tt.args.symbol, tt.args.interval, tt.args.limit, tt.args.startTime, tt.args.endTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMarketPriceCandlestick() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMarketPriceCandlestick() got = %v, want %v", got, tt.want)
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
		{"TestMarketClient_GetOldTradeLookUp", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "",
			formID:    0,
			startTime: 0,
			endTime:   0,
			limit:     0,
		}, OlderTradeLookUpResponse{}, true},
		{"TestMarketClient_GetOldTradeLookUp", fields{
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{
			symbol:    "test",
			formID:    10000,
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, nil, false},
		{"TestMarketClient_GetOldTradeLookUp", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "test",
			formID:    10000,
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetOldTradeLookUp(tt.args.symbol, tt.args.limit, tt.args.formID, tt.args.startTime, tt.args.endTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOldTradeLookUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOldTradeLookUp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetOpenInterest(t *testing.T) {
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
		{"TestMarketClient_GetOpenInterest", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol: "",
		}, OpenInterestResponse{}, true},
		{"TestMarketClient_GetOpenInterest", fields{
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{
			symbol: "test",
		}, nil, false},
		{"TestMarketClient_GetOpenInterest", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol: "test",
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetOpenInterest(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOpenInterest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOpenInterest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetOpenInterestStatistics(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	type args struct {
		symbol    string
		period    string
		startTime int64
		endTime   int64
		limit     int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestMarketClient_GetOpenInterestStatistics", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "",
			period:    "",
			startTime: 0,
			endTime:   0,
			limit:     0,
		}, OpenInterestStatisticsResponse{}, true},
		{"TestMarketClient_GetOpenInterestStatistics", fields{
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{
			symbol:    "test",
			period:    "test",
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, nil, false},
		{"TestMarketClient_GetOpenInterestStatistics", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "test",
			period:    "test",
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetOpenInterestStatistics(tt.args.symbol, tt.args.period, tt.args.startTime, tt.args.endTime, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOpenInterestStatistics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOpenInterestStatistics() got = %v, want %v", got, tt.want)
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
		}, OrderBookResponse{}, true},
		{"TestMarketClient_GetOrderBook", fields{
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{
			symbol: "test",
			limit:  10000,
		}, nil, false},
		{"TestMarketClient_GetOrderBook", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol: "test",
			limit:  10000,
		}, nil, true},
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
			if reflect.DeepEqual(got, tt.want) {
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
		{"TestMarketClient_GetOrderBook", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol: "",
			limit:  0,
		}, OrderBookResponse{}, true},
		{"TestMarketClient_GetOrderBook", fields{
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{
			symbol: "test",
			limit:  10000,
		}, nil, false},
		{"TestMarketClient_GetOrderBook", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol: "test",
			limit:  10000,
		}, nil, true},
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
			if reflect.DeepEqual(got, tt.want) {
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
		}, args{
			symbol: "",
		}, nil, true},
		{"TestMarketClient_GetSymbolOrderBookTicker", fields{
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{
			symbol: "test",
		}, nil, false},
		{"TestMarketClient_GetSymbolOrderBookTicker", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol: "test",
		}, nil, true},
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
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSymbolOrderBookTicker() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetSymbolTicker(t *testing.T) {
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
		{"TestMarketClient_GetSymbolTicker", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{""}, []SymbolOrderBookTickerResponse{}, true},
		{"TestMarketClient_GetSymbolTicker", fields{
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{"test"}, model.APIErrorResponse{}, false},
		{"TestMarketClient_GetSymbolTicker", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{"test"}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetSymbolTicker(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSymbolTicker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSymbolTicker() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetTakerBuySellVolume(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	type args struct {
		symbol    string
		period    string
		startTime int64
		endTime   int64
		limit     int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestMarketClient_GetTakerBuySellVolume", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "",
			period:    "",
			startTime: 0,
			endTime:   0,
			limit:     0,
		}, TakerBuySellVolumeResponse{}, true},
		{"TestMarketClient_GetTakerBuySellVolume", fields{
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{
			symbol:    "test",
			period:    "test",
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, model.APIErrorResponse{}, false},
		{"TestMarketClient_GetTakerBuySellVolume", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "test",
			period:    "test",
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetTakerBuySellVolume(tt.args.symbol, tt.args.period, tt.args.startTime, tt.args.endTime, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTakerBuySellVolume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTakerBuySellVolume() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetTopTradeAccountsRatio(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	type args struct {
		symbol    string
		period    string
		startTime int64
		endTime   int64
		limit     int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestMarketClient_GetTopTradeAccountsRatio", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "",
			period:    "",
			startTime: 0,
			endTime:   0,
			limit:     0,
		}, nil, true},
		{"TestMarketClient_GetTopTradeAccountsRatio", fields{
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{
			symbol:    "test",
			period:    "test",
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, model.APIErrorResponse{Code: -1130, Message: "parameter 'period' is invalid."}, false},
		{"TestMarketClient_GetTopTradeAccountsRatio", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "test",
			period:    "test",
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, TopTraderAccountsRatioResponse{}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetTopTradeAccountsRatio(tt.args.symbol, tt.args.period, tt.args.startTime, tt.args.endTime, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTopTradeAccountsRatio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTopTradeAccountsRatio() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketClient_GetTopTradePositionsRatio(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
		AppKey  string
	}
	type args struct {
		symbol    string
		period    string
		startTime int64
		endTime   int64
		limit     int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestMarketClient_GetTopTradePositionsRatio", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "",
			period:    "",
			startTime: 0,
			endTime:   0,
			limit:     0,
		}, nil, true},
		{"TestMarketClient_GetTopTradePositionsRatio", fields{
			Builder: NewMarketClient(config.USDFuturesRestHost, "").Builder,
			AppKey:  "",
		}, args{
			symbol:    "test",
			period:    "test",
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, model.APIErrorResponse{Code: -1130, Message: "parameter 'period' is invalid."}, false},
		{"TestMarketClient_GetTopTradePositionsRatio", fields{
			Builder: &binance.PublicUrlBuilder{},
			AppKey:  "",
		}, args{
			symbol:    "test",
			period:    "test",
			startTime: 10000,
			endTime:   10000,
			limit:     10000,
		}, TopTraderPositionsRatioResponse{}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarketClient{
				Builder: tt.fields.Builder,
				AppKey:  tt.fields.AppKey,
			}
			got, err := m.GetTopTradePositionsRatio(tt.args.symbol, tt.args.period, tt.args.startTime, tt.args.endTime, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTopTradePositionsRatio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTopTradePositionsRatio() got = %v, want %v", got, tt.want)
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
