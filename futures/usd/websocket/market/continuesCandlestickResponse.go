package futuresclient

// ContinuesCandlestickResponse json parser
type ContinuesCandlestickResponse struct {
	EventType    string `json:"e"`  // Event type
	EventTime    int64  `json:"E"`  // Event time stamp
	Pair         string `json:"ps"` // Pair
	ContractType string `json:"ct"` // Contract type
	KLine        struct {
		StartTime        int64  `json:"t"` // K line start time
		CloseTime        int64  `json:"T"` // K line close time
		Interval         string `json:"i"` // Interval
		FirstTradeID     int64  `json:"f"` // First Trade ID
		LastTradeID      int64  `json:"L"` // Last Trade ID
		OpenPrice        string `json:"o"` // Open Price
		ClosePrice       string `json:"c"` // Close Price
		HighPrice        string `json:"h"` // Highest price
		LowPrice         string `json:"l"` // Lowest price
		BaseVolume       string `json:"v"` // Base asset volume
		TradeNumbers     int64  `json:"n"` // Number of trades
		Closed           bool   `json:"x"` // Is this K line closed
		QuoteVolume      string `json:"q"` // Quote asset volume
		TakerBaseVolume  string `json:"V"` // Taker buy base asset volume
		TakerQuoteVolume string `json:"Q"` // Taker buy quote asset volume
		OtherParam       string `json:"B"` // Ignore
	} `json:"k"` // K line
}

// ContinuesCandlestickCombinedResponse json parser
type ContinuesCandlestickCombinedResponse struct {
	StreamName string                       `json:"stream"` // Stream Name
	Data       ContinuesCandlestickResponse `json:"data"`   // data
}
