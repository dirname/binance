package futuresclient

// CompositeIndexResponse json parser
type CompositeIndexResponse struct {
	EventType   string `json:"e"` // Event type
	EventTime   int64  `json:"E"` // Event time stamp
	Symbol      string `json:"s"` // Symbol
	Price       string `json:"p"` // Price
	Composition []struct {
		BaseAsset        string `json:"b"`
		QuantityWeight   string `json:"w"` // Weight in quantity
		PercentageWeight string `json:"W"` // Weight in percentage
	} `json:"c"` // composition
}

// CompositeIndexCombinedResponse json parser
type CompositeIndexCombinedResponse struct {
	StreamName string                 `json:"stream"` // Stream Name
	Data       CompositeIndexResponse `json:"data"`   // data
}
