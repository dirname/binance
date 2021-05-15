package futuresclient

const (
	MissParameters                 = "miss parameters"
	SymbolEmpty                    = "symbol cannot be empty"
	IntervalEmpty                  = "interval cannot be empty"
	PairEmpty                      = "pair cannot be empty"
	ContractEmpty                  = "contract cannot be empty"
	PeriodEmpty                    = "period cannot be empty"
	CallBackRateInvalid            = "callback cannot less than or qual zero"
	CoinEmpty                      = "coin cannot be empty"
	AddressEmpty                   = "address cannot be empty"
	AmountInvalid                  = "amount cannot less than or equal zero"
	AssetEmpty                     = "asset cannot be empty"
	TimeInForce                    = "timeInForce cannot be empty"
	PriceInvalid                   = "price cannot less than or equal zero"
	StopPriceInvalid               = "stopPrice cannot less than or equal zero"
	QuantityInvalid                = "quantity cannot less than or equal zero"
	QuantityOrQuoteOrderQtyInvalid = "quantity or quoteOrderQty cannot less than or equal zero"
	InvalidParameters              = "invalid parameters"
	ReduceOnlyWithClosePosition    = "closePosition does not accept parameters reduceOnly and quantity"
	OrderTypeWithClosePosition     = "closePosition only accepts order types STOP_MARKET and TAKE_PROFIT_MARKET"
)
