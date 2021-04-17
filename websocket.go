package binance

import (
	"bytes"
	"fmt"
	"github.com/dirname/Binance/config"
	logger "github.com/dirname/Binance/logging"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

// ConnectedHandler invoked after websocket connected
type ConnectedHandler func()

// MessageHandler invoked after valid message received
type MessageHandler func(message []byte) (interface{}, error)

// ResponseHandler invoked after response is parsed
type ResponseHandler func(response interface{})

// PingHandler invoked after ws received ping frame
type PingHandler func(appData string) error

// PongHandler invoked after ws received pong frame
type PongHandler func(appData string) error

// WebsocketClient websocket client class that get data from websocket
type WebsocketClient struct {
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

// Init Initializer websocket client
func (w *WebsocketClient) Init(host string, stream ...string) {
	w.host = host
	if len(stream) == 0 {
		logger.Error("Stream cannot be empty")
		return
	}
	streams := "/ws/" + stream[0]
	if len(stream) > 1 {
		streams = "/stream?streams="
		var buffer bytes.Buffer
		for _, v := range stream {
			buffer.WriteString(v + "/")
		}
		streams = streams + buffer.String()[:buffer.Len()-1]
	}
	w.stream = streams
	w.stopReadChannel = make(chan int, 1)
	w.stopTickerChannel = make(chan int, 1)
	w.stopKeepAliveChannel = make(chan int, 1)
	w.sendMutex = &sync.Mutex{}
	switch host {
	case config.USDTFuturesWssHost:
		w.keepAliveInterval = 5 * time.Minute
	case config.CoinFuturesWssHost:
		w.keepAliveInterval = 5 * time.Minute
	case config.SpotWssHost:
		w.keepAliveInterval = 3 * time.Minute
	default:
		w.keepAliveInterval = 3 * time.Minute
	}
	logger.Debug("Stream splicing: %s", streams)
}

// SetConnectedHandler set client connected handler
func (w *WebsocketClient) SetConnectedHandler(handler ConnectedHandler) {
	w.connectedHandler = handler
}

// SetMessageHandler set client on message handler
func (w *WebsocketClient) SetMessageHandler(handler MessageHandler) {
	w.messageHandler = handler
}

// SetResponseHandler set client on response handler
func (w *WebsocketClient) SetResponseHandler(handler ResponseHandler) {
	w.responseHandler = handler
}

// SetReconnectWaitTime set time wait ws reconnect
func (w *WebsocketClient) SetReconnectWaitTime(time time.Duration) {
	w.reconnectWaitTime = time
}

// SetReadTimerInterval set time wait ws to received
func (w *WebsocketClient) SetReadTimerInterval(time time.Duration) {
	w.readTimerInterval = time
}

// SetKeepAliveInterval set time to keep alive
func (w *WebsocketClient) SetKeepAliveInterval(time time.Duration) {
	w.keepAliveInterval = time
}

// SetPingHandler set client on ping handler
func (w *WebsocketClient) SetPingHandler(handler PingHandler) {
	if w.conn == nil {
		logger.Error("Websocket sent error: No established websocket connection")
		return
	}
	w.conn.SetPingHandler(handler)
}

// SetPongHandler set client on pong handler
func (w *WebsocketClient) SetPongHandler(handler PongHandler) {
	if w.conn == nil {
		logger.Error("Websocket sent error: No established websocket connection")
		return
	}
	w.conn.SetPongHandler(handler)
}

// Connect connect to websocket server
func (w *WebsocketClient) Connect(autoReconnect bool) {
	w.connectWebsocket()
	w.startKeepAliveTicker()
	if autoReconnect {
		w.startTicker()
	}
}

// Send send the message to ws client
func (w *WebsocketClient) Send(data string) {
	if w.conn == nil {
		logger.Error("Websocket sent error: No established websocket connection")
		return
	}

	w.sendMutex.Lock()
	err := w.conn.WriteMessage(websocket.TextMessage, []byte(data))
	w.sendMutex.Unlock()

	if err != nil {
		logger.Error("Websocket sent %s error: %s", data, err.Error())
	}
}

// SendJSON send the json to ws client
func (w *WebsocketClient) SendJSON(data interface{}) {
	if w.conn == nil {
		logger.Error("Websocket sent error: No established websocket connection")
		return
	}

	w.sendMutex.Lock()
	err := w.conn.WriteJSON(data)
	w.sendMutex.Unlock()

	if err != nil {
		logger.Error("Websocket sent json %v error: %s", data, err.Error())
	}
}

// Close close the connection to server
func (w *WebsocketClient) Close() {
	w.stopKeepAliveTicker()
	w.stopTicker()
	w.disconnectWebsocket()
}

func (w *WebsocketClient) connectWebsocket() {
	var err error
	url := fmt.Sprintf("wss://%s%s", w.host, w.stream)
	logger.Debug("Start connecting %s...", url)
	w.conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		logger.Error("Websocket connection error: %s", err.Error())
		return
	}
	logger.Info("%s connected", url)
	w.lastReceivedTime = time.Now()
	w.establishmentTime = time.Now()
	w.startReadLoop()
	if w.connectedHandler != nil {
		w.connectedHandler()
	}
}

func (w *WebsocketClient) disconnectWebsocket() {
	if w.conn == nil {
		return
	}

	go w.stopReadLoop()
	logger.Debug("Websocket disconnecting...")
	err := w.conn.Close()
	if err != nil {
		logger.Error("Websocket failed to disconnecting error: %s", err.Error())
		return
	}

	logger.Info("Websocket disconnected")
}

func (w *WebsocketClient) startTicker() {
	w.ticker = time.NewTicker(w.readTimerInterval)
	go w.tickerLoop()
}

func (w *WebsocketClient) stopTicker() {
	w.ticker.Stop()
	w.stopTickerChannel <- 1
}

func (w *WebsocketClient) tickerLoop() {
	logger.Debug("Ticker loop started")
	for {
		select {
		case <-w.stopTickerChannel:
			logger.Debug("Ticker loop stopped")
			return

		case <-w.ticker.C:
			elapsedSecond := time.Now().Sub(w.lastReceivedTime).Microseconds()
			//logger.Debug("Last received data was %f sec ago", elapsedSecond)
			if elapsedSecond > w.reconnectWaitTime.Microseconds() {
				logger.Info("Websocket started reconnect...")
				w.disconnectWebsocket()
				w.connectWebsocket()
			}
		}
	}
}

func (w *WebsocketClient) startKeepAliveTicker() {
	w.keepAliveTicker = time.NewTicker(w.keepAliveInterval)
	go w.keepAliveLoop()
}

func (w *WebsocketClient) stopKeepAliveTicker() {
	w.keepAliveTicker.Stop()
	w.stopKeepAliveChannel <- 1
}

func (w *WebsocketClient) keepAliveLoop() {
	logger.Debug("Keep alive loop started")
	for {
		select {
		case <-w.stopKeepAliveChannel:
			logger.Debug("Keep alive stopped")
			return

		case <-w.keepAliveTicker.C:
			deadline := time.Now().Add(10 * time.Second)
			err := w.conn.WriteControl(websocket.PongMessage, []byte{}, deadline)
			if err != nil {
				logger.Warn("Response to Pong failed: %s reconnect....", err.Error())
				w.disconnectWebsocket()
				w.connectWebsocket()
			}

			elapsedSecond := time.Now().Sub(w.establishmentTime).Seconds()
			if elapsedSecond >= 86400 {
				logger.Info("Exceed the maximum time of connection valid, start to reconnect...")
				w.disconnectWebsocket()
				w.connectWebsocket()
			}
		}
	}
}

func (w *WebsocketClient) startReadLoop() {
	go w.readLoop()
}

func (w *WebsocketClient) stopReadLoop() {
	w.stopReadChannel <- 1
}

func (w *WebsocketClient) readLoop() {
	logger.Debug("Read loop started")
	for {
		select {
		case <-w.stopReadChannel:
			logger.Debug("Read loop stopped")
			return
		default:
			if w.conn == nil {
				logger.Error("Read error: No established websocket connection")
				time.Sleep(w.readTimerInterval)
				continue
			}

			msgType, buf, err := w.conn.ReadMessage()

			if err != nil {
				logger.Error("Read error: %s", err)
				time.Sleep(w.readTimerInterval)
				continue
			}

			w.lastReceivedTime = time.Now()

			if msgType == websocket.TextMessage {
				result, err := w.messageHandler(buf)
				if err != nil {
					logger.Error("Handle message error: %s", err)
				}
				if w.responseHandler != nil {
					w.responseHandler(result)
				}
			}
		}
	}
}
