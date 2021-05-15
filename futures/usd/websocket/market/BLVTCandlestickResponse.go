package futuresclient

// BLVTCandlestickResponse json parser
type BLVTCandlestickResponse struct {
	EventType string `json:"e"` // Event type
	EventTime int64  `json:"E"` // Event time stamp
	Symbol    string `json:"s"` // BLVT name
	KLine     struct {
		StartTime        int64  `json:"t"` // K line start time
		CloseTime        int64  `json:"T"` // K line close time
		Symbol           string `json:"s"` // BLVT name
		Interval         string `json:"i"` // Interval
		FirstTradeID     int64  `json:"f"` // First NAV update time
		LastTradeID      int64  `json:"L"` // Last NAV update time
		OpenPrice        string `json:"o"` // Open NAV Price
		ClosePrice       string `json:"c"` // Close NAV Price
		HighPrice        string `json:"h"` // Highest NAV price
		LowPrice         string `json:"l"` // Lowest NAV price
		RealLeverage     string `json:"v"` // Real leverage
		NAVNumber        int64  `json:"n"` // Number of NAV update
		Closed           bool   `json:"x"` // Ignore
		QuoteVolume      string `json:"q"` // Ignore
		TakerBaseVolume  string `json:"V"` // Ignore
		TakerQuoteVolume string `json:"Q"` // Ignore
		OtherParam       string `json:"B"` // Ignore
	} `json:"k"`
}

// BLVTCandlestickCombinedResponse json parser
type BLVTCandlestickCombinedResponse struct {
	StreamName string                  `json:"stream"` // Stream Name
	Data       BLVTCandlestickResponse `json:"data"`   // data
}
