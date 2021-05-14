package futuresclient

//DualSidePositionResponse Get user's position mode (Hedge Mode or One-way Mode ) on EVERY symbol
type DualSidePositionResponse struct {
	DualSidePosition bool `json:"dualSidePosition"`
}

//MultiAssetsMarginResponse Get user's Multi-Assets mode (Multi-Assets Mode or Single-Asset Mode) on Every symbol
type MultiAssetsMarginResponse struct {
	MultiAssetsMargin bool `json:"multiAssetsMargin"`
}

