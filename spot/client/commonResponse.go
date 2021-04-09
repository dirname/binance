package spotclient

// SystemStatusResponse system status
type SystemStatusResponse struct {
	Status  int    `json:"status"` // Status 0 is normal 1 system maintenance
	Message string `json:"msg"`
}

// ExchangeInfoResponse exchange info
type ExchangeInfoResponse struct {
	TimeZone   string `json:"timezone"`   // timezone
	ServerTime int64  `json:"serverTime"` // server time
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"` // rateLimitType
		Interval      string `json:"interval"`      // interval
		IntervalNum   int64  `json:"intervalNum"`   // intervalNum
		Limit         int64  `json:"limit"`         // limit
	} `json:"rateLimits"` // rate limits
	ExchangeFilters []interface{} `json:"exchangeFilters"` // exchangeFilters
	Symbols         []struct {
		Symbol                 string   `json:"symbol"`                 // symbol
		Status                 string   `json:"status"`                 // status
		BaseAsset              string   `json:"baseAsset"`              // baseAsset
		BaseAssetPrecision     int64    `json:"baseAssetPrecision"`     // baseAssetPrecision
		QuoteAsset             string   `json:"quoteAsset"`             // quoteAsset
		QuotePrecision         int64    `json:"quotePrecision"`         // quotePrecision
		QuoteAssetPrecision    int64    `json:"quoteAssetPrecision"`    // quoteAssetPrecision
		OrderTypes             []string `json:"orderTypes"`             // order types
		IcebergAllowed         bool     `json:"icebergAllowed"`         // icebergAllowed
		OCOAllowed             bool     `json:"ocoAllowed"`             // ocoAllowed
		IsSpotTradingAllowed   bool     `json:"isSpotTradingAllowed"`   // isSpotTradingAllowed
		IsMarginTradingAllowed bool     `json:"isMarginTradingAllowed"` // isMarginTradingAllowed
		Filters                []struct {
			FilterType string `json:"filterType"` // filterType
			MinPrice   string `json:"minPrice"`   // minPrice
			MaxPrice   string `json:"maxPrice"`   // maxPrice
			TickSize   string `json:"tickSize"`   // tickSize
		} `json:"filters"` // Filters
		Permissions []string `json:"permissions"` // Permissions
	} `json:"symbols"` // symbols
}
