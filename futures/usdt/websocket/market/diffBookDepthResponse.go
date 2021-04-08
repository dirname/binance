package futuresusdt

// DiffBookDepthResponse json parser
type DiffBookDepthResponse struct {
	EventType                 string     `json:"e"`  // Event type
	EventTime                 int64      `json:"E"`  // Event time stamp
	TransactionTime           int64      `json:"T"`  // Transaction time stamp
	Symbol                    string     `json:"s"`  // Symbol
	FirstUpdateID             int64      `json:"U"`  // First update ID in event
	FinalUpdateID             int64      `json:"u"`  // Final update ID in event
	FinalUpdateIDInLastStream int64      `json:"Pu"` // Final update Id in last stream(ie `u` in last stream)
	Bids                      [][]string `json:"b"`  // Bids to be updated, in addition, there will be two elements in the array, the first is Price Level to be, and the second is Quantity
	Asks                      [][]string `json:"a"`  // Asks to be updated, in addition, there will be two elements in the array, the first is Price Level to be, and the second is Quantity
}

// DiffBookDepthCombinedResponse json parser
type DiffBookDepthCombinedResponse struct {
	StreamName string                `json:"stream"` // Stream Name
	Data       DiffBookDepthResponse `json:"data"`   // data
}
