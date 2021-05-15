package futuresclient

// AllBookTickerResponse json parser
type AllBookTickerResponse struct {
	EventType       string `json:"e"` // Event type
	OrderUpdateID   int64  `json:"u"` // Order book updateId
	EventTime       int64  `json:"E"` // Event time stamp
	TransactionTime int64  `json:"T"` // Transaction time stamp
	Symbol          string `json:"s"` // Symbol
	BestPrice       string `json:"b"` // Best bid price
	BestQTY         string `json:"B"` // Best bid qty
	BestAskPrice    string `json:"a"` // Best ask price
	BestAskQTY      string `json:"A"` // Best ask qty
}

// AllBookTickerCombinedResponse json parser
type AllBookTickerCombinedResponse struct {
	StreamName string                `json:"stream"` // Stream Name
	Data       AllBookTickerResponse `json:"data"`   // data
}
