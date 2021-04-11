package spotclient

import (
	"github.com/dirname/Binance/spot/client/orderRespType"
	"github.com/dirname/Binance/spot/client/orderSide"
	"github.com/dirname/Binance/spot/client/orderType"
	"github.com/shopspring/decimal"
	"testing"
)

func Test_buildIcebergQty(t *testing.T) {
	type args struct {
		icebergQty decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_buildIcebergQty", args{icebergQty: decimal.NewFromInt(0)}, ""},
		{"Test_buildIcebergQty", args{icebergQty: decimal.NewFromInt(1)}, "&icebergQty=1"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildIcebergQty(tt.args.icebergQty); got != tt.want {
				t.Errorf("buildIcebergQty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildNewClientOderID(t *testing.T) {
	type args struct {
		newClientOderID string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_buildNewClientOderID", args{newClientOderID: ""}, ""},
		{"Test_buildNewClientOderID", args{newClientOderID: "123"}, "&newClientOrderId=123"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildNewClientOderID(tt.args.newClientOderID); got != tt.want {
				t.Errorf("buildNewClientOderID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildNewOrderRespType(t *testing.T) {
	type args struct {
		newOrderRespType string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_buildNewOrderRespType", args{newOrderRespType: orderRespType.Result}, "&newOrderRespType=RESULT"},
		{"Test_buildNewOrderRespType", args{newOrderRespType: ""}, ""},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildNewOrderRespType(tt.args.newOrderRespType); got != tt.want {
				t.Errorf("buildNewOrderRespType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildOrder(t *testing.T) {
	type args struct {
		symbol           string
		side             string
		ordersType       string
		timeInForce      string
		newClientOderID  string
		newOrderRespType string
		quantity         decimal.Decimal
		quoteOrderQTY    decimal.Decimal
		price            decimal.Decimal
		stopPrice        decimal.Decimal
		icebergQTY       decimal.Decimal
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Test_buildOrder", args{
			symbol:           "test",
			side:             "sell",
			ordersType:       orderType.Limit,
			timeInForce:      "",
			newClientOderID:  "",
			newOrderRespType: "",
			quantity:         decimal.Decimal{},
			quoteOrderQTY:    decimal.Decimal{},
			price:            decimal.Decimal{},
			stopPrice:        decimal.Decimal{},
			icebergQTY:       decimal.Decimal{},
		}, "", true},
		{"Test_buildOrder", args{
			symbol:           "",
			side:             "sell",
			ordersType:       orderType.Limit,
			timeInForce:      "",
			newClientOderID:  "",
			newOrderRespType: "",
			quantity:         decimal.Decimal{},
			quoteOrderQTY:    decimal.Decimal{},
			price:            decimal.Decimal{},
			stopPrice:        decimal.Decimal{},
			icebergQTY:       decimal.Decimal{},
		}, "", true},
		{"Test_buildOrder", args{
			symbol:           "test",
			side:             orderSide.Buy,
			ordersType:       orderType.Limit,
			timeInForce:      "test",
			newClientOderID:  "",
			newOrderRespType: "",
			quantity:         decimal.NewFromFloat(0.123),
			quoteOrderQTY:    decimal.Decimal{},
			price:            decimal.NewFromFloat(0.7878),
			stopPrice:        decimal.Decimal{},
			icebergQTY:       decimal.Decimal{},
		}, "symbol=test&side=BUY&type=LIMIT&timeInForce=test&quantity=0.123&price=0.7878", false},
		{"Test_buildOrder", args{
			symbol:           "test",
			side:             "sell",
			ordersType:       orderType.Limit,
			timeInForce:      "",
			newClientOderID:  "",
			newOrderRespType: "",
			quantity:         decimal.Decimal{},
			quoteOrderQTY:    decimal.Decimal{},
			price:            decimal.Decimal{},
			stopPrice:        decimal.Decimal{},
			icebergQTY:       decimal.Decimal{},
		}, "", true},
		{"Test_buildOrder", args{
			symbol:           "test",
			side:             "sell",
			ordersType:       orderType.LimitMarker,
			timeInForce:      "",
			newClientOderID:  "",
			newOrderRespType: "",
			quantity:         decimal.Decimal{},
			quoteOrderQTY:    decimal.Decimal{},
			price:            decimal.NewFromFloat(0.000123),
			stopPrice:        decimal.Decimal{},
			icebergQTY:       decimal.Decimal{},
		}, "", true},
		{"Test_buildOrder", args{
			symbol:           "test",
			side:             "sell",
			ordersType:       orderType.LimitMarker,
			timeInForce:      "",
			newClientOderID:  "",
			newOrderRespType: "",
			quantity:         decimal.NewFromFloat(0.000123),
			quoteOrderQTY:    decimal.Decimal{},
			price:            decimal.Decimal{},
			stopPrice:        decimal.Decimal{},
			icebergQTY:       decimal.Decimal{},
		}, "", true},
		{"Test_buildOrder", args{
			symbol:           "test",
			side:             "sell",
			ordersType:       orderType.TakeProfit,
			timeInForce:      "",
			newClientOderID:  "",
			newOrderRespType: "",
			quantity:         decimal.NewFromFloat(0.000123),
			quoteOrderQTY:    decimal.Decimal{},
			price:            decimal.Decimal{},
			stopPrice:        decimal.Decimal{},
			icebergQTY:       decimal.Decimal{},
		}, "", true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildOrder(tt.args.symbol, tt.args.side, tt.args.ordersType, tt.args.timeInForce, tt.args.newClientOderID, tt.args.newOrderRespType, tt.args.quantity, tt.args.quoteOrderQTY, tt.args.price, tt.args.stopPrice, tt.args.icebergQTY)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("buildOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildPrice(t *testing.T) {
	type args struct {
		price decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_buildPrice", args{decimal.NewFromInt(0)}, ""},
		{"Test_buildPrice", args{decimal.NewFromInt(100)}, "&price=100"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildPrice(tt.args.price); got != tt.want {
				t.Errorf("buildPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildQuantity(t *testing.T) {
	type args struct {
		quantity decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_buildQuantity", args{decimal.NewFromInt(0)}, ""},
		{"Test_buildQuantity", args{decimal.NewFromInt(100)}, "&quantity=100"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildQuantity(tt.args.quantity); got != tt.want {
				t.Errorf("buildQuantity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildQuoteOrderQty(t *testing.T) {
	type args struct {
		qty decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_buildQuoteOrderQty", args{decimal.NewFromInt(0)}, ""},
		{"Test_buildQuoteOrderQty", args{decimal.NewFromFloat(0.123)}, "&quoteOrderQty=0.123"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildQuoteOrderQty(tt.args.qty); got != tt.want {
				t.Errorf("buildQuoteOrderQty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildStopPrice(t *testing.T) {
	type args struct {
		stopPrice decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_buildStopPrice", args{decimal.NewFromInt(0)}, ""},
		{"Test_buildStopPrice", args{decimal.NewFromInt(123)}, "&stopPrice=123"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildStopPrice(tt.args.stopPrice); got != tt.want {
				t.Errorf("buildStopPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildTimeInForce(t *testing.T) {
	type args struct {
		timeInForce string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_buildTimeInForce", args{""}, ""},
		{"Test_buildTimeInForce", args{"test"}, "&timeInForce=test"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTimeInForce(tt.args.timeInForce); got != tt.want {
				t.Errorf("buildTimeInForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isInArray(t *testing.T) {
	type args struct {
		target   string
		strArray []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test_isInArray", args{
			target:   orderType.Market,
			strArray: mustTimeInForce,
		}, false},
		{"Test_isInArray", args{
			target:   orderType.Limit,
			strArray: mustTimeInForce,
		}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInArray(tt.args.target, tt.args.strArray); got != tt.want {
				t.Errorf("isInArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildDecimalParam(t *testing.T) {
	type args struct {
		paramName string
		param     decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_buildDecimalParam", args{
			paramName: "",
			param:     decimal.NewFromInt(0),
		}, ""},
		{"Test_buildDecimalParam", args{
			paramName: "",
			param:     decimal.NewFromInt(1),
		}, "&=1"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildDecimalParam(tt.args.paramName, tt.args.param); got != tt.want {
				t.Errorf("buildDecimalParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildStringParam(t *testing.T) {
	type args struct {
		paramName string
		param     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_buildStringParam", args{
			paramName: "",
			param:     "",
		}, ""},
		{"Test_buildStringParam", args{
			paramName: "",
			param:     "1",
		}, "&=1"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildStringParam(tt.args.paramName, tt.args.param); got != tt.want {
				t.Errorf("buildStringParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildInt64Param(t *testing.T) {
	type args struct {
		paramName string
		param     int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_buildInt64Param", args{
			paramName: "",
			param:     0,
		}, ""},
		{"Test_buildInt64Param", args{
			paramName: "",
			param:     1,
		}, "&=1"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildInt64Param(tt.args.paramName, tt.args.param); got != tt.want {
				t.Errorf("buildInt64Param() = %v, want %v", got, tt.want)
			}
		})
	}
}