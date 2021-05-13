package spotclient

import (
	binance "github.com/dirname/binance"
	"github.com/dirname/binance/config"
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
	"time"
)

func TestNewWalletClient(t *testing.T) {
	type args struct {
		host      string
		appKey    string
		appSecret string
	}
	tests := []struct {
		name string
		args args
		want *WalletClient
	}{
		{"TestNewWalletClient", args{
			host:      "",
			appKey:    "",
			appSecret: "",
		}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWalletClient(tt.args.host, tt.args.appKey, tt.args.appSecret); reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWalletClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletClient_AccountDividendRecord(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		asset     string
		startTime int64
		endTime   int64
		limit     int
		recv      time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestWalletClient_AccountDividendRecord", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			asset:     "",
			startTime: 0,
			endTime:   0,
			limit:     0,
			recv:      0,
		}, nil, false},
		{"TestWalletClient_AccountDividendRecord", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			asset:     "test",
			startTime: 10,
			endTime:   10,
			limit:     10,
			recv:      10,
		}, nil, false},
		{"TestWalletClient_AccountDividendRecord", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			asset:     "test",
			startTime: 10,
			endTime:   10,
			limit:     1000,
			recv:      10,
		}, nil, false},
		{"TestWalletClient_AccountDividendRecord", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			asset:     "test",
			startTime: 10,
			endTime:   10,
			limit:     -10,
			recv:      10,
		}, nil, false},
		{"TestWalletClient_AccountDividendRecord", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			asset:     "test",
			startTime: 10,
			endTime:   10,
			limit:     -10,
			recv:      10,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletClient{
				Builder: tt.fields.Builder,
			}
			got, err := w.AccountDividendRecord(tt.args.asset, tt.args.startTime, tt.args.endTime, tt.args.limit, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountDividendRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountDividendRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletClient_DepositAddressNetwork(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		coin    string
		network string
		recv    time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestWalletClient_DepositAddressNetwork", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			coin:    "BTC",
			network: "",
			recv:    0,
		}, nil, false},
		{"TestWalletClient_DepositAddressNetwork", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			coin:    "BTC",
			network: "test",
			recv:    0,
		}, nil, false},
		{"TestWalletClient_DepositAddressNetwork", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			coin:    "",
			network: "",
			recv:    0,
		}, true, true},
		{"TestWalletClient_DepositAddressNetwork", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			coin:    "test",
			network: "test",
			recv:    0,
		}, true, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletClient{
				Builder: tt.fields.Builder,
			}
			got, err := w.DepositAddressNetwork(tt.args.coin, tt.args.network, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("DepositAddressNetwork() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("DepositAddressNetwork() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletClient_DepositHistoryNetwork(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		coin      string
		status    int32
		offset    int32
		limit     int32
		startTime int64
		endTime   int64
		recv      time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestWalletClient_DepositHistoryNetwork", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			coin:      "",
			status:    0,
			offset:    0,
			limit:     0,
			startTime: 0,
			endTime:   0,
			recv:      0,
		}, nil, false},
		{"TestWalletClient_DepositHistoryNetwork", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			coin:      "test",
			status:    10,
			offset:    10,
			limit:     10,
			startTime: 10,
			endTime:   10,
			recv:      10,
		}, nil, false},
		{"TestWalletClient_DepositHistoryNetwork", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			coin:      "test",
			status:    10,
			offset:    10,
			limit:     10,
			startTime: 10,
			endTime:   10,
			recv:      10,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletClient{
				Builder: tt.fields.Builder,
			}
			got, err := w.DepositHistoryNetwork(tt.args.coin, tt.args.status, tt.args.offset, tt.args.limit, tt.args.startTime, tt.args.endTime, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("DepositHistoryNetwork() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("DepositHistoryNetwork() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletClient_DustTransfer(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		asset []string
		recv  time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestWalletClient_DustTransfer", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			asset: nil,
			recv:  0,
		}, nil, false},
		{"TestWalletClient_DustTransfer", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			asset: []string{"test", "test"},
			recv:  0,
		}, nil, false},
		{"TestWalletClient_DustTransfer", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			asset: []string{"test", "test"},
			recv:  0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletClient{
				Builder: tt.fields.Builder,
			}
			got, err := w.DustTransfer(tt.args.asset, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("DustTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("DustTransfer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletClient_FastWithdrawSwitch(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		status bool
		recv   time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestWalletClient_FastWithdrawSwitch", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			status: false,
			recv:   0,
		}, nil, false},
		{"TestWalletClient_FastWithdrawSwitch", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			status: true,
			recv:   0,
		}, nil, false},
		{"TestWalletClient_FastWithdrawSwitch", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			status: true,
			recv:   0,
		}, false, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletClient{
				Builder: tt.fields.Builder,
			}
			got, err := w.FastWithdrawSwitch(tt.args.status, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("FastWithdrawSwitch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("FastWithdrawSwitch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletClient_GetAllCoinsInfo(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		recv time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestWalletClient_GetAllCoinsInfo", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{0}, nil, false},
		{"TestWalletClient_GetAllCoinsInfo", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{0}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletClient{
				Builder: tt.fields.Builder,
			}
			got, err := w.GetAllCoinsInfo(tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllCoinsInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllCoinsInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletClient_GetFuturesSnapshot(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		startTime int64
		endTime   int64
		limit     int32
		recv      time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestWalletClient_GetFuturesSnapshot", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			startTime: 0,
			endTime:   0,
			limit:     0,
			recv:      0,
		}, nil, false},
		{"TestWalletClient_GetFuturesSnapshot", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			startTime: 1230,
			endTime:   12230,
			limit:     110,
			recv:      0,
		}, nil, false},
		{"TestWalletClient_GetFuturesSnapshot", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			startTime: 1230,
			endTime:   12230,
			limit:     110,
			recv:      0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletClient{
				Builder: tt.fields.Builder,
			}
			got, err := w.GetFuturesSnapshot(tt.args.startTime, tt.args.endTime, tt.args.limit, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFuturesSnapshot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFuturesSnapshot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletClient_GetMarginSnapshot(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		startTime int64
		endTime   int64
		limit     int32
		recv      time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestWalletClient_GetMarginSnapshot", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			startTime: 0,
			endTime:   0,
			limit:     0,
			recv:      0,
		}, nil, false},
		{"TestWalletClient_GetMarginSnapshot", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			startTime: 123210,
			endTime:   12320,
			limit:     110,
			recv:      0,
		}, nil, false},
		{"TestWalletClient_GetMarginSnapshot", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			startTime: 123210,
			endTime:   12320,
			limit:     110,
			recv:      0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletClient{
				Builder: tt.fields.Builder,
			}
			got, err := w.GetMarginSnapshot(tt.args.startTime, tt.args.endTime, tt.args.limit, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMarginSnapshot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMarginSnapshot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletClient_GetSpotSnapshot(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		startTime int64
		endTime   int64
		limit     int32
		recv      time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestWalletClient_GetSpotSnapshot", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			startTime: 0,
			endTime:   0,
			limit:     0,
			recv:      0,
		}, nil, false},
		{"TestWalletClient_GetSpotSnapshot", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			startTime: 128093,
			endTime:   128039,
			limit:     0,
			recv:      0,
		}, nil, false},
		{"TestWalletClient_GetSpotSnapshot", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			startTime: 128093,
			endTime:   128039,
			limit:     50,
			recv:      0,
		}, nil, false},
		{"TestWalletClient_GetSpotSnapshot", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			startTime: 128093,
			endTime:   128039,
			limit:     50,
			recv:      0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletClient{
				Builder: tt.fields.Builder,
			}
			got, err := w.GetSpotSnapshot(tt.args.startTime, tt.args.endTime, tt.args.limit, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSpotSnapshot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSpotSnapshot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletClient_SAPIAccountAPIStatus(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		recv time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestWalletClient_SAPIAccountAPIStatus", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{0}, nil, false},
		{"TestWalletClient_SAPIAccountAPIStatus", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{0}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletClient{
				Builder: tt.fields.Builder,
			}
			got, err := w.SAPIAccountAPIStatus(tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("SAPIAccountAPIStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("SAPIAccountAPIStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletClient_SAPIAccountStatus(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		recv time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestWalletClient_SAPIAccountStatus", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{0}, nil, false},
		{"TestWalletClient_SAPIAccountStatus", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{0}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletClient{
				Builder: tt.fields.Builder,
			}
			got, err := w.SAPIAccountStatus(tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("SAPIAccountStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("SAPIAccountStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletClient_SAPIAssetDetail(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		recv time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestWalletClient_SAPIAssetDetail", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{0}, nil, false},
		{"TestWalletClient_SAPIAssetDetail", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{0}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletClient{
				Builder: tt.fields.Builder,
			}
			got, err := w.SAPIAssetDetail(tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("SAPIAssetDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("SAPIAssetDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletClient_SAPIDustLog(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		recv time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestWalletClient_SAPIDustLog", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{0}, nil, false},
		{"TestWalletClient_SAPIDustLog", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{0}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletClient{
				Builder: tt.fields.Builder,
			}
			got, err := w.SAPIDustLog(tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("SAPIDustLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("SAPIDustLog() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletClient_SAPITradeFee(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol string
		recv   time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestWalletClient_SAPITradeFee", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			symbol: "",
			recv:   0,
		}, nil, false},
		{"TestWalletClient_SAPITradeFee", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			symbol: "test",
			recv:   0,
		}, nil, false},
		{"TestWalletClient_SAPITradeFee", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol: "test",
			recv:   0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletClient{
				Builder: tt.fields.Builder,
			}
			got, err := w.SAPITradeFee(tt.args.symbol, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("SAPITradeFee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("SAPITradeFee() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletClient_SAPIWithdraw(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		coin               string
		clientID           string
		network            string
		address            string
		addressTag         string
		name               string
		amount             decimal.Decimal
		transactionFeeFlag bool
		recv               time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestWalletClient_SAPIWithdraw", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			coin:               "BTC",
			clientID:           "",
			network:            "",
			address:            "test",
			addressTag:         "",
			name:               "",
			amount:             decimal.NewFromInt(1),
			transactionFeeFlag: false,
			recv:               0,
		}, nil, false},
		{"TestWalletClient_SAPIWithdraw", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			coin:               "",
			clientID:           "",
			network:            "",
			address:            "test",
			addressTag:         "",
			name:               "",
			amount:             decimal.NewFromInt(1),
			transactionFeeFlag: false,
			recv:               0,
		}, true, true},
		{"TestWalletClient_SAPIWithdraw", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			coin:               "BTC",
			clientID:           "test",
			network:            "test",
			address:            "test",
			addressTag:         "test",
			name:               "test",
			amount:             decimal.NewFromInt(1),
			transactionFeeFlag: false,
			recv:               0,
		}, true, false},
		{"TestWalletClient_SAPIWithdraw", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			coin:               "BTC",
			clientID:           "test",
			network:            "test",
			address:            "",
			addressTag:         "test",
			name:               "test",
			amount:             decimal.NewFromInt(1),
			transactionFeeFlag: false,
			recv:               0,
		}, true, true},
		{"TestWalletClient_SAPIWithdraw", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			coin:               "BTC",
			clientID:           "test",
			network:            "test",
			address:            "test",
			addressTag:         "test",
			name:               "test",
			amount:             decimal.Decimal{},
			transactionFeeFlag: false,
			recv:               0,
		}, true, true},
		{"TestWalletClient_SAPIWithdraw", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			coin:               "BTC",
			clientID:           "test",
			network:            "test",
			address:            "test",
			addressTag:         "test",
			name:               "test",
			amount:             decimal.Decimal{},
			transactionFeeFlag: false,
			recv:               0,
		}, true, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletClient{
				Builder: tt.fields.Builder,
			}
			got, err := w.SAPIWithdraw(tt.args.coin, tt.args.clientID, tt.args.network, tt.args.address, tt.args.addressTag, tt.args.name, tt.args.amount, tt.args.transactionFeeFlag, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("SAPIWithdraw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("SAPIWithdraw() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletClient_UniversalTransfer(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		typeName string
		asset    string
		amount   string
		recv     time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestWalletClient_UniversalTransfer", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			typeName: "test",
			asset:    "test",
			amount:   "test",
			recv:     0,
		}, nil, false},
		{"TestWalletClient_UniversalTransfer", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			typeName: "",
			asset:    "",
			amount:   "",
			recv:     0,
		}, true, true},
		{"TestWalletClient_UniversalTransfer", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			typeName: "test",
			asset:    "test",
			amount:   "test",
			recv:     0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletClient{
				Builder: tt.fields.Builder,
			}
			got, err := w.UniversalTransfer(tt.args.typeName, tt.args.asset, tt.args.amount, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("UniversalTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniversalTransfer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletClient_UniversalTransferRecord(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		typeName  string
		startTime int64
		endTime   int64
		current   int32
		size      int32
		recv      time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestWalletClient_UniversalTransferRecord", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			typeName:  "test",
			startTime: 0,
			endTime:   0,
			current:   0,
			size:      0,
			recv:      0,
		}, nil, false},
		{"TestWalletClient_UniversalTransferRecord", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			typeName:  "",
			startTime: 10,
			endTime:   10,
			current:   10,
			size:      10,
			recv:      10,
		}, true, true},
		{"TestWalletClient_UniversalTransferRecord", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			typeName:  "test",
			startTime: 110,
			endTime:   110,
			current:   110,
			size:      110,
			recv:      110,
		}, nil, false},
		{"TestWalletClient_UniversalTransferRecord", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			typeName:  "test",
			startTime: 0,
			endTime:   0,
			current:   -10,
			size:      -50,
			recv:      0,
		}, nil, false},
		{"TestWalletClient_UniversalTransferRecord", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			typeName:  "test",
			startTime: 0,
			endTime:   0,
			current:   -10,
			size:      -50,
			recv:      0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletClient{
				Builder: tt.fields.Builder,
			}
			got, err := w.UniversalTransferRecord(tt.args.typeName, tt.args.startTime, tt.args.endTime, tt.args.current, tt.args.size, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("UniversalTransferRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniversalTransferRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletClient_WithdrawHistoryNetwork(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		coin      string
		status    int32
		offset    int32
		limit     int32
		startTime int64
		endTime   int64
		recv      time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestWalletClient_WithdrawHistoryNetwork", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			coin:      "",
			status:    0,
			offset:    0,
			limit:     0,
			startTime: 0,
			endTime:   0,
			recv:      0,
		}, nil, false},
		{"TestWalletClient_WithdrawHistoryNetwork", fields{binance.NewPrivateUrlBuilder(config.SpotRestHost, "", "")}, args{
			coin:      "test",
			status:    10,
			offset:    10,
			limit:     10,
			startTime: 10,
			endTime:   10,
			recv:      0,
		}, nil, false},
		{"TestWalletClient_WithdrawHistoryNetwork", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			coin:      "test",
			status:    10,
			offset:    10,
			limit:     10,
			startTime: 10,
			endTime:   10,
			recv:      0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletClient{
				Builder: tt.fields.Builder,
			}
			got, err := w.WithdrawHistoryNetwork(tt.args.coin, tt.args.status, tt.args.offset, tt.args.limit, tt.args.startTime, tt.args.endTime, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithdrawHistoryNetwork() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithdrawHistoryNetwork() got = %v, want %v", got, tt.want)
			}
		})
	}
}
