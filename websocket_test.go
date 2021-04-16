package binance

import (
	"github.com/dirname/Binance/config"
	"github.com/gorilla/websocket"
	"sync"
	"testing"
	"time"
)

func TestWebsocketClient_Close(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"TestWebsocketClient_Close", fields{
			host:                 "",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
			}
			u.Close()
		})
	}
}

func TestWebsocketClient_Connect(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
		reconnectWaitTime    time.Duration
		readTimerInterval    time.Duration
	}
	type args struct {
		autoReconnect bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestWebsocketClient_Connect", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    10 * time.Second,
			readTimerInterval:    1 * time.Second,
			reconnectWaitTime:    1 * time.Second,
		}, args{autoReconnect: false}},
		{"TestWebsocketClient_Connect", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    10 * time.Second,
			readTimerInterval:    1 * time.Second,
			reconnectWaitTime:    1 * time.Second,
		}, args{autoReconnect: true}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
				readTimerInterval:    1 * time.Second,
				reconnectWaitTime:    1 * time.Second,
			}
			u.Connect(tt.args.autoReconnect)
			u.Close()
		})
	}
}

func TestWebsocketClient_Init(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
	}
	type args struct {
		host   string
		stream []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "TestWebsocketClient_Init",
			fields: fields{},
			args:   args{},
		},
		{
			name:   "TestWebsocketClient_Init",
			fields: fields{},
			args: args{
				host:   config.SpotWssHost,
				stream: []string{"test", "test"},
			},
		},
		{
			name:   "TestWebsocketClient_Init",
			fields: fields{},
			args: args{
				host:   config.CoinFuturesWssHost,
				stream: []string{"test", "test"},
			},
		},
		{
			name:   "TestWebsocketClient_Init",
			fields: fields{},
			args: args{
				host:   config.USDTFuturesWssHost,
				stream: []string{"test", "test"},
			},
		},
		{
			name:   "TestWebsocketClient_Init",
			fields: fields{},
			args: args{
				host:   "echo.websocket.org",
				stream: []string{"test", "test"},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{}
			u.Init(tt.args.host, tt.args.stream...)
		})
	}
}

func TestWebsocketClient_Send(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
	}
	type args struct {
		data string
	}
	conn, _, _ := websocket.DefaultDialer.Dial("wss://echo.websocket.org", nil)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestWebsocketClient_Send", fields{
			host:                 "",
			stream:               "",
			conn:                 conn,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    10 * time.Second,
		}, args{"test"}},
		{"TestWebsocketClient_Send", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    10 * time.Second,
		}, args{"test"}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
			}
			u.Send(tt.args.data)
		})
	}
}

func TestWebsocketClient_SendJSON(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
	}
	type args struct {
		data interface{}
	}
	conn, _, _ := websocket.DefaultDialer.Dial("wss://echo.websocket.org", nil)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestWebsocketClient_SendJSON", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    10 * time.Second,
		}, args{"{\"msg\":\"test\"}"}},
		{"TestWebsocketClient_SendJSON", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 conn,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    10 * time.Second,
		}, args{"{\"msg\":\"test\"}"}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
			}
			u.SendJSON(tt.args.data)
		})
	}
}

func TestWebsocketClient_SetConnectedHandler(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
	}
	type args struct {
		handler ConnectedHandler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestWebsocketClient_SetConnectedHandler", fields{
			host:                 "",
			stream:               "",
			conn:                 &websocket.Conn{},
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      nil,
			stopTickerChannel:    nil,
			stopKeepAliveChannel: nil,
			keepAliveTicker:      nil,
			ticker:               nil,
			sendMutex:            nil,
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    0,
		}, args{nil}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
			}
			u.SetConnectedHandler(func() {
				return
			})
		})
	}
}

func TestWebsocketClient_SetMessageHandler(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
	}
	type args struct {
		handler MessageHandler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestWebsocketClient_SetMessageHandler", fields{
			host:                 "",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      nil,
			stopTickerChannel:    nil,
			stopKeepAliveChannel: nil,
			keepAliveTicker:      nil,
			ticker:               nil,
			sendMutex:            nil,
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    0,
		}, args{nil}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
			}
			u.SetMessageHandler(func(message []byte) (interface{}, error) {
				return nil, nil
			})
		})
	}
}

func TestWebsocketClient_SetPingHandler(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
	}
	type args struct {
		handler PingHandler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestWebsocketClient_SetPingHandler", fields{
			host:                 "",
			stream:               "",
			conn:                 &websocket.Conn{},
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      nil,
			stopTickerChannel:    nil,
			stopKeepAliveChannel: nil,
			keepAliveTicker:      nil,
			ticker:               nil,
			sendMutex:            nil,
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    0,
		}, args{nil}},
		{"TestWebsocketClient_SetPingHandler", fields{
			host:                 "",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      nil,
			stopTickerChannel:    nil,
			stopKeepAliveChannel: nil,
			keepAliveTicker:      nil,
			ticker:               nil,
			sendMutex:            nil,
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    0,
		}, args{nil}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
			}
			u.SetPingHandler(func(appData string) error {
				return nil
			})
		})
	}
}

func TestWebsocketClient_SetPongHandler(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
	}
	type args struct {
		handler PongHandler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestWebsocketClient_SetPongHandler", fields{
			host:                 "",
			stream:               "",
			conn:                 &websocket.Conn{},
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      nil,
			stopTickerChannel:    nil,
			stopKeepAliveChannel: nil,
			keepAliveTicker:      nil,
			ticker:               nil,
			sendMutex:            nil,
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    0,
		}, args{nil}},
		{"TestWebsocketClient_SetPongHandler", fields{
			host:                 "",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      nil,
			stopTickerChannel:    nil,
			stopKeepAliveChannel: nil,
			keepAliveTicker:      nil,
			ticker:               nil,
			sendMutex:            nil,
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    0,
		}, args{nil}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
			}
			u.SetPongHandler(func(appData string) error {
				return nil
			})
		})
	}
}

func TestWebsocketClient_SetResponseHandler(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
	}
	type args struct {
		handler ResponseHandler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestWebsocketClient_SetResponseHandler", fields{
			host:                 "",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      nil,
			stopTickerChannel:    nil,
			stopKeepAliveChannel: nil,
			keepAliveTicker:      nil,
			ticker:               nil,
			sendMutex:            nil,
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    0,
		}, args{nil}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
			}
			u.SetResponseHandler(func(response interface{}) {
				return
			})
		})
	}
}

func TestWebsocketClient_connectWebsocket(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"TestWebsocketClient_connectWebsocket", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    10 * time.Second,
		}},
		{"TestWebsocketClient_connectWebsocket", fields{
			host:   "echo.websocket.org",
			stream: "",
			conn:   nil,
			connectedHandler: func() {
				return
			},
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    10 * time.Second,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
			}
			u.connectWebsocket()
			u.Close()
		})
	}
}

func TestWebsocketClient_disconnectWebsocket(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"TestWebsocketClient_disconnectWebsocket", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    10 * time.Second,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
			}
			u.Connect(false)
			u.disconnectWebsocket()
			u.Close()
		})
	}
}

func TestWebsocketClient_keepAliveLoop(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"TestWebsocketClient_keepAliveLoop", fields{
			host:                 "",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Now(),
			establishmentTime:    time.Now(),
			keepAliveInterval:    10 * time.Second,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
			}
			go u.keepAliveLoop()
			u.stopKeepAliveTicker()
		})
	}
}

func TestWebsocketClient_readLoop(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"TestWebsocketClient_readLoop", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Now(),
			establishmentTime:    time.Now(),
			keepAliveInterval:    10 * time.Second,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
			}
			go u.readLoop()
			u.stopReadLoop()
		})
	}
}

func TestWebsocketClient_startKeepAliveTicker(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"TestWebsocketClient_startKeepAliveTicker", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Now(),
			establishmentTime:    time.Now(),
			keepAliveInterval:    10 * time.Second,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
			}
			u.startKeepAliveTicker()
			time.Sleep(2 * time.Second)
			u.stopKeepAliveTicker()
		})
	}
}

func TestWebsocketClient_startReadLoop(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"TestWebsocketClient_startReadLoop", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Now(),
			establishmentTime:    time.Now(),
			keepAliveInterval:    10 * time.Second,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
			}
			u.startReadLoop()
			u.stopReadLoop()
		})
	}
}

func TestWebsocketClient_startTicker(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
		readTimerInterval    time.Duration
		reconnectWaitTime    time.Duration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"TestWebsocketClient_startTicker", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Now(),
			establishmentTime:    time.Now(),
			keepAliveInterval:    10 * time.Second,
			readTimerInterval:    1 * time.Second,
			reconnectWaitTime:    1 * time.Second,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
				readTimerInterval:    tt.fields.readTimerInterval,
				reconnectWaitTime:    tt.fields.reconnectWaitTime,
			}
			u.startTicker()
			u.stopTicker()
		})
	}
}

func TestWebsocketClient_stopKeepAliveTicker(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"TestWebsocketClient_stopKeepAliveTicker", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Now(),
			establishmentTime:    time.Now(),
			keepAliveInterval:    10 * time.Second,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
			}
			u.startKeepAliveTicker()
			u.stopKeepAliveTicker()
		})
	}
}

func TestWebsocketClient_stopReadLoop(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"TestWebsocketClient_stopReadLoop", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Now(),
			establishmentTime:    time.Now(),
			keepAliveInterval:    10 * time.Second,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
			}
			u.connectWebsocket()
			u.startReadLoop()
			u.stopReadLoop()
			u.disconnectWebsocket()
		})
	}
}

func TestWebsocketClient_stopTicker(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
		readTimerInterval    time.Duration
		reconnectWaitTime    time.Duration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"TestWebsocketClient_stopTicker", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Now(),
			establishmentTime:    time.Now(),
			keepAliveInterval:    10 * time.Second,
			readTimerInterval:    1 * time.Second,
			reconnectWaitTime:    1 * time.Second,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
				readTimerInterval:    tt.fields.readTimerInterval,
				reconnectWaitTime:    tt.fields.reconnectWaitTime,
			}
			u.startTicker()
			u.stopTicker()
		})
	}
}

func TestWebsocketClient_tickerLoop(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"TestWebsocketClient_tickerLoop", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Now(),
			establishmentTime:    time.Now(),
			keepAliveInterval:    10 * time.Second,
		}},
		{"TestWebsocketClient_tickerLoop", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Unix(0, (time.Now().Unix()-860000)*int64(time.Millisecond)),
			establishmentTime:    time.Now(),
			keepAliveInterval:    10 * time.Second,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
			}
			go u.tickerLoop()
			time.Sleep(2 * time.Second)
			u.stopTicker()
		})
	}
}

func TestWebsocketClient_SetReadTimerInterval(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
		reconnectWaitTime    time.Duration
		readTimerInterval    time.Duration
	}
	type args struct {
		time time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestWebsocketClient_SetReadTimerInterval", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Now(),
			establishmentTime:    time.Now(),
			keepAliveInterval:    10 * time.Second,
		}, args{10 * time.Microsecond}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
				reconnectWaitTime:    tt.fields.reconnectWaitTime,
				readTimerInterval:    tt.fields.readTimerInterval,
			}
			w.SetReadTimerInterval(tt.args.time)
		})
	}
}

func TestWebsocketClient_SetReconnectWaitTime(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
		reconnectWaitTime    time.Duration
		readTimerInterval    time.Duration
	}
	type args struct {
		time time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestWebsocketClient_SetReconnectWaitTime", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Now(),
			establishmentTime:    time.Now(),
			keepAliveInterval:    10 * time.Second,
		}, args{10 * time.Microsecond}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
				reconnectWaitTime:    tt.fields.reconnectWaitTime,
				readTimerInterval:    tt.fields.readTimerInterval,
			}
			w.SetReconnectWaitTime(tt.args.time)
		})
	}
}

func TestWebsocketClient_SetKeepAliveInterval(t *testing.T) {
	type fields struct {
		host                 string
		stream               string
		conn                 *websocket.Conn
		connectedHandler     ConnectedHandler
		messageHandler       MessageHandler
		responseHandler      ResponseHandler
		stopReadChannel      chan int
		stopTickerChannel    chan int
		stopKeepAliveChannel chan int
		keepAliveTicker      *time.Ticker
		ticker               *time.Ticker
		sendMutex            *sync.Mutex
		lastReceivedTime     time.Time
		establishmentTime    time.Time
		keepAliveInterval    time.Duration
		reconnectWaitTime    time.Duration
		readTimerInterval    time.Duration
	}
	type args struct {
		time time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestWebsocketClient_SetKeepAliveInterval", fields{
			host:                 "echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(1 * time.Second),
			ticker:               time.NewTicker(1 * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Now(),
			establishmentTime:    time.Now(),
			keepAliveInterval:    10 * time.Second,
		}, args{10 * time.Microsecond}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WebsocketClient{
				host:                 tt.fields.host,
				stream:               tt.fields.stream,
				conn:                 tt.fields.conn,
				connectedHandler:     tt.fields.connectedHandler,
				messageHandler:       tt.fields.messageHandler,
				responseHandler:      tt.fields.responseHandler,
				stopReadChannel:      tt.fields.stopReadChannel,
				stopTickerChannel:    tt.fields.stopTickerChannel,
				stopKeepAliveChannel: tt.fields.stopKeepAliveChannel,
				keepAliveTicker:      tt.fields.keepAliveTicker,
				ticker:               tt.fields.ticker,
				sendMutex:            tt.fields.sendMutex,
				lastReceivedTime:     tt.fields.lastReceivedTime,
				establishmentTime:    tt.fields.establishmentTime,
				keepAliveInterval:    tt.fields.keepAliveInterval,
				reconnectWaitTime:    tt.fields.reconnectWaitTime,
				readTimerInterval:    tt.fields.readTimerInterval,
			}
			w.SetKeepAliveInterval(tt.args.time)
		})
	}
}
