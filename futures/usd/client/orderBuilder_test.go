package futuresclient

import (
	"github.com/dirname/binance/futures/usd/client/define/orderType"
	"github.com/shopspring/decimal"
	"testing"
)

func Test_buildActivationPrice(t *testing.T) {
	type args struct {
		activationPrice decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_buildActivationPrice", args{decimal.NewFromInt(1)}, "&activationPrice=1"},
		{"Test_buildActivationPrice", args{decimal.Decimal{}}, ""},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildActivationPrice(tt.args.activationPrice); got != tt.want {
				t.Errorf("buildActivationPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildCallbackRate(t *testing.T) {
	type args struct {
		callbackRate decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_buildCallbackRate", args{decimal.NewFromInt(1)}, "&callbackRate=1"},
		{"Test_buildCallbackRate", args{decimal.Decimal{}}, ""},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildCallbackRate(tt.args.callbackRate); got != tt.want {
				t.Errorf("buildCallbackRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildClosePosition(t *testing.T) {
	type args struct {
		closePosition string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_buildClosePosition", args{"test"}, "&closePosition=test"},
		{"Test_buildClosePosition", args{""}, ""},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildClosePosition(tt.args.closePosition); got != tt.want {
				t.Errorf("buildClosePosition() = %v, want %v", got, tt.want)
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
		{"Test_buildNewClientOderID", args{"test"}, "&newClientOrderId=test"},
		{"Test_buildNewClientOderID", args{""}, ""},
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
		{"Test_buildNewOrderRespType", args{"test"}, "&newOrderRespType=test"},
		{"Test_buildNewOrderRespType", args{""}, ""},
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

func Test_buildPositionSide(t *testing.T) {
	type args struct {
		positionSide string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_buildPositionSide", args{"test"}, "&positionSide=test"},
		{"Test_buildPositionSide", args{""}, ""},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildPositionSide(tt.args.positionSide); got != tt.want {
				t.Errorf("buildPositionSide() = %v, want %v", got, tt.want)
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
		{"Test_buildPrice", args{decimal.NewFromInt(1)}, "&price=1"},
		{"Test_buildPrice", args{decimal.Decimal{}}, ""},
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

func Test_buildPriceProtect(t *testing.T) {
	type args struct {
		priceProtect string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_buildPriceProtect", args{"test"}, "&priceProtect=test"},
		{"Test_buildPriceProtect", args{""}, ""},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildPriceProtect(tt.args.priceProtect); got != tt.want {
				t.Errorf("buildPriceProtect() = %v, want %v", got, tt.want)
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
		{"Test_buildQuantity", args{decimal.NewFromInt(1)}, "&quantity=1"},
		{"Test_buildQuantity", args{decimal.Decimal{}}, ""},
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

func Test_buildReduceOnly(t *testing.T) {
	type args struct {
		reduceOnly string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_buildReduceOnly", args{"test"}, "&reduceOnly=test"},
		{"Test_buildReduceOnly", args{""}, ""},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildReduceOnly(tt.args.reduceOnly); got != tt.want {
				t.Errorf("buildReduceOnly() = %v, want %v", got, tt.want)
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
		{"Test_buildStopPrice", args{decimal.NewFromInt(1)}, "&stopPrice=1"},
		{"Test_buildStopPrice", args{decimal.Decimal{}}, ""},
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
		{"Test_buildTimeInForce", args{"test"}, "&timeInForce=test"},
		{"Test_buildTimeInForce", args{""}, ""},
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

func Test_buildWorkingType(t *testing.T) {
	type args struct {
		workingType string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_buildWorkingType", args{"test"}, "&workingType=test"},
		{"Test_buildWorkingType", args{""}, ""},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildWorkingType(tt.args.workingType); got != tt.want {
				t.Errorf("buildWorkingType() = %v, want %v", got, tt.want)
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
			strArray: mustPrice,
		}, false},
		{"Test_isInArray", args{
			target:   orderType.Limit,
			strArray: mustPrice,
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

func Test_buildOrder(t *testing.T) {
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
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Test_buildOrder", args{
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
		}, "", true},
		{"Test_buildOrder", args{
			symbol:           "test",
			side:             "test",
			positionSide:     "",
			ordersType:       "test",
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
		}, "symbol=test&side=test&type=test", false},
		{"Test_buildOrder", args{
			symbol:           "test",
			side:             "test",
			positionSide:     "",
			ordersType:       "test",
			reduceOnly:       "test",
			newClientOrderID: "",
			closePosition:    "test",
			timeInForce:      "",
			workingType:      "",
			priceProtect:     "",
			newOrderRespType: "",
			quantity:         decimal.Decimal{},
			price:            decimal.Decimal{},
			stopPrice:        decimal.Decimal{},
			activationPrice:  decimal.Decimal{},
			callbackRate:     decimal.Decimal{},
		}, "", true},
		{"Test_buildOrder", args{
			symbol:           "test",
			side:             "test",
			positionSide:     "",
			ordersType:       "test",
			reduceOnly:       "",
			newClientOrderID: "",
			closePosition:    "test",
			timeInForce:      "",
			workingType:      "",
			priceProtect:     "",
			newOrderRespType: "",
			quantity:         decimal.Decimal{},
			price:            decimal.Decimal{},
			stopPrice:        decimal.Decimal{},
			activationPrice:  decimal.Decimal{},
			callbackRate:     decimal.Decimal{},
		}, "", true},
		{"Test_buildOrder", args{
			symbol:           "test",
			side:             "test",
			positionSide:     "",
			ordersType:       orderType.Limit,
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
		}, "", true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildOrder(tt.args.symbol, tt.args.side, tt.args.positionSide, tt.args.ordersType, tt.args.reduceOnly, tt.args.newClientOrderID, tt.args.closePosition, tt.args.timeInForce, tt.args.workingType, tt.args.priceProtect, tt.args.newOrderRespType, tt.args.quantity, tt.args.price, tt.args.stopPrice, tt.args.activationPrice, tt.args.callbackRate)
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

func Test_checker(t *testing.T) {
	type args struct {
		ordersType   string
		timeInForce  string
		quantity     decimal.Decimal
		price        decimal.Decimal
		stopPrice    decimal.Decimal
		callbackRate decimal.Decimal
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test_checker", args{
			ordersType:   orderType.Limit,
			timeInForce:  "",
			quantity:     decimal.Decimal{},
			price:        decimal.Decimal{},
			stopPrice:    decimal.Decimal{},
			callbackRate: decimal.Decimal{},
		}, true},
		{"Test_checker", args{
			ordersType:   orderType.TrailingStopMarket,
			timeInForce:  "",
			quantity:     decimal.Decimal{},
			price:        decimal.Decimal{},
			stopPrice:    decimal.Decimal{},
			callbackRate: decimal.Decimal{},
		}, true},
		{"Test_checker", args{
			ordersType:   orderType.Limit,
			timeInForce:  "test",
			quantity:     decimal.Decimal{},
			price:        decimal.Decimal{},
			stopPrice:    decimal.Decimal{},
			callbackRate: decimal.Decimal{},
		}, true},
		{"Test_checker", args{
			ordersType:   orderType.Limit,
			timeInForce:  "test",
			quantity:     decimal.Decimal{},
			price:        decimal.NewFromInt(1),
			stopPrice:    decimal.Decimal{},
			callbackRate: decimal.Decimal{},
		}, true},
		{"Test_checker", args{
			ordersType:   orderType.Stop,
			timeInForce:  "test",
			quantity:     decimal.Decimal{},
			price:        decimal.NewFromInt(1),
			stopPrice:    decimal.Decimal{},
			callbackRate: decimal.Decimal{},
		}, true},
		{"Test_checker", args{
			ordersType:   orderType.Limit,
			timeInForce:  "test",
			quantity:     decimal.Decimal{},
			price:        decimal.NewFromInt(1),
			stopPrice:    decimal.NewFromInt(1),
			callbackRate: decimal.Decimal{},
		}, true},
		{"Test_checker", args{
			ordersType:   orderType.Limit,
			timeInForce:  "test",
			quantity:     decimal.NewFromInt(1),
			price:        decimal.NewFromInt(1),
			stopPrice:    decimal.NewFromInt(1),
			callbackRate: decimal.Decimal{},
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checker(tt.args.ordersType, tt.args.timeInForce, tt.args.quantity, tt.args.price, tt.args.stopPrice, tt.args.callbackRate); (err != nil) != tt.wantErr {
				t.Errorf("checker() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
