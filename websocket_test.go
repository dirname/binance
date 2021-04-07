package binance

import (
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
			keepAliveTicker:      time.NewTicker(TimerIntervalSecond * time.Second),
			ticker:               time.NewTicker(TimerIntervalSecond * time.Second),
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
			host:                 "wss://echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(TimerIntervalSecond * time.Second),
			ticker:               time.NewTicker(TimerIntervalSecond * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    10,
		}, args{autoReconnect: false}},
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
			u.Init("wss://echo.websocket.org", "")
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
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestWebsocketClient_Send", fields{
			host:                 "wss://echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(TimerIntervalSecond * time.Second),
			ticker:               time.NewTicker(TimerIntervalSecond * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    10,
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
			u.Connect(false)
			u.Send(tt.args.data)
			u.Close()
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
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestWebsocketClient_SendJSON", fields{
			host:                 "wss://echo.websocket.org",
			stream:               "",
			conn:                 nil,
			connectedHandler:     nil,
			messageHandler:       nil,
			responseHandler:      nil,
			stopReadChannel:      make(chan int, 1),
			stopTickerChannel:    make(chan int, 1),
			stopKeepAliveChannel: make(chan int, 1),
			keepAliveTicker:      time.NewTicker(TimerIntervalSecond * time.Second),
			ticker:               time.NewTicker(TimerIntervalSecond * time.Second),
			sendMutex:            &sync.Mutex{},
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    10,
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
			u.Connect(false)
			u.SendJSON(tt.args.data)
			u.Close()
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
			u.SetConnectedHandler(func() {

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
			_ = &WebsocketClient{
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
			_ = &WebsocketClient{
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
			_ = &WebsocketClient{
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
			_ = &WebsocketClient{
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
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &WebsocketClient{
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
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &WebsocketClient{
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
			stopReadChannel:      nil,
			stopTickerChannel:    nil,
			stopKeepAliveChannel: nil,
			keepAliveTicker:      nil,
			ticker:               nil,
			sendMutex:            nil,
			lastReceivedTime:     time.Time{},
			establishmentTime:    time.Time{},
			keepAliveInterval:    0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &WebsocketClient{
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
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &WebsocketClient{
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
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &WebsocketClient{
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
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &WebsocketClient{
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
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"TestWebsocketClient_startTicker", fields{
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
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &WebsocketClient{
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
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &WebsocketClient{
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
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &WebsocketClient{
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
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"TestWebsocketClient_stopTicker", fields{
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
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &WebsocketClient{
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
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &WebsocketClient{
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
		})
	}
}
