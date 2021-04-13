# Binance API Go Language SDK

[![Go Report Card](https://goreportcard.com/badge/github.com/dirname/Binance?style=flat-square)](https://goreportcard.com/report/github.com/dirname/Binance)
[![Build Status](https://img.shields.io/travis/dirname/Binance?style=flat-square)](https://travis-ci.org/dirname/Binance)
[![codecov](https://img.shields.io/codecov/c/gh/dirname/binance/main?style=flat-square&token=A6U52MYCXN)](https://codecov.io/gh/dirname/Binance)
[![license](https://img.shields.io/github/license/dirname/Binance?style=flat-square)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/dirname/Binance?style=flat-square)](https://golang.org)
[![GoDoc](https://img.shields.io/badge/reference-007d9c?style=flat-square&logo=Go&logoColor=F9F9F9&labelColor=5C5C5C&labelWidth=80)](https://pkg.go.dev/github.com/dirname/Binance)
[![Release](https://img.shields.io/github/release/dirname/Binance.svg?style=flat-square)](https://github.com/dirname/Binance/releases)
[![Release Date](https://img.shields.io/github/release-date/dirname/Binance?style=flat-square)](https://github.com/dirname/Binance/releases)
[![Commit](https://img.shields.io/github/last-commit/dirname/Binance?style=flat-square)](https://github.com/dirname/Binance/commits)

This is a Binance Go language sdk that uses a method similar
to [HuobiRDCenter/huobi_Golang](https://github.com/huobirdcenter/huobi_golang)

# Official Documents

Please make sure you have read the [Binance API document](https://binance-docs.github.io/apidocs/)
before continuing.

# Precautions

You are very welcome to submit issues or pull requests to share with this project to make the project more perfect

- *Remove all WAPI*

  *As Binance has planned to remove all WAPI endpoints from the Binance API at 2021-08-01 2:00 AM (UTC). So this
  warehouse has removed all WAPI endpoints from the Binance API in advance*
  *[More Information](https://www.binance.com/en/support/announcement/f45dde7da58b473aa885349946bed269)*

- [spot/client/market.go:Line185](https://github.com/dirname/Binance/blob/main/spot/client/market.go#L185)
  `GetCandlestick()`

  *It is not recommended using this method, this method may have type errors, please use websocket subscription to
  obtain data.*

# List of implemented APIs

The following table shows the functions included in this SDK

## Spot/Margin/Savings/Mining

Category | Client | Access Type
:------------: | :------------: | :------------:
Common | System Status (SAPI) | RESTful API
| | Test Connectivity | RESTful API
| | Check ServerTime | RESTful API
Wallet | All Coin's Information | RESTful API
| | Daily Account Snapshot | RESTful API
| | Disabled/Enabled Fast Withdraw | RESTful API
| | Withdraw (SAPI) | RESTful API
| | Deposit History(supporting network) | RESTful API
| | Withdraw History(supporting network) | RESTful API
| | Deposit Address(supporting network) | RESTful API
| | Account Status (SAPI) | RESTful API
| | Account API Trading Status (SAPI) | RESTful API
| | DustLog (SAPI) | RESTful API
| | Dust Transfer | RESTful API
| | Asset Dividend Record | RESTful API
| | Asset Detail (SAPI) | RESTful API
| | Trade Fee (SAPI) | RESTful API
| | User Universal Transfer | RESTful API
| | Query User Universal Transfer History | RESTful API
Market | Order Book | RESTful API
| | Recent Trades List | RESTful API
| | Old Trade Lookup | RESTful API
| | Aggregate Trade | RESTful API
| | Candlestick | RESTful API
| | Current Average Price | RESTful API
| | 24h Ticker Price Change | RESTful API
| | Symbol Price Ticker | RESTful API
| | Symbol Order Book Ticker | RESTful API
| | Aggregate Trade | WebSocket
| | Trade | WebSocket
| | Candlestick | WebSocket
| | Individual Symbol Mini Ticker | WebSocket
| | All Market Mini Tickers | WebSocket
| | Individual Symbol Ticker | WebSocket
| | All Market Tickers | WebSocket
| | Individual Symbol Book Ticker | WebSocket
| | All Book Tickers | WebSocket
| | Partial Book Depth | WebSocket
| | Diff. Book Depth | WebSocket
Trade | Test New Order | RESTful API
| | New order | RESTful API
| | Cancel order | RESTful API
| | Cancel all order | RESTful API
| | Query order | RESTful API
| | Current open orders | RESTful API
| | All orders | RESTful API
| | New OCO | RESTful API
| | Cancel OCO | RESTful API
| | Query OCO | RESTful API
| | Query all OCO | RESTful API
| | Query Open OCO | RESTful API
| | Account Information | RESTful API
| | Account Trade List | RESTful API

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

## RESTful wallet transfer

Examples of user transfers in the universal

```go
package spotclient

import (
	binance "github.com/dirname/Binance"
	"github.com/dirname/Binance/config"
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
	"time"
)

func init() {
	logger.Enable(false)
	walletClient = spotclient.NewWalletClient(config.SpotRestHost, config.AppKey, config.AppSecret)
}

func universalTransfer() {
	response, err := walletClient.UniversalTransfer("UMFUTURE_MAIN", "USDT", "10", 0)
	if err != nil {
		logger.Error("universalTransfer err: %s", err.Error())
	}
	switch response.(type) {
	case model.APIErrorResponse:
		logger.Info("universalTransfer API error: %v", response.(model.APIErrorResponse))
	case spotclient.UniversalTransferResponse:
		logger.Info("universalTransfer: %v", response.(spotclient.UniversalTransferResponse))
	default:
		logger.Info("universalTransfer Unknown response: %v", response)
	}
}
```

## WebSocket market quotes

Take the Kline of the USDⓈ-M Futures as an example

```go
package main

import (
	"fmt"
	"github.com/dirname/Binance/futures/usdt/websocket/market"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
)

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