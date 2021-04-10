package spotclient

// PartialBookDepthResponse json parser
type PartialBookDepthResponse struct {
	LastUpdateID int64      `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"` // Bids to be updated, in addition, there will be two elements in the array, the first is Price Level to be, and the second is Quantity
	Asks         [][]string `json:"asks"` // Asks to be updated, in addition, there will be two elements in the array, the first is Price Level to be, and the second is Quantity
}

// PartialBookDepthCombinedResponse json parser
type PartialBookDepthCombinedResponse struct {
	StreamName string                   `json:"stream"` // Stream Name
	Data       PartialBookDepthResponse `json:"data"`   // data
}
