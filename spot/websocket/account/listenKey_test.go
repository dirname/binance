package spotclient

import (
	binance "github.com/dirname/Binance"
	"github.com/dirname/Binance/config"
	"reflect"
	"testing"
)

func TestNewListenKeyBuilder(t *testing.T) {
	type args struct {
		host      string
		appKey    string
		appSecret string
	}
	tests := []struct {
		name string
		args args
		want *listenKeyBuilder
	}{
		{"TestNewListenKeyBuilder", args{
			host:      "",
			appKey:    "",
			appSecret: "",
		}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListenKeyBuilder(tt.args.host, tt.args.appKey, tt.args.appSecret); reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListenKeyBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listenKeyBuilder_CreateIsolatedListenKey(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{"Test_listenKeyBuilder_CreateIsolatedListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, args{symbol: ""}, "", true},
		{"Test_listenKeyBuilder_CreateIsolatedListenKey", fields{Builder: NewListenKeyBuilder(config.SpotRestHost, "", "").Builder}, args{symbol: ""}, "", true},
		{"Test_listenKeyBuilder_CreateIsolatedListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, args{symbol: "test"}, "", true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &listenKeyBuilder{
				Builder: tt.fields.Builder,
			}
			got, err := b.CreateIsolatedListenKey(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateIsolatedListenKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateIsolatedListenKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listenKeyBuilder_CreateMarginListenKey(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"Test_listenKeyBuilder_CreateMarginListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, "", true},
		{"Test_listenKeyBuilder_CreateMarginListenKey", fields{Builder: NewListenKeyBuilder(config.SpotRestHost, "", "").Builder}, "API-key format invalid.", true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &listenKeyBuilder{
				Builder: tt.fields.Builder,
			}
			got, err := b.CreateMarginListenKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMarginListenKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateMarginListenKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listenKeyBuilder_CreateSpotListenKey(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"Test_listenKeyBuilder_CreateSpotListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, "", true},
		{"Test_listenKeyBuilder_CreateSpotListenKey", fields{Builder: NewListenKeyBuilder(config.SpotRestHost, "", "").Builder}, "API-key format invalid.", true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &listenKeyBuilder{
				Builder: tt.fields.Builder,
			}
			got, err := b.CreateSpotListenKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSpotListenKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateSpotListenKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listenKeyBuilder_DeleteIsolatedListenKey(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol string
		key    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"Test_listenKeyBuilder_DeleteIsolatedListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, args{
			symbol: "test",
			key:    "test",
		}, nil, true},
		{"Test_listenKeyBuilder_DeleteIsolatedListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, args{
			symbol: "",
			key:    "",
		}, nil, true},
		{"Test_listenKeyBuilder_DeleteIsolatedListenKey", fields{Builder: NewListenKeyBuilder(config.SpotRestHost, "", "").Builder}, args{
			symbol: "",
			key:    "",
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &listenKeyBuilder{
				Builder: tt.fields.Builder,
			}
			got, err := b.DeleteIsolatedListenKey(tt.args.symbol, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteIsolatedListenKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteIsolatedListenKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listenKeyBuilder_DeleteMarginListenKey(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"Test_listenKeyBuilder_DeleteMarginListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, args{
			key: "test",
		}, nil, true},
		{"Test_listenKeyBuilder_DeleteMarginListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, args{
			key: "",
		}, nil, true},
		{"Test_listenKeyBuilder_DeleteMarginListenKey", fields{Builder: NewListenKeyBuilder(config.SpotRestHost, "", "").Builder}, args{
			key: "",
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &listenKeyBuilder{
				Builder: tt.fields.Builder,
			}
			got, err := b.DeleteMarginListenKey(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteMarginListenKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteMarginListenKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listenKeyBuilder_DeleteSpotListenKey(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"Test_listenKeyBuilder_DeleteSpotListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, args{
			key: "test",
		}, nil, true},
		{"Test_listenKeyBuilder_DeleteSpotListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, args{
			key: "",
		}, nil, true},
		{"Test_listenKeyBuilder_DeleteSpotListenKey", fields{Builder: NewListenKeyBuilder(config.SpotRestHost, "", "").Builder}, args{
			key: "",
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &listenKeyBuilder{
				Builder: tt.fields.Builder,
			}
			got, err := b.DeleteSpotListenKey(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteSpotListenKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteSpotListenKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listenKeyBuilder_PingIsolatedListenKey(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol string
		key    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"Test_listenKeyBuilder_PingIsolatedListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, args{
			symbol: "test",
			key:    "test",
		}, nil, true},
		{"Test_listenKeyBuilder_PingIsolatedListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, args{
			symbol: "",
			key:    "",
		}, nil, true},
		{"Test_listenKeyBuilder_PingIsolatedListenKey", fields{Builder: NewListenKeyBuilder(config.SpotRestHost, "", "").Builder}, args{
			symbol: "",
			key:    "",
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &listenKeyBuilder{
				Builder: tt.fields.Builder,
			}
			got, err := b.PingIsolatedListenKey(tt.args.symbol, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("PingIsolatedListenKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PingIsolatedListenKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listenKeyBuilder_PingMarginListenKey(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"Test_listenKeyBuilder_PingMarginListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, args{
			key: "test",
		}, nil, true},
		{"Test_listenKeyBuilder_PingMarginListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, args{
			key: "",
		}, nil, true},
		{"Test_listenKeyBuilder_PingMarginListenKey", fields{Builder: NewListenKeyBuilder(config.SpotRestHost, "", "").Builder}, args{
			key: "",
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &listenKeyBuilder{
				Builder: tt.fields.Builder,
			}
			got, err := b.PingMarginListenKey(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("PingMarginListenKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PingMarginListenKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listenKeyBuilder_PingSpotListenKey(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"Test_listenKeyBuilder_PingSpotListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, args{
			key: "test",
		}, nil, true},
		{"Test_listenKeyBuilder_PingSpotListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, args{
			key: "",
		}, nil, true},
		{"Test_listenKeyBuilder_PingSpotListenKey", fields{Builder: NewListenKeyBuilder(config.SpotRestHost, "", "").Builder}, args{
			key: "",
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &listenKeyBuilder{
				Builder: tt.fields.Builder,
			}
			got, err := b.PingSpotListenKey(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("PingSpotListenKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PingSpotListenKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}
