package orderStatus

const (
	New             = "NEW"              // The order has been accepted by the engine.
	PartiallyFilled = "PARTIALLY_FILLED" // A part of the order has been filled.
	Filled          = "FILLED"           // The order has been completed.
	Canceled        = "CANCELED"         // The order has been canceled by the user.
	PendingCanceled = "PENDING_CANCEL"   // Currently unused
	Rejected        = "REJECTED"         // The order was not accepted by the engine and not processed.
	Expired         = "EXPIRED"          // The order was canceled according to the order type's rules (e.g. LIMIT FOK orders with no fill, LIMIT IOC or MARKET orders that partially fill) or by the exchange, (e.g. orders canceled during liquidation, orders canceled during maintenance)
)
