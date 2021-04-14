package spotclient

import (
	"errors"
	"fmt"
	"github.com/dirname/Binance/spot/client/orderType"
	"github.com/shopspring/decimal"
	"sort"
	"strings"
)

var mustTimeInForce, mustPrice, mustStopPrice, mustQuantity []string

func init() {
	mustTimeInForce = []string{orderType.Limit, orderType.StopLossLimit, orderType.TakeProfitLimit}
	mustPrice = []string{orderType.Limit, orderType.StopLossLimit, orderType.TakeProfitLimit, orderType.LimitMarker}
	mustStopPrice = []string{orderType.StopLoss, orderType.StopLossLimit, orderType.TakeProfit, orderType.TakeProfitLimit}
	mustQuantity = []string{orderType.Limit, orderType.StopLossLimit, orderType.StopLoss, orderType.TakeProfitLimit, orderType.TakeProfit, orderType.LimitMarker}
}

func isInArray(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

func buildOrder(symbol, side, ordersType, timeInForce, newClientOderID, newOrderRespType string, quantity, quoteOrderQTY, price, stopPrice, icebergQTY decimal.Decimal) (string, error) {
	var err error
	if symbol == "" || side == "" || ordersType == "" {
		err = errors.New(MissParameters)
		return "", err
	}
	params := strings.Join([]string{"symbol=", symbol, "&side=", side, "&type=", ordersType}, "")
	if timeInForce == "" && isInArray(ordersType, mustTimeInForce) {
		err = errors.New(ordersType + ". " + TimeInForce)
		return "", err
	}
	if price.LessThanOrEqual(decimal.NewFromInt(0)) && isInArray(ordersType, mustPrice) {
		err = errors.New(ordersType + ". " + PriceInvalid)
		return "", err
	}
	if stopPrice.LessThanOrEqual(decimal.NewFromInt(0)) && isInArray(ordersType, mustStopPrice) {
		err = errors.New(ordersType + ". " + PriceInvalid)
		return "", err
	}
	if quantity.LessThanOrEqual(decimal.NewFromInt(0)) && isInArray(ordersType, mustQuantity) {
		err = errors.New(ordersType + ". " + QuantityInvalid)
		return "", err
	}
	if ordersType == orderType.Market && quantity.LessThanOrEqual(decimal.NewFromInt(0)) && quoteOrderQTY.LessThanOrEqual(decimal.NewFromInt(0)) {
		err = errors.New(ordersType + ". " + QuantityOrQuoteOrderQtyInvalid)
	}
	s := []string{buildTimeInForce(timeInForce), buildNewClientOderID(newClientOderID), buildNewOrderRespType(newOrderRespType), buildQuantity(quantity), buildQuoteOrderQty(quoteOrderQTY), buildPrice(price), buildStopPrice(stopPrice), buildIcebergQty(icebergQTY)}
	params += strings.Join(s, "")
	return params, err
}

func buildNewClientOderID(newClientOderID string) string {
	if newClientOderID != "" {
		return fmt.Sprintf("&newClientOrderId=%s", newClientOderID)
	}
	return ""
}

func buildTimeInForce(timeInForce string) string {
	if timeInForce != "" {
		return fmt.Sprintf("&timeInForce=%s", timeInForce)
	}
	return ""
}

func buildQuantity(quantity decimal.Decimal) string {
	if quantity.LessThanOrEqual(decimal.NewFromInt(0)) {
		return ""
	}
	return fmt.Sprintf("&quantity=%s", quantity)
}

func buildQuoteOrderQty(qty decimal.Decimal) string {
	if qty.LessThanOrEqual(decimal.NewFromInt(0)) {
		return ""
	}
	return fmt.Sprintf("&quoteOrderQty=%s", qty)
}

func buildPrice(price decimal.Decimal) string {
	if price.LessThanOrEqual(decimal.NewFromInt(0)) {
		return ""
	}
	return fmt.Sprintf("&price=%s", price)
}

func buildStopPrice(stopPrice decimal.Decimal) string {
	if stopPrice.LessThanOrEqual(decimal.NewFromInt(0)) {
		return ""
	}
	return fmt.Sprintf("&stopPrice=%s", stopPrice)
}

func buildIcebergQty(icebergQty decimal.Decimal) string {
	if icebergQty.LessThanOrEqual(decimal.NewFromInt(0)) {
		return ""
	}
	return fmt.Sprintf("&icebergQty=%s", icebergQty)
}

func buildNewOrderRespType(newOrderRespType string) string {
	if newOrderRespType != "" {
		return fmt.Sprintf("&newOrderRespType=%s", newOrderRespType)
	}
	return ""
}

func buildStringParam(paramName, param string) string {
	if param != "" {
		return fmt.Sprintf("&%s=%s", paramName, param)
	}
	return ""
}

func buildDecimalParam(paramName string, param decimal.Decimal) string {
	if param.LessThanOrEqual(decimal.NewFromInt(0)) {
		return ""
	}
	return fmt.Sprintf("&%s=%s", paramName, param)
}

func buildInt64Param(paramName string, param int64) string {
	if param > 0 {
		return fmt.Sprintf("&%s=%d", paramName, param)
	}
	return ""
}
