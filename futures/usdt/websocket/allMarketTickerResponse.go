package futuresusdt

// AllMarketTickerResponse json parser
type AllMarketTickerResponse struct {
	EventType          string `json:"e"` // Event type
	EventTime          int64  `json:"E"` // Event time stamp
	Symbol             string `json:"s"` // Symbol
	PriceChange        string `json:"p"` // Price change
	PriceChangePercent string `json:"P"` // Price change percent
	AverageWeighted    string `json:"w"` // Weighted average price
	LastPrice          string `json:"c"` // Last price
	LastQuantity       string `json:"Q"` // Last quantity
	OpenPrice          string `json:"o"` // Open price
	HighPrice          string `json:"h"` // High price
	LowPrice           string `json:"l"` // Low price
	BaseVolume         string `json:"v"` // Total traded base asset volume
	QuoteVolume        string `json:"q"` // Total traded quote asset volume
	OpenTime           int64  `json:"O"` // Statistics open time
	CloseTime          int64  `json:"C"` // Statistics close time
	FirstTradeID       int64  `json:"F"` // First trade id
	LastTradeID        int64  `json:"L"` // Last trade id
	TotalTrades        int64  `json:"n"` // Total number of trades
}

// AllMarketTickerCombinedResponse json parser
type AllMarketTickerCombinedResponse struct {
	StreamName string                    `json:"stream"` // Stream Name
	Data       []AllMarketTickerResponse `json:"data"`   // data
}
