package futuresclient

import (
	"github.com/dirname/binance"
	"github.com/dirname/binance/config"
	"github.com/dirname/binance/model"
	"github.com/dirname/binance/spot/client/define/order"
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
		positionSide     string
		ordersType       string
		reduceOnly       string
		newClientOrderID string
		closePosition    string
		timeInForce      string
		workingType      string
		priceProtect     string
		newOrderRespType string
		quantity         decimal.Decimal
		price            decimal.Decimal
		stopPrice        decimal.Decimal
		activationPrice  decimal.Decimal
		callbackRate     decimal.Decimal
		recv             time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestNewOrder", fields{Builder: binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:           "",
			side:             "",
			positionSide:     "",
			ordersType:       "",
			reduceOnly:       "",
			newClientOrderID: "",
			closePosition:    "",
			timeInForce:      "",
			workingType:      "",
			priceProtect:     "",
			newOrderRespType: "",
			quantity:         decimal.Decimal{},
			price:            decimal.Decimal{},
			stopPrice:        decimal.Decimal{},
			activationPrice:  decimal.Decimal{},
			callbackRate:     decimal.Decimal{},
			recv:             0,
		}, nil, true},
		{"TestNewOrder", fields{Builder: binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:           "test",
			side:             "test",
			positionSide:     "",
			ordersType:       order.Limit,
			reduceOnly:       "",
			newClientOrderID: "",
			closePosition:    "",
			timeInForce:      "",
			workingType:      "",
			priceProtect:     "",
			newOrderRespType: "",
			quantity:         decimal.Decimal{},
			price:            decimal.Decimal{},
			stopPrice:        decimal.Decimal{},
			activationPrice:  decimal.Decimal{},
			callbackRate:     decimal.Decimal{},
			recv:             0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.TestNewOrder(tt.args.symbol, tt.args.side, tt.args.positionSide, tt.args.ordersType, tt.args.reduceOnly, tt.args.newClientOrderID, tt.args.closePosition, tt.args.timeInForce, tt.args.workingType, tt.args.priceProtect, tt.args.newOrderRespType, tt.args.quantity, tt.args.price, tt.args.stopPrice, tt.args.activationPrice, tt.args.callbackRate, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("TestNewOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
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

func TestTradeClient_ADLQuantile(t1 *testing.T) {
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
		{"TestTradeClient_ADLQuantile", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol: "",
			recv:   0,
		}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.ADLQuantile(tt.args.symbol, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("ADLQuantile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("ADLQuantile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_BatchDeleteOrder(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol            string
		orderIDList       []int64
		origClientOrderID []string
		recv              time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestTradeClient_BatchDeleteOrder", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:            "",
			orderIDList:       nil,
			origClientOrderID: nil,
			recv:              0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.BatchDeleteOrder(tt.args.symbol, tt.args.orderIDList, tt.args.origClientOrderID, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("BatchDeleteOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("BatchDeleteOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_BatchNewOrder(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		batchOrders OrderBatch
		recv        time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestTradeClient_BatchNewOrder", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			batchOrders: nil,
			recv:        0,
		}, nil, true},
		{"TestTradeClient_BatchNewOrder", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			batchOrders: OrderBatch{},
			recv:        0,
		}, BatchNewOrderResponse{}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.BatchNewOrder(tt.args.batchOrders, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("BatchNewOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("BatchNewOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_ChangeMultiAssetsMode(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		margin bool
		recv   time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestTradeClient_ChangeMultiAssetsMode", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			margin: false,
			recv:   0,
		}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		{"TestTradeClient_ChangeMultiAssetsMode", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			margin: true,
			recv:   0,
		}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.ChangeMultiAssetsMode(tt.args.margin, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("ChangeMultiAssetsMode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("ChangeMultiAssetsMode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_CommissionRate(t1 *testing.T) {
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
		{"TestTradeClient_CommissionRate", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol: "",
			recv:   0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.CommissionRate(tt.args.symbol, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("CommissionRate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("CommissionRate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_CountDownDeleteAllOrder(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol        string
		countDownTime int64
		recv          time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestTradeClient_CountDownDeleteAllOrder", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:        "",
			countDownTime: 0,
			recv:          0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.CountDownDeleteAllOrder(tt.args.symbol, tt.args.countDownTime, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("CountDownDeleteAllOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("CountDownDeleteAllOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_DeleteAllOrder(t1 *testing.T) {
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
		{"TestTradeClient_DeleteAllOrder", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol: "",
			recv:   0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.DeleteAllOrder(tt.args.symbol, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("DeleteAllOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("DeleteAllOrder() got = %v, want %v", got, tt.want)
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
		{"TestTradeClient_DeleteOrder", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:            "",
			origClientOrderID: "",
			orderID:           0,
			recv:              0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.DeleteOrder(tt.args.symbol, tt.args.origClientOrderID, tt.args.orderID, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("DeleteOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("DeleteOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_ForceOrders(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol        string
		autoCloseType string
		startTime     int64
		endTime       int64
		limit         int32
		recv          time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestTradeClient_ForceOrders", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:        "",
			autoCloseType: "",
			startTime:     0,
			endTime:       0,
			limit:         0,
			recv:          0,
		}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.ForceOrders(tt.args.symbol, tt.args.autoCloseType, tt.args.startTime, tt.args.endTime, tt.args.limit, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("ForceOrders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("ForceOrders() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_GetAllOpenOrder(t1 *testing.T) {
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
		{"TestTradeClient_GetAllOpenOrder", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol: "",
			recv:   0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.GetAllOpenOrder(tt.args.symbol, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetAllOpenOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetAllOpenOrder() got = %v, want %v", got, tt.want)
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
		{"TestTradeClient_GetAllOpenOrder", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:    "",
			orderID:   0,
			startTime: 0,
			endTime:   0,
			limit:     0,
			recv:      0,
		}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
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
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetAllOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_GetBalance(t1 *testing.T) {
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
		{"TestTradeClient_GetBalance", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{0}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.GetBalance(tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetBalance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_GetMultiAssetsMargin(t1 *testing.T) {
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
		{"TestTradeClient_GetMultiAssetsMargin", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{0}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		{"TestTradeClient_GetMultiAssetsMargin", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{0}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.GetMultiAssetsMargin(tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetMultiAssetsMargin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetMultiAssetsMargin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_GetOpenOrder(t1 *testing.T) {
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
		{"TestTradeClient_GetOpenOrder", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:            "",
			origClientOrderID: "",
			orderID:           0,
			recv:              0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.GetOpenOrder(tt.args.symbol, tt.args.origClientOrderID, tt.args.orderID, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetOpenOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
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
		{"TestTradeClient_GetOpenOrder", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:            "",
			origClientOrderID: "",
			orderID:           0,
			recv:              0,
		}, nil, true},
		{"TestTradeClient_GetOrder", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:            "test",
			origClientOrderID: "",
			orderID:           0,
			recv:              0,
		}, nil, true},
		{"TestTradeClient_GetOrder", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:            "test",
			origClientOrderID: "test",
			orderID:           0,
			recv:              0,
		}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		{"TestTradeClient_GetOrder", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:            "test",
			origClientOrderID: "",
			orderID:           10,
			recv:              0,
		}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		{"TestTradeClient_GetOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:            "test",
			origClientOrderID: "",
			orderID:           10,
			recv:              0,
		}, OrderBookResponse{}, true},
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
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_GetPositionSideDual(t1 *testing.T) {
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
		{"TestTradeClient_GetPositionSideDual", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{0}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		{"TestTradeClient_GetPositionSideDual", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{0}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.GetPositionSideDual(tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetPositionSideDual() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetPositionSideDual() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_GetUserInfo(t1 *testing.T) {
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
		{"TestTradeClient_GetUserInfo", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{0}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.GetUserInfo(tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetUserInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_Incomes(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol     string
		incomeType string
		startTime  int64
		endTime    int64
		limit      int32
		recv       time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestTradeClient_Incomes", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:     "",
			incomeType: "",
			startTime:  0,
			endTime:    0,
			limit:      0,
			recv:       0,
		}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.Incomes(tt.args.symbol, tt.args.incomeType, tt.args.startTime, tt.args.endTime, tt.args.limit, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Incomes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Incomes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_IsolatedPositionHistory(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol    string
		sideType  int32
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
		{"TestTradeClient_IsolatedPositionHistory", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:    "",
			sideType:  0,
			limit:     0,
			startTime: 0,
			endTime:   0,
			recv:      0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.IsolatedPositionHistory(tt.args.symbol, tt.args.sideType, tt.args.limit, tt.args.startTime, tt.args.endTime, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("IsolatedPositionHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("IsolatedPositionHistory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_IsolatedPositionMargin(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol       string
		positionSide string
		amount       decimal.Decimal
		sideType     int32
		recv         time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestTradeClient_IsolatedPositionMargin", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:       "",
			positionSide: "",
			amount:       decimal.Decimal{},
			sideType:     0,
			recv:         0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.IsolatedPositionMargin(tt.args.symbol, tt.args.positionSide, tt.args.amount, tt.args.sideType, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("IsolatedPositionMargin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("IsolatedPositionMargin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_Leverage(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol   string
		leverage int32
		recv     time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestTradeClient_Leverage", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:   "",
			leverage: 0,
			recv:     0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.Leverage(tt.args.symbol, tt.args.leverage, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Leverage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Leverage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_LeverageBracket(t1 *testing.T) {
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
		{"TestTradeClient_LeverageBracket", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol: "",
			recv:   0,
		}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.LeverageBracket(tt.args.symbol, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("LeverageBracket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("LeverageBracket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_MarginType(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		symbol     string
		marginType string
		recv       time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestTradeClient_MarginType", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:     "",
			marginType: "",
			recv:       0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.MarginType(tt.args.symbol, tt.args.marginType, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("MarginType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("MarginType() got = %v, want %v", got, tt.want)
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
		positionSide     string
		ordersType       string
		reduceOnly       string
		newClientOrderID string
		closePosition    string
		timeInForce      string
		workingType      string
		priceProtect     string
		newOrderRespType string
		quantity         decimal.Decimal
		price            decimal.Decimal
		stopPrice        decimal.Decimal
		activationPrice  decimal.Decimal
		callbackRate     decimal.Decimal
		recv             time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestTradeClient_NewOrder", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:           "",
			side:             "",
			positionSide:     "",
			ordersType:       "",
			reduceOnly:       "",
			newClientOrderID: "",
			closePosition:    "",
			timeInForce:      "",
			workingType:      "",
			priceProtect:     "",
			newOrderRespType: "",
			quantity:         decimal.Decimal{},
			price:            decimal.Decimal{},
			stopPrice:        decimal.Decimal{},
			activationPrice:  decimal.Decimal{},
			callbackRate:     decimal.Decimal{},
			recv:             0,
		}, nil, true},
		{"TestTradeClient_NewOrder", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:           "test",
			side:             order.Buy,
			positionSide:     "",
			ordersType:       order.Limit,
			reduceOnly:       "",
			newClientOrderID: "",
			closePosition:    "",
			timeInForce:      "test",
			workingType:      "",
			priceProtect:     "",
			newOrderRespType: "",
			quantity:         decimal.NewFromInt(1),
			price:            decimal.NewFromInt(1),
			stopPrice:        decimal.Decimal{},
			activationPrice:  decimal.Decimal{},
			callbackRate:     decimal.Decimal{},
			recv:             0,
		}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		{"TestTradeClient_NewOrder", fields{binance.NewPrivateUrlBuilder("", "", "")}, args{
			symbol:           "test",
			side:             order.Buy,
			positionSide:     "",
			ordersType:       order.Limit,
			reduceOnly:       "",
			newClientOrderID: "",
			closePosition:    "",
			timeInForce:      "test",
			workingType:      "",
			priceProtect:     "",
			newOrderRespType: order.Result,
			quantity:         decimal.NewFromInt(1),
			price:            decimal.NewFromInt(1),
			stopPrice:        decimal.Decimal{},
			activationPrice:  decimal.Decimal{},
			callbackRate:     decimal.Decimal{},
			recv:             0,
		}, NewOrderResponseResult{}, true},
		{"TestTradeClient_NewOrder", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:           "test",
			side:             order.Buy,
			positionSide:     "",
			ordersType:       order.Limit,
			reduceOnly:       "",
			newClientOrderID: "",
			closePosition:    "",
			timeInForce:      "test",
			workingType:      "",
			priceProtect:     "",
			newOrderRespType: "",
			quantity:         decimal.NewFromInt(1),
			price:            decimal.NewFromInt(1),
			stopPrice:        decimal.Decimal{},
			activationPrice:  decimal.Decimal{},
			callbackRate:     decimal.Decimal{},
			recv:             0,
		}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.NewOrder(tt.args.symbol, tt.args.side, tt.args.positionSide, tt.args.ordersType, tt.args.reduceOnly, tt.args.newClientOrderID, tt.args.closePosition, tt.args.timeInForce, tt.args.workingType, tt.args.priceProtect, tt.args.newOrderRespType, tt.args.quantity, tt.args.price, tt.args.stopPrice, tt.args.activationPrice, tt.args.callbackRate, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("NewOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_PositionRisk(t1 *testing.T) {
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
		{"TestTradeClient_PositionRisk", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol: "",
			recv:   0,
		}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.PositionRisk(tt.args.symbol, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("PositionRisk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("PositionRisk() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_PositionSideDual(t1 *testing.T) {
	type fields struct {
		Builder *binance.PrivateUrlBuilder
	}
	type args struct {
		dualSidePosition bool
		recv             time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestTradeClient_PositionSideDual", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			dualSidePosition: false,
			recv:             0,
		}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		{"TestTradeClient_PositionSideDual", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			dualSidePosition: true,
			recv:             0,
		}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.PositionSideDual(tt.args.dualSidePosition, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("PositionSideDual() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("PositionSideDual() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_QuantitativeTradingRules(t1 *testing.T) {
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
		{"TestTradeClient_QuantitativeTradingRules", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol: "",
			recv:   0,
		}, model.APIErrorResponse{
			Code:    -2014,
			Message: "API-key format invalid.",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.QuantitativeTradingRules(tt.args.symbol, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("QuantitativeTradingRules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("QuantitativeTradingRules() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeClient_TradeLists(t1 *testing.T) {
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
		{"TestTradeClient_TradeLists", fields{binance.NewPrivateUrlBuilder(config.USDFuturesRestHost, "", "")}, args{
			symbol:    "",
			startTime: 0,
			endTime:   0,
			formID:    0,
			limit:     0,
			recv:      0,
		}, nil, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TradeClient{
				Builder: tt.fields.Builder,
			}
			got, err := t.TradeLists(tt.args.symbol, tt.args.startTime, tt.args.endTime, tt.args.formID, tt.args.limit, tt.args.recv)
			if (err != nil) != tt.wantErr {
				t1.Errorf("TradeLists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("TradeLists() got = %v, want %v", got, tt.want)
			}
		})
	}
}
