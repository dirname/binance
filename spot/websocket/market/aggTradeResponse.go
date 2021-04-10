package spotclient

// AggTradeResponse json parser
type AggTradeResponse struct {
	EventType           string `json:"e"` // Event type
	EventTime           int64  `json:"E"` // Event time stamp
	Symbol              string `json:"s"` // Symbol
	AggregateTradeID    int64  `json:"a"` // Aggregate trade ID
	Price               string `json:"p"` // Price
	Quantity            string `json:"q"` // Quantity
	FirstTradeID        int64  `json:"f"` // First trade ID
	LastTradeID         int64  `json:"l"` // Last trade ID
	TradeTime           int64  `json:"T"` // Trade time
	IsBuyerMarketMarker bool   `json:"m"` // Is the buyer the market maker
	UppercaseM          bool   `json:"M"` // Ignore
}

// AggTradeCombinedResponse json parser
type AggTradeCombinedResponse struct {
	StreamName string           `json:"stream"` // Stream Name
	Data       AggTradeResponse `json:"data"`   // data
}
