package spotclient

// SymbolBookTickerResponse json parser
type SymbolBookTickerResponse struct {
	OrderUpdateID int64  `json:"u"` // Order book updateId
	Symbol        string `json:"s"` // Symbol
	BestPrice     string `json:"b"` // Best bid price
	BestQTY       string `json:"B"` // Best bid qty
	BestAskPrice  string `json:"a"` // Best ask price
	BestAskQTY    string `json:"A"` // Best ask qty
}

// SymbolBookTickerCombinedResponse json parser
type SymbolBookTickerCombinedResponse struct {
	StreamName string                   `json:"stream"` // Stream Name
	Data       SymbolBookTickerResponse `json:"data"`   // data
}
