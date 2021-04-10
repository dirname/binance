package spotclient

// TradeResponse json parser
type TradeResponse struct {
	EventType           string `json:"e"` // Event type
	EventTime           int64  `json:"E"` // Event time stamp
	Symbol              string `json:"s"` // Symbol
	TradeID             int    `json:"t"` // Trade ID
	Price               string `json:"p"` // Price
	Quantity            string `json:"q"` // Quantity
	BuyerOrderID        int    `json:"b"` // BuyerOrderID
	SellerOrderID       int    `json:"a"` // SellerOrderID
	TradeTime           int    `json:"T"` // Trade time
	IsBuyerMarketMarker bool   `json:"m"` // Is the buyer the market maker
	UppercaseM          bool   `json:"M"` // Ignore
}

// TradeCombinedResponse json parser
type TradeCombinedResponse struct {
	StreamName string        `json:"stream"` // Stream Name
	Data       TradeResponse `json:"data"`   // data
}
