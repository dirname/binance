package futuresclient

type ExchangeInfoResponse struct {
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	RateLimits      []struct {
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
		RateLimitType string `json:"rateLimitType"`
	} `json:"rateLimits"`
	ServerTime int64 `json:"serverTime"`
	Assets     []struct {
		Asset             string      `json:"asset"`
		MarginAvailable   bool        `json:"marginAvailable"`
		AutoAssetExchange interface{} `json:"autoAssetExchange"`
	} `json:"assets"`
	Symbols []struct {
		Symbol                   string                   `json:"symbol"`
		Pair                     string                   `json:"pair"`
		ContractType             string                   `json:"contractType"`
		DeliveryDate             int64                    `json:"deliveryDate"`
		OnboardDate              int64                    `json:"onboardDate"`
		Status                   string                   `json:"status"`
		MaintenanceMarginPercent string                   `json:"maintMarginPercent"`
		RequiredMarginPercent    string                   `json:"requiredMarginPercent"`
		BaseAsset                string                   `json:"baseAsset"`
		QuoteAsset               string                   `json:"quoteAsset"`
		MarginAsset              string                   `json:"marginAsset"`
		PricePrecision           int                      `json:"pricePrecision"`
		QuantityPrecision        int                      `json:"quantityPrecision"`
		BaseAssetPrecision       int                      `json:"baseAssetPrecision"`
		QuotePrecision           int                      `json:"quotePrecision"`
		UnderlyingType           string                   `json:"underlyingType"`
		UnderlyingSubType        []string                 `json:"underlyingSubType"`
		SettlePlan               int                      `json:"settlePlan"`
		TriggerProtect           string                   `json:"triggerProtect"`
		Filters                  []map[string]interface{} `json:"filters"`
		OrderType                []string                 `json:"OrderType"`
		TimeInForce              []string                 `json:"timeInForce"`
	} `json:"symbols"`
	Timezone string `json:"timezone"`
}

//OrderBookResponse depth response
type OrderBookResponse struct {
	LastUpdateID int64      `json:"lastUpdateId"`
	EventTime    int64      `json:"E"`
	EngineTIme   int64      `json:"T"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

//RecentTradesListResponse Get recent trades.
type RecentTradesListResponse []struct {
	ID           int64  `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	QuoteQty     string `json:"quoteQty"`
	Time         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
}

// OlderTradeLookUpResponse Get older market trades.
type OlderTradeLookUpResponse []struct {
	ID           int64  `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	QuoteQty     string `json:"quoteQty"`
	Time         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
}

// AggregateTradeResponse Get compressed, aggregate trades. Trades that fill at the time, from the same order, with the same price will have the quantity aggregated.
type AggregateTradeResponse []struct {
	TradeID       int64  `json:"a"`
	Price         string `json:"p"`
	Quantity      string `json:"q"`
	FirstTradeID  int64  `json:"f"`
	LastTradeID   int64  `json:"l"`
	TimeStamp     int64  `json:"T"`
	IsBuyerMarker bool   `json:"m"`
}

// CandlestickResponse kline response
type CandlestickResponse []interface{}

//ContractCandlestickResponse contract kline response
type ContractCandlestickResponse []interface{}

//IndexCandlestickResponse index kline response
type IndexCandlestickResponse []interface{}

//MarketPriceCandlestickResponse market price kline response
type MarketPriceCandlestickResponse []interface{}

//MarketPriceResponse market price response
type MarketPriceResponse struct {
	Symbol          string `json:"symbol"`
	MarkPrice       string `json:"markPrice"`
	IndexPrice      string `json:"indexPrice"`
	LastFundingRate string `json:"lastFundingRate"`
	NextFundingTime int64  `json:"nextFundingTime"`
	InterestRate    string `json:"interestRate"`
	Time            int64  `json:"time"`
}

//FundingRateResponse funding rate history
type FundingRateResponse []struct {
	Symbol      string `json:"symbol"`
	FundingRate string `json:"fundingRate"`
	FundingTime int64  `json:"fundingTime"`
}

//TickerPriceChangeStatisticsResponse ticker price change statistics
type TickerPriceChangeStatisticsResponse struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstID            int64  `json:"firstId"`
	LastID             int64  `json:"lastId"`
	Count              int32  `json:"count"`
}

//SymbolPriceTickerResponse Latest price for a symbol or symbols.
type SymbolPriceTickerResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
	Time   int64  `json:"time"`
}

//SymbolOrderBookTickerResponse Best price/qty on the order book for a symbol or symbols.
type SymbolOrderBookTickerResponse struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
	Time     int64  `json:"time"`
}

//OpenInterestResponse Get present open interest of a specific symbol.
type OpenInterestResponse struct {
	OpenInterest string `json:"openInterest"`
	Symbol       string `json:"symbol"`
	Time         int64  `json:"time"`
}

//OpenInterestStatisticsResponse Get open interest statistics
type OpenInterestStatisticsResponse []struct {
	Symbol               string `json:"symbol"`
	SumOpenInterest      string `json:"sumOpenInterest"`
	SumOpenInterestValue string `json:"sumOpenInterestValue"`
	Timestamp            string `json:"timestamp"`
}

//TopTraderAccountsRatioResponse Get top trader accounts ratio response
type TopTraderAccountsRatioResponse []struct {
	Symbol         string `json:"symbol"`
	LongShortRatio string `json:"longShortRatio"`
	LongAccount    string `json:"longAccount"`
	ShortAccount   string `json:"shortAccount"`
	Timestamp      string `json:"timestamp"`
}

//TopTraderPositionsRatioResponse Get top trader positions ratio response
type TopTraderPositionsRatioResponse []struct {
	Symbol         string `json:"symbol"`
	LongShortRatio string `json:"longShortRatio"`
	LongAccount    string `json:"longAccount"`
	ShortAccount   string `json:"shortAccount"`
	Timestamp      string `json:"timestamp"`
}

//LongShortRatioResponse get long short ratio
type LongShortRatioResponse []struct {
	Symbol         string `json:"symbol"`
	LongShortRatio string `json:"longShortRatio"`
	LongAccount    string `json:"longAccount"`
	ShortAccount   string `json:"shortAccount"`
	Timestamp      string `json:"timestamp"`
}

//TakerBuySellVolumeResponse get tacker buy/sell volume
type TakerBuySellVolumeResponse struct {
	BuySellRatio string `json:"buySellRatio"`
	BuyVol       string `json:"buyVol"`
	SellVol      string `json:"sellVol"`
	Timestamp    string `json:"timestamp"`
}

//HistoricalBLVTNavCandlestickResponse get BLVT Nav candlestick
type HistoricalBLVTNavCandlestickResponse []interface{}

//CompositeIndexSymbolResponse composite index symbol response
type CompositeIndexSymbolResponse struct {
	Symbol        string `json:"symbol"`
	Time          int64  `json:"time"`
	BaseAssetList []struct {
		BaseAsset          string `json:"baseAsset"`
		WeightInQuantity   string `json:"weightInQuantity"`
		WeightInPercentage string `json:"weightInPercentage"`
	} `json:"baseAssetList"`
}
