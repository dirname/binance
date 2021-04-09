package binance

import (
	"github.com/dirname/Binance/config"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestHttpRequest(t *testing.T) {
	type args struct {
		request *http.Request
	}
	req, _ := http.NewRequest(http.MethodGet, "https://"+config.SpotRestHost+"/api/v3/ping", nil)
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{"TestHttpRequest", args{request: req}, []byte("{}"), false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HttpRequest(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HttpRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPrivateUrlBuilder(t *testing.T) {
	type args struct {
		host      string
		appKey    string
		appSecret string
	}
	tests := []struct {
		name string
		args args
		want *PrivateUrlBuilder
	}{
		{"TestNewPrivateUrlBuilder", args{
			host:      "",
			appKey:    "",
			appSecret: "",
		}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPrivateUrlBuilder(tt.args.host, tt.args.appKey, tt.args.appSecret); reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPrivateUrlBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPublicUrlBuilder(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name string
		args args
		want *PublicUrlBuilder
	}{
		{"TestNewPublicUrlBuilder", args{host: ""}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPublicUrlBuilder(tt.args.host); reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPublicUrlBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrivateUrlBuilder_SetAPIKey(t *testing.T) {
	type fields struct {
		host      string
		appKey    string
		appSecret string
	}
	type args struct {
		appKey    string
		appSecret string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestPrivateUrlBuilder_SetAPIKey", fields{
			host:      "",
			appKey:    "",
			appSecret: "",
		}, args{
			appKey:    "",
			appSecret: "",
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PrivateUrlBuilder{
				host:      tt.fields.host,
				appKey:    tt.fields.appKey,
				appSecret: tt.fields.appSecret,
			}
			p.SetAPIKey(tt.args.appKey, tt.args.appSecret)
		})
	}
}

func TestPublicUrlBuilder_Build(t *testing.T) {
	type fields struct {
		host string
	}
	type args struct {
		method string
		path   string
		params string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Request
		wantErr bool
	}{
		{"TestPublicUrlBuilder_Build", fields{host: ""}, args{
			method: "",
			path:   "",
			params: "",
		}, nil, false},
		{
			"TestPublicUrlBuilder_Build", fields{host: ""}, args{
			method: "",
			path:   "",
			params: "test",
		}, nil, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PublicUrlBuilder{
				host: tt.fields.host,
			}
			got, err := p.Build(tt.args.method, tt.args.path, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSigner_Sum(t *testing.T) {
	type fields struct {
		key []byte
	}
	type args struct {
		queryString string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"TestSigner_Sum", fields{key: []byte("NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j")}, args{queryString: "symbol=LTCBTC&side=BUY&type=LIMIT&timeInForce=GTC&quantity=1&price=0.1&recvWindow=5000&timestamp=1499827319559"}, "c8db56825ae71d6d79447849e617115f4a920fa2acdcab2b053c4b2838bd6b71"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Signer{
				Key: tt.fields.key,
			}
			if got := s.Sum(tt.args.queryString); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrivateUrlBuilder_Build(t *testing.T) {
	type fields struct {
		host      string
		appKey    string
		appSecret string
		signer    *Signer
	}
	type args struct {
		method    string
		path      string
		params    string
		sign      bool
		timeStamp bool
		recv      time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Request
		wantErr bool
	}{
		{"TestPrivateUrlBuilder_Build", fields{
			host:      "",
			appKey:    "",
			appSecret: "",
			signer:    nil,
		}, args{
			method:    "",
			path:      "",
			params:    "",
			sign:      false,
			timeStamp: false,
			recv:      0,
		}, nil, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PrivateUrlBuilder{
				host:      tt.fields.host,
				appKey:    tt.fields.appKey,
				appSecret: tt.fields.appSecret,
				signer:    tt.fields.signer,
			}
			got, err := p.Build(tt.args.method, tt.args.path, tt.args.params, tt.args.sign, tt.args.timeStamp, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() got = %v, want %v", got, tt.want)
			}
		})
	}
}
