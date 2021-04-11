package spotclient



// NewOrderResponseACK Send in a new order.
type NewOrderResponseACK struct {
	Symbol        string `json:"symbol"`
	OrderID       int64  `json:"orderId"`
	OrderListID   int64  `json:"orderListId"`
	ClientOrderID string `json:"clientOrderId"`
	TransactTime  int64  `json:"transactTime"`
}

// NewOrderResponseResult Send in a new order.
type NewOrderResponseResult struct {
	Symbol             string `json:"symbol"`
	OrderID            int    `json:"orderId"`
	OrderListID        int    `json:"orderListId"`
	ClientOrderID      string `json:"clientOrderId"`
	TransactTime       int64  `json:"transactTime"`
	Price              string `json:"price"`
	OrigQty            string `json:"origQty"`
	ExecutedQty        string `json:"executedQty"`
	CumulativeQuoteQTY string `json:"cummulativeQuoteQty"`
	Status             string `json:"status"`
	TimeInForce        string `json:"timeInForce"`
	Type               string `json:"type"`
	Side               string `json:"side"`
}

// NewOrderResponseFull send in a new order.
type NewOrderResponseFull struct {
	Symbol             string `json:"symbol"`
	OrderID            int    `json:"orderId"`
	OrderListID        int    `json:"orderListId"`
	ClientOrderID      string `json:"clientOrderId"`
	TransactTime       int64  `json:"transactTime"`
	Price              string `json:"price"`
	OrigQTY            string `json:"origQty"`
	ExecutedQty        string `json:"executedQty"`
	CumulativeQuoteQTY string `json:"cummulativeQuoteQty"`
	Status             string `json:"status"`
	TimeInForce        string `json:"timeInForce"`
	Type               string `json:"type"`
	Side               string `json:"side"`
	Fills              []struct {
		Price           string `json:"price"`
		QTY             string `json:"qty"`
		Commission      string `json:"commission"`
		CommissionAsset string `json:"commissionAsset"`
	} `json:"fills"`
}
