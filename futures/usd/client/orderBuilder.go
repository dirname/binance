package futuresclient

import (
	"errors"
	"fmt"
	"github.com/dirname/binance/futures/usd/client/define/order"
	"github.com/shopspring/decimal"
	"sort"
	"strings"
)

//OrderBatch batch order model
type OrderBatch []struct {
	Symbol           string          `json:"symbol"`
	Side             string          `json:"side"`
	PositionSide     string          `json:"positionSide,omitempty"`
	Type             string          `json:"type"`
	ReduceOnly       string          `json:"reduceOnly,omitempty"`
	Quantity         decimal.Decimal `json:"quantity"`
	Price            decimal.Decimal `json:"price,omitempty"`
	NewClientOrderID string          `json:"newClientOrderId,omitempty"`
	StopPrice        decimal.Decimal `json:"stopPrice,omitempty"`
	ActivationPrice  decimal.Decimal `json:"activationPrice,omitempty"`
	CallbackRate     decimal.Decimal `json:"callbackRate,omitempty"`
	TimeInForce      string          `json:"timeInForce,omitempty"`
	WorkingType      string          `json:"workingType,omitempty"`
	PriceProtect     string          `json:"priceProtect,omitempty"`
	NewOrderRespType string          `json:"newOrderRespType,omitempty"`
}

var mustPrice, mustStopPrice, mustQuantity, mustTimeInForce, mustCallbackRate []string

func init() {
	mustPrice = []string{order.Limit, order.Stop, order.TakeProfit}
	mustQuantity = []string{order.Limit, order.Market, order.Stop, order.TakeProfit}
	mustStopPrice = []string{order.Stop, order.TakeProfit, order.StopMarket, order.TakeProfitMarket}
	mustTimeInForce = []string{order.Limit}
	mustCallbackRate = []string{order.TrailingStopMarket}
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
	err = checker(ordersType, timeInForce, quantity, price, stopPrice, callbackRate)
	if err != nil {
		return "", err
	}
	if (closePosition != "" && reduceOnly != "") || (closePosition != "" && !quantity.LessThanOrEqual(decimal.NewFromInt(0))) {
		err = errors.New(ReduceOnlyWithClosePosition)
		return "", err
	}
	if closePosition != "" && (ordersType != order.StopMarket && ordersType != order.TakeProfitMarket) {
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

func checker(ordersType, timeInForce string, quantity, price, stopPrice, callbackRate decimal.Decimal) error {
	var err error
	if timeInForce == "" && isInArray(ordersType, mustTimeInForce) {
		err = errors.New(ordersType + ". " + TimeInForce)
		return err
	}
	if callbackRate.LessThanOrEqual(decimal.NewFromInt(0)) && isInArray(ordersType, mustCallbackRate) {
		err = errors.New(ordersType + ". " + CallBackRateInvalid)
		return err
	}
	if price.LessThanOrEqual(decimal.NewFromInt(0)) && isInArray(ordersType, mustPrice) {
		err = errors.New(ordersType + ". " + PriceInvalid)
		return err
	}
	if stopPrice.LessThanOrEqual(decimal.NewFromInt(0)) && isInArray(ordersType, mustStopPrice) {
		err = errors.New(ordersType + ". " + PriceInvalid)
		return err
	}
	if quantity.LessThanOrEqual(decimal.NewFromInt(0)) && isInArray(ordersType, mustQuantity) {
		err = errors.New(ordersType + ". " + QuantityInvalid)
		return err
	}
	return nil
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
