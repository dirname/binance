package spotclient

import (
	binance "github.com/dirname/binance"
	"github.com/dirname/binance/config"
	"reflect"
	"testing"
	"time"
)

func TestCommonClient_GetSAPISystemStatus(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{"TestCommonClient_GetSAPISystemStatus", fields{Builder: binance.NewPublicUrlBuilder(config.SpotRestHost)}, SystemStatusResponse{
			Status:  0,
			Message: "normal",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommonClient{
				Builder: tt.fields.Builder,
			}
			got, err := c.GetSAPISystemStatus()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSAPISystemStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSAPISystemStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommonClient_GetServerTime(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		{"TestCommonClient_GetServerTime", fields{Builder: binance.NewPublicUrlBuilder(config.SpotRestHost)}, time.Now().UnixNano() / 1e6, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommonClient{
				Builder: tt.fields.Builder,
			}
			got, err := c.GetServerTime()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetServerTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got >= tt.want+5000 {
				t.Errorf("GetServerTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommonClient_Ping(t *testing.T) {
	type fields struct {
		Builder *binance.PublicUrlBuilder
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{"TestCommonClient_Ping", fields{Builder: binance.NewPublicUrlBuilder(config.SpotRestHost)}, nil, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommonClient{
				Builder: tt.fields.Builder,
			}
			got, err := c.Ping()
			if (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ping() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCommonClient(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name string
		args args
		want *CommonClient
	}{
		{"TestNewCommonClient", args{host: config.SpotRestHost}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCommonClient(tt.args.host); reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommonClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
