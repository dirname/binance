package futuresusdt

// AllMarketTickerResponse json parser
type AllMarketTickerResponse struct {
	EventType       string `json:"e"` // Event type
	EventTime       int64  `json:"E"` // Event time stamp
	Symbol          string `json:"s"` // Symbol
	MarketPrice     string `json:"p"` // Market price
	IndexPrice      string `json:"i"` // Index price
	EstimatedPrice  string `json:"P"` // Estimated Settle Price, only useful in the last hour before the settlement starts
	FundingRate     string `json:"r"` // Funding rate
	NextFundingTime int64  `json:"T"` // Next funding time
}

// AllMarketPriceCombinedResponse json parser
type AllMarketTickerCombinedResponse struct {
	StreamName string                    `json:"stream"` // Stream Name
	Data       []AllMarketTickerResponse `json:"data"`   // data
}
