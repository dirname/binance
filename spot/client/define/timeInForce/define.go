package timeInForce

const (
	GTC = "GTC" // Good Til Canceled An order will be on the book unless the order is canceled.
	IOC = "IOC" // Immediate Or Cancel An order will try to fill the order as much as it can before the order expires.
	FOK = "FOK" // Fill or Kill An order will expire if the full order cannot be filled upon execution.
)
