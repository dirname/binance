# Binance API Go Language SDK

[![Go Report Card](https://goreportcard.com/badge/github.com/dirname/Binance?style=flat-square)](https://goreportcard.com/report/github.com/dirname/Binance)
[![Build Status](https://img.shields.io/travis/dirname/Binance?style=flat-square)](https://travis-ci.org/dirname/Binance)
[![codecov](https://img.shields.io/codecov/c/gh/dirname/binance/main?style=flat-square&token=A6U52MYCXN)](https://codecov.io/gh/dirname/Binance)
[![license](https://img.shields.io/github/license/dirname/Binance?style=flat-square)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/dirname/Binance?style=flat-square)](https://golang.org)
[![GoDoc](https://img.shields.io/badge/reference-007d9c?style=flat-square&logo=Go&logoColor=F9F9F9&labelColor=5C5C5C&labelWidth=80)](https://pkg.go.dev/github.com/dirname/Binance)
[![Release](https://img.shields.io/github/release/dirname/Binance.svg?style=flat-square)](https://github.com/dirname/Binance/releases)

This is a Binance Go language sdk that uses a method similar
to [HuobiRDCenter/huobi_Golang](https://github.com/huobirdcenter/huobi_golang)

# Official Documents

Please make sure you have read the [Binance API document](https://binance-docs.github.io/apidocs/)
before continuing.

# List of implemented APIs

The following table shows the functions included in this SDK

## USDⓈ-M Futures

Category | Client | Access Type
:------------: | :------------: | :------------:
Market | Aggregate Trade | WebSocket
| | Market Price | WebSocket
| | All Market Price | WebSocket
| | Kline/Candlestick | WebSocket
| | Continuous Candlestick | WebSocket
| | Individual Symbol Mini Ticker | WebSocket
| | All Market Mini Tickers | WebSocket
| | Individual Symbol Ticker | WebSocket
| | All Book Tickers | WebSocket
| | Liquidation Order | WebSocket
| | All Market Liquidation Orders | WebSocket
| | Partial Book Depth | WebSocket
| | Diff. Book Depth | WebSocket
| | BLVT Info | WebSocket
| | BLVT NAV Kline/Candlestick | WebSocket
| | Composite Index Symbol | WebSocket

# Usage

Below are some examples of usage

## WebSocket market quotes

Take the Kline of the USDⓈ-M Futures as an example

```go
func main() {
	client := futuresusdt.NewUSDTFuturesCandlestickWebsocketClient("btcusdt@kline_1m")
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@kline_1m", "btcusdt@kline_5m")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresusdt.CandlestickResponse:
			logger.Info("Candlestick Response: %v", response.(futuresusdt.CandlestickResponse))
		case futuresusdt.CandlestickCombinedResponse:
			logger.Info("CandlestickCombinedResponse: %v", response.(futuresusdt.CandlestickCombinedResponse))
		case model.WebsocketCommonResponse:
			logger.Info("Websocket Response: %v", response.(model.WebsocketCommonResponse))
		case model.WebsocketErrorResponse:
			logger.Info("Websocket Error Response: %v", response.(model.WebsocketErrorResponse))
		default:
			logger.Info("Unknown Response: %v", response)
		}
	})
	client.Connect(true)
	fmt.Scanln()

	client.Unsubscribe(123, "btcusdt@kline_1m", "btcusdt@kline_5m")
	client.Close()
	logger.Info("Client closed")
}
```

# Disclaimer

This SDK is for **personal use** and has not been **officially produced**. They are only used to help users become
familiar with API endpoint. Please use it with caution, and expand the R&D according to your own situation.

**In addition**, the performance and stability of this sdk need to be tested by yourself, and the loss due to
performance or stability shall be **borne by you**