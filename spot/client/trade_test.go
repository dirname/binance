package spotclient

import (
	binance "github.com/dirname/Binance"
	"github.com/dirname/Binance/spot/client/orderRespType"
	"github.com/dirname/Binance/spot/client/orderType"
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
	"time"
)

func TestNewOrder(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol           string
		side             string
		orderType        string
		timeInForce      string
		newClientOderID  string
		newOrderRespType string
		quantity         decimal.Decimal
		quoteOrderQTY    decimal.Decimal
		price            decimal.Decimal
		stopPrice        decimal.Decimal
		icebergQTY       decimal.Decimal
		recv             time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestNewOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:           "",
			side:             "",
			orderType:        "",
			timeInForce:      "",
			newClientOderID:  "",
			newOrderRespType: "",
			quantity:         decimal.Decimal{},
			quoteOrderQTY:    decimal.Decimal{},
			price:            decimal.Decimal{},
			stopPrice:        decimal.Decimal{},
			icebergQTY:       decimal.Decimal{},
			recv:             0,
		}, false, true},
		{"TestNewOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:           "test",
			side:             "test",
			orderType:        orderType.LimitMarker,
			timeInForce:      "test",
			newClientOderID:  "test",
			newOrderRespType: "test",
			quantity:         decimal.NewFromInt(100),
			quoteOrderQTY:    decimal.NewFromInt(100),
			price:            decimal.NewFromInt(100),
			stopPrice:        decimal.NewFromInt(100),
			icebergQTY:       decimal.NewFromInt(100),
			recv:             0,
		}, false, true},
		{"TestNewOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:           "test",
			side:             "test",
			orderType:        orderType.Limit,
			timeInForce:      "test",
			newClientOderID:  "test",
			newOrderRespType: "test",
			quantity:         decimal.NewFromInt(100),
			quoteOrderQTY:    decimal.NewFromInt(100),
			price:            decimal.NewFromInt(100),
			stopPrice:        decimal.NewFromInt(100),
			icebergQTY:       decimal.NewFromInt(100),
			recv:             0,
		}, false, true},
		{"TestNewOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:           "test",
			side:             "test",
			orderType:        orderType.Market,
			timeInForce:      "test",
			newClientOderID:  "test",
			newOrderRespType: "test",
			quantity:         decimal.NewFromInt(100),
			quoteOrderQTY:    decimal.NewFromInt(100),
			price:            decimal.NewFromInt(100),
			stopPrice:        decimal.NewFromInt(100),
			icebergQTY:       decimal.NewFromInt(100),
			recv:             0,
		}, false, true},
		{"TestNewOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:           "test",
			side:             "test",
			orderType:        orderType.StopLoss,
			timeInForce:      "test",
			newClientOderID:  "test",
			newOrderRespType: "test",
			quantity:         decimal.NewFromInt(100),
			quoteOrderQTY:    decimal.NewFromInt(100),
			price:            decimal.NewFromInt(100),
			stopPrice:        decimal.NewFromInt(100),
			icebergQTY:       decimal.NewFromInt(100),
			recv:             0,
		}, false, true},
		{"TestNewOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:           "test",
			side:             "test",
			orderType:        orderType.StopLossLimit,
			timeInForce:      "test",
			newClientOderID:  "test",
			newOrderRespType: "test",
			quantity:         decimal.NewFromInt(100),
			quoteOrderQTY:    decimal.NewFromInt(100),
			price:            decimal.NewFromInt(100),
			stopPrice:        decimal.NewFromInt(100),
			icebergQTY:       decimal.NewFromInt(100),
			recv:             0,
		}, false, true},
		{"TestNewOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:           "test",
			side:             "test",
			orderType:        orderType.TakeProfitLimit,
			timeInForce:      "test",
			newClientOderID:  "test",
			newOrderRespType: "test",
			quantity:         decimal.NewFromInt(100),
			quoteOrderQTY:    decimal.NewFromInt(100),
			price:            decimal.NewFromInt(100),
			stopPrice:        decimal.NewFromInt(100),
			icebergQTY:       decimal.NewFromInt(100),
			recv:             0,
		}, false, true},
		{"TestNewOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:           "test",
			side:             "test",
			orderType:        orderType.TakeProfit,
			timeInForce:      "test",
			newClientOderID:  "test",
			newOrderRespType: "test",
			quantity:         decimal.NewFromInt(100),
			quoteOrderQTY:    decimal.NewFromInt(100),
			price:            decimal.NewFromInt(100),
			stopPrice:        decimal.NewFromInt(100),
			icebergQTY:       decimal.NewFromInt(100),
			recv:             0,
		}, false, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.TestNewOrder(tt.args.symbol, tt.args.side, tt.args.orderType, tt.args.timeInForce, tt.args.newClientOderID, tt.args.newOrderRespType, tt.args.quantity, tt.args.quoteOrderQTY, tt.args.price, tt.args.stopPrice, tt.args.icebergQTY, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("TestNewOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t1.Errorf("TestNewOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTradeClient(t *testing.T) {
	type args struct {
		host      string
		appKey    string
		appSecret string
	}
	tests := []struct {
		name string
		args args
		want *TradeClient
	}{
		{"TestNewTradeClient", args{
			host:      "",
			appKey:    "",
			appSecret: "",
		}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTradeClient(tt.args.host, tt.args.appKey, tt.args.appSecret); reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTradeClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_DeleteOCO(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol            string
		listClientOrderID string
		newClientOrderID  string
		orderListID       int64
		recv              time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestTradeClient_DeleteOCO", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:            "",
			listClientOrderID: "",
			newClientOrderID:  "",
			orderListID:       0,
			recv:              0,
		}, nil, true},
		{"TestTradeClient_DeleteOCO", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:            "test",
			listClientOrderID: "test",
			newClientOrderID:  "",
			orderListID:       0,
			recv:              0,
		}, DeleteOCOResponse{}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.DeleteOCO(tt.args.symbol, tt.args.listClientOrderID, tt.args.newClientOrderID, tt.args.orderListID, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("DeleteOCO() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("DeleteOCO() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_DeleteOpenOrders(t1 *testing.T) {
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
		{"TestTradeClient_DeleteOpenOrders", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol: "",
			recv:   0,
		}, false, true},
		{"TestTradeClient_DeleteOpenOrders", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol: "test",
			recv:   0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.DeleteOpenOrders(tt.args.symbol, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("DeleteOpenOrders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t1.Errorf("DeleteOpenOrders() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_DeleteOrder(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol            string
		origClientOrderID string
		newClientOrderID  string
		orderID           int64
		recv              time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestTradeClient_DeleteOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:            "",
			origClientOrderID: "",
			newClientOrderID:  "",
			orderID:           0,
			recv:              0,
		}, true, true},
		{"TestTradeClient_DeleteOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:            "test",
			origClientOrderID: "test",
			newClientOrderID:  "test",
			orderID:           1,
			recv:              0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.DeleteOrder(tt.args.symbol, tt.args.origClientOrderID, tt.args.newClientOrderID, tt.args.orderID, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("DeleteOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t1.Errorf("DeleteOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_GetAccountInfo(t1 *testing.T) {
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
		{"TestTradeClient_GetAccountInfo", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{0}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.GetAccountInfo(tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetAccountInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetAccountInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_GetAccountTradeList(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol    string
		startTime int64
		endTime   int64
		formID    int64
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
		{"TestTradeClient_GetAccountTradeList", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:    "",
			startTime: 0,
			endTime:   0,
			formID:    0,
			limit:     0,
			recv:      0,
		}, nil, true},
		{"TestTradeClient_GetAccountTradeList", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:    "test",
			startTime: 10000,
			endTime:   10000,
			formID:    10000,
			limit:     10000,
			recv:      10000,
		}, AccountTradesListResponse{}, true},
		{"TestTradeClient_GetAccountTradeList", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:    "test",
			startTime: 10000,
			endTime:   10000,
			formID:    10000,
			limit:     -10000,
			recv:      10000,
		}, AccountTradesListResponse{}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.GetAccountTradeList(tt.args.symbol, tt.args.startTime, tt.args.endTime, tt.args.formID, tt.args.limit, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetAccountTradeList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetAccountTradeList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_GetAllOCOOrder(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		formID    int64
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
		{"TestTradeClient_GetAllOCOOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			formID:    0,
			startTime: 0,
			endTime:   0,
			limit:     0,
			recv:      0,
		}, GetAllOCOResponse{}, true},
		{"TestTradeClient_GetAllOCOOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			formID:    1000,
			startTime: 1000,
			endTime:   1000,
			limit:     1000,
			recv:      1000,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.GetAllOCOOrder(tt.args.formID, tt.args.startTime, tt.args.endTime, tt.args.limit, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetAllOCOOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetAllOCOOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_GetAllOrder(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol    string
		orderID   int64
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
		{"TestTradeClient_GetAllOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:    "",
			orderID:   0,
			startTime: 0,
			endTime:   0,
			limit:     0,
			recv:      0,
		}, false, true},
		{"TestTradeClient_GetAllOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:    "test",
			orderID:   1000,
			startTime: 1000,
			endTime:   1000,
			limit:     1000,
			recv:      1000,
		}, false, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.GetAllOrder(tt.args.symbol, tt.args.orderID, tt.args.startTime, tt.args.endTime, tt.args.limit, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetAllOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetAllOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_GetOCOOrder(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		origClientOrderID string
		orderListID       int64
		recv              time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestTradeClient_GetOCOOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			origClientOrderID: "",
			orderListID:       0,
			recv:              0,
		}, nil, true},
		{"TestTradeClient_GetOCOOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			origClientOrderID: "test",
			orderListID:       10,
			recv:              10,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.GetOCOOrder(tt.args.origClientOrderID, tt.args.orderListID, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetOCOOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetOCOOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_GetOpenOCOOrder(t1 *testing.T) {
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
		{"TestTradeClient_GetOpenOCOOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{0}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.GetOpenOCOOrder(tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetOpenOCOOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetOpenOCOOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_GetOpenOrder(t1 *testing.T) {
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
		{"TestTradeClient_GetOpenOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol: "",
			recv:   0,
		}, false, true},
		{"TestTradeClient_GetOpenOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol: "test",
			recv:   0,
		}, false, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.GetOpenOrder(tt.args.symbol, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetOpenOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetOpenOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_GetOrder(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol            string
		origClientOrderID string
		orderID           int64
		recv              time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestTradeClient_GetOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:            "",
			origClientOrderID: "",
			orderID:           0,
			recv:              0,
		}, false, true},
		{"TestTradeClient_GetOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:            "test",
			origClientOrderID: "test",
			orderID:           10,
			recv:              0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.GetOrder(tt.args.symbol, tt.args.origClientOrderID, tt.args.orderID, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_NewOCO(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol               string
		listClientOrderID    string
		side                 string
		limitClientOrderId   string
		stopClientOrderId    string
		stopLimitTimeInForce string
		newOrderRespType     string
		quantity             decimal.Decimal
		price                decimal.Decimal
		limitIcebergQty      decimal.Decimal
		stopPrice            decimal.Decimal
		stopLimitPrice       decimal.Decimal
		stopIcebergQty       decimal.Decimal
		recv                 time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestTradeClient_NewOCO", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:               "",
			listClientOrderID:    "",
			side:                 "",
			limitClientOrderId:   "",
			stopClientOrderId:    "",
			stopLimitTimeInForce: "",
			newOrderRespType:     "",
			quantity:             decimal.Decimal{},
			price:                decimal.Decimal{},
			limitIcebergQty:      decimal.Decimal{},
			stopPrice:            decimal.Decimal{},
			stopLimitPrice:       decimal.Decimal{},
			stopIcebergQty:       decimal.Decimal{},
			recv:                 0,
		}, false, true},
		{"TestTradeClient_NewOCO", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:               "test",
			listClientOrderID:    "test",
			side:                 "test",
			limitClientOrderId:   "test",
			stopClientOrderId:    "test",
			stopLimitTimeInForce: "test",
			newOrderRespType:     "test",
			quantity:             decimal.Decimal{},
			price:                decimal.Decimal{},
			limitIcebergQty:      decimal.Decimal{},
			stopPrice:            decimal.Decimal{},
			stopLimitPrice:       decimal.Decimal{},
			stopIcebergQty:       decimal.Decimal{},
			recv:                 0,
		}, false, true},
		{"TestTradeClient_NewOCO", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:               "test",
			listClientOrderID:    "test",
			side:                 "test",
			limitClientOrderId:   "test",
			stopClientOrderId:    "test",
			stopLimitTimeInForce: "test",
			newOrderRespType:     "test",
			quantity:             decimal.NewFromInt(10),
			price:                decimal.NewFromInt(10),
			limitIcebergQty:      decimal.NewFromInt(10),
			stopPrice:            decimal.Decimal{},
			stopLimitPrice:       decimal.NewFromInt(10),
			stopIcebergQty:       decimal.Decimal{},
			recv:                 0,
		}, false, true},
		{"TestTradeClient_NewOCO", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:               "test",
			listClientOrderID:    "test",
			side:                 "test",
			limitClientOrderId:   "test",
			stopClientOrderId:    "test",
			stopLimitTimeInForce: "",
			newOrderRespType:     "test",
			quantity:             decimal.Decimal{},
			price:                decimal.Decimal{},
			limitIcebergQty:      decimal.Decimal{},
			stopPrice:            decimal.Decimal{},
			stopLimitPrice:       decimal.NewFromInt(100),
			stopIcebergQty:       decimal.Decimal{},
			recv:                 0,
		}, false, true},
		{"TestTradeClient_NewOCO", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:               "test",
			listClientOrderID:    "test",
			side:                 "test",
			limitClientOrderId:   "test",
			stopClientOrderId:    "test",
			stopLimitTimeInForce: "test",
			newOrderRespType:     "test",
			quantity:             decimal.NewFromInt(100),
			price:                decimal.NewFromInt(100),
			limitIcebergQty:      decimal.NewFromInt(100),
			stopPrice:            decimal.NewFromInt(100),
			stopLimitPrice:       decimal.NewFromInt(100),
			stopIcebergQty:       decimal.NewFromInt(100),
			recv:                 0,
		}, false, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.NewOCO(tt.args.symbol, tt.args.listClientOrderID, tt.args.side, tt.args.limitClientOrderId, tt.args.stopClientOrderId, tt.args.stopLimitTimeInForce, tt.args.newOrderRespType, tt.args.quantity, tt.args.price, tt.args.limitIcebergQty, tt.args.stopPrice, tt.args.stopLimitPrice, tt.args.stopIcebergQty, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("NewOCO() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewOCO() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_NewOrder(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol           string
		side             string
		orderType        string
		timeInForce      string
		newClientOderID  string
		newOrderRespType string
		quantity         decimal.Decimal
		quoteOrderQTY    decimal.Decimal
		price            decimal.Decimal
		stopPrice        decimal.Decimal
		icebergQTY       decimal.Decimal
		recv             time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestNewOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:           "",
			side:             "",
			orderType:        "",
			timeInForce:      "",
			newClientOderID:  "",
			newOrderRespType: "",
			quantity:         decimal.Decimal{},
			quoteOrderQTY:    decimal.Decimal{},
			price:            decimal.Decimal{},
			stopPrice:        decimal.Decimal{},
			icebergQTY:       decimal.Decimal{},
			recv:             0,
		}, false, true},
		{"TestNewOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:           "test",
			side:             "test",
			orderType:        orderType.LimitMarker,
			timeInForce:      "test",
			newClientOderID:  "test",
			newOrderRespType: orderRespType.Result,
			quantity:         decimal.NewFromInt(100),
			quoteOrderQTY:    decimal.NewFromInt(100),
			price:            decimal.NewFromInt(100),
			stopPrice:        decimal.NewFromInt(100),
			icebergQTY:       decimal.NewFromInt(100),
			recv:             0,
		}, false, true},
		{"TestNewOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:           "test",
			side:             "test",
			orderType:        orderType.Limit,
			timeInForce:      "test",
			newClientOderID:  "test",
			newOrderRespType: orderRespType.ACK,
			quantity:         decimal.NewFromInt(100),
			quoteOrderQTY:    decimal.NewFromInt(100),
			price:            decimal.NewFromInt(100),
			stopPrice:        decimal.NewFromInt(100),
			icebergQTY:       decimal.NewFromInt(100),
			recv:             0,
		}, false, true},
		{"TestNewOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:           "test",
			side:             "test",
			orderType:        orderType.Market,
			timeInForce:      "test",
			newClientOderID:  "test",
			newOrderRespType: orderRespType.Full,
			quantity:         decimal.NewFromInt(100),
			quoteOrderQTY:    decimal.NewFromInt(100),
			price:            decimal.NewFromInt(100),
			stopPrice:        decimal.NewFromInt(100),
			icebergQTY:       decimal.NewFromInt(100),
			recv:             0,
		}, false, true},
		{"TestNewOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:           "test",
			side:             "test",
			orderType:        orderType.StopLoss,
			timeInForce:      "test",
			newClientOderID:  "test",
			newOrderRespType: "test",
			quantity:         decimal.NewFromInt(100),
			quoteOrderQTY:    decimal.NewFromInt(100),
			price:            decimal.NewFromInt(100),
			stopPrice:        decimal.NewFromInt(100),
			icebergQTY:       decimal.NewFromInt(100),
			recv:             0,
		}, false, true},
		{"TestNewOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:           "test",
			side:             "test",
			orderType:        orderType.StopLossLimit,
			timeInForce:      "test",
			newClientOderID:  "test",
			newOrderRespType: "test",
			quantity:         decimal.NewFromInt(100),
			quoteOrderQTY:    decimal.NewFromInt(100),
			price:            decimal.NewFromInt(100),
			stopPrice:        decimal.NewFromInt(100),
			icebergQTY:       decimal.NewFromInt(100),
			recv:             0,
		}, false, true},
		{"TestNewOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:           "test",
			side:             "test",
			orderType:        orderType.TakeProfitLimit,
			timeInForce:      "test",
			newClientOderID:  "test",
			newOrderRespType: "test",
			quantity:         decimal.NewFromInt(100),
			quoteOrderQTY:    decimal.NewFromInt(100),
			price:            decimal.NewFromInt(100),
			stopPrice:        decimal.NewFromInt(100),
			icebergQTY:       decimal.NewFromInt(100),
			recv:             0,
		}, false, true},
		{"TestNewOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:           "test",
			side:             "test",
			orderType:        orderType.TakeProfit,
			timeInForce:      "test",
			newClientOderID:  "test",
			newOrderRespType: "test",
			quantity:         decimal.NewFromInt(100),
			quoteOrderQTY:    decimal.NewFromInt(100),
			price:            decimal.NewFromInt(100),
			stopPrice:        decimal.NewFromInt(100),
			icebergQTY:       decimal.NewFromInt(100),
			recv:             0,
		}, false, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.NewOrder(tt.args.symbol, tt.args.side, tt.args.orderType, tt.args.timeInForce, tt.args.newClientOderID, tt.args.newOrderRespType, tt.args.quantity, tt.args.quoteOrderQTY, tt.args.price, tt.args.stopPrice, tt.args.icebergQTY, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("NewOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
