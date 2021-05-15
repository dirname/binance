package futuresclient

import (
	"errors"
	"fmt"
	"github.com/dirname/binance/futures/usd/client/orderType"
	"github.com/shopspring/decimal"
	"sort"
	"strings"
)

var mustPrice, mustStopPrice, mustQuantity, mustTimeInForce, mustCallbackRate []string

func init() {
	mustPrice = []string{orderType.Limit, orderType.Stop, orderType.TakeProfit}
	mustQuantity = []string{orderType.Limit, orderType.Market, orderType.Stop, orderType.TakeProfit}
	mustStopPrice = []string{orderType.Stop, orderType.TakeProfit, orderType.StopMarket, orderType.TakeProfitMarket}
	mustTimeInForce = []string{orderType.Limit}
	mustCallbackRate = []string{orderType.TrailingStopMarket}
}

func isInArray(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

func buildOrder(symbol, side, positionSide, ordersType, reduceOnly, newClientOrderID, closePosition, timeInForce, workingType, priceProtect, newOrderRespType string, quantity, price, stopPrice, activationPrice, callbackRate decimal.Decimal) (string, error) {
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
	if callbackRate.LessThanOrEqual(decimal.NewFromInt(0)) && isInArray(ordersType, mustCallbackRate) {
		err = errors.New(ordersType + ". " + CallBackRateInvalid)
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
	if (closePosition != "" && reduceOnly != "") || (closePosition != "" && !quantity.LessThanOrEqual(decimal.NewFromInt(0))) {
		err = errors.New(ReduceOnlyWithClosePosition)
		return "", err
	}
	if closePosition != "" && (ordersType != orderType.StopMarket && ordersType != orderType.TakeProfitMarket) {
		err = errors.New(OrderTypeWithClosePosition)
		return "", err
	}
	s := []string{
		buildPositionSide(positionSide), buildReduceOnly(reduceOnly), buildNewClientOderID(newClientOrderID),
		buildClosePosition(closePosition), buildTimeInForce(timeInForce), buildWorkingType(workingType), buildPriceProtect(priceProtect),
		buildNewOrderRespType(newOrderRespType), buildQuantity(quantity), buildPrice(price), buildStopPrice(stopPrice), buildActivationPrice(activationPrice),
		buildCallbackRate(callbackRate)}
	params += strings.Join(s, "")
	return params, err
}

func buildPositionSide(positionSide string) string {
	if positionSide != "" {
		return fmt.Sprintf("&positionSide=%s", positionSide)
	}
	return ""
}

func buildReduceOnly(reduceOnly string) string {
	if reduceOnly != "" {
		return fmt.Sprintf("&reduceOnly=%s", reduceOnly)
	}
	return ""
}

func buildNewClientOderID(newClientOderID string) string {
	if newClientOderID != "" {
		return fmt.Sprintf("&newClientOrderId=%s", newClientOderID)
	}
	return ""
}

func buildClosePosition(closePosition string) string {
	if closePosition != "" {
		return fmt.Sprintf("&closePosition=%s", closePosition)
	}
	return ""
}

func buildTimeInForce(timeInForce string) string {
	if timeInForce != "" {
		return fmt.Sprintf("&timeInForce=%s", timeInForce)
	}
	return ""
}

func buildWorkingType(workingType string) string {
	if workingType != "" {
		return fmt.Sprintf("&workingType=%s", workingType)
	}
	return ""
}

func buildPriceProtect(priceProtect string) string {
	if priceProtect != "" {
		return fmt.Sprintf("&priceProtect=%s", priceProtect)
	}
	return ""
}

func buildNewOrderRespType(newOrderRespType string) string {
	if newOrderRespType != "" {
		return fmt.Sprintf("&newOrderRespType=%s", newOrderRespType)
	}
	return ""
}

func buildQuantity(quantity decimal.Decimal) string {
	if quantity.LessThanOrEqual(decimal.NewFromInt(0)) {
		return ""
	}
	return fmt.Sprintf("&quantity=%s", quantity)
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

func buildActivationPrice(activationPrice decimal.Decimal) string {
	if activationPrice.LessThanOrEqual(decimal.NewFromInt(0)) {
		return ""
	}
	return fmt.Sprintf("&activationPrice=%s", activationPrice)
}

func buildCallbackRate(callbackRate decimal.Decimal) string {
	if callbackRate.LessThanOrEqual(decimal.NewFromInt(0)) {
		return ""
	}
	return fmt.Sprintf("&callbackRate=%s", callbackRate)
}
