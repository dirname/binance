package model

// APIErrorResponse API and SAPI error response
type APIErrorResponse struct {
	Code    int    `json:"code"` // error code
	Message string `json:"msg"`  // error msg
}

// WAPIErrorResponse WAPI error response
type WAPIErrorResponse struct {
	Success bool   `json:"success"` // is success
	Message string `json:"msg"`
}
