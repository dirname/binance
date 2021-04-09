package spotclient

// AllCoinInfoResponse all coin info response
type AllCoinInfoResponse []struct {
	Coin             string `json:"coin"`
	DepositAllEnable bool   `json:"depositAllEnable"`
	Free             string `json:"free"`
	Freeze           string `json:"freeze"`
	IPOAble          string `json:"ipoable"`
	IPOIng           string `json:"ipoing"`
	IsLegalMoney     bool   `json:"isLegalMoney"`
	Locked           string `json:"locked"`
	Name             string `json:"name"`
	NetworkList      []struct {
		AddressRegex            string `json:"addressRegex"`
		Coin                    string `json:"coin"`
		DepositDesc             string `json:"depositDesc,omitempty"`
		DepositEnable           bool   `json:"depositEnable"`
		IsDefault               bool   `json:"isDefault"`
		MemoRegex               string `json:"memoRegex"`
		MinConfirm              int    `json:"minConfirm"`
		Name                    string `json:"name"`
		Network                 string `json:"network"`
		ResetAddressStatus      bool   `json:"resetAddressStatus"`
		SpecialTips             string `json:"specialTips"`
		UnLockConfirm           int    `json:"unLockConfirm"`
		WithdrawDesc            string `json:"withdrawDesc,omitempty"`
		WithdrawEnable          bool   `json:"withdrawEnable"`
		WithdrawFee             string `json:"withdrawFee"`
		WithdrawMin             string `json:"withdrawMin"`
		InsertTime              int64  `json:"insertTime,omitempty"`
		UpdateTime              int64  `json:"updateTime,omitempty"`
		WithdrawIntegerMultiple string `json:"withdrawIntegerMultiple,omitempty"`
	} `json:"networkList"`
	Storage           string `json:"storage"`
	Trading           bool   `json:"trading"`
	WithdrawAllEnable bool   `json:"withdrawAllEnable"`
	Withdrawing       string `json:"withdrawing"`
}

// SpotSnapshotResponse spot snapshot response
type SpotSnapshotResponse struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	SnapshotVos []struct {
		Data struct {
			Balances []struct {
				Asset  string `json:"asset"`
				Free   string `json:"free"`
				Locked string `json:"locked"`
			} `json:"balances"`
			TotalAssetOfBtc string `json:"totalAssetOfBtc"`
		} `json:"data"`
		Type       string `json:"type"`
		UpdateTime int64  `json:"updateTime"`
	} `json:"snapshotVos"`
}

// FuturesSnapshotResponse futures snapshot response
type FuturesSnapshotResponse struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	SnapshotVos []struct {
		Data struct {
			Assets []struct {
				Asset         string `json:"asset"`
				MarginBalance string `json:"marginBalance"`
				WalletBalance string `json:"walletBalance"`
			} `json:"assets"`
			Position []struct {
				EntryPrice       string `json:"entryPrice"`
				MarkPrice        string `json:"markPrice"`
				PositionAmt      string `json:"positionAmt"`
				Symbol           string `json:"symbol"`
				UnRealizedProfit string `json:"unRealizedProfit"`
			} `json:"position"`
		} `json:"data"`
		Type       string `json:"type"`
		UpdateTime int64  `json:"updateTime"`
	} `json:"snapshotVos"`
}

// MarginSnapshotResponse margin snapshot response
type MarginSnapshotResponse struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	SnapshotVos []struct {
		Data struct {
			MarginLevel         string `json:"marginLevel"`
			TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
			TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
			TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`
			UserAssets          []struct {
				Asset    string `json:"asset"`
				Borrowed string `json:"borrowed"`
				Free     string `json:"free"`
				Interest string `json:"interest"`
				Locked   string `json:"locked"`
				NetAsset string `json:"netAsset"`
			} `json:"userAssets"`
		} `json:"data"`
		Type       string `json:"type"`
		UpdateTime int64  `json:"updateTime"`
	} `json:"snapshotVos"`
}

// SAPIWithdrawResponse sapi withdraw response
type SAPIWithdrawResponse struct {
	ID string `json:"id"`
}

// WAPIWithdrawResponse wapi withdraw response
type WAPIWithdrawResponse struct {
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
	ID      string `json:"id"`
}

// DepositHistoryNetworkResponse DepositHistoryNetwork response
type DepositHistoryNetworkResponse []struct {
	Amount       string `json:"amount"`
	Coin         string `json:"coin"`
	Network      string `json:"network"`
	Status       int    `json:"status"`
	Address      string `json:"address"`
	AddressTag   string `json:"addressTag"`
	TxID         string `json:"txId"`
	InsertTime   int64  `json:"insertTime"`
	TransferType int    `json:"transferType"`
	ConfirmTimes string `json:"confirmTimes"`
}

// DepositHistoryResponse DepositHistory response
type DepositHistoryResponse struct {
	DepositList []struct {
		InsertTime int64   `json:"insertTime"`
		Amount     float64 `json:"amount"`
		Asset      string  `json:"asset"`
		Address    string  `json:"address"`
		TxID       string  `json:"txId"`
		Status     int     `json:"status"`
		AddressTag string  `json:"addressTag,omitempty"`
	} `json:"depositList"`
	Success bool `json:"success"`
}

// WithdrawHistoryNetworkResponse withdraw network response
type WithdrawHistoryNetworkResponse []struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	ApplyTime       string `json:"applyTime"`
	Coin            string `json:"coin"`
	ID              string `json:"id"`
	WithdrawOrderID string `json:"withdrawOrderId,omitempty"`
	Network         string `json:"network"`
	TransferType    int    `json:"transferType,omitempty"`
	Status          int    `json:"status"`
	TxID            string `json:"txId"`
}

// WithdrawHistoryResponse withdraw response
type WithdrawHistoryResponse struct {
	WithdrawList []struct {
		ID              string      `json:"id"`
		WithdrawOrderID interface{} `json:"withdrawOrderId"`
		Amount          float64     `json:"amount"`
		TransactionFee  float64     `json:"transactionFee"`
		Address         string      `json:"address"`
		Asset           string      `json:"asset"`
		TxID            string      `json:"txId"`
		ApplyTime       int64       `json:"applyTime"`
		Status          int         `json:"status"`
		AddressTag      string      `json:"addressTag,omitempty"`
	} `json:"withdrawList"`
	Success bool `json:"success"`
}

// DepositAddressNetworkResponse deposit address response
type DepositAddressNetworkResponse struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	Url     string `json:"url"`
}

// DepositAddressResponse deposit address response
type DepositAddressResponse struct {
	Address    string `json:"address"`
	Success    bool   `json:"success"`
	AddressTag string `json:"addressTag"`
	Asset      string `json:"asset"`
}
