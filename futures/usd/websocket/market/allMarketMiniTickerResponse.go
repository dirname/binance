package futuresclient

// AllMarketMiniTickerResponse json parser
type AllMarketMiniTickerResponse []struct {
	EventType   string `json:"e"` // Event type
	EventTime   int64  `json:"E"` // Event time stamp
	Symbol      string `json:"s"` // Symbol
	ClosePrice  string `json:"c"` // Close price
	OpenPrice   string `json:"o"` // Open price
	HighPrice   string `json:"h"` // Highest price
	LowPrice    string `json:"l"` // Lowest price
	BaseVolume  string `json:"v"` // Total traded base asset volume
	QuoteVolume string `json:"q"` // Total traded quote asset volume
}

// AllMarketMiniTickerCombinedResponse json parser
type AllMarketMiniTickerCombinedResponse struct {
	StreamName string                      `json:"stream"` // Stream Name
	Data       AllMarketMiniTickerResponse `json:"data"`   // data
}
