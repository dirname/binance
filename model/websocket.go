package model

// WebsocketCommon websocket common (property subscribe property etc.)
type WebsocketCommon struct {
	Method string      `json:"method"`           // Method method
	Params interface{} `json:"params,omitempty"` // Params params
	ID     uint        `json:"id"`               // ID id
}

// WebsocketCommonResponse websocket common (subscribe unsubscribe property etc.)
type WebsocketCommonResponse struct {
	Result interface{} `json:"result"` // Result result
	ID     uint        `json:"id"`     // ID id
}

// WebsocketErrorResponse websocket error message
type WebsocketErrorResponse struct {
	Code    int    `json:"code"` // Code error code
	Message string `json:"msg"`  // Message error message
}
