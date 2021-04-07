package futuresusdt

// SymbolMiniTickerResponse json parser
type SymbolMiniTickerResponse struct {
	EventType   string `json:"e"` // Event type
	EventTime   int64  `json:"E"` // Event time stamp
	Symbol      string `json:"s"` // Symbol
	ClosePrice  string `json:"c"` // Close price
	OpenPrice   string `json:"o"` // Open price
	HighPrice   string `json:"h"` // High price
	LowPrice    string `json:"l"` // Low price
	BaseVolume  string `json:"v"` // Total traded base asset volume
	QuoteVolume string `json:"q"` // Total traded quote asset volume
}

// SymbolMiniTickerCombinedResponse json parser
type SymbolMiniTickerCombinedResponse struct {
	StreamName string                   `json:"stream"` // Stream Name
	Data       SymbolMiniTickerResponse `json:"data"`   // data
}
