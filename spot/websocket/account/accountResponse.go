package spotclient

// AccountPosition account position
type AccountPosition struct {
	EventType  string `json:"e"`
	EventTime  int64  `json:"E"`
	UpdateTime int64  `json:"u"`
	Balance    []struct {
		Asset            string `json:"a"`
		AvailableBalance string `json:"f"`
		FreezeBalance    string `json:"l"`
	} `json:"B"`
}

// AccountCombinedPosition account position
type AccountCombinedPosition struct {
	StreamName string          `json:"stream"` // Stream Name
	Data       AccountPosition `json:"data"`   // data
}

// BalanceUpdate balance update
type BalanceUpdate struct {
	EventType    string `json:"e"`
	EventTime    int64  `json:"E"`
	Asset        string `json:"a"`
	BalanceDelta string `json:"d"`
	ClearTime    int64  `json:"T"`
}

// BalanceCombinedUpdate balance update
type BalanceCombinedUpdate struct {
	StreamName string        `json:"stream"` // Stream Name
	Data       BalanceUpdate `json:"data"`   // data
}

// ExecutionReport execution report
type ExecutionReport struct {
	EventType                        string      `json:"e"` // Event type
	EventTime                        int64       `json:"E"` // Event time
	Symbol                           string      `json:"s"` // Symbol
	ClientOrderID                    string      `json:"c"` // Client order ID
	Side                             string      `json:"S"` // Side
	OrderType                        string      `json:"o"` // Order type
	TimeInForce                      string      `json:"f"` // Time in force
	OrderQuantity                    string      `json:"q"` // Order quantity
	OrderPrice                       string      `json:"p"` // Order price
	StopPrice                        string      `json:"P"` // Stop price
	IcebergQuantity                  string      `json:"F"` // Iceberg quantity
	OrderListID                      int64       `json:"g"` // OrderListId
	ClientID                         string      `json:"C"` // Original client order ID; This is the ID of the order being canceled
	ExecutionType                    string      `json:"x"` // Current execution type
	OrderStatus                      string      `json:"X"` // Current order status
	OrderRejectReason                string      `json:"r"` // Order reject reason; will be an error code.
	OrderID                          int         `json:"i"` // Order ID
	LastExecutedQuantity             string      `json:"l"` // Last executed quantity
	CumulativeFilledQuantity         string      `json:"z"` // Cumulative filled quantity
	LastExecutedPrice                string      `json:"L"` // Last executed price
	CommissionAmount                 string      `json:"n"` // Commission amount
	CommissionAsset                  interface{} `json:"N"` // Commission asset
	TransactionTime                  int64       `json:"T"` // Transaction time
	TradeID                          int         `json:"t"` // Trade ID
	UppercaseI                       int         `json:"I"` // Ignore
	IsOrderOnBook                    bool        `json:"w"` // Is the order on the book?
	IsMakerSide                      bool        `json:"m"` // Is this trade the maker side?
	UppercaseM                       bool        `json:"M"` // Ignore
	OrderCreationTime                int64       `json:"O"` // Order creation time
	CumulativeQuoteAsset             string      `json:"Z"` // Cumulative quote asset transacted quantity
	LastQuoteAssetTransactedQuantity string      `json:"Y"` // Last quote asset transacted quantity (i.e. lastPrice * lastQty)
	QuoteOrderQty                    string      `json:"Q"` // Quote Order Qty
}

// ExecutionCombinedReport execution report
type ExecutionCombinedReport struct {
	StreamName string          `json:"stream"` // Stream Name
	Data       ExecutionReport `json:"data"`   // data
}

// ListStatus list status
type ListStatus struct {
	EventType         string `json:"e"`
	EventTime         int64  `json:"E"`
	Symbol            string `json:"s"`
	OrderListID       int64  `json:"g"`
	ContingencyType   string `json:"c"`
	ListStatusType    string `json:"l"`
	ListOrderType     string `json:"L"`
	ListRejectType    string `json:"r"`
	ListClientOrderID string `json:"C"`
	TransactionTime   int64  `json:"T"`
	Objects           []struct {
		Symbol        string `json:"s"`
		OrderID       int64  `json:"i"`
		ClientOrderID string `json:"c"`
	} `json:"O"`
}

// ListCombinedStatus list combined status
type ListCombinedStatus struct {
	StreamName string     `json:"stream"` // Stream Name
	Data       ListStatus `json:"data"`   // data
}

// OrderJudge order judge
type OrderJudge struct {
	EventType interface{} `json:"e"`
}

// OrderCombinedJudge order combined judge
type OrderCombinedJudge struct {
	StreamName string     `json:"stream"` // Stream Name
	Data       OrderJudge `json:"data"`   // data
}
