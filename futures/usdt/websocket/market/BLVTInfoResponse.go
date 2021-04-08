package futuresusdt

import "github.com/shopspring/decimal"

// BLVTInfoResponse json parser
type BLVTInfoResponse struct {
	EventType   string          `json:"e"` // Event type
	EventTime   int64           `json:"E"` // Event time stamp
	Symbol      string          `json:"s"` // BLVT name
	IssuedToken decimal.Decimal `json:"m"` // Token issued
	Baskets     []struct {
		Symbol  string          `json:"s"` // Futures symbol
		Numbers decimal.Decimal `json:"n"` // position
	} `json:"b"`                              // Baskets
	NetValue       decimal.Decimal `json:"n"` // BLVT NAV
	RealLeverage   decimal.Decimal `json:"l"` // Real leverage
	TargetLeverage decimal.Decimal `json:"t"` // Target leverage
	FundingRatio   decimal.Decimal `json:"f"` // FundingRatio
}

// BLVTInfoCombinedResponse json parser
type BLVTInfoCombinedResponse struct {
	StreamName string           `json:"stream"` // Stream Name
	Data       BLVTInfoResponse `json:"data"`   // data
}
