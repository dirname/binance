package account

import (
	"github.com/dirname/binance"
	"github.com/dirname/binance/config"
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

func Test_listenKeyBuilder_CreateListenKey(t *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"Test_listenKeyBuilder_CreateListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, "", true},
		{"Test_listenKeyBuilder_CreateListenKey", fields{Builder: NewListenKeyBuilder(config.USDFuturesRestHost, "", "").Builder}, "API-key format invalid.", true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &listenKeyBuilder{
				Builder: tt.fields.Builder,
			}
			got, err := b.CreateListenKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateListenKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateListenKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listenKeyBuilder_DeleteListenKey(t *testing.T) {
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
		{"Test_listenKeyBuilder_DeleteListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, args{
			key: "test",
		}, nil, true},
		{"Test_listenKeyBuilder_DeleteListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, args{
			key: "",
		}, nil, true},
		{"Test_listenKeyBuilder_DeleteListenKey", fields{Builder: NewListenKeyBuilder(config.USDFuturesRestHost, "", "").Builder}, args{
			key: "test",
		}, "API-key format invalid.", false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &listenKeyBuilder{
				Builder: tt.fields.Builder,
			}
			got, err := b.DeleteListenKey(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteListenKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteListenKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listenKeyBuilder_PingListenKey(t *testing.T) {
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
		{"Test_listenKeyBuilder_PingListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, args{
			key: "test",
		}, nil, true},
		{"Test_listenKeyBuilder_PingListenKey", fields{Builder: NewListenKeyBuilder("", "", "").Builder}, args{
			key: "",
		}, nil, true},
		{"Test_listenKeyBuilder_PingListenKey", fields{Builder: NewListenKeyBuilder(config.USDFuturesRestHost, "", "").Builder}, args{
			key: "test",
		}, "API-key format invalid.", false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &listenKeyBuilder{
				Builder: tt.fields.Builder,
			}
			got, err := b.PingListenKey(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("PingListenKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PingListenKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}
