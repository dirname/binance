package order

const (
	TRUE               = "TRUE"                 //TRUE upper string true
	FALSE              = "FALSE"                //FALSE upper string false
	False              = "false"                //False lower string false
	True               = "true"                 //True lower string true
	Result             = "RESULT"               //Result result response
	ACK                = "ACK"                  //ACK ack response
	Buy                = "BUY"                  //Buy buy side
	Sell               = "SELL"                 //Sell sell side
	Long               = "LONG"                 //Long long side
	Both               = "Both"                 //Both both side
	Short              = "SHORT"                //Short short side
	Limit              = "LIMIT"                //Limit limit
	Market             = "MARKET"               //Market market
	Stop               = "STOP"                 //Stop stop
	TakeProfit         = "TAKE_PROFIT"          //TakeProfit take profit
	StopMarket         = "STOP_MARKET"          //StopMarket stop market
	TakeProfitMarket   = "TAKE_PROFIT_MARKET"   //TakeProfitMarket take profit market price
	TrailingStopMarket = "TRAILING_STOP_MARKET" //TrailingStopMarket trailing stop market
	GTC                = "GTC"                  //GTC Good Til Canceled An order will be on the book unless the order is canceled.
	IOC                = "IOC"                  //IOC Immediate Or Cancel An order will try to fill the order as much as it can before the order expires.
	FOK                = "FOK"                  //FOK Fill or Kill An order will expire if the full order cannot be filled upon execution.
)
