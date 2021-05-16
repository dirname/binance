package order

const (
	ACK             = "ACK"               //ACK ACK
	Result          = "RESULT"            //Result Result
	Full            = "FULL"              //Full Full
	Buy             = "BUY"               //Buy buy
	Sell            = "SELL"              //Sell sell
	New             = "NEW"               //New The order has been accepted by the engine.
	PartiallyFilled = "PARTIALLY_FILLED"  //PartiallyFilled A part of the order has been filled.
	Filled          = "FILLED"            //Filled The order has been completed.
	Canceled        = "CANCELED"          //Canceled The order has been canceled by the user.
	PendingCanceled = "PENDING_CANCEL"    //PendingCanceled Currently unused
	Rejected        = "REJECTED"          //Rejected The order was not accepted by the engine and not processed.
	Expired         = "EXPIRED"           //Expired The order was canceled according to the order type's rules (e.g. LIMIT FOK orders with no fill, LIMIT IOC or MARKET orders that partially fill) or by the exchange, (e.g. orders canceled during liquidation, orders canceled during maintenance)
	Limit           = "LIMIT"             //Limit LIMIT
	Market          = "MARKET"            //Market MARKET
	StopLoss        = "STOP_LOSS"         //StopLoss STOP_LOSS
	StopLossLimit   = "STOP_LOSS_LIMIT"   //StopLossLimit STOP_LOSS_LIMIT
	TakeProfit      = "TAKE_PROFIT"       //TakeProfit TAKE_PROFIT
	TakeProfitLimit = "TAKE_PROFIT_LIMIT" //TakeProfitLimit TAKE_PROFIT_LIMIT
	LimitMarker     = "LIMIT_MAKER"       //LimitMarker LIMIT_MARKER
	GTC             = "GTC"               //GTC Good Til Canceled An order will be on the book unless the order is canceled.
	IOC             = "IOC"               //IOC Immediate Or Cancel An order will try to fill the order as much as it can before the order expires.
	FOK             = "FOK"               //FOK Fill or Kill An order will expire if the full order cannot be filled upon execution.
)
