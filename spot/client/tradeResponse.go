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
	OrderID            int64  `json:"orderId"`
	OrderListID        int64  `json:"orderListId"`
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
	OrderID            int64  `json:"orderId"`
	OrderListID        int64  `json:"orderListId"`
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

// DeleteOrderResponse delete a order
type DeleteOrderResponse struct {
	Symbol             string `json:"symbol"`
	OrigClientOrderID  string `json:"origClientOrderId"`
	OrderID            int64  `json:"orderId"`
	OrderListID        int64  `json:"orderListId"`
	ClientOrderID      string `json:"clientOrderId"`
	Price              string `json:"price"`
	OrigQty            string `json:"origQty"`
	ExecutedQty        string `json:"executedQty"`
	CumulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status             string `json:"status"`
	TimeInForce        string `json:"timeInForce"`
	Type               string `json:"type"`
	Side               string `json:"side"`
}

// DeleteOpenOrderResponse delete all open orders
type DeleteOpenOrderResponse []struct {
	Symbol             string `json:"symbol"`
	OrigClientOrderID  string `json:"origClientOrderId,omitempty"`
	OrderID            int64  `json:"orderId,omitempty"`
	OrderListID        int64  `json:"orderListId"`
	ClientOrderID      string `json:"clientOrderId,omitempty"`
	Price              string `json:"price,omitempty"`
	OrigQty            string `json:"origQty,omitempty"`
	ExecutedQty        string `json:"executedQty,omitempty"`
	CumulativeQuoteQty string `json:"cummulativeQuoteQty,omitempty"`
	Status             string `json:"status,omitempty"`
	TimeInForce        string `json:"timeInForce,omitempty"`
	Type               string `json:"type,omitempty"`
	Side               string `json:"side,omitempty"`
	ContingencyType    string `json:"contingencyType,omitempty"`
	ListStatusType     string `json:"listStatusType,omitempty"`
	ListOrderStatus    string `json:"listOrderStatus,omitempty"`
	ListClientOrderID  string `json:"listClientOrderId,omitempty"`
	TransactionTime    int64  `json:"transactionTime,omitempty"`
	Orders             []struct {
		Symbol        string `json:"symbol"`
		OrderID       int64  `json:"orderId"`
		ClientOrderID string `json:"clientOrderId"`
	} `json:"orders,omitempty"`
	OrderReports []struct {
		Symbol             string `json:"symbol"`
		OrigClientOrderID  string `json:"origClientOrderId"`
		OrderID            int64  `json:"orderId"`
		OrderListID        int64  `json:"orderListId"`
		ClientOrderID      string `json:"clientOrderId"`
		Price              string `json:"price"`
		OrigQty            string `json:"origQty"`
		ExecutedQty        string `json:"executedQty"`
		CumulativeQuoteQty string `json:"cummulativeQuoteQty"`
		Status             string `json:"status"`
		TimeInForce        string `json:"timeInForce"`
		Type               string `json:"type"`
		Side               string `json:"side"`
		StopPrice          string `json:"stopPrice,omitempty"`
		IcebergQty         string `json:"icebergQty"`
	} `json:"orderReports,omitempty"`
}

// GetOrderResponse get order info
type GetOrderResponse struct {
	Symbol             string `json:"symbol"`
	OrderID            int64  `json:"orderId"`
	OrderListID        int64  `json:"orderListId"`
	ClientOrderID      string `json:"clientOrderId"`
	Price              string `json:"price"`
	OrigQty            string `json:"origQty"`
	ExecutedQty        string `json:"executedQty"`
	CumulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status             string `json:"status"`
	TimeInForce        string `json:"timeInForce"`
	Type               string `json:"type"`
	Side               string `json:"side"`
	StopPrice          string `json:"stopPrice"`
	IcebergQty         string `json:"icebergQty"`
	Time               int64  `json:"time"`
	UpdateTime         int64  `json:"updateTime"`
	IsWorking          bool   `json:"isWorking"`
	OrigQuoteOrderQty  string `json:"origQuoteOrderQty"`
}

// GetOpenOrdersResponse get all open orders
type GetOpenOrdersResponse []struct {
	Symbol             string `json:"symbol"`
	OrderID            int64  `json:"orderId"`
	OrderListID        int64  `json:"orderListId"`
	ClientOrderID      string `json:"clientOrderId"`
	Price              string `json:"price"`
	OrigQty            string `json:"origQty"`
	ExecutedQty        string `json:"executedQty"`
	CumulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status             string `json:"status"`
	TimeInForce        string `json:"timeInForce"`
	Type               string `json:"type"`
	Side               string `json:"side"`
	StopPrice          string `json:"stopPrice"`
	IcebergQty         string `json:"icebergQty"`
	Time               int64  `json:"time"`
	UpdateTime         int64  `json:"updateTime"`
	IsWorking          bool   `json:"isWorking"`
	OrigQuoteOrderQty  string `json:"origQuoteOrderQty"`
}

// GetAllOrdersResponse get all orders
type GetAllOrdersResponse []struct {
	Symbol             string `json:"symbol"`
	OrderID            int64  `json:"orderId"`
	OrderListID        int64  `json:"orderListId"`
	ClientOrderID      string `json:"clientOrderId"`
	Price              string `json:"price"`
	OrigQty            string `json:"origQty"`
	ExecutedQty        string `json:"executedQty"`
	CumulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status             string `json:"status"`
	TimeInForce        string `json:"timeInForce"`
	Type               string `json:"type"`
	Side               string `json:"side"`
	StopPrice          string `json:"stopPrice"`
	IcebergQty         string `json:"icebergQty"`
	Time               int64  `json:"time"`
	UpdateTime         int64  `json:"updateTime"`
	IsWorking          bool   `json:"isWorking"`
	OrigQuoteOrderQty  string `json:"origQuoteOrderQty"`
}

// NewOCOOrderResponse new oco order
type NewOCOOrderResponse struct {
	OrderListID       int64  `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderID string `json:"listClientOrderId"`
	TransactionTime   int64  `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderID       int64  `json:"orderId"`
		ClientOrderID string `json:"clientOrderId"`
	} `json:"orders"`
	OrderReports []struct {
		Symbol             string `json:"symbol"`
		OrderID            int64  `json:"orderId"`
		OrderListID        int64  `json:"orderListId"`
		ClientOrderID      string `json:"clientOrderId"`
		TransactTime       int64  `json:"transactTime"`
		Price              string `json:"price"`
		OrigQty            string `json:"origQty"`
		ExecutedQty        string `json:"executedQty"`
		CumulativeQuoteQty string `json:"cummulativeQuoteQty"`
		Status             string `json:"status"`
		TimeInForce        string `json:"timeInForce"`
		Type               string `json:"type"`
		Side               string `json:"side"`
		StopPrice          string `json:"stopPrice,omitempty"`
	} `json:"orderReports"`
}

// DeleteOCOResponse delete oco order response
type DeleteOCOResponse struct {
	OrderListID       int64  `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderID string `json:"listClientOrderId"`
	TransactionTime   int64  `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderID       int64  `json:"orderId"`
		ClientOrderID string `json:"clientOrderId"`
	} `json:"orders"`
	OrderReports []struct {
		Symbol             string `json:"symbol"`
		OrigClientOrderID  string `json:"origClientOrderId"`
		OrderID            int64  `json:"orderId"`
		OrderListID        int64  `json:"orderListId"`
		ClientOrderID      string `json:"clientOrderId"`
		Price              string `json:"price"`
		OrigQty            string `json:"origQty"`
		ExecutedQty        string `json:"executedQty"`
		CumulativeQuoteQty string `json:"cummulativeQuoteQty"`
		Status             string `json:"status"`
		TimeInForce        string `json:"timeInForce"`
		Type               string `json:"type"`
		Side               string `json:"side"`
		StopPrice          string `json:"stopPrice,omitempty"`
	} `json:"orderReports"`
}

// GetOCOResponse get oco order
type GetOCOResponse struct {
	OrderListID       int64  `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderID string `json:"listClientOrderId"`
	TransactionTime   int64  `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderID       int64  `json:"orderId"`
		ClientOrderID string `json:"clientOrderId"`
	} `json:"orders"`
}

// GetAllOCOResponse get all oco orders
type GetAllOCOResponse []struct {
	OrderListID       int64  `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderID string `json:"listClientOrderId"`
	TransactionTime   int64  `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderID       int64  `json:"orderId"`
		ClientOrderID string `json:"clientOrderId"`
	} `json:"orders"`
}

// GetOpenOCOResponse get open oco orders
type GetOpenOCOResponse []struct {
	OrderListID       int64  `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderID string `json:"listClientOrderId"`
	TransactionTime   int64  `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderID       int64  `json:"orderId"`
		ClientOrderID string `json:"clientOrderId"`
	} `json:"orders"`
}

// AccountInfoResponse account information
type AccountInfoResponse struct {
	MakerCommission  int64  `json:"makerCommission"`
	TakerCommission  int64  `json:"takerCommission"`
	BuyerCommission  int64  `json:"buyerCommission"`
	SellerCommission int64  `json:"sellerCommission"`
	CanTrade         bool   `json:"canTrade"`
	CanWithdraw      bool   `json:"canWithdraw"`
	CanDeposit       bool   `json:"canDeposit"`
	UpdateTime       int    `json:"updateTime"`
	AccountType      string `json:"accountType"`
	Balances         []struct {
		Asset  string `json:"asset"`
		Free   string `json:"free"`
		Locked string `json:"locked"`
	} `json:"balances"`
	Permissions []string `json:"permissions"`
}

// AccountTradesListResponse account trades list
type AccountTradesListResponse []struct {
	Symbol          string `json:"symbol"`
	ID              int64  `json:"id"`
	OrderID         int64  `json:"orderId"`
	OrderListID     int64  `json:"orderListId"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	QuoteQty        string `json:"quoteQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsBestMatch     bool   `json:"isBestMatch"`
}
