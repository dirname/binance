package client

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
		Asset             string `json:"asset"`
		MarginAvailable   bool   `json:"marginAvailable"`
		AutoAssetExchange *int   `json:"autoAssetExchange"`
	} `json:"assets"`
	Symbols []struct {
		Symbol                string                   `json:"symbol"`
		Pair                  string                   `json:"pair"`
		ContractType          string                   `json:"contractType"`
		DeliveryDate          int64                    `json:"deliveryDate"`
		OnboardDate           int64                    `json:"onboardDate"`
		Status                string                   `json:"status"`
		MaintMarginPercent    string                   `json:"maintMarginPercent"`
		RequiredMarginPercent string                   `json:"requiredMarginPercent"`
		BaseAsset             string                   `json:"baseAsset"`
		QuoteAsset            string                   `json:"quoteAsset"`
		MarginAsset           string                   `json:"marginAsset"`
		PricePrecision        int                      `json:"pricePrecision"`
		QuantityPrecision     int                      `json:"quantityPrecision"`
		BaseAssetPrecision    int                      `json:"baseAssetPrecision"`
		QuotePrecision        int                      `json:"quotePrecision"`
		UnderlyingType        string                   `json:"underlyingType"`
		UnderlyingSubType     []string                 `json:"underlyingSubType"`
		SettlePlan            int                      `json:"settlePlan"`
		TriggerProtect        string                   `json:"triggerProtect"`
		Filters               []map[string]interface{} `json:"filters"`
		OrderType             []string                 `json:"OrderType"`
		TimeInForce           []string                 `json:"timeInForce"`
	} `json:"symbols"`
	Timezone string `json:"timezone"`
}
