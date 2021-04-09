package spotclient

import binance "github.com/dirname/Binance"

// AccountClient responsible to get common information
type AccountClient struct {
	Builder *binance.PublicUrlBuilder
}

// NewAccountClient Initializer
func NewAccountClient(host string) *AccountClient {
	return &AccountClient{
		Builder: binance.NewPublicUrlBuilder(host),
	}
}
