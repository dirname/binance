package spotclient

// SystemStatusResponse system status
type SystemStatusResponse struct {
	Status  int    `json:"status"` // Status 0 is normal 1 system maintenance
	Message string `json:"msg"`
}
