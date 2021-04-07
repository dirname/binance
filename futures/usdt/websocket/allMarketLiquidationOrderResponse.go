package futuresusdt

// AllMarketLiquidationOrderResponse json parser
type AllMarketLiquidationOrderResponse struct {
	EventType string `json:"e"` // Event type
	EventTime int64  `json:"E"` // Event time stamp
	Order     struct {
		Symbol                    string `json:"s"`  // Symbol
		Side                      string `json:"S"`  // Side
		OrderType                 string `json:"o"`  // Order Type
		ForceTime                 string `json:"f"`  // Time in force
		OriginalQuantity          string `json:"q"`  // Original Quantity
		Price                     string `json:"p"`  // Price
		AveragePrice              string `json:"ap"` // Average price
		OrderStatus               string `json:"X"`  // Order status
		LastFilledQuantity        string `json:"l"`  // Order last filled quantity
		FilledAccumulatedQuantity string `json:"z"`  // Order filled accumulated quantity
		OrderTradeTime            int64  `json:"T"`  // Order time
	} `json:"o"` // Order
}

// AllMarketLiquidationOrderCombinedResponse json parser
type AllMarketLiquidationOrderCombinedResponse struct {
	StreamName string                            `json:"stream"` // Stream Name
	Data       AllMarketLiquidationOrderResponse `json:"data"`   // data
}
