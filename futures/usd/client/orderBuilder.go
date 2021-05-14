package futuresclient

import "github.com/dirname/binance/spot/client/orderType"

var mustPrice, mustStopPrice, mustQuantity []string

func init() {
	mustPrice = []string{orderType.Limit, orderType.StopLossLimit, orderType.TakeProfitLimit, orderType.LimitMarker}
	mustStopPrice = []string{orderType.StopLoss, orderType.StopLossLimit, orderType.TakeProfit, orderType.TakeProfitLimit}
	mustQuantity = []string{orderType.Limit, orderType.StopLossLimit, orderType.StopLoss, orderType.TakeProfitLimit, orderType.TakeProfit, orderType.LimitMarker}
}
