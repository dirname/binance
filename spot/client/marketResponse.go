package spotclient

// ExchangeInfoResponse Current exchange trading rules and symbol information
type ExchangeInfoResponse struct {
	TimeZone   string `json:"timezone"`   // timezone
	ServerTime int64  `json:"serverTime"` // server time
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"` // rateLimitType
		Interval      string `json:"interval"`      // interval
		IntervalNum   int64  `json:"intervalNum"`   // intervalNum
		Limit         int64  `json:"limit"`         // limit
	} `json:"rateLimits"`                                  // rate limits
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
		} `json:"filters"`                        // Filters
		Permissions []string `json:"permissions"` // Permissions
	} `json:"symbols"` // symbols
}

// OrderBookResponse order book (depth) info The first in the array is price and the second is qty
type OrderBookResponse struct {
	LastUpdateId int64      `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

// RecentTradesListResponse Get recent trades.
type RecentTradesListResponse []struct {
	ID           int64  `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	QuoteQty     string `json:"quoteQty"`
	Time         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

// OlderTradeLookUpResponse Get older market trades.
type OlderTradeLookUpResponse []struct {
	ID           int64  `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	QuoteQty     string `json:"quoteQty"`
	Time         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

// AggregateTradeResponse Get compressed, aggregate trades. Trades that fill at the time, from the same order, with the same price will have the quantity aggregated.
type AggregateTradeResponse []struct {
	TradeID          int64  `json:"a"` // Aggregate tradeId
	Price            string `json:"p"` // Price
	Quantity         string `json:"q"` // Quantity
	FirstTradeID     int64  `json:"f"` // First tradeId
	LastTradeID      int64  `json:"l"` // Last tradeId
	TimeStamp        int64  `json:"T"` // Timestamp
	IsBuyerMarker    bool   `json:"m"` // Was the buyer the maker?
	IsBestPriceMatch bool   `json:"M"` // Was the trade the best price match?
}

// CandlestickResponse kline response
type CandlestickResponse [][]interface{}

// CurrentAveragePriceResponse Current average price for a symbol.
type CurrentAveragePriceResponse struct {
	Min   int    `json:"mins"`
	Price string `json:"price"`
}

// TickerPriceChangeStatisticsResponse 24 hour rolling window price change statistics. Careful when accessing this with no symbol.
type TickerPriceChangeStatisticsResponse struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQTY            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	AskPrice           string `json:"askPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstID            int64  `json:"firstId"`
	LastID             int64  `json:"lastId"`
	Count              int    `json:"count"`
}

// SymbolPriceTickerResponse Latest price for a symbol or symbols.
type SymbolPriceTickerResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// SymbolOrderBookTickerResponse Best price/qty on the order book for a symbol or symbols.
type SymbolOrderBookTickerResponse struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQTY   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQTY   string `json:"askQty"`
}
